#!/usr/bin/env bash
# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit

EXIT_CODE=0

function checkIfClassExist(){
    local externalName=$1
    local classesNames=$2

    for class in ${classesNames}
    do
        if [[ "${class}" = "${externalName}" ]]; then
            return 0
        fi
    done

    return 1
}

function checkIfClassExistForInstance(){
    serviceClassesNames=$(kubectl get serviceclasses --all-namespaces -ojsonpath="{.items[*].spec.externalName}")
    clusterServiceClassesNames=$(kubectl get clusterserviceclasses -ojsonpath="{.items[*].spec.externalName}")

    for className in $(kubectl get serviceinstances --all-namespaces -ojsonpath="{.items[*].status.userSpecifiedClassName}")
    do
        type=$(echo "${className}" | cut -d'/' -f1)
        externalName=$(echo "${className}" | cut -d'/' -f2)

        set +o errexit
        if [[ "${type}" = "ClusterServiceClass" ]]; then
            checkIfClassExist "${externalName}" "${clusterServiceClassesNames}"
            if [[ $? -eq 1 ]]; then
                echo "${className} not exist in the cluster for the ServiceInstance anymore"
                EXIT_CODE=1
            fi
        fi
        if [[ "${type}" = "ServiceClass" ]]; then
            checkIfClassExist "${externalName}" "${serviceClassesNames}"
            if [[ $? -eq 1 ]]; then
                echo "${className} not exist in the cluster for the ServiceInstance anymore"
                EXIT_CODE=1
            fi
        fi
    done
    set -o errexit
}

#
# Check if any class/plans were removed from broker's catalog
#
CSC=$(kubectl get clusterserviceclasses -ojsonpath="{.items[*].metadata.deletionTimestamp}")
if [[ -n ${CSC} ]]; then
    echo "There are ClusterServiceClasses with deletionTimestamp set"
    EXIT_CODE=1
fi
for status in $(kubectl get clusterserviceclasses -ojsonpath="{.items[*].status.removedFromBrokerCatalog}")
do
if [[ -n "${status}" ]] && ${status}; then
    echo "There are removedFromBrokerCatalog ClusterServiceClasses"
    EXIT_CODE=1
fi
done

SC=$(kubectl get serviceclasses --all-namespaces -ojsonpath="{.items[*].metadata.deletionTimestamp}")
if [[ -n ${SC} ]]; then
    echo "There are ServiceClasses with deletionTimestamp set"
    EXIT_CODE=1
fi
for status in $(kubectl get serviceclasses --all-namespaces -ojsonpath="{.items[*].status.removedFromBrokerCatalog}")
do
if [[ -n "${status}" ]] && ${status}; then
    echo "There are removedFromBrokerCatalog ServiceClasses"
    EXIT_CODE=1
fi
done

CSP=$(kubectl get clusterserviceplans -ojsonpath="{.items[*].metadata.deletionTimestamp}")
if [[ -n ${CSP} ]]; then
    echo "There are ClusterServicePlans with deletionTimestamp set"
    EXIT_CODE=1
fi
for status in $(kubectl get clusterserviceplans -ojsonpath="{.items[*].status.removedFromBrokerCatalog}")
do
if [[ -n "${status}" ]] && ${status}; then
    echo "There are removedFromBrokerCatalog ClusterServicePlans"
    EXIT_CODE=1
fi
done

SP=$(kubectl get serviceplans --all-namespaces -ojsonpath="{.items[*].metadata.deletionTimestamp}")
if [[ -n ${SP} ]]; then
    echo "There are ServicePlans with deletionTimestamp set"
    EXIT_CODE=1
fi
for status in $(kubectl get serviceplans --all-namespaces -ojsonpath="{.items[*].status.removedFromBrokerCatalog}")
do
if [[ -n "${status}" ]] && ${status}; then
    echo "There are removedFromBrokerCatalog ServicePlans"
    EXIT_CODE=1
fi
done

#
# Check if any instance/binding is in progress or is being deleted
#
SI=$(kubectl get serviceinstances --all-namespaces -ojsonpath="{.items[*].metadata.deletionTimestamp}")
if [[ -n "${SI}" ]]; then
    echo "There are ServiceInstances with deletionTimestamp set"
    EXIT_CODE=1
fi
for status in $(kubectl get serviceinstances --all-namespaces -ojsonpath="{.items[*].status.asyncOpInProgress}")
do
if [[ -n "${status}" ]] && ${status}; then
    echo "There are ServiceInstance in progress"
    EXIT_CODE=1
fi
done
# checks if there are some instances with not existing classes
checkIfClassExistForInstance

SBI=$(kubectl get servicebindings --all-namespaces -ojsonpath="{.items[*].metadata.deletionTimestamp}")
if [[ -n ${SBI} ]]; then
    echo "There are ServiceBindings with deletionTimestamp set"
    EXIT_CODE=1
fi
for status in $(kubectl get servicebindings --all-namespaces -ojsonpath="{.items[*].status.asyncOpInProgress}")
do
if [[ -n "${status}" ]] && ${status}; then
    echo "There are ServiceBinding in progress"
    EXIT_CODE=1
fi
done

#
# Check if any broker is being deleted
#
SB=$(kubectl get servicebrokers --all-namespaces -ojsonpath="{.items[*].metadata.deletionTimestamp}")
if [[ -n ${SB} ]]; then
    echo "There are being deleted ServiceBrokers"
    EXIT_CODE=1
fi
CSB=$(kubectl get clusterservicebrokers -ojsonpath="{.items[*].metadata.deletionTimestamp}")
if [[ -n ${CSB} ]]; then
    echo "There are being deleted ClusterServiceBrokers"
    EXIT_CODE=1
fi

if [[ ${EXIT_CODE} -eq 0 ]]; then
    echo "Your Service Catalog resources are ready to migrate!"
else
    echo "Please prepare your Service Catalog resources before migration. Check docs/migration-apiserver-to-crds.md#implementation-details"
fi

exit ${EXIT_CODE}
