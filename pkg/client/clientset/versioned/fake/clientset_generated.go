/*
Copyright 2020 The KubeSphere Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
	clientset "kubesphere.io/kubesphere/pkg/client/clientset/versioned"
	applicationv1alpha1 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/application/v1alpha1"
	fakeapplicationv1alpha1 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/application/v1alpha1/fake"
	auditingv1alpha1 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/auditing/v1alpha1"
	fakeauditingv1alpha1 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/auditing/v1alpha1/fake"
	clusterv1alpha1 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/cluster/v1alpha1"
	fakeclusterv1alpha1 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/cluster/v1alpha1/fake"
	devopsv1alpha1 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/devops/v1alpha1"
	fakedevopsv1alpha1 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/devops/v1alpha1/fake"
	devopsv1alpha3 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/devops/v1alpha3"
	fakedevopsv1alpha3 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/devops/v1alpha3/fake"
	iamv1alpha2 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/iam/v1alpha2"
	fakeiamv1alpha2 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/iam/v1alpha2/fake"
	networkv1alpha1 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/network/v1alpha1"
	fakenetworkv1alpha1 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/network/v1alpha1/fake"
	notificationv2beta1 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/notification/v2beta1"
	fakenotificationv2beta1 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/notification/v2beta1/fake"
	notificationv2beta2 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/notification/v2beta2"
	fakenotificationv2beta2 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/notification/v2beta2/fake"
	quotav1alpha2 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/quota/v1alpha2"
	fakequotav1alpha2 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/quota/v1alpha2/fake"
	servicemeshv1alpha2 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/servicemesh/v1alpha2"
	fakeservicemeshv1alpha2 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/servicemesh/v1alpha2/fake"
	storagev1alpha1 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/storage/v1alpha1"
	fakestoragev1alpha1 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/storage/v1alpha1/fake"
	tenantv1alpha1 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/tenant/v1alpha1"
	faketenantv1alpha1 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/tenant/v1alpha1/fake"
	tenantv1alpha2 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/tenant/v1alpha2"
	faketenantv1alpha2 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/tenant/v1alpha2/fake"
	typesv1beta1 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/types/v1beta1"
	faketypesv1beta1 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/types/v1beta1/fake"
	typesv1beta2 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/types/v1beta2"
	faketypesv1beta2 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/types/v1beta2/fake"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var _ clientset.Interface = &Clientset{}

// ApplicationV1alpha1 retrieves the ApplicationV1alpha1Client
func (c *Clientset) ApplicationV1alpha1() applicationv1alpha1.ApplicationV1alpha1Interface {
	return &fakeapplicationv1alpha1.FakeApplicationV1alpha1{Fake: &c.Fake}
}

// AuditingV1alpha1 retrieves the AuditingV1alpha1Client
func (c *Clientset) AuditingV1alpha1() auditingv1alpha1.AuditingV1alpha1Interface {
	return &fakeauditingv1alpha1.FakeAuditingV1alpha1{Fake: &c.Fake}
}

// ClusterV1alpha1 retrieves the ClusterV1alpha1Client
func (c *Clientset) ClusterV1alpha1() clusterv1alpha1.ClusterV1alpha1Interface {
	return &fakeclusterv1alpha1.FakeClusterV1alpha1{Fake: &c.Fake}
}

// DevopsV1alpha1 retrieves the DevopsV1alpha1Client
func (c *Clientset) DevopsV1alpha1() devopsv1alpha1.DevopsV1alpha1Interface {
	return &fakedevopsv1alpha1.FakeDevopsV1alpha1{Fake: &c.Fake}
}

// DevopsV1alpha3 retrieves the DevopsV1alpha3Client
func (c *Clientset) DevopsV1alpha3() devopsv1alpha3.DevopsV1alpha3Interface {
	return &fakedevopsv1alpha3.FakeDevopsV1alpha3{Fake: &c.Fake}
}

// IamV1alpha2 retrieves the IamV1alpha2Client
func (c *Clientset) IamV1alpha2() iamv1alpha2.IamV1alpha2Interface {
	return &fakeiamv1alpha2.FakeIamV1alpha2{Fake: &c.Fake}
}

// NetworkV1alpha1 retrieves the NetworkV1alpha1Client
func (c *Clientset) NetworkV1alpha1() networkv1alpha1.NetworkV1alpha1Interface {
	return &fakenetworkv1alpha1.FakeNetworkV1alpha1{Fake: &c.Fake}
}

// NotificationV2beta1 retrieves the NotificationV2beta1Client
func (c *Clientset) NotificationV2beta1() notificationv2beta1.NotificationV2beta1Interface {
	return &fakenotificationv2beta1.FakeNotificationV2beta1{Fake: &c.Fake}
}

// NotificationV2beta2 retrieves the NotificationV2beta2Client
func (c *Clientset) NotificationV2beta2() notificationv2beta2.NotificationV2beta2Interface {
	return &fakenotificationv2beta2.FakeNotificationV2beta2{Fake: &c.Fake}
}

// QuotaV1alpha2 retrieves the QuotaV1alpha2Client
func (c *Clientset) QuotaV1alpha2() quotav1alpha2.QuotaV1alpha2Interface {
	return &fakequotav1alpha2.FakeQuotaV1alpha2{Fake: &c.Fake}
}

// ServicemeshV1alpha2 retrieves the ServicemeshV1alpha2Client
func (c *Clientset) ServicemeshV1alpha2() servicemeshv1alpha2.ServicemeshV1alpha2Interface {
	return &fakeservicemeshv1alpha2.FakeServicemeshV1alpha2{Fake: &c.Fake}
}

// StorageV1alpha1 retrieves the StorageV1alpha1Client
func (c *Clientset) StorageV1alpha1() storagev1alpha1.StorageV1alpha1Interface {
	return &fakestoragev1alpha1.FakeStorageV1alpha1{Fake: &c.Fake}
}

// TenantV1alpha1 retrieves the TenantV1alpha1Client
func (c *Clientset) TenantV1alpha1() tenantv1alpha1.TenantV1alpha1Interface {
	return &faketenantv1alpha1.FakeTenantV1alpha1{Fake: &c.Fake}
}

// TenantV1alpha2 retrieves the TenantV1alpha2Client
func (c *Clientset) TenantV1alpha2() tenantv1alpha2.TenantV1alpha2Interface {
	return &faketenantv1alpha2.FakeTenantV1alpha2{Fake: &c.Fake}
}

// TypesV1beta1 retrieves the TypesV1beta1Client
func (c *Clientset) TypesV1beta1() typesv1beta1.TypesV1beta1Interface {
	return &faketypesv1beta1.FakeTypesV1beta1{Fake: &c.Fake}
}

// TypesV1beta2 retrieves the TypesV1beta2Client
func (c *Clientset) TypesV1beta2() typesv1beta2.TypesV1beta2Interface {
	return &faketypesv1beta2.FakeTypesV1beta2{Fake: &c.Fake}
}
