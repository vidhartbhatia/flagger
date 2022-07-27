/*
Copyright 2021 The KEDA Authors

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
	autoscalingv2beta2 "k8s.io/api/autoscaling/v2beta2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=scaledobjects,scope=Namespaced,shortName=so
// +kubebuilder:printcolumn:name="ScaleTargetKind",type="string",JSONPath=".status.scaleTargetKind"
// +kubebuilder:printcolumn:name="ScaleTargetName",type="string",JSONPath=".spec.scaleTargetRef.name"
// +kubebuilder:printcolumn:name="Min",type="integer",JSONPath=".spec.minReplicaCount"
// +kubebuilder:printcolumn:name="Max",type="integer",JSONPath=".spec.maxReplicaCount"
// +kubebuilder:printcolumn:name="Triggers",type="string",JSONPath=".spec.triggers[*].type"
// +kubebuilder:printcolumn:name="Authentication",type="string",JSONPath=".spec.triggers[*].authenticationRef.name"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status"
// +kubebuilder:printcolumn:name="Active",type="string",JSONPath=".status.conditions[?(@.type==\"Active\")].status"
// +kubebuilder:printcolumn:name="Fallback",type="string",JSONPath=".status.conditions[?(@.type==\"Fallback\")].status"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ScaledObject is a specification for a ScaledObject resource
type ScaledObject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ScaledObjectSpec `json:"spec"`
	// +optional
	Status ScaledObjectStatus `json:"status,omitempty"`
}

// HealthStatus is the status for a ScaledObject's health
type HealthStatus struct {
	// +optional
	NumberOfFailures *int32 `json:"numberOfFailures,omitempty"`
	// +optional
	Status HealthStatusType `json:"status,omitempty"`
}

// HealthStatusType is an indication of whether the health status is happy or failing
type HealthStatusType string

const (
	// HealthStatusHappy means the status of the health object is happy
	HealthStatusHappy HealthStatusType = "Happy"

	// HealthStatusFailing means the status of the health object is failing
	HealthStatusFailing HealthStatusType = "Failing"
)

// ScaledObjectSpec is the spec for a ScaledObject resource
type ScaledObjectSpec struct {
	ScaleTargetRef *ScaleTarget `json:"scaleTargetRef"`
	// +optional
	PollingInterval *int32 `json:"pollingInterval,omitempty"`
	// +optional
	CooldownPeriod *int32 `json:"cooldownPeriod,omitempty"`
	// +optional
	IdleReplicaCount *int32 `json:"idleReplicaCount,omitempty"`
	// +optional
	MinReplicaCount *int32 `json:"minReplicaCount,omitempty"`
	// +optional
	MaxReplicaCount *int32 `json:"maxReplicaCount,omitempty"`
	// +optional
	Advanced *AdvancedConfig `json:"advanced,omitempty"`

	Triggers []ScaleTriggers `json:"triggers"`
	// +optional
	Fallback *Fallback `json:"fallback,omitempty"`
}

// Fallback is the spec for fallback options
type Fallback struct {
	FailureThreshold int32 `json:"failureThreshold"`
	Replicas         int32 `json:"replicas"`
}

// AdvancedConfig specifies advance scaling options
type AdvancedConfig struct {
	// +optional
	HorizontalPodAutoscalerConfig *HorizontalPodAutoscalerConfig `json:"horizontalPodAutoscalerConfig,omitempty"`
	// +optional
	RestoreToOriginalReplicaCount bool `json:"restoreToOriginalReplicaCount,omitempty"`
}

// HorizontalPodAutoscalerConfig specifies horizontal scale config
type HorizontalPodAutoscalerConfig struct {
	// +optional
	Behavior *autoscalingv2beta2.HorizontalPodAutoscalerBehavior `json:"behavior,omitempty"`
	// +optional
	Name string `json:"name,omitempty"`
}

// ScaleTarget holds the a reference to the scale target Object
type ScaleTarget struct {
	Name string `json:"name"`
	// +optional
	APIVersion string `json:"apiVersion,omitempty"`
	// +optional
	Kind string `json:"kind,omitempty"`
	// +optional
	EnvSourceContainerName string `json:"envSourceContainerName,omitempty"`
}

// ScaleTriggers reference the scaler that will be used
type ScaleTriggers struct {
	Type string `json:"type"`
	// +optional
	Name     string            `json:"name,omitempty"`
	Metadata map[string]string `json:"metadata"`
	// +optional
	AuthenticationRef *ScaledObjectAuthRef `json:"authenticationRef,omitempty"`
	// +optional
	MetricType autoscalingv2beta2.MetricTargetType `json:"metricType,omitempty"`
}

// +k8s:openapi-gen=true

// ScaledObjectStatus is the status for a ScaledObject resource
// +optional
type ScaledObjectStatus struct {
	// +optional
	ScaleTargetKind string `json:"scaleTargetKind,omitempty"`
	// +optional
	ScaleTargetGVKR *GroupVersionKindResource `json:"scaleTargetGVKR,omitempty"`
	// +optional
	OriginalReplicaCount *int32 `json:"originalReplicaCount,omitempty"`
	// +optional
	LastActiveTime *metav1.Time `json:"lastActiveTime,omitempty"`
	// +optional
	ExternalMetricNames []string `json:"externalMetricNames,omitempty"`
	// +optional
	ResourceMetricNames []string `json:"resourceMetricNames,omitempty"`
	// +optional
	Conditions Conditions `json:"conditions,omitempty"`
	// +optional
	Health map[string]HealthStatus `json:"health,omitempty"`
	// +optional
	PausedReplicaCount *int32 `json:"pausedReplicaCount,omitempty"`
	// +optional
	HpaName string `json:"hpaName,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ScaledObjectList is a list of ScaledObject resources
type ScaledObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ScaledObject `json:"items"`
}

// ScaledObjectAuthRef points to the TriggerAuthentication or ClusterTriggerAuthentication object that
// is used to authenticate the scaler with the environment
type ScaledObjectAuthRef struct {
	Name string `json:"name"`
	// Kind of the resource being referred to. Defaults to TriggerAuthentication.
	// +optional
	Kind string `json:"kind,omitempty"`
}

// GroupVersionKindResource provides unified structure for schema.GroupVersionKind and Resource
type GroupVersionKindResource struct {
	Group    string `json:"group"`
	Version  string `json:"version"`
	Kind     string `json:"kind"`
	Resource string `json:"resource"`
}

// ConditionType specifies the available conditions for the resource
type ConditionType string

const (
	// ConditionReady specifies that the resource is ready.
	// For long-running resources.
	ConditionReady ConditionType = "Ready"
	// ConditionActive specifies that the resource has finished.
	// For resource which run to completion.
	ConditionActive ConditionType = "Active"
	// ConditionFallback specifies that the resource has a fallback active.
	ConditionFallback ConditionType = "Fallback"
)

const (
	// ScaledObjectConditionReadySucccesReason defines the default Reason for correct ScaledObject
	ScaledObjectConditionReadySucccesReason = "ScaledObjectReady"
	// ScaledObjectConditionReadySuccessMessage defines the default Message for correct ScaledObject
	ScaledObjectConditionReadySuccessMessage = "ScaledObject is defined correctly and is ready for scaling"
)

// Condition to store the condition state
type Condition struct {
	// Type of condition
	// +required
	Type ConditionType `json:"type" description:"type of status condition"`

	// Status of the condition, one of True, False, Unknown.
	// +required
	Status metav1.ConditionStatus `json:"status" description:"status of the condition, one of True, False, Unknown"`

	// The reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty" description:"one-word CamelCase reason for the condition's last transition"`

	// A human readable message indicating details about the transition.
	// +optional
	Message string `json:"message,omitempty" description:"human-readable message indicating details about last transition"`
}

// Conditions an array representation to store multiple Conditions
type Conditions []Condition

const PausedReplicasAnnotation = "autoscaling.keda.sh/paused-replicas"
