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

package servicecatalog

import (
	"errors"
	"fmt"
	"strings"

	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	appierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
)

const (
	// FieldExternalPlanName is the jsonpath to a plan's external name.
	FieldExternalPlanName = "spec.externalName"

	// FieldClusterServiceClassRef is the jsonpath to a plan's associated class name.
	FieldClusterServiceClassRef = "spec.clusterServiceClassRef.name"

	// FieldServiceClassRef is the jsonpath to a plan's associated class name.
	FieldServiceClassRef = "spec.serviceClassRef.name"
)

// RetrievePlanOptions allows to specify which plans will be retrieved
type RetrievePlanOptions struct {
	ClassID   string
	Namespace string
	Scope     Scope
}

// CreatePlanOptions allows specifying what options to use for creating a plan.
type CreatePlanOptions struct {
	Name      string
	Scope     Scope
	Namespace string
	From      string
}

// Plan provides a unifying layer of cluster and namespace scoped plan resources.
type Plan interface {

	// GetName returns the plan's name.
	GetName() string

	// GetShortStatus returns the plan's status.
	GetShortStatus() string

	// GetNamespace returns the plan's namespace, or "" if it's cluster-scoped.
	GetNamespace() string

	// GetExternalName returns the plan's external name.
	GetExternalName() string

	// GetDescription returns the plan description.
	GetDescription() string

	// GetFree returns if the plan is free.
	GetFree() bool

	// GetClassID returns the plan's class name.
	GetClassID() string

	// GetStatus returns the plan's status.
	GetStatus() string

	// GetSpec returns the plan's spec.
	GetSpec() v1beta1.CommonServicePlanSpec
}

// RetrievePlans lists all plans defined in the cluster.
func (sdk *SDK) RetrievePlans(classID string, opts ScopeOptions) ([]Plan, error) {
	plans, err := sdk.retrievePlansByListOptions(opts, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	if classID == "" {
		return plans, nil
	}

	var filtered []Plan
	for _, p := range plans {
		if p.GetClassID() == classID {
			filtered = append(filtered, p)
		}
	}

	return filtered, nil
}

func (sdk *SDK) retrievePlansByListOptions(opts ScopeOptions, listOpts metav1.ListOptions) ([]Plan, error) {
	var plans []Plan

	if opts.Scope.Matches(ClusterScope) {
		csp, err := sdk.ServiceCatalog().ClusterServicePlans().List(metav1.ListOptions{})
		if err != nil {
			return nil, fmt.Errorf("unable to list cluster-scoped plans (%s)", err)
		}

		for _, p := range csp.Items {
			plan := p
			plans = append(plans, &plan)
		}
	}

	if opts.Scope.Matches(NamespaceScope) {
		sp, err := sdk.ServiceCatalog().ServicePlans(opts.Namespace).List(metav1.ListOptions{})
		if err != nil {
			// Gracefully handle when the feature-flag for namespaced broker resources isn't enabled on the server.
			if appierrors.IsNotFound(err) {
				return plans, nil
			}
			return nil, fmt.Errorf("unable to list plans in %q (%s)", opts.Namespace, err)
		}

		for _, p := range sp.Items {
			plan := p
			plans = append(plans, &plan)
		}
	}

	return plans, nil
}

// RetrievePlanByName gets a plan by its external name.
func (sdk *SDK) RetrievePlanByName(name string, opts ScopeOptions) (Plan, error) {
	var searchResults []Plan
	lopts := metav1.ListOptions{
		FieldSelector: fields.OneTermEqualSelector(FieldExternalPlanName, name).String(),
	}
	if opts.Scope.Matches(ClusterScope) {
		cspList, err := sdk.ServiceCatalog().ClusterServicePlans().List(lopts)
		if err != nil {
			return nil, fmt.Errorf("unable to search plans by name '%s', (%s)", name, err)
		}
		for _, p := range cspList.Items {
			searchResults = append(searchResults, &p)
		}
	}
	if opts.Scope.Matches(NamespaceScope) {
		spList, err := sdk.ServiceCatalog().ServicePlans(opts.Namespace).List(lopts)
		if err != nil {
			return nil, fmt.Errorf("unable to search plans by name '%s', (%s)", name, err)
		}
		for _, p := range spList.Items {
			searchResults = append(searchResults, &p)
		}
	}

	listOpts := metav1.ListOptions{
		FieldSelector: fields.OneTermEqualSelector(FieldExternalPlanName, name).String(),
	}

	return sdk.retrieveSinglePlanByListOptions(name, opts, listOpts)
}

// RetrievePlanByClassAndName gets a plan by its external name and class name combination.
func (sdk *SDK) RetrievePlanByClassAndName(className, planName string, opts ScopeOptions) (Plan, error) {
	if opts.Scope == AllScope {
		return nil, errors.New("invalid scope: all")
	}

	class, err := sdk.RetrieveClassByName(className, opts)
	if err != nil {
		return nil, err
	}

	var classRefSelector fields.Selector
	if opts.Scope.Matches(ClusterScope) {
		classRefSelector = fields.OneTermEqualSelector(FieldClusterServiceClassRef, class.GetName())
	} else {
		classRefSelector = fields.OneTermEqualSelector(FieldServiceClassRef, class.GetName())
	}

	listOpts := metav1.ListOptions{
		FieldSelector: fields.AndSelectors(
			classRefSelector,
			fields.OneTermEqualSelector(FieldExternalPlanName, planName),
		).String(),
	}

	ss := []string{class.GetName(), planName}
	return sdk.retrieveSinglePlanByListOptions(strings.Join(ss, "/"), opts, listOpts)
}

func (sdk *SDK) retrieveSinglePlanByListOptions(name string, scopeOpts ScopeOptions, listOpts metav1.ListOptions) (Plan, error) {
	plans, err := sdk.retrievePlansByListOptions(scopeOpts, listOpts)
	if err != nil {
		return nil, err
	}
	if len(plans) == 0 {
		return nil, fmt.Errorf("plan not found '%s'", name)
	}
	if len(plans) > 1 {
		return nil, fmt.Errorf("more than one matching plan found for '%s'", name)
	}
	return plans[0], nil
}

// RetrievePlanByID gets a plan by its UUID.
func (sdk *SDK) RetrievePlanByID(uuid string, opts ScopeOptions) (Plan, error) {
	if opts.Scope == AllScope {
		return nil, errors.New("invalid scope: all")
	}

	if opts.Scope.Matches(ClusterScope) {
		p, err := sdk.ServiceCatalog().ClusterServicePlans().Get(uuid, metav1.GetOptions{})
		if err != nil {
			return nil, fmt.Errorf("unable to get cluster-scoped plan by uuid '%s' (%s)", uuid, err)
		}
		return p, nil
	}

	if opts.Scope.Matches(NamespaceScope) {
		p, err := sdk.ServiceCatalog().ServicePlans(opts.Namespace).Get(uuid, metav1.GetOptions{})
		if err != nil {
			return nil, fmt.Errorf("unable to get plan by uuid '%s' (%s)", uuid, err)
		}
		return p, nil
	}

	return nil, fmt.Errorf("unable to get plan by uuid '%s'", uuid)
}

// CreatePlan creates a new plan.
func (sdk *SDK) CreatePlan(opts CreatePlanOptions,
) (Plan, error) {
	if opts.Scope == AllScope {
		return nil, errors.New("invalid scope: all")
	}
	fromPlan, err := sdk.RetrievePlanByName(opts.From, ScopeOptions{
		Scope:     opts.Scope,
		Namespace: opts.Namespace,
	})
	if err != nil {
		return nil, err
	}
	if opts.Scope.Matches(NamespaceScope) {
		var plan = fromPlan.(*v1beta1.ServicePlan)
		return sdk.createServicePlan(plan, opts.Name, opts.Namespace)
	}
	var plan = fromPlan.(*v1beta1.ClusterServicePlan)
	return sdk.createClusterServicePlan(plan, opts.Name)
}

func (sdk *SDK) createClusterServicePlan(fromPlan *v1beta1.ClusterServicePlan, name string,
) (*v1beta1.ClusterServicePlan, error) {
	var plan = &v1beta1.ClusterServicePlan{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: fromPlan.Spec,
	}
	plan.Spec.ExternalName = name
	return sdk.ServiceCatalog().ClusterServicePlans().Create(plan)
}

func (sdk *SDK) createServicePlan(fromPlan *v1beta1.ServicePlan, name string,
	namespace string) (*v1beta1.ServicePlan, error) {
	var plan = &v1beta1.ServicePlan{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: fromPlan.Spec,
	}
	plan.Spec.ExternalName = name
	return sdk.ServiceCatalog().ServicePlans(plan.GetNamespace()).Create(plan)
}
