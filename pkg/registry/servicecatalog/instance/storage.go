/*
Copyright 2016 The Kubernetes Authors.

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

package instance

import (
	"context"
	"errors"
	"fmt"

	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage"
)

var (
	errNotAnServiceInstance = errors.New("not an instance")
)

// NewSingular returns a new shell of a service instance, according to the given namespace and
// name
func NewSingular(ns, name string) runtime.Object {
	return &servicecatalog.ServiceInstance{
		TypeMeta: metav1.TypeMeta{
			Kind: "ServiceInstance",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: ns,
			Name:      name,
		},
	}
}

// EmptyObject returns an empty instance
func EmptyObject() runtime.Object {
	return &servicecatalog.ServiceInstance{}
}

// NewList returns a new shell of an instance list
func NewList() runtime.Object {
	return &servicecatalog.ServiceInstanceList{
		TypeMeta: metav1.TypeMeta{
			Kind: "ServiceInstanceList",
		},
		Items: []servicecatalog.ServiceInstance{},
	}
}

// CheckObject returns a non-nil error if obj is not an instance object
func CheckObject(obj runtime.Object) error {
	_, ok := obj.(*servicecatalog.ServiceInstance)
	if !ok {
		return errNotAnServiceInstance
	}
	return nil
}

// Match determines whether an ServiceInstance matches a field and label
// selector.
func Match(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// toSelectableFields returns a field set that represents the object for matching purposes.
func toSelectableFields(instance *servicecatalog.ServiceInstance) fields.Set {
	// If you add a new selectable field, you also need to modify
	// pkg/apis/servicecatalog/v1beta1/conversion[_test].go
	objectMetaFieldsSet := generic.ObjectMetaFieldsSet(&instance.ObjectMeta, true)

	specFieldSet := make(fields.Set, 3)

	if instance.Spec.ClusterServiceClassRef != nil {
		specFieldSet["spec.clusterServiceClassRef.name"] = instance.Spec.ClusterServiceClassRef.Name
	}

	if instance.Spec.ClusterServicePlanRef != nil {
		specFieldSet["spec.clusterServicePlanRef.name"] = instance.Spec.ClusterServicePlanRef.Name
	}

	if instance.Spec.ExternalID != "" {
		specFieldSet["spec.externalID"] = instance.Spec.ExternalID
	}

	return generic.MergeFieldsSets(objectMetaFieldsSet, specFieldSet)
}

// GetAttrs returns labels and fields of a given object for filtering purposes.
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
	instance, ok := obj.(*servicecatalog.ServiceInstance)
	if !ok {
		return nil, nil, false, fmt.Errorf("given object is not an ServiceInstance")
	}
	return labels.Set(instance.ObjectMeta.Labels), toSelectableFields(instance), instance.Initializers != nil, nil
}

// NewStorage creates a new rest.Storage responsible for accessing ServiceInstance
// resources
func NewStorage(optsGetter generic.RESTOptionsGetter) (instancePlans, instanceStatus, instanceReference rest.Storage, err error) {
	store := registry.Store{
		NewFunc:                  EmptyObject,
		NewListFunc:              NewList,
		PredicateFunc:            Match,
		DefaultQualifiedResource: servicecatalog.Resource("serviceinstances"),

		CreateStrategy:          instanceRESTStrategies,
		UpdateStrategy:          instanceRESTStrategies,
		DeleteStrategy:          instanceRESTStrategies,
		EnableGarbageCollection: true,
	}

	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, nil, nil, err
	}

	statusStore := store
	statusStore.UpdateStrategy = instanceStatusUpdateStrategy

	referenceStore := store
	referenceStore.UpdateStrategy = instanceReferenceUpdateStrategy

	return &store, &StatusREST{&statusStore}, &ReferenceREST{&referenceStore}, nil
}

// StatusREST defines the REST operations for the status subresource via
// implementation of various rest interfaces.  It supports the http verbs GET,
// PATCH, and PUT.
type StatusREST struct {
	store *registry.Store
}

// New returns a new ServiceInstance
func (r *StatusREST) New() runtime.Object {
	return &servicecatalog.ServiceInstance{}
}

// Get retrieves the object from the storage. It is required to support Patch
// and to implement the rest.Getter interface.
func (r *StatusREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	return r.store.Get(ctx, name, options)
}

var (
	_ rest.Storage = &StatusREST{}
	_ rest.Getter  = &StatusREST{}
	_ rest.Updater = &StatusREST{}
)

// Update alters the status subset of an object and it
// implements rest.Updater interface
func (r *StatusREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc) (runtime.Object, bool, error) {
	return r.store.Update(ctx, name, objInfo, createValidation, updateValidation)
}

// ReferenceREST defines the REST operations for the reference subresource.
type ReferenceREST struct {
	store *registry.Store
}

// New returns a new ServiceInstance
func (r *ReferenceREST) New() runtime.Object {
	return &servicecatalog.ServiceInstance{}
}

// Get retrieves the object from the storage. It is required to support Patch
// and to implement the rest.Getter interface.
func (r *ReferenceREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	return r.store.Get(ctx, name, options)
}

// Update alters the reference subset of an object and it
// implements rest.Updater interface
func (r *ReferenceREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc) (runtime.Object, bool, error) {
	return r.store.Update(ctx, name, objInfo, createValidation, updateValidation)
}
