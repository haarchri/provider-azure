/*
Copyright 2019 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance With the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package test

import (
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2017-06-01/storage"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	runtimev1alpha1 "github.com/crossplaneio/crossplane-runtime/apis/core/v1alpha1"

	storagev1alpha2 "github.com/crossplaneio/stack-azure/apis/storage/v1alpha2"
)

// MockAccount builder for testing account object
type MockAccount struct {
	*storagev1alpha2.Account
}

// NewMockAccount creates new account wrapper
func NewMockAccount(name string) *MockAccount {
	return &MockAccount{Account: &storagev1alpha2.Account{
		ObjectMeta: metav1.ObjectMeta{Name: name},
	}}
}

// WithTypeMeta sets TypeMeta value
func (ta *MockAccount) WithTypeMeta(tm metav1.TypeMeta) *MockAccount {
	ta.TypeMeta = tm
	return ta
}

// WithObjectMeta sets ObjectMeta value
func (ta *MockAccount) WithObjectMeta(om metav1.ObjectMeta) *MockAccount {
	ta.ObjectMeta = om
	return ta
}

// WithUID sets UID value
func (ta *MockAccount) WithUID(uid string) *MockAccount {
	ta.ObjectMeta.UID = types.UID(uid)
	return ta
}

// WithDeleteTimestamp sets metadata deletion timestamp
func (ta *MockAccount) WithDeleteTimestamp(t metav1.Time) *MockAccount {
	ta.Account.ObjectMeta.DeletionTimestamp = &t
	return ta
}

// WithFinalizer sets finalizer
func (ta *MockAccount) WithFinalizer(f string) *MockAccount {
	ta.Account.ObjectMeta.Finalizers = append(ta.Account.ObjectMeta.Finalizers, f)
	return ta
}

// WithFinalizers sets finalizers list
func (ta *MockAccount) WithFinalizers(f []string) *MockAccount {
	ta.Account.ObjectMeta.Finalizers = f
	return ta
}

// WithSpecClassRef set class reference
func (ta *MockAccount) WithSpecClassRef(ref *corev1.ObjectReference) *MockAccount {
	ta.Spec.ClassReference = ref
	return ta
}

// WithSpecClaimRef set class reference
func (ta *MockAccount) WithSpecClaimRef(ref *corev1.ObjectReference) *MockAccount {
	ta.Spec.ClaimReference = ref
	return ta
}

// WithSpecProvider set a provider
func (ta *MockAccount) WithSpecProvider(name string) *MockAccount {
	ta.Spec.ProviderReference = &corev1.ObjectReference{Name: name}
	return ta
}

// WithSpecReclaimPolicy sets resource reclaim policy
func (ta *MockAccount) WithSpecReclaimPolicy(policy runtimev1alpha1.ReclaimPolicy) *MockAccount {
	ta.Spec.ReclaimPolicy = policy
	return ta
}

// WithSpecStorageAccountName sets spec value
func (ta *MockAccount) WithSpecStorageAccountName(name string) *MockAccount {
	ta.Spec.StorageAccountName = name
	return ta
}

// WithSpecStorageAccountSpec sets storage account specs
func (ta *MockAccount) WithSpecStorageAccountSpec(spec *storagev1alpha2.StorageAccountSpec) *MockAccount {
	ta.Spec.StorageAccountSpec = spec
	return ta
}

// WithStorageAccountStatus set storage account status
func (ta *MockAccount) WithStorageAccountStatus(status *storagev1alpha2.StorageAccountStatus) *MockAccount {
	ta.Status.StorageAccountStatus = status
	return ta
}

// WithSpecStatusFromProperties set storage account spec status from storage properties
func (ta *MockAccount) WithSpecStatusFromProperties(props *storage.AccountProperties) *MockAccount {
	acct := &storage.Account{
		AccountProperties: props,
	}
	ta.WithSpecStorageAccountSpec(storagev1alpha2.NewStorageAccountSpec(acct)).
		WithStorageAccountStatus(storagev1alpha2.NewStorageAccountStatus(acct))
	return ta
}

// WithSpecWriteConnectionSecretToReference sets where the storage account will write its
// connection secret.
func (ta *MockAccount) WithSpecWriteConnectionSecretToReference(ns, name string) *MockAccount {
	ta.Spec.WriteConnectionSecretToReference = &runtimev1alpha1.SecretReference{
		Namespace: ns,
		Name:      name,
	}
	return ta
}

// WithStatusConditions sets the storage account's conditioned status.
func (ta *MockAccount) WithStatusConditions(c ...runtimev1alpha1.Condition) *MockAccount {
	ta.Status.SetConditions(c...)
	return ta
}

// WithStatusBindingPhase sets the storage account's binding phase.
func (ta *MockAccount) WithStatusBindingPhase(p runtimev1alpha1.BindingPhase) *MockAccount {
	ta.Status.SetBindingPhase(p)
	return ta
}
