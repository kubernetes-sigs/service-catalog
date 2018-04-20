/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen

package internalversion

import (
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServicePlanLister helps list ServicePlans.
type ServicePlanLister interface {
	// List lists all ServicePlans in the indexer.
	List(selector labels.Selector) (ret []*servicecatalog.ServicePlan, err error)
	// ServicePlans returns an object that can list and get ServicePlans.
	ServicePlans(namespace string) ServicePlanNamespaceLister
	ServicePlanListerExpansion
}

// servicePlanLister implements the ServicePlanLister interface.
type servicePlanLister struct {
	indexer cache.Indexer
}

// NewServicePlanLister returns a new ServicePlanLister.
func NewServicePlanLister(indexer cache.Indexer) ServicePlanLister {
	return &servicePlanLister{indexer: indexer}
}

// List lists all ServicePlans in the indexer.
func (s *servicePlanLister) List(selector labels.Selector) (ret []*servicecatalog.ServicePlan, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*servicecatalog.ServicePlan))
	})
	return ret, err
}

// ServicePlans returns an object that can list and get ServicePlans.
func (s *servicePlanLister) ServicePlans(namespace string) ServicePlanNamespaceLister {
	return servicePlanNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServicePlanNamespaceLister helps list and get ServicePlans.
type ServicePlanNamespaceLister interface {
	// List lists all ServicePlans in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*servicecatalog.ServicePlan, err error)
	// Get retrieves the ServicePlan from the indexer for a given namespace and name.
	Get(name string) (*servicecatalog.ServicePlan, error)
	ServicePlanNamespaceListerExpansion
}

// servicePlanNamespaceLister implements the ServicePlanNamespaceLister
// interface.
type servicePlanNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServicePlans in the indexer for a given namespace.
func (s servicePlanNamespaceLister) List(selector labels.Selector) (ret []*servicecatalog.ServicePlan, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*servicecatalog.ServicePlan))
	})
	return ret, err
}

// Get retrieves the ServicePlan from the indexer for a given namespace and name.
func (s servicePlanNamespaceLister) Get(name string) (*servicecatalog.ServicePlan, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(servicecatalog.Resource("serviceplan"), name)
	}
	return obj.(*servicecatalog.ServicePlan), nil
}
