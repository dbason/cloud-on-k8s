// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package v1beta1

import (
	commonv1 "github.com/elastic/cloud-on-k8s/pkg/apis/common/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const EnterpriseSearchContainerName = "enterprise-search"

// EnterpriseSearchSpec holds the specification of an Enterprise Search resource.
type EnterpriseSearchSpec struct {
	// Version of Enterprise Search.
	Version string `json:"version,omitempty"`

	// Image is the Enterprise Search Docker image to deploy.
	Image string `json:"image,omitempty"`

	// Count of Enterprise Search instances to deploy.
	Count int32 `json:"count,omitempty"`

	// Config holds the Enterprise Search configuration.
	Config *commonv1.Config `json:"config,omitempty"`

	// HTTP holds the HTTP layer configuration for Enterprise Search resource.
	HTTP commonv1.HTTPConfig `json:"http,omitempty"`

	// ElasticsearchRef is a reference to the output Elasticsearch cluster running in the same Kubernetes cluster.
	ElasticsearchRef commonv1.ObjectSelector `json:"elasticsearchRef,omitempty"`

	// PodTemplate provides customisation options (labels, annotations, affinity rules, resource requests, and so on)
	// for the Enterprise Search pods.
	// +kubebuilder:validation:Optional
	PodTemplate corev1.PodTemplateSpec `json:"podTemplate,omitempty"`

	// ServiceAccountName is used to check access from the current resource to a resource (eg. Elasticsearch) in a different namespace.
	// Can only be used if ECK is enforcing RBAC on references.
	// +optional
	ServiceAccountName string `json:"serviceAccountName,omitempty"`
}


func (ents *EnterpriseSearch) ServiceAccountName() string {
	return ents.Spec.ServiceAccountName
}
//
//// En expresses the status of the Apm Server instances.
//type ApmServerHealth string
//
//const (
//	// ApmServerRed means no instance is currently available.
//	ApmServerRed ApmServerHealth = "red"
//	// ApmServerGreen means at least one instance is available.
//	ApmServerGreen ApmServerHealth = "green"
//)

// EnterpriseSearchStatus defines the observed state of EnterpriseSearch
type EnterpriseSearchStatus struct {
	commonv1.ReconcilerStatus `json:",inline"`
	//Health                         ApmServerHealth `json:"health,omitempty"`
	// ExternalService is the name of the service associated to the Enterprise Search Pods.
	ExternalService string `json:"service,omitempty"`
	// Association is the status of any auto-linking to Elasticsearch clusters.
	Association commonv1.AssociationStatus `json:"associationStatus,omitempty"`
}
//
//// IsDegraded returns true if the current status is worse than the previous.
//func (as ApmServerStatus) IsDegraded(prev ApmServerStatus) bool {
//	return prev.Health == ApmServerGreen && as.Health != ApmServerGreen
//}

// +kubebuilder:object:root=true

// EnterpriseSearch represents an Enterprise Search resource in a Kubernetes cluster.
// +kubebuilder:resource:categories=elastic,shortName=entsearch
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="health",type="string",JSONPath=".status.health"
// +kubebuilder:printcolumn:name="nodes",type="integer",JSONPath=".status.availableNodes",description="Available nodes"
// +kubebuilder:printcolumn:name="version",type="string",JSONPath=".spec.version",description="Enterprise Search version"
// +kubebuilder:printcolumn:name="age",type="date",JSONPath=".metadata.creationTimestamp"
type EnterpriseSearch struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec      EnterpriseSearchSpec                  `json:"spec,omitempty"`
	Status    EnterpriseSearchStatus                `json:"status,omitempty"`
	assocConf *commonv1.AssociationConf `json:"-"` //nolint:govet
}

// +kubebuilder:object:root=true

// EnterpriseSearchList contains a list of EnterpriseSearch
type EnterpriseSearchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EnterpriseSearch `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EnterpriseSearch{}, &EnterpriseSearchList{})
}

// IsMarkedForDeletion returns true if the EnterpriseSearch is going to be deleted
func (ents *EnterpriseSearch) IsMarkedForDeletion() bool {
	return !ents.DeletionTimestamp.IsZero()
}

func (ents *EnterpriseSearch) ElasticsearchRef() commonv1.ObjectSelector {
	return ents.Spec.ElasticsearchRef
}


func (ents *EnterpriseSearch) AssociationConf() *commonv1.AssociationConf {
	return ents.assocConf
}

func (ents *EnterpriseSearch) SetAssociationConf(assocConf *commonv1.AssociationConf) {
	ents.assocConf = assocConf
}