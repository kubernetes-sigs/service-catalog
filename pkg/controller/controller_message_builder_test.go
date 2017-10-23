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

package controller

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

// MessageBuilder builds Event Type messages to help unit tests create these strings
// Example usage:
// mb := new(MessageBuilder).warning().reson("ReasonForError).msg("Error: hello world.")
// ms.String()
type MessageBuilder struct {
	eventMessage  string
	reasonMessage string
	message       string
}

func (mb *MessageBuilder) warning() *MessageBuilder {
	mb.eventMessage = corev1.EventTypeWarning
	return mb
}

func (mb *MessageBuilder) normal() *MessageBuilder {
	mb.eventMessage = corev1.EventTypeNormal
	return mb
}

func (mb *MessageBuilder) reason(reason string) *MessageBuilder {
	mb.reasonMessage = reason
	return mb
}

// Appends a message to the message builder.
func (mb *MessageBuilder) msg(msg string) *MessageBuilder {
	space := ""
	if mb.message > "" {
		space = " "
	}
	mb.message = fmt.Sprintf(`%s%s%s`, mb.message, space, msg)
	return mb
}

func (mb *MessageBuilder) msgf(format string, a ...interface{}) *MessageBuilder {
	msg := fmt.Sprintf(format, a...)
	return mb.msg(msg)
}

func (mb *MessageBuilder) StringArr() []string {
	return []string{mb.String()}
}

// msgCreateServiceBindingError Adds a message in the following form:
// "Error creating ServiceBinding for ServiceInstance %q of ClusterServiceClass (K8S: %q ExternalName: %q) at ClusterServiceBroker %q:"
func (mb *MessageBuilder) msgCreateServiceBindingError(serviceInstance, serviceClassK8S, serviceClassExternalName, broker string) *MessageBuilder {
	msg := fmt.Sprintf("Error creating ServiceBinding for ServiceInstance %q of ClusterServiceClass (K8S: %q ExternalName: %q) at ClusterServiceBroker %q:",
		serviceInstance, serviceClassK8S, serviceClassExternalName, broker)
	return mb.msg(msg)
}

// msgUnbindingError Adds a message in the following form:
// "Error unbinding from ServiceInstance %q of ClusterServiceClass (K8S: %q ExternalName: %q) at ClusterServiceBroker %q:"
func (mb *MessageBuilder) msgUnbindingError(serviceInstance, serviceClassK8S, serviceClassExternalName, broker string) *MessageBuilder {
	msg := fmt.Sprintf("Error unbinding from ServiceInstance %q of ClusterServiceClass (K8S: %q ExternalName: %q) at ClusterServiceBroker %q:",
		serviceInstance, serviceClassK8S, serviceClassExternalName, broker)
	return mb.msg(msg)
}

// msgNonBindableClusterServiceClass Adds a message in the following form:
// "References a non-bindable ClusterServiceClass (K8S: %q ExternalName: %q) and Plan (%q)"
func (mb *MessageBuilder) msgNonBindableClusterServiceClass(serviceClassK8S, serviceClassExternalName, plan string) *MessageBuilder {
	msg := fmt.Sprintf("References a non-bindable ClusterServiceClass (K8S: %q ExternalName: %q) and Plan (%q) combination",
		serviceClassK8S, serviceClassExternalName, plan)
	return mb.msg(msg)
}

func (mb *MessageBuilder) String() string {
	s := ""
	space := ""
	if mb.eventMessage > "" {
		s += fmt.Sprintf("%s%s", space, mb.eventMessage)
		space = " "
	}
	if mb.reasonMessage > "" {
		s += fmt.Sprintf("%s%s", space, mb.reasonMessage)
		space = " "
	}
	if mb.message > "" {
		s += fmt.Sprintf("%s%s", space, mb.message)
		space = " "
	}
	return s
}
