/*
Copyright 2019 The Crossplane Authors.

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

package v1alpha3

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group   = "storage.azure.crossplane.io"
	Version = "v1alpha3"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)

// Account type metadata.
var (
	AccountKind             = reflect.TypeOf(Account{}).Name()
	AccountGroupKind        = schema.GroupKind{Group: Group, Kind: AccountKind}.String()
	AccountKindAPIVersion   = AccountKind + "." + SchemeGroupVersion.String()
	AccountGroupVersionKind = SchemeGroupVersion.WithKind(AccountKind)
)

// Container type metadata.
var (
	ContainerKind             = reflect.TypeOf(Container{}).Name()
	ContainerGroupKind        = schema.GroupKind{Group: Group, Kind: ContainerKind}.String()
	ContainerKindAPIVersion   = ContainerKind + "." + SchemeGroupVersion.String()
	ContainerGroupVersionKind = SchemeGroupVersion.WithKind(ContainerKind)
)

// ContainerClass type metadata.
var (
	ContainerClassKind             = reflect.TypeOf(ContainerClass{}).Name()
	ContainerClassGroupKind        = schema.GroupKind{Group: Group, Kind: ContainerClassKind}.String()
	ContainerClassKindAPIVersion   = ContainerClassKind + "." + SchemeGroupVersion.String()
	ContainerClassGroupVersionKind = SchemeGroupVersion.WithKind(ContainerClassKind)
)

func init() {
	SchemeBuilder.Register(&Account{}, &AccountList{})
	SchemeBuilder.Register(&Container{}, &ContainerList{})
	SchemeBuilder.Register(&ContainerClass{}, &ContainerClassList{})
}
