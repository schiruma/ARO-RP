package cluster

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/util/feature"
)

func (m *manager) setHyperthreadingToggle(ctx context.Context) error {
	var err error

	hyperthreadingField := ""
	if feature.IsRegisteredForFeature(m.subscriptionDoc.Subscription.Properties, api.FeatureFlagHyperthreadingToggle) {
		hyperthreadingField = string(api.HyperthreadingDisabled)
	} else {
		hyperthreadingField = string(api.HyperthreadingEnabled)
	}

	m.doc, err = m.db.PatchWithLease(ctx, m.doc.Key, func(doc *api.OpenShiftClusterDocument) error {
		doc.OpenShiftCluster.Properties.MasterProfile.HyperthreadingField = string(hyperthreadingField)
		return nil
	})
	return err

}
