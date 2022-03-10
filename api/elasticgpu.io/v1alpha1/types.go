/*
Copyright 2022.

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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ElasticGPUSpec defines the desired state of ElasticGPU
type ElasticGPUSpec struct {
	Capacity         ResourceList `json:"capacity"`
	ElasticGPUSource `json:"elasticGPUSource"`
	ClaimRef         v1.ObjectReference `json:"claimRef"`
	NodeAffinity     GPUNodeAffinity    `json:"nodeAffinity"`
}

type GPUNodeAffinity struct {
	Required *v1.NodeSelector `json:"required"`
}

// ElasticGPUStatus defines the observed state of ElasticGPU
type ElasticGPUStatus struct {
	Phase ElasticGPUPhase `json:"phase,omitempty" protobuf:"bytes,1,opt,name=phase,casttype=ElasticGPUPhase"`
	// A human-readable message indicating details about why the volume is in this state.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,2,opt,name=message"`
	// Reason is a brief CamelCase string that describes any failure and is meant
	// for machine parsing and tidy display in the CLI.
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,3,opt,name=reason"`
}

type ElasticGPUPhase string

const (
	GPUPending   ElasticGPUPhase = "Pending"
	GPUAvailable ElasticGPUPhase = "Available"
	GPUBound     ElasticGPUPhase = "Bound"
	GPUReleased  ElasticGPUPhase = "Released"
	GPUFailed    ElasticGPUPhase = "Failed"
)

//+genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ElasticGPU is the Schema for the elasticgpus API
type ElasticGPU struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ElasticGPUSpec   `json:"spec,omitempty"`
	Status ElasticGPUStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ElasticGPUList contains a list of ElasticGPU
type ElasticGPUList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticGPU `json:"items"`
}

type ElasticGPUSource struct {
	qGPU        *qGPU        `json:"qGPU"`
	PhysicalGPU *PhysicalGPU `json:"physicalGPU"`
	GPUShare    *GPUShare    `json:"gpuShare"`
}

type qGPU struct {
	DeviceName string `json:"DeviceName"`
	GPUIndex   int    `json:"gpuIndex"`
}

type PhysicalGPU struct {
	GPUIndex int `json:"gpuIndex"`
}

type GPUShare struct {
	GPUIndex string `json:"gpuIndex"`
}

type ResourceList map[ResourceName]resource.Quantity
type ResourceName string

const (
	ResourceGPUCore   ResourceName = "gpu-core"
	ResourceGPUMemory ResourceName = "gpu-memory"
)

func init() {
	SchemeBuilder.Register(&ElasticGPU{}, &ElasticGPUList{})
}
