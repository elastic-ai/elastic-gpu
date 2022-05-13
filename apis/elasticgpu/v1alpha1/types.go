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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ElasticGPUSpec defines the desired state of ElasticGPU
type ElasticGPUSpec struct {
	Capacity         v1.ResourceList `json:"capacity,omitempty"`
	ElasticGPUSource `json:",inline"`
	ClaimRef         v1.ObjectReference `json:"claimRef,omitempty"`
	NodeAffinity     GPUNodeAffinity    `json:"nodeAffinity,omitempty"`
}

type GPUNodeAffinity struct {
	Required *v1.NodeSelector `json:"required,omitempty"`
}

// ElasticGPUStatus defines the observed state of ElasticGPU
type ElasticGPUStatus struct {
	Phase ElasticGPUPhase `json:"phase,omitempty"`
	// A human-readable message indicating details about why the elastic gpu is in this state.
	// +optional
	Message string `json:"message,omitempty"`
	// Reason is a brief CamelCase string that describes any failure and is meant
	// for machine parsing and tidy display in the CLI.
	// +optional
	Reason string `json:"reason,omitempty"`
}

type ElasticGPUPhase string

const (
	GPUPending   ElasticGPUPhase = "Pending"
	GPUAvailable ElasticGPUPhase = "Available"
	GPUBound     ElasticGPUPhase = "Bound"
	GPUReleased  ElasticGPUPhase = "Released"
	GPUFailed    ElasticGPUPhase = "Failed"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ElasticGPU is the Schema for the elasticgpus API
type ElasticGPU struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticGPUSpec   `json:"spec,omitempty"`
	Status            ElasticGPUStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ElasticGPUList is a list of ElasticGPU items
type ElasticGPUList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticGPU `json:"items"`
}

type ElasticGPUSource struct {
	QGPU        *QGPUSource        `json:"qGPU,omitempty"`
	PhysicalGPU *PhysicalGPUSource `json:"physicalGPU,omitempty"`
	GPUShare    *GPUShareSource    `json:"gpuShare,omitempty"`
}

type QGPUSource struct {
	GPUName    string   `json:"gpuName,omitempty"`
	DeviceName string   `json:"deviceName,omitempty"`
	Paths      []string `json:"paths,omitempty"`
}

type PhysicalGPUSource struct {
	GPUNames []string `json:"gpuNames"`
}

type GPUShareSource struct {
	GPUName string `json:"gpuName,omitempty"`
}

const (
	ResourceGPUCore         v1.ResourceName = "elasticgpu.com/gpu-core"
	ResourceGPUMemory       v1.ResourceName = "elasticgpu.com/gpu-memory"
	ResourceQGPUCore        v1.ResourceName = "tke.cloud.tencent.com/qgpu-core"
	ResourceQGPUOfflineCore v1.ResourceName = "tke.cloud.tencent.com/qgpu-core-greedy"
	ResourceQGPUMemory      v1.ResourceName = "tke.cloud.tencent.com/qgpu-memory"
	ResourcePGPU            v1.ResourceName = "nvidia.com/gpu"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ElasticGPUClaim is a user's request for and claim to a ElasticGPU
type ElasticGPUClaim struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticGPUClaimSpec   `json:"spec,omitempty"`
	Status            ElasticGPUClaimStatus `json:"status,omitempty"`
}

// ElasticGPUClaimSpec is the specification of a ElasticGPUClaim
type ElasticGPUClaimSpec struct {
	Resources           v1.ResourceRequirements `json:"resources,omitempty"`
	ElasticGPUName      string                  `json:"elasticGPUName,omitempty"`
	ElasticGPUClassName *string                 `json:"elasticGPUClassName,omitempty"`
}

// ElasticGPUClaimStatus is the current status of a ElasticGPUClaim
type ElasticGPUClaimStatus struct {
	Phase ElasticGPUClaimPhase `json:"phase,omitempty"`
}

type ElasticGPUClaimPhase string

const (
	ClaimPending ElasticGPUClaimPhase = "Pending"
	ClaimBound   ElasticGPUClaimPhase = "Bound"
	ClaimLost    ElasticGPUClaimPhase = "Lost"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ElasticGPUClaimList is a list of ElasticGPUClaim items
type ElasticGPUClaimList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticGPUClaim `json:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:object:root=true

// ElasticGPUClass is non-namespaced
type ElasticGPUClass struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Provisioner       string            `json:"provisioner"`
	Parameters        map[string]string `json:"parameters,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ElasticGPUClassList is a list of ElasticGPUClass items
type ElasticGPUClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticGPUClass `json:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type GPU struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GPUSpec   `json:"spec,omitempty"`
	Status            GPUStatus `json:"status,omitempty"`
}

type GPUSpec struct {
	Index    int    `json:"index"`
	UUID     string `json:"uuid,omitempty"`
	Model    string `json:"model,omitempty"`
	Path     string `json:"path,omitempty"`
	Memory   uint64 `json:"memory,omitempty"`
	NodeName string `json:"nodeName,omitempty"`
}

type GPUStatus struct {
	State       string                  `json:"state,omitempty"`
	Capacity    v1.ResourceList         `json:"capacity,omitempty"`
	Allocatable v1.ResourceList         `json:"allocatable,omitempty"`
	Allocated   map[string]*PodResource `json:"allocated,omitempty"`
}

type PodResource struct {
	Namespace  string              `json:"namespace,omitempty"`
	Pod        string              `json:"pod,omitempty"`
	Containers []ContainerResource `json:"containers,omitempty"`
}

type ContainerResource struct {
	Container string          `json:"container,omitempty"`
	Resource  v1.ResourceList `json:"resource,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GPUList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GPU `json:"items"`
}
