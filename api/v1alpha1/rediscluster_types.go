/*
Copyright 2026.

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
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Common Kubernetes configuration for all components of the Redis cluster.
// +k8s:deepcopy-gen=true
type KubernetesConfig struct {
	Image                              string                               `json:"image,omitempty"`
	ImagePullPolicy                    corev1.PullPolicy                    `json:"imagePullPolicy,omitempty"`
	Resources                          corev1.ResourceRequirements          `json:"resources,omitempty"`
	ImagePullSecrets                   []corev1.LocalObjectReference        `json:"imagePullSecrets,omitempty"`
	UpdateStrategy                     appsv1.StatefulSetUpdateStrategy     `json:"updateStrategy,omitempty"`
	PersistentVolumeClaimReclaimPolicy corev1.PersistentVolumeReclaimPolicy `json:"persistentVolumeClaimReclaimPolicy,omitempty"`
	Service                            *ServiceConfig                       `json:"service,omitempty"`
	IgnoreAnnotations                  []string                             `json:"ignoreAnnotations,omitempty"`
	MinReadySeconds                    int32                                `json:"minReadySeconds,omitempty"`
}

// ServiceConfig define the type of service to be created and its annotations
// +k8s:deepcopy-gen=true
type ServiceConfig struct {
	// +kubebuilder:validation:Enum=LoadBalancer;NodePort;ClusterIP
	ServiceType        string            `json:"serviceType,omitempty"`
	ServiceAnnotations map[string]string `json:"annotations,omitempty"`
	// IncludeBusPort when set to true, it will add bus port to the service, such as 16379.
	// This field is only used for Redis cluster mode.
	IncludeBusPort *bool `json:"includeBusPort,omitempty"`
	// Headless config for which suffix is -headless service
	Headless *Service `json:"headless,omitempty"`
	// Additional config for which suffix is -additional service
	Additional *Service `json:"additional,omitempty"`
}

// Service is the struct to define the service type and its annotations
// +k8s:deepcopy-gen=true
type Service struct {
	// +kubebuilder:validation:Enum=LoadBalancer;NodePort;ClusterIP
	// +kubebuilder:default:=ClusterIP
	Type                  string            `json:"type,omitempty"`
	AdditionalAnnotations map[string]string `json:"additionalAnnotations,omitempty"`
	// IncludeBusPort when set to true, it will add bus port to the service, such as 16379.
	// This field is only used for Redis cluster mode.
	IncludeBusPort *bool `json:"includeBusPort,omitempty"`
	// +kubebuilder:default:=true
	Enabled *bool `json:"enabled,omitempty"`
}

type RedisConfig struct {
	MaxMemoryPercentOfLimit *int32   `json:"maxMemoryPercentOfLimit,omitempty"`
	DynamicConfig           []string `json:"dynamicConfig,omitempty"`
	AdditionalRedisConfig   *string  `json:"additionalRedisConfig,omitempty"`
}

type RedisLeader struct {
	Replicas     *int32              `json:"replicas,omitempty"`
	RedisConfig  *RedisConfig        `json:"redisConfig,omitempty"`
	Affinity     *corev1.Affinity    `json:"affinity,omitempty"`
	NodeSelector map[string]string   `json:"nodeSelector,omitempty"`
	Tolerations  []corev1.Toleration `json:"tolerations,omitempty"`
}

// RedisClusterSpec defines the desired state of RedisCluster.
type RedisClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of RedisCluster. Edit rediscluster_types.go to remove/update
	ClusterSize      *int32            `json:"clusterSize,omitempty"`
	KubernetesConfig *KubernetesConfig `json:"kubernetesConfig,omitempty"`
	HostNetwork      *bool             `json:"hostNetwork,omitempty"`
	Port             *int32            `json:"port,omitempty"`
	ClusterVersion   *string           `json:"clusterVersion,omitempty"`
	RedisLeader
}

// RedisClusterStatus defines the observed state of RedisCluster.
type RedisClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// RedisCluster is the Schema for the redisclusters API.
type RedisCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedisClusterSpec   `json:"spec,omitempty"`
	Status RedisClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedisClusterList contains a list of RedisCluster.
type RedisClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RedisCluster{}, &RedisClusterList{})
}
