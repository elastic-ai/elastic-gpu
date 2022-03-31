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
	Capacity         v1.ResourceList `json:"capacity,omitempty" protobuf:"bytes,1,rep,name=capacity,casttype=ResourceList,castkey=ResourceName"`
	ElasticGPUSource `json:",inline" protobuf:"bytes,2,opt,name=elasticGPUSource"`
	ClaimRef         v1.ObjectReference `json:"claimRef,omitempty" protobuf:"bytes,3,opt,name=claimRef"`
	NodeAffinity     GPUNodeAffinity    `json:"nodeAffinity,omitempty" protobuf:"bytes,4,opt,name=nodeAffinity"`
	NodeName         string             `json:"nodeName,omitempty" protobuf:"bytes,5,opt,name=nodeName"`
}

type GPUNodeAffinity struct {
	Required *v1.NodeSelector `json:"required,omitempty" protobuf:"bytes,1,opt,name=required"`
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
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              ElasticGPUSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            ElasticGPUStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

//+kubebuilder:object:root=true

// ElasticGPUList contains a list of ElasticGPU
type ElasticGPUList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []ElasticGPU `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type ElasticGPUSource struct {
	QGPU        *QGPUElasticGPUSource        `json:"qGPU,omitempty" protobuf:"bytes,1,opt,name=qGPU"`
	PhysicalGPU *PhysicalGPUElasticGPUSource `json:"physicalGPU,omitempty" protobuf:"bytes,2,opt,name=physicalGPU"`
	GPUShare    *GPUShareElasticGPUSource    `json:"gpuShare,omitempty" protobuf:"bytes,3,opt,name=gpuShare"`
}

type BaseGPUSource struct {
	Index string `json:"index" protobuf:"bytes,1,opt,name=index"`
	UUID  string `json:"uuid,omitempty" protobuf:"bytes,2,opt,name=uuid"`
}

type QGPUElasticGPUSource struct {
	BaseGPUSource `json:",inline" protobuf:"bytes,1,opt,name=baseGPUSource"`
	DeviceName    string   `json:"DeviceName,omitempty" protobuf:"bytes,2,opt,name=deviceName"`
	Paths         []string `json:"paths,omitempty" protobuf:"bytes,3,rep,name=paths"`
}

type PhysicalGPUElasticGPUSource struct {
	BaseGPUSource `json:",inline" protobuf:"bytes,1,opt,name=baseGPUSource"`
}

type GPUShareElasticGPUSource struct {
	BaseGPUSource `json:",inline" protobuf:"bytes,1,opt,name=baseGPUSource"`
}

const (
	ResourceGPUCore         v1.ResourceName = "elasticgpu.com/gpu-core"
	ResourceGPUMemory       v1.ResourceName = "elasticgpu.com/gpu-memory"
	ResourceQGPUCore        v1.ResourceName = "tke.cloud.tencent.com/qgpu-core"
	ResourceQGPUOfflineCore v1.ResourceName = "tke.cloud.tencent.com/qgpu-core-greedy"
	ResourceQGPUMemory      v1.ResourceName = "tke.cloud.tencent.com/qgpu-memory"
	ResourcePGPU            v1.ResourceName = "nvidia.com/gpu"
)

func init() {
	SchemeBuilder.Register(&ElasticGPU{}, &ElasticGPUList{})
}
