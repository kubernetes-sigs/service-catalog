// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_BasicAuthConfig, InType: reflect.TypeOf(&BasicAuthConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_BearerTokenAuthConfig, InType: reflect.TypeOf(&BearerTokenAuthConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ClusterServiceBroker, InType: reflect.TypeOf(&ClusterServiceBroker{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ClusterServiceBrokerList, InType: reflect.TypeOf(&ClusterServiceBrokerList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ClusterServiceBrokerSpec, InType: reflect.TypeOf(&ClusterServiceBrokerSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ClusterServiceClass, InType: reflect.TypeOf(&ClusterServiceClass{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ClusterServiceClassList, InType: reflect.TypeOf(&ClusterServiceClassList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ClusterServiceClassSpec, InType: reflect.TypeOf(&ClusterServiceClassSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ClusterServicePlan, InType: reflect.TypeOf(&ClusterServicePlan{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ClusterServicePlanList, InType: reflect.TypeOf(&ClusterServicePlanList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ClusterServicePlanSpec, InType: reflect.TypeOf(&ClusterServicePlanSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_DesiredPlan, InType: reflect.TypeOf(&DesiredPlan{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ParametersFromSource, InType: reflect.TypeOf(&ParametersFromSource{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_SecretKeyReference, InType: reflect.TypeOf(&SecretKeyReference{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServiceBinding, InType: reflect.TypeOf(&ServiceBinding{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServiceBindingCondition, InType: reflect.TypeOf(&ServiceBindingCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServiceBindingList, InType: reflect.TypeOf(&ServiceBindingList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServiceBindingPropertiesState, InType: reflect.TypeOf(&ServiceBindingPropertiesState{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServiceBindingSpec, InType: reflect.TypeOf(&ServiceBindingSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServiceBindingStatus, InType: reflect.TypeOf(&ServiceBindingStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServiceBrokerAuthInfo, InType: reflect.TypeOf(&ServiceBrokerAuthInfo{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServiceBrokerCondition, InType: reflect.TypeOf(&ServiceBrokerCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServiceBrokerStatus, InType: reflect.TypeOf(&ServiceBrokerStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServiceClassStatus, InType: reflect.TypeOf(&ServiceClassStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServiceInstance, InType: reflect.TypeOf(&ServiceInstance{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServiceInstanceCondition, InType: reflect.TypeOf(&ServiceInstanceCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServiceInstanceList, InType: reflect.TypeOf(&ServiceInstanceList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServiceInstancePropertiesState, InType: reflect.TypeOf(&ServiceInstancePropertiesState{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServiceInstanceSpec, InType: reflect.TypeOf(&ServiceInstanceSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServiceInstanceStatus, InType: reflect.TypeOf(&ServiceInstanceStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServicePlanStatus, InType: reflect.TypeOf(&ServicePlanStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_UserInfo, InType: reflect.TypeOf(&UserInfo{})},
	)
}

// DeepCopy_v1alpha1_BasicAuthConfig is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_BasicAuthConfig(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BasicAuthConfig)
		out := out.(*BasicAuthConfig)
		*out = *in
		if in.SecretRef != nil {
			in, out := &in.SecretRef, &out.SecretRef
			*out = new(v1.ObjectReference)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_v1alpha1_BearerTokenAuthConfig is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_BearerTokenAuthConfig(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BearerTokenAuthConfig)
		out := out.(*BearerTokenAuthConfig)
		*out = *in
		if in.SecretRef != nil {
			in, out := &in.SecretRef, &out.SecretRef
			*out = new(v1.ObjectReference)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ClusterServiceBroker is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ClusterServiceBroker(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterServiceBroker)
		out := out.(*ClusterServiceBroker)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*meta_v1.ObjectMeta)
		}
		if err := DeepCopy_v1alpha1_ClusterServiceBrokerSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1alpha1_ServiceBrokerStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ClusterServiceBrokerList is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ClusterServiceBrokerList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterServiceBrokerList)
		out := out.(*ClusterServiceBrokerList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ClusterServiceBroker, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_ClusterServiceBroker(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ClusterServiceBrokerSpec is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ClusterServiceBrokerSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterServiceBrokerSpec)
		out := out.(*ClusterServiceBrokerSpec)
		*out = *in
		if in.AuthInfo != nil {
			in, out := &in.AuthInfo, &out.AuthInfo
			*out = new(ServiceBrokerAuthInfo)
			if err := DeepCopy_v1alpha1_ServiceBrokerAuthInfo(*in, *out, c); err != nil {
				return err
			}
		}
		if in.CABundle != nil {
			in, out := &in.CABundle, &out.CABundle
			*out = make([]byte, len(*in))
			copy(*out, *in)
		}
		if in.RelistDuration != nil {
			in, out := &in.RelistDuration, &out.RelistDuration
			*out = new(meta_v1.Duration)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ClusterServiceClass is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ClusterServiceClass(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterServiceClass)
		out := out.(*ClusterServiceClass)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*meta_v1.ObjectMeta)
		}
		if err := DeepCopy_v1alpha1_ClusterServiceClassSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ClusterServiceClassList is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ClusterServiceClassList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterServiceClassList)
		out := out.(*ClusterServiceClassList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ClusterServiceClass, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_ClusterServiceClass(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ClusterServiceClassSpec is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ClusterServiceClassSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterServiceClassSpec)
		out := out.(*ClusterServiceClassSpec)
		*out = *in
		if in.ExternalMetadata != nil {
			in, out := &in.ExternalMetadata, &out.ExternalMetadata
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*runtime.RawExtension)
			}
		}
		if in.Tags != nil {
			in, out := &in.Tags, &out.Tags
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Requires != nil {
			in, out := &in.Requires, &out.Requires
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ClusterServicePlan is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ClusterServicePlan(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterServicePlan)
		out := out.(*ClusterServicePlan)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*meta_v1.ObjectMeta)
		}
		if err := DeepCopy_v1alpha1_ClusterServicePlanSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ClusterServicePlanList is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ClusterServicePlanList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterServicePlanList)
		out := out.(*ClusterServicePlanList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ClusterServicePlan, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_ClusterServicePlan(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ClusterServicePlanSpec is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ClusterServicePlanSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterServicePlanSpec)
		out := out.(*ClusterServicePlanSpec)
		*out = *in
		if in.Bindable != nil {
			in, out := &in.Bindable, &out.Bindable
			*out = new(bool)
			**out = **in
		}
		if in.ExternalMetadata != nil {
			in, out := &in.ExternalMetadata, &out.ExternalMetadata
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*runtime.RawExtension)
			}
		}
		if in.ServiceInstanceCreateParameterSchema != nil {
			in, out := &in.ServiceInstanceCreateParameterSchema, &out.ServiceInstanceCreateParameterSchema
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*runtime.RawExtension)
			}
		}
		if in.ServiceInstanceUpdateParameterSchema != nil {
			in, out := &in.ServiceInstanceUpdateParameterSchema, &out.ServiceInstanceUpdateParameterSchema
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*runtime.RawExtension)
			}
		}
		if in.ServiceBindingCreateParameterSchema != nil {
			in, out := &in.ServiceBindingCreateParameterSchema, &out.ServiceBindingCreateParameterSchema
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*runtime.RawExtension)
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_DesiredPlan is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_DesiredPlan(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DesiredPlan)
		out := out.(*DesiredPlan)
		*out = *in
		return nil
	}
}

// DeepCopy_v1alpha1_ParametersFromSource is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ParametersFromSource(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ParametersFromSource)
		out := out.(*ParametersFromSource)
		*out = *in
		if in.SecretKeyRef != nil {
			in, out := &in.SecretKeyRef, &out.SecretKeyRef
			*out = new(SecretKeyReference)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_v1alpha1_SecretKeyReference is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_SecretKeyReference(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SecretKeyReference)
		out := out.(*SecretKeyReference)
		*out = *in
		return nil
	}
}

// DeepCopy_v1alpha1_ServiceBinding is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ServiceBinding(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceBinding)
		out := out.(*ServiceBinding)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*meta_v1.ObjectMeta)
		}
		if err := DeepCopy_v1alpha1_ServiceBindingSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1alpha1_ServiceBindingStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ServiceBindingCondition is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ServiceBindingCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceBindingCondition)
		out := out.(*ServiceBindingCondition)
		*out = *in
		if newVal, err := c.DeepCopy(&in.LastTransitionTime); err != nil {
			return err
		} else {
			out.LastTransitionTime = *newVal.(*meta_v1.Time)
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ServiceBindingList is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ServiceBindingList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceBindingList)
		out := out.(*ServiceBindingList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ServiceBinding, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_ServiceBinding(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ServiceBindingPropertiesState is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ServiceBindingPropertiesState(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceBindingPropertiesState)
		out := out.(*ServiceBindingPropertiesState)
		*out = *in
		if in.Parameters != nil {
			in, out := &in.Parameters, &out.Parameters
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*runtime.RawExtension)
			}
		}
		if in.UserInfo != nil {
			in, out := &in.UserInfo, &out.UserInfo
			*out = new(UserInfo)
			if err := DeepCopy_v1alpha1_UserInfo(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ServiceBindingSpec is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ServiceBindingSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceBindingSpec)
		out := out.(*ServiceBindingSpec)
		*out = *in
		if in.Parameters != nil {
			in, out := &in.Parameters, &out.Parameters
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*runtime.RawExtension)
			}
		}
		if in.ParametersFrom != nil {
			in, out := &in.ParametersFrom, &out.ParametersFrom
			*out = make([]ParametersFromSource, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_ParametersFromSource(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.UserInfo != nil {
			in, out := &in.UserInfo, &out.UserInfo
			*out = new(UserInfo)
			if err := DeepCopy_v1alpha1_UserInfo(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ServiceBindingStatus is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ServiceBindingStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceBindingStatus)
		out := out.(*ServiceBindingStatus)
		*out = *in
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]ServiceBindingCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_ServiceBindingCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.OperationStartTime != nil {
			in, out := &in.OperationStartTime, &out.OperationStartTime
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*meta_v1.Time)
			}
		}
		if in.InProgressProperties != nil {
			in, out := &in.InProgressProperties, &out.InProgressProperties
			*out = new(ServiceBindingPropertiesState)
			if err := DeepCopy_v1alpha1_ServiceBindingPropertiesState(*in, *out, c); err != nil {
				return err
			}
		}
		if in.ExternalProperties != nil {
			in, out := &in.ExternalProperties, &out.ExternalProperties
			*out = new(ServiceBindingPropertiesState)
			if err := DeepCopy_v1alpha1_ServiceBindingPropertiesState(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ServiceBrokerAuthInfo is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ServiceBrokerAuthInfo(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceBrokerAuthInfo)
		out := out.(*ServiceBrokerAuthInfo)
		*out = *in
		if in.Basic != nil {
			in, out := &in.Basic, &out.Basic
			*out = new(BasicAuthConfig)
			if err := DeepCopy_v1alpha1_BasicAuthConfig(*in, *out, c); err != nil {
				return err
			}
		}
		if in.Bearer != nil {
			in, out := &in.Bearer, &out.Bearer
			*out = new(BearerTokenAuthConfig)
			if err := DeepCopy_v1alpha1_BearerTokenAuthConfig(*in, *out, c); err != nil {
				return err
			}
		}
		if in.BasicAuthSecret != nil {
			in, out := &in.BasicAuthSecret, &out.BasicAuthSecret
			*out = new(v1.ObjectReference)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ServiceBrokerCondition is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ServiceBrokerCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceBrokerCondition)
		out := out.(*ServiceBrokerCondition)
		*out = *in
		if newVal, err := c.DeepCopy(&in.LastTransitionTime); err != nil {
			return err
		} else {
			out.LastTransitionTime = *newVal.(*meta_v1.Time)
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ServiceBrokerStatus is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ServiceBrokerStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceBrokerStatus)
		out := out.(*ServiceBrokerStatus)
		*out = *in
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]ServiceBrokerCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_ServiceBrokerCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.OperationStartTime != nil {
			in, out := &in.OperationStartTime, &out.OperationStartTime
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*meta_v1.Time)
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ServiceClassStatus is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ServiceClassStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceClassStatus)
		out := out.(*ServiceClassStatus)
		*out = *in
		return nil
	}
}

// DeepCopy_v1alpha1_ServiceInstance is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ServiceInstance(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceInstance)
		out := out.(*ServiceInstance)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*meta_v1.ObjectMeta)
		}
		if err := DeepCopy_v1alpha1_ServiceInstanceSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1alpha1_ServiceInstanceStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ServiceInstanceCondition is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ServiceInstanceCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceInstanceCondition)
		out := out.(*ServiceInstanceCondition)
		*out = *in
		if newVal, err := c.DeepCopy(&in.LastTransitionTime); err != nil {
			return err
		} else {
			out.LastTransitionTime = *newVal.(*meta_v1.Time)
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ServiceInstanceList is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ServiceInstanceList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceInstanceList)
		out := out.(*ServiceInstanceList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ServiceInstance, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_ServiceInstance(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ServiceInstancePropertiesState is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ServiceInstancePropertiesState(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceInstancePropertiesState)
		out := out.(*ServiceInstancePropertiesState)
		*out = *in
		if in.Parameters != nil {
			in, out := &in.Parameters, &out.Parameters
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*runtime.RawExtension)
			}
		}
		if in.UserInfo != nil {
			in, out := &in.UserInfo, &out.UserInfo
			*out = new(UserInfo)
			if err := DeepCopy_v1alpha1_UserInfo(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ServiceInstanceSpec is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ServiceInstanceSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceInstanceSpec)
		out := out.(*ServiceInstanceSpec)
		*out = *in
		if in.ClusterServiceClassRef != nil {
			in, out := &in.ClusterServiceClassRef, &out.ClusterServiceClassRef
			*out = new(v1.ObjectReference)
			**out = **in
		}
		if in.ClusterServicePlanRef != nil {
			in, out := &in.ClusterServicePlanRef, &out.ClusterServicePlanRef
			*out = new(v1.ObjectReference)
			**out = **in
		}
		if in.Parameters != nil {
			in, out := &in.Parameters, &out.Parameters
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*runtime.RawExtension)
			}
		}
		if in.ParametersFrom != nil {
			in, out := &in.ParametersFrom, &out.ParametersFrom
			*out = make([]ParametersFromSource, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_ParametersFromSource(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.UserInfo != nil {
			in, out := &in.UserInfo, &out.UserInfo
			*out = new(UserInfo)
			if err := DeepCopy_v1alpha1_UserInfo(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ServiceInstanceStatus is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ServiceInstanceStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceInstanceStatus)
		out := out.(*ServiceInstanceStatus)
		*out = *in
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]ServiceInstanceCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_ServiceInstanceCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.LastOperation != nil {
			in, out := &in.LastOperation, &out.LastOperation
			*out = new(string)
			**out = **in
		}
		if in.DashboardURL != nil {
			in, out := &in.DashboardURL, &out.DashboardURL
			*out = new(string)
			**out = **in
		}
		if in.OperationStartTime != nil {
			in, out := &in.OperationStartTime, &out.OperationStartTime
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*meta_v1.Time)
			}
		}
		if in.InProgressProperties != nil {
			in, out := &in.InProgressProperties, &out.InProgressProperties
			*out = new(ServiceInstancePropertiesState)
			if err := DeepCopy_v1alpha1_ServiceInstancePropertiesState(*in, *out, c); err != nil {
				return err
			}
		}
		if in.ExternalProperties != nil {
			in, out := &in.ExternalProperties, &out.ExternalProperties
			*out = new(ServiceInstancePropertiesState)
			if err := DeepCopy_v1alpha1_ServiceInstancePropertiesState(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ServicePlanStatus is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ServicePlanStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServicePlanStatus)
		out := out.(*ServicePlanStatus)
		*out = *in
		return nil
	}
}

// DeepCopy_v1alpha1_UserInfo is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_UserInfo(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UserInfo)
		out := out.(*UserInfo)
		*out = *in
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Extra != nil {
			in, out := &in.Extra, &out.Extra
			*out = make(map[string]ExtraValue)
			for key, val := range *in {
				if newVal, err := c.DeepCopy(&val); err != nil {
					return err
				} else {
					(*out)[key] = *newVal.(*ExtraValue)
				}
			}
		}
		return nil
	}
}
