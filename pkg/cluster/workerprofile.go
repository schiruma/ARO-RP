package cluster

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/util/feature"
)

func (m *manager) setWorkerHyperthreadingToggle(ctx context.Context) error {
	var err error

	hyperthreadingField := api.HyperthreadingEnabled
	if feature.IsRegisteredForFeature(m.subscriptionDoc.Subscription.Properties, api.FeatureFlagHyperthreadingToggle) {
		hyperthreadingField = api.HyperthreadingDisabled
	}

	m.doc, err = m.db.PatchWithLease(ctx, m.doc.Key, func(doc *api.OpenShiftClusterDocument) error {
		doc.OpenShiftCluster.Properties.MasterProfile.HyperthreadingField = hyperthreadingField
		return nil
	})
	return err

}
