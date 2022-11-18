package guardrails

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"bytes"
	"context"
	"embed"
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"

	arov1alpha1 "github.com/Azure/ARO-RP/pkg/operator/apis/aro.openshift.io/v1alpha1"
	aroclient "github.com/Azure/ARO-RP/pkg/operator/clientset/versioned"
	"github.com/Azure/ARO-RP/pkg/operator/controllers/guardrails/config"
	"github.com/Azure/ARO-RP/pkg/util/deployer"
	"github.com/Azure/ARO-RP/pkg/util/dynamichelper"
	"github.com/Azure/ARO-RP/pkg/util/version"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	kruntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"sigs.k8s.io/yaml"
)

const (
	ControllerName               = "GuardRails"
	controllerEnabled            = "aro.guardrails.enabled"        // boolean, false by default
	controllerNamespace          = "aro.guardrails.namespace"      // string
	controllerManaged            = "aro.guardrails.deploy.managed" // trinary, do-nothing by default
	controllerPullSpec           = "aro.guardrails.deploy.pullspec"
	controllerManagerRequestsCPU = "aro.guardrails.deploy.manager.requests.cpu"
	controllerManagerRequestsMem = "aro.guardrails.deploy.manager.requests.mem"
	controllerManagerLimitCPU    = "aro.guardrails.deploy.manager.limit.cpu"
	controllerManagerLimitMem    = "aro.guardrails.deploy.manager.limit.mem"
	controllerAuditRequestsCPU   = "aro.guardrails.deploy.audit.requests.cpu"
	controllerAuditRequestsMem   = "aro.guardrails.deploy.audit.requests.mem"
	controllerAuditLimitCPU      = "aro.guardrails.deploy.audit.limit.cpu"
	controllerAuditLimitMem      = "aro.guardrails.deploy.audit.limit.mem"

	controllerValidatingWebhookFailurePolicy = "aro.guardrails.validatingwebhook.managed"
	controllerValidatingWebhookTimeout       = "aro.guardrails.validatingwebhook.timeoutSeconds"
	controllerMutatingWebhookFailurePolicy   = "aro.guardrails.mutatingwebhook.managed"
	controllerMutatingWebhookTimeout         = "aro.guardrails.mutatingwebhook.timeoutSeconds"

	controllerReconciliationMinutes     = "aro.guardrails.reconciliationMinutes"                  // int, 60 by default.
	controllerPolicyMachineDenyManaged  = "aro.guardrails.policies.aro-machines-deny.managed"     // trinary, do-nothing by default
	controllerPolicyMachineDenyEnforced = "aro.guardrails.policies.aro-machines-deny.enforcement" // [“warn”, “dryrun”, “deny”], “dryrun” by default/on invalid input

	defaultNamespace = "openshift-azure-guardrails"

	defaultManagerRequestsCPU = "100m"
	defaultManagerLimitCPU    = "1000m"
	defaultManagerRequestsMem = "256Mi"
	defaultManagerLimitMem    = "512Mi"
	defaultAuditRequestsCPU   = "100m"
	defaultAuditLimitCPU      = "1000m"
	defaultAuditRequestsMem   = "256Mi"
	defaultAuditLimitMem      = "512Mi"

	defaultReconciliationMinutes = "60"

	defaultValidatingWebhookFailurePolicy = "Ignore"
	defaultValidatingWebhookTimeout       = "3"
	defaultMutatingWebhookFailurePolicy   = "Ignore"
	defaultMutatingWebhookTimeout         = "1"

	gkDeploymentPath  = "staticresources"
	gkTemplatePath    = "gktemplates"
	gkConstraintsPath = "gkconstraints"
)

//go:embed staticresources
var staticFiles embed.FS

//go:embed gktemplates
var gkPolicyTemplates embed.FS

//go:embed gkconstraints
var gkPolicyConraints embed.FS

var pullSecretName = types.NamespacedName{Name: "pull-secret", Namespace: "openshift-config"}

type Reconciler struct {
	arocli             aroclient.Interface
	kubernetescli      kubernetes.Interface
	deployer           deployer.Deployer
	gkPolicyTemplate   deployer.Deployer
	gkPolicyConstraint deployer.Deployer

	readinessPollTime     time.Duration
	readinessTimeout      time.Duration
	dh                    dynamichelper.Interface
	policyTickerDone      chan bool
	reconciliationMinutes int
}

func NewReconciler(arocli aroclient.Interface, kubernetescli kubernetes.Interface, dh dynamichelper.Interface) *Reconciler {
	return &Reconciler{
		arocli:             arocli,
		kubernetescli:      kubernetescli,
		deployer:           deployer.NewDeployer(kubernetescli, dh, staticFiles, gkDeploymentPath),
		gkPolicyTemplate:   deployer.NewDeployer(kubernetescli, dh, gkPolicyTemplates, gkTemplatePath),
		gkPolicyConstraint: deployer.NewDeployer(kubernetescli, dh, gkPolicyConraints, gkConstraintsPath),
		dh:                 dh,

		readinessPollTime: 10 * time.Second,
		readinessTimeout:  5 * time.Minute,
	}
}

func (r *Reconciler) Reconcile(ctx context.Context, request ctrl.Request) (ctrl.Result, error) {
	instance, err := r.arocli.AroV1alpha1().Clusters().Get(ctx, arov1alpha1.SingletonClusterName, metav1.GetOptions{})
	if err != nil {
		return reconcile.Result{}, err
	}

	// how to handle the enable/disable sequence of enabled and managed?
	if !instance.Spec.OperatorFlags.GetSimpleBoolean(controllerEnabled) {
		// controller is disabled
		return reconcile.Result{}, nil
	}

	managed := instance.Spec.OperatorFlags.GetWithDefault(controllerManaged, "")

	// If enabled and managed=true, install GuardRails
	// If enabled and managed=false, remove the GuardRails deployment
	// If enabled and managed is missing, do nothing
	if strings.EqualFold(managed, "true") {
		deployConfig := getDefaultDeployConfig(ctx, instance)
		// no need to CreateOrUpdate if gatekeeper is already ready, it costs ~3.5min
		if ready, err := r.gatekeeperIsReady(ctx, deployConfig); err != nil || !ready {
			// Deploy the GateKeeper manifests and config
			err = r.deployer.CreateOrUpdate(ctx, instance, deployConfig)
			if err != nil {
				logrus.Printf("\x1b[%dm guardrails:: reconcile error updating %s\x1b[0m", 31, err.Error())
				return reconcile.Result{}, err
			}

			// Check that GuardRails has become ready, wait up to readinessTimeout (default 5min)
			timeoutCtx, cancel := context.WithTimeout(ctx, r.readinessTimeout)
			defer cancel()

			err := wait.PollImmediateUntil(r.readinessPollTime, func() (bool, error) {
				return r.gatekeeperIsReady(ctx, deployConfig)
			}, timeoutCtx.Done())
			if err != nil {
				return reconcile.Result{}, fmt.Errorf("GateKeeper deployment timed out on Ready: %w", err)
			}
		}
		policyConfig := &config.GuardRailsPolicyConfig{}
		if r.gkPolicyTemplate != nil && r.gkPolicyConstraint != nil {
			logrus.Printf("\x1b[%dm guardrails:: CreateOrUpdate gkPolicyTemplate for %v\x1b[0m", 31, r.gkPolicyTemplate)
			// Deploy the GateKeeper policies templates
			err = r.gkPolicyTemplate.CreateOrUpdate(ctx, instance, policyConfig)
			if err != nil {
				logrus.Printf("\x1b[%dm guardrails:: reconcile error setup template %s\x1b[0m", 31, err.Error())
				return reconcile.Result{}, err
			}
			logrus.Printf("\x1b[%dm guardrails:: ensuring gkPolicyConstraint for %v\x1b[0m", 31, r.gkPolicyConstraint)
			// // Deploy the GateKeeper policies contraints
			// err = r.gkPolicyConstraint.CreateOrUpdate(ctx, instance, policyConfig)
			// if err != nil {
			// 	logrus.Printf("\x1b[%dm guardrails:: reconcile error setup constraints %s\x1b[0m", 31, err.Error())
			// 	return reconcile.Result{}, err
			// }

			err = r.ensurePolicy(ctx, gkPolicyConraints, gkConstraintsPath)
			if err != nil {
				logrus.Printf("\x1b[%dm guardrails:: reconcile error ensure constraints %s\x1b[0m", 31, err.Error())
				return reconcile.Result{}, err
			}
		}

		// temp := templatesv1beta1.ConstraintTemplate{}

		// TODO: need to find a way to check if gatekeeper policies have been deployed successfully
		// Check that GuardRails policies has become ready, wait up to readinessTimeout (default 5min)
		// timeoutPolicyCtx, cancel := context.WithTimeout(ctx, r.readinessTimeout)
		// defer cancel()

		// err = wait.PollImmediateUntil(r.readinessPollTime, func() (bool, error) {
		// 	if ready, err := r.gkPolicyTemplate.IsReady(ctx, "", "arodenylabels"); !ready || err != nil { //  ConstraintTemplate
		// 		return ready, err
		// 	}
		// 	return r.gkPolicyConstraint.IsReady(ctx, "", "aro-machines-deny") // arodenylabels
		// }, timeoutPolicyCtx.Done())
		// if err != nil {
		// 	return reconcile.Result{}, fmt.Errorf("GateKeeper policy timed out on Ready: %w", err)
		// }

		// start a ticker to periodically re-enforce gatekeeper policies, in case they are deleted by users
		minutes := instance.Spec.OperatorFlags.GetWithDefault(controllerReconciliationMinutes, defaultReconciliationMinutes)
		min, err := strconv.Atoi(minutes)
		if err != nil {
			// invalid user config here, use default value
			min, _ = strconv.Atoi(defaultReconciliationMinutes)
		}
		if r.reconciliationMinutes != min && r.policyTickerDone != nil {
			// trigger ticker reset
			r.reconciliationMinutes = min
			r.policyTickerDone <- false
			// logrus.Printf("\x1b[%dm guardrails:: reconcile ticker check min %d recon %d ticker %v\x1b[0m", 31, min, r.reconciliationMinutes, r.policyTickerDone)
		}

		// make sure only one ticker started
		if r.policyTickerDone == nil {
			go r.policyTicker(ctx, instance)
		}

	} else if strings.EqualFold(managed, "false") {
		logrus.Printf("\x1b[%dm guardrails:: reconcile starts removing guardrails\x1b[0m", 31)
		if r.gkPolicyTemplate != nil {
			// stop the gatekeeper policies re-enforce ticker
			if r.policyTickerDone != nil {
				r.policyTickerDone <- true
				close(r.policyTickerDone)
			}

			err = r.removePolicy(ctx, gkPolicyConraints, gkConstraintsPath)
			if err != nil {
				logrus.Printf("\x1b[%dm guardrails:: reconcile error removing constraints %s\x1b[0m", 31, err.Error())
				return reconcile.Result{}, err
			}

			err = r.gkPolicyTemplate.Remove(ctx, config.GuardRailsPolicyConfig{})
			if err != nil {
				return reconcile.Result{}, err
			}
		}
		err = r.deployer.Remove(ctx, config.GuardRailsDeploymentConfig{Namespace: instance.Spec.OperatorFlags.GetWithDefault(controllerNamespace, defaultNamespace)})
		if err != nil {
			logrus.Printf("\x1b[%dm guardrails:: reconcile error removing deployment %s\x1b[0m", 31, err.Error())
			return reconcile.Result{}, err
		}
	}

	return reconcile.Result{}, nil
}

func (r *Reconciler) policyTicker(ctx context.Context, instance *arov1alpha1.Cluster) {
	r.policyTickerDone = make(chan bool)
	var err error

	minutes := instance.Spec.OperatorFlags.GetWithDefault(controllerReconciliationMinutes, defaultReconciliationMinutes)
	r.reconciliationMinutes, err = strconv.Atoi(minutes)
	if err != nil {
		// invalid user config here, use default value
		r.reconciliationMinutes, _ = strconv.Atoi(defaultReconciliationMinutes)
	}

	ticker := time.NewTicker(time.Duration(r.reconciliationMinutes) * time.Minute)
	defer ticker.Stop()
	for {
		select {
		case done := <-r.policyTickerDone:
			if done {
				r.policyTickerDone = nil
				return
			}
			// false to trigger a ticker reset
			logrus.Printf("\x1b[%dm guardrails:: ticker reset to %d min \x1b[0m", 31, r.reconciliationMinutes)
			ticker.Reset(time.Duration(r.reconciliationMinutes) * time.Minute)
		case <-ticker.C:
			err = r.ensurePolicy(ctx, gkPolicyConraints, gkConstraintsPath)
			if err != nil {
				logrus.Printf("\x1b[%dm guardrails::Ticker ensurePolicy error ensure constraints %s\x1b[0m", 31, err.Error())
			}
		}

	}
}
func (r *Reconciler) gatekeeperIsReady(ctx context.Context, deployConfig *config.GuardRailsDeploymentConfig) (bool, error) {
	if ready, err := r.deployer.IsReady(ctx, deployConfig.Namespace, "gatekeeper-audit"); !ready || err != nil {
		return ready, err
	}
	return r.deployer.IsReady(ctx, deployConfig.Namespace, "gatekeeper-controller-manager")
}

// SetupWithManager setup our manager
func (r *Reconciler) SetupWithManager(mgr ctrl.Manager) error {

	pullSecretPredicate := predicate.NewPredicateFuncs(func(o client.Object) bool {
		return (o.GetName() == pullSecretName.Name && o.GetNamespace() == pullSecretName.Namespace)
	})

	aroClusterPredicate := predicate.NewPredicateFuncs(func(o client.Object) bool {
		return o.GetName() == arov1alpha1.SingletonClusterName
	})

	grBuilder := ctrl.NewControllerManagedBy(mgr).
		For(&arov1alpha1.Cluster{}, builder.WithPredicates(aroClusterPredicate)).
		Watches(
			&source.Kind{Type: &corev1.Secret{}},
			&handler.EnqueueRequestForObject{},
			builder.WithPredicates(pullSecretPredicate),
		)

	resources, err := r.deployer.Template(&config.GuardRailsDeploymentConfig{}, staticFiles)
	if err != nil {
		return err
	}

	for _, i := range resources {
		o, ok := i.(client.Object)
		if ok {
			grBuilder.Owns(o)
		}
	}

	// kgTemplateResources, err := r.gkPolicyTemplate.Template(&config.GuardRailsPolicyConfig{}, gkPolicyTemplates)
	// if err != nil {
	// 	return err
	// }

	// for _, i := range kgTemplateResources {
	// 	o, ok := i.(client.Object)
	// 	if ok {
	// 		grBuilder.Owns(o)
	// 	}
	// }

	// we won't listen for changes on policies, since we only want to reconcile on a timer anyway
	if err := grBuilder.
		WithEventFilter(predicate.Or(predicate.GenerationChangedPredicate{}, predicate.AnnotationChangedPredicate{}, predicate.LabelChangedPredicate{})).
		Named(ControllerName).
		Complete(r); err != nil {
		logrus.Printf("\x1b[%dm guardrails::SetupWithManager deployment failed %v 0\x1b[0m", 31, err)
		return err
	}
	return nil
}

func (r *Reconciler) getPolicyConfig(ctx context.Context, na string) (string, string, error) {

	parts := strings.Split(na, ".")
	if len(parts) < 1 {
		return "", "", errors.New("unrecognised name: " + na)
	}
	name := parts[0]

	instance, err := r.arocli.AroV1alpha1().Clusters().Get(ctx, arov1alpha1.SingletonClusterName, metav1.GetOptions{})
	if err != nil {
		return "", "", err
	}

	managedPath := fmt.Sprintf("aro.guardrails.policies.%s.managed", name)
	managed := instance.Spec.OperatorFlags.GetWithDefault(managedPath, "false")

	enforcementPath := fmt.Sprintf("aro.guardrails.policies.%s.enforcement", name)
	enforcement := instance.Spec.OperatorFlags.GetWithDefault(enforcementPath, "dryrun")

	return managed, enforcement, nil
}

func (r *Reconciler) ensurePolicy(ctx context.Context, fs embed.FS, path string) error {
	template, err := template.ParseFS(fs, filepath.Join(path, "*"))
	if err != nil {
		return err
	}

	creates := make([]kruntime.Object, 0)
	buffer := new(bytes.Buffer)
	for _, templ := range template.Templates() {
		managed, enforcement, err := r.getPolicyConfig(ctx, templ.Name())
		if err != nil {
			return err
		}
		policyConfig := &config.GuardRailsPolicyConfig{
			Enforcement: enforcement,
		}
		err = templ.Execute(buffer, policyConfig)
		if err != nil {
			return err
		}
		data := buffer.Bytes()

		obj := &unstructured.Unstructured{}
		json, err := yaml.YAMLToJSON(data)
		if err != nil {
			return err
		}
		err = obj.UnmarshalJSON(json)
		if err != nil {
			return err
		}

		if managed != "true" {
			err := r.dh.EnsureDeletedGVR(ctx, obj.GroupVersionKind().GroupKind().String(), obj.GetNamespace(), obj.GetName(), obj.GroupVersionKind().Version)
			if err != nil && !strings.Contains(err.Error(), "NotFound") { //!kerrors.IsNotFound(err) {
				logrus.Printf("\x1b[%dm guardrails:: ensurePolicy failed to remove UnstructuredObj kind %s name %s\x1b[0m", 31, obj.GroupVersionKind().GroupKind().String(), obj.GetName())
				return err
			}
			continue
		}

		uns := dynamichelper.NewUnstructuredObj(*obj)
		creates = append(creates, *uns)
	}
	err = r.dh.Ensure(ctx, creates...)
	if err != nil {
		return err
	}
	return nil
}

func (r *Reconciler) removePolicy(ctx context.Context, fs embed.FS, path string) error {
	template, err := template.ParseFS(fs, filepath.Join(path, "*"))
	if err != nil {
		return err
	}

	buffer := new(bytes.Buffer)
	for _, templ := range template.Templates() {
		err := templ.Execute(buffer, nil)
		if err != nil {
			return err
		}
		data := buffer.Bytes()

		obj := &unstructured.Unstructured{}
		json, err := yaml.YAMLToJSON(data)
		if err != nil {
			return err
		}
		err = obj.UnmarshalJSON(json)
		if err != nil {
			return err
		}

		err = r.dh.EnsureDeletedGVR(ctx, obj.GroupVersionKind().GroupKind().String(), obj.GetNamespace(), obj.GetName(), obj.GroupVersionKind().Version)
		if err != nil && !strings.Contains(err.Error(), "NotFound") { //!kerrors.IsNotFound(err) {
			logrus.Printf("\x1b[%dm guardrails:: removePolicy failed removing UnstructuredObj kind %s name %s\x1b[0m", 31, obj.GroupVersionKind().GroupKind().String(), obj.GetName())
			return err
		}
	}
	return nil
}

func getDefaultDeployConfig(ctx context.Context, instance *arov1alpha1.Cluster) *config.GuardRailsDeploymentConfig {
	// apply the default value if the flag is empty or missing
	deployConfig := &config.GuardRailsDeploymentConfig{
		Pullspec:  instance.Spec.OperatorFlags.GetWithDefault(controllerPullSpec, version.GateKeeperImage(instance.Spec.ACRDomain)),
		Namespace: instance.Spec.OperatorFlags.GetWithDefault(controllerNamespace, defaultNamespace),

		ManagerRequestsCPU: instance.Spec.OperatorFlags.GetWithDefault(controllerManagerRequestsCPU, defaultManagerRequestsCPU),
		ManagerLimitCPU:    instance.Spec.OperatorFlags.GetWithDefault(controllerManagerLimitCPU, defaultManagerLimitCPU),
		ManagerRequestsMem: instance.Spec.OperatorFlags.GetWithDefault(controllerManagerRequestsMem, defaultManagerRequestsMem),
		ManagerLimitMem:    instance.Spec.OperatorFlags.GetWithDefault(controllerManagerLimitMem, defaultManagerLimitMem),

		AuditRequestsCPU: instance.Spec.OperatorFlags.GetWithDefault(controllerAuditRequestsCPU, defaultAuditRequestsCPU),
		AuditLimitCPU:    instance.Spec.OperatorFlags.GetWithDefault(controllerAuditLimitCPU, defaultAuditLimitCPU),
		AuditRequestsMem: instance.Spec.OperatorFlags.GetWithDefault(controllerAuditRequestsMem, defaultAuditRequestsMem),
		AuditLimitMem:    instance.Spec.OperatorFlags.GetWithDefault(controllerAuditLimitMem, defaultAuditLimitMem),

		ValidatingWebhookTimeout:       instance.Spec.OperatorFlags.GetWithDefault(controllerValidatingWebhookTimeout, defaultValidatingWebhookTimeout),
		ValidatingWebhookFailurePolicy: instance.Spec.OperatorFlags.GetWithDefault(controllerValidatingWebhookFailurePolicy, defaultValidatingWebhookFailurePolicy),
		MutatingWebhookTimeout:         instance.Spec.OperatorFlags.GetWithDefault(controllerMutatingWebhookTimeout, defaultMutatingWebhookTimeout),
		MutatingWebhookFailurePolicy:   instance.Spec.OperatorFlags.GetWithDefault(controllerMutatingWebhookFailurePolicy, defaultMutatingWebhookFailurePolicy),
	}
	validatingManaged := instance.Spec.OperatorFlags.GetWithDefault(controllerValidatingWebhookFailurePolicy, "")
	switch {
	case validatingManaged == "":
		deployConfig.ValidatingWebhookFailurePolicy = defaultValidatingWebhookFailurePolicy
	case strings.EqualFold(validatingManaged, "true"):
		deployConfig.ValidatingWebhookFailurePolicy = "Fail"
	case strings.EqualFold(validatingManaged, "false"):
		deployConfig.ValidatingWebhookFailurePolicy = "Ignore"
	}
	mutatingManaged := instance.Spec.OperatorFlags.GetWithDefault(controllerMutatingWebhookFailurePolicy, "")
	switch {
	case mutatingManaged == "":
		deployConfig.MutatingWebhookFailurePolicy = defaultMutatingWebhookFailurePolicy
	case strings.EqualFold(mutatingManaged, "true"):
		deployConfig.MutatingWebhookFailurePolicy = "Fail"
	case strings.EqualFold(mutatingManaged, "false"):
		deployConfig.MutatingWebhookFailurePolicy = "Ignore"
	}
	return deployConfig
}
