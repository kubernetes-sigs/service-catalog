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

package servicecatalog

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
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*BasicAuthConfig).DeepCopyInto(out.(*BasicAuthConfig))
			return nil
		}, InType: reflect.TypeOf(&BasicAuthConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*BearerTokenAuthConfig).DeepCopyInto(out.(*BearerTokenAuthConfig))
			return nil
		}, InType: reflect.TypeOf(&BearerTokenAuthConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterServiceBroker).DeepCopyInto(out.(*ClusterServiceBroker))
			return nil
		}, InType: reflect.TypeOf(&ClusterServiceBroker{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterServiceBrokerList).DeepCopyInto(out.(*ClusterServiceBrokerList))
			return nil
		}, InType: reflect.TypeOf(&ClusterServiceBrokerList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterServiceBrokerSpec).DeepCopyInto(out.(*ClusterServiceBrokerSpec))
			return nil
		}, InType: reflect.TypeOf(&ClusterServiceBrokerSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterServiceClass).DeepCopyInto(out.(*ClusterServiceClass))
			return nil
		}, InType: reflect.TypeOf(&ClusterServiceClass{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterServiceClassList).DeepCopyInto(out.(*ClusterServiceClassList))
			return nil
		}, InType: reflect.TypeOf(&ClusterServiceClassList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterServiceClassSpec).DeepCopyInto(out.(*ClusterServiceClassSpec))
			return nil
		}, InType: reflect.TypeOf(&ClusterServiceClassSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterServicePlan).DeepCopyInto(out.(*ClusterServicePlan))
			return nil
		}, InType: reflect.TypeOf(&ClusterServicePlan{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterServicePlanList).DeepCopyInto(out.(*ClusterServicePlanList))
			return nil
		}, InType: reflect.TypeOf(&ClusterServicePlanList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterServicePlanSpec).DeepCopyInto(out.(*ClusterServicePlanSpec))
			return nil
		}, InType: reflect.TypeOf(&ClusterServicePlanSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ParametersFromSource).DeepCopyInto(out.(*ParametersFromSource))
			return nil
		}, InType: reflect.TypeOf(&ParametersFromSource{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PlanReference).DeepCopyInto(out.(*PlanReference))
			return nil
		}, InType: reflect.TypeOf(&PlanReference{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SecretKeyReference).DeepCopyInto(out.(*SecretKeyReference))
			return nil
		}, InType: reflect.TypeOf(&SecretKeyReference{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ServiceBinding).DeepCopyInto(out.(*ServiceBinding))
			return nil
		}, InType: reflect.TypeOf(&ServiceBinding{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ServiceBindingCondition).DeepCopyInto(out.(*ServiceBindingCondition))
			return nil
		}, InType: reflect.TypeOf(&ServiceBindingCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ServiceBindingList).DeepCopyInto(out.(*ServiceBindingList))
			return nil
		}, InType: reflect.TypeOf(&ServiceBindingList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ServiceBindingPropertiesState).DeepCopyInto(out.(*ServiceBindingPropertiesState))
			return nil
		}, InType: reflect.TypeOf(&ServiceBindingPropertiesState{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ServiceBindingSpec).DeepCopyInto(out.(*ServiceBindingSpec))
			return nil
		}, InType: reflect.TypeOf(&ServiceBindingSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ServiceBindingStatus).DeepCopyInto(out.(*ServiceBindingStatus))
			return nil
		}, InType: reflect.TypeOf(&ServiceBindingStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ServiceBrokerAuthInfo).DeepCopyInto(out.(*ServiceBrokerAuthInfo))
			return nil
		}, InType: reflect.TypeOf(&ServiceBrokerAuthInfo{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ServiceBrokerCondition).DeepCopyInto(out.(*ServiceBrokerCondition))
			return nil
		}, InType: reflect.TypeOf(&ServiceBrokerCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ServiceBrokerStatus).DeepCopyInto(out.(*ServiceBrokerStatus))
			return nil
		}, InType: reflect.TypeOf(&ServiceBrokerStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ServiceClassStatus).DeepCopyInto(out.(*ServiceClassStatus))
			return nil
		}, InType: reflect.TypeOf(&ServiceClassStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ServiceInstance).DeepCopyInto(out.(*ServiceInstance))
			return nil
		}, InType: reflect.TypeOf(&ServiceInstance{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ServiceInstanceCondition).DeepCopyInto(out.(*ServiceInstanceCondition))
			return nil
		}, InType: reflect.TypeOf(&ServiceInstanceCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ServiceInstanceList).DeepCopyInto(out.(*ServiceInstanceList))
			return nil
		}, InType: reflect.TypeOf(&ServiceInstanceList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ServiceInstancePropertiesState).DeepCopyInto(out.(*ServiceInstancePropertiesState))
			return nil
		}, InType: reflect.TypeOf(&ServiceInstancePropertiesState{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ServiceInstanceSpec).DeepCopyInto(out.(*ServiceInstanceSpec))
			return nil
		}, InType: reflect.TypeOf(&ServiceInstanceSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ServiceInstanceStatus).DeepCopyInto(out.(*ServiceInstanceStatus))
			return nil
		}, InType: reflect.TypeOf(&ServiceInstanceStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ServicePlanStatus).DeepCopyInto(out.(*ServicePlanStatus))
			return nil
		}, InType: reflect.TypeOf(&ServicePlanStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*UserInfo).DeepCopyInto(out.(*UserInfo))
			return nil
		}, InType: reflect.TypeOf(&UserInfo{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasicAuthConfig) DeepCopyInto(out *BasicAuthConfig) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.ObjectReference)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasicAuthConfig.
func (in *BasicAuthConfig) DeepCopy() *BasicAuthConfig {
	if in == nil {
		return nil
	}
	out := new(BasicAuthConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BearerTokenAuthConfig) DeepCopyInto(out *BearerTokenAuthConfig) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.ObjectReference)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BearerTokenAuthConfig.
func (in *BearerTokenAuthConfig) DeepCopy() *BearerTokenAuthConfig {
	if in == nil {
		return nil
	}
	out := new(BearerTokenAuthConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterServiceBroker) DeepCopyInto(out *ClusterServiceBroker) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterServiceBroker.
func (in *ClusterServiceBroker) DeepCopy() *ClusterServiceBroker {
	if in == nil {
		return nil
	}
	out := new(ClusterServiceBroker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterServiceBroker) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterServiceBrokerList) DeepCopyInto(out *ClusterServiceBrokerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterServiceBroker, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterServiceBrokerList.
func (in *ClusterServiceBrokerList) DeepCopy() *ClusterServiceBrokerList {
	if in == nil {
		return nil
	}
	out := new(ClusterServiceBrokerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterServiceBrokerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterServiceBrokerSpec) DeepCopyInto(out *ClusterServiceBrokerSpec) {
	*out = *in
	if in.AuthInfo != nil {
		in, out := &in.AuthInfo, &out.AuthInfo
		if *in == nil {
			*out = nil
		} else {
			*out = new(ServiceBrokerAuthInfo)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.CABundle != nil {
		in, out := &in.CABundle, &out.CABundle
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.RelistDuration != nil {
		in, out := &in.RelistDuration, &out.RelistDuration
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.Duration)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterServiceBrokerSpec.
func (in *ClusterServiceBrokerSpec) DeepCopy() *ClusterServiceBrokerSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterServiceBrokerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterServiceClass) DeepCopyInto(out *ClusterServiceClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterServiceClass.
func (in *ClusterServiceClass) DeepCopy() *ClusterServiceClass {
	if in == nil {
		return nil
	}
	out := new(ClusterServiceClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterServiceClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterServiceClassList) DeepCopyInto(out *ClusterServiceClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterServiceClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterServiceClassList.
func (in *ClusterServiceClassList) DeepCopy() *ClusterServiceClassList {
	if in == nil {
		return nil
	}
	out := new(ClusterServiceClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterServiceClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterServiceClassSpec) DeepCopyInto(out *ClusterServiceClassSpec) {
	*out = *in
	if in.ExternalMetadata != nil {
		in, out := &in.ExternalMetadata, &out.ExternalMetadata
		if *in == nil {
			*out = nil
		} else {
			*out = new(runtime.RawExtension)
			(*in).DeepCopyInto(*out)
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
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterServiceClassSpec.
func (in *ClusterServiceClassSpec) DeepCopy() *ClusterServiceClassSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterServiceClassSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterServicePlan) DeepCopyInto(out *ClusterServicePlan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterServicePlan.
func (in *ClusterServicePlan) DeepCopy() *ClusterServicePlan {
	if in == nil {
		return nil
	}
	out := new(ClusterServicePlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterServicePlan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterServicePlanList) DeepCopyInto(out *ClusterServicePlanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterServicePlan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterServicePlanList.
func (in *ClusterServicePlanList) DeepCopy() *ClusterServicePlanList {
	if in == nil {
		return nil
	}
	out := new(ClusterServicePlanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterServicePlanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterServicePlanSpec) DeepCopyInto(out *ClusterServicePlanSpec) {
	*out = *in
	if in.Bindable != nil {
		in, out := &in.Bindable, &out.Bindable
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.ExternalMetadata != nil {
		in, out := &in.ExternalMetadata, &out.ExternalMetadata
		if *in == nil {
			*out = nil
		} else {
			*out = new(runtime.RawExtension)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.ServiceInstanceCreateParameterSchema != nil {
		in, out := &in.ServiceInstanceCreateParameterSchema, &out.ServiceInstanceCreateParameterSchema
		if *in == nil {
			*out = nil
		} else {
			*out = new(runtime.RawExtension)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.ServiceInstanceUpdateParameterSchema != nil {
		in, out := &in.ServiceInstanceUpdateParameterSchema, &out.ServiceInstanceUpdateParameterSchema
		if *in == nil {
			*out = nil
		} else {
			*out = new(runtime.RawExtension)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.ServiceBindingCreateParameterSchema != nil {
		in, out := &in.ServiceBindingCreateParameterSchema, &out.ServiceBindingCreateParameterSchema
		if *in == nil {
			*out = nil
		} else {
			*out = new(runtime.RawExtension)
			(*in).DeepCopyInto(*out)
		}
	}
	out.ClusterServiceClassRef = in.ClusterServiceClassRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterServicePlanSpec.
func (in *ClusterServicePlanSpec) DeepCopy() *ClusterServicePlanSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterServicePlanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParametersFromSource) DeepCopyInto(out *ParametersFromSource) {
	*out = *in
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		if *in == nil {
			*out = nil
		} else {
			*out = new(SecretKeyReference)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParametersFromSource.
func (in *ParametersFromSource) DeepCopy() *ParametersFromSource {
	if in == nil {
		return nil
	}
	out := new(ParametersFromSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanReference) DeepCopyInto(out *PlanReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanReference.
func (in *PlanReference) DeepCopy() *PlanReference {
	if in == nil {
		return nil
	}
	out := new(PlanReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeyReference) DeepCopyInto(out *SecretKeyReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeyReference.
func (in *SecretKeyReference) DeepCopy() *SecretKeyReference {
	if in == nil {
		return nil
	}
	out := new(SecretKeyReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBinding) DeepCopyInto(out *ServiceBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBinding.
func (in *ServiceBinding) DeepCopy() *ServiceBinding {
	if in == nil {
		return nil
	}
	out := new(ServiceBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingCondition) DeepCopyInto(out *ServiceBindingCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingCondition.
func (in *ServiceBindingCondition) DeepCopy() *ServiceBindingCondition {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingList) DeepCopyInto(out *ServiceBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingList.
func (in *ServiceBindingList) DeepCopy() *ServiceBindingList {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingPropertiesState) DeepCopyInto(out *ServiceBindingPropertiesState) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		if *in == nil {
			*out = nil
		} else {
			*out = new(runtime.RawExtension)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.UserInfo != nil {
		in, out := &in.UserInfo, &out.UserInfo
		if *in == nil {
			*out = nil
		} else {
			*out = new(UserInfo)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingPropertiesState.
func (in *ServiceBindingPropertiesState) DeepCopy() *ServiceBindingPropertiesState {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingPropertiesState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingSpec) DeepCopyInto(out *ServiceBindingSpec) {
	*out = *in
	out.ServiceInstanceRef = in.ServiceInstanceRef
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		if *in == nil {
			*out = nil
		} else {
			*out = new(runtime.RawExtension)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.ParametersFrom != nil {
		in, out := &in.ParametersFrom, &out.ParametersFrom
		*out = make([]ParametersFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UserInfo != nil {
		in, out := &in.UserInfo, &out.UserInfo
		if *in == nil {
			*out = nil
		} else {
			*out = new(UserInfo)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingSpec.
func (in *ServiceBindingSpec) DeepCopy() *ServiceBindingSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingStatus) DeepCopyInto(out *ServiceBindingStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ServiceBindingCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OperationStartTime != nil {
		in, out := &in.OperationStartTime, &out.OperationStartTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.InProgressProperties != nil {
		in, out := &in.InProgressProperties, &out.InProgressProperties
		if *in == nil {
			*out = nil
		} else {
			*out = new(ServiceBindingPropertiesState)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.ExternalProperties != nil {
		in, out := &in.ExternalProperties, &out.ExternalProperties
		if *in == nil {
			*out = nil
		} else {
			*out = new(ServiceBindingPropertiesState)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingStatus.
func (in *ServiceBindingStatus) DeepCopy() *ServiceBindingStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBrokerAuthInfo) DeepCopyInto(out *ServiceBrokerAuthInfo) {
	*out = *in
	if in.Basic != nil {
		in, out := &in.Basic, &out.Basic
		if *in == nil {
			*out = nil
		} else {
			*out = new(BasicAuthConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Bearer != nil {
		in, out := &in.Bearer, &out.Bearer
		if *in == nil {
			*out = nil
		} else {
			*out = new(BearerTokenAuthConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.BasicAuthSecret != nil {
		in, out := &in.BasicAuthSecret, &out.BasicAuthSecret
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.ObjectReference)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBrokerAuthInfo.
func (in *ServiceBrokerAuthInfo) DeepCopy() *ServiceBrokerAuthInfo {
	if in == nil {
		return nil
	}
	out := new(ServiceBrokerAuthInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBrokerCondition) DeepCopyInto(out *ServiceBrokerCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBrokerCondition.
func (in *ServiceBrokerCondition) DeepCopy() *ServiceBrokerCondition {
	if in == nil {
		return nil
	}
	out := new(ServiceBrokerCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBrokerStatus) DeepCopyInto(out *ServiceBrokerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ServiceBrokerCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OperationStartTime != nil {
		in, out := &in.OperationStartTime, &out.OperationStartTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBrokerStatus.
func (in *ServiceBrokerStatus) DeepCopy() *ServiceBrokerStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceBrokerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceClassStatus) DeepCopyInto(out *ServiceClassStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceClassStatus.
func (in *ServiceClassStatus) DeepCopy() *ServiceClassStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceClassStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceInstance) DeepCopyInto(out *ServiceInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceInstance.
func (in *ServiceInstance) DeepCopy() *ServiceInstance {
	if in == nil {
		return nil
	}
	out := new(ServiceInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceInstanceCondition) DeepCopyInto(out *ServiceInstanceCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceInstanceCondition.
func (in *ServiceInstanceCondition) DeepCopy() *ServiceInstanceCondition {
	if in == nil {
		return nil
	}
	out := new(ServiceInstanceCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceInstanceList) DeepCopyInto(out *ServiceInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceInstanceList.
func (in *ServiceInstanceList) DeepCopy() *ServiceInstanceList {
	if in == nil {
		return nil
	}
	out := new(ServiceInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceInstancePropertiesState) DeepCopyInto(out *ServiceInstancePropertiesState) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		if *in == nil {
			*out = nil
		} else {
			*out = new(runtime.RawExtension)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.UserInfo != nil {
		in, out := &in.UserInfo, &out.UserInfo
		if *in == nil {
			*out = nil
		} else {
			*out = new(UserInfo)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceInstancePropertiesState.
func (in *ServiceInstancePropertiesState) DeepCopy() *ServiceInstancePropertiesState {
	if in == nil {
		return nil
	}
	out := new(ServiceInstancePropertiesState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceInstanceSpec) DeepCopyInto(out *ServiceInstanceSpec) {
	*out = *in
	out.PlanReference = in.PlanReference
	if in.ClusterServiceClassRef != nil {
		in, out := &in.ClusterServiceClassRef, &out.ClusterServiceClassRef
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.ObjectReference)
			**out = **in
		}
	}
	if in.ClusterServicePlanRef != nil {
		in, out := &in.ClusterServicePlanRef, &out.ClusterServicePlanRef
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.ObjectReference)
			**out = **in
		}
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		if *in == nil {
			*out = nil
		} else {
			*out = new(runtime.RawExtension)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.ParametersFrom != nil {
		in, out := &in.ParametersFrom, &out.ParametersFrom
		*out = make([]ParametersFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UserInfo != nil {
		in, out := &in.UserInfo, &out.UserInfo
		if *in == nil {
			*out = nil
		} else {
			*out = new(UserInfo)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceInstanceSpec.
func (in *ServiceInstanceSpec) DeepCopy() *ServiceInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceInstanceStatus) DeepCopyInto(out *ServiceInstanceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ServiceInstanceCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastOperation != nil {
		in, out := &in.LastOperation, &out.LastOperation
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.DashboardURL != nil {
		in, out := &in.DashboardURL, &out.DashboardURL
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.OperationStartTime != nil {
		in, out := &in.OperationStartTime, &out.OperationStartTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.InProgressProperties != nil {
		in, out := &in.InProgressProperties, &out.InProgressProperties
		if *in == nil {
			*out = nil
		} else {
			*out = new(ServiceInstancePropertiesState)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.ExternalProperties != nil {
		in, out := &in.ExternalProperties, &out.ExternalProperties
		if *in == nil {
			*out = nil
		} else {
			*out = new(ServiceInstancePropertiesState)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceInstanceStatus.
func (in *ServiceInstanceStatus) DeepCopy() *ServiceInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicePlanStatus) DeepCopyInto(out *ServicePlanStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicePlanStatus.
func (in *ServicePlanStatus) DeepCopy() *ServicePlanStatus {
	if in == nil {
		return nil
	}
	out := new(ServicePlanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserInfo) DeepCopyInto(out *UserInfo) {
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Extra != nil {
		in, out := &in.Extra, &out.Extra
		*out = make(map[string]ExtraValue, len(*in))
		for key, val := range *in {
			(*out)[key] = make(ExtraValue, len(val))
			copy((*out)[key], val)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserInfo.
func (in *UserInfo) DeepCopy() *UserInfo {
	if in == nil {
		return nil
	}
	out := new(UserInfo)
	in.DeepCopyInto(out)
	return out
}
