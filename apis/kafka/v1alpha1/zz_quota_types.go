/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type QuotaObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type QuotaParameters struct {

	// A map of string k/v properties.
	// +kubebuilder:validation:Optional
	Config map[string]*float64 `json:"config,omitempty" tf:"config,omitempty"`

	// The name of the entity
	// +kubebuilder:validation:Required
	EntityName *string `json:"entityName" tf:"entity_name,omitempty"`

	// The type of the entity (client-id, user, ip)
	// +kubebuilder:validation:Required
	EntityType *string `json:"entityType" tf:"entity_type,omitempty"`
}

// QuotaSpec defines the desired state of Quota
type QuotaSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QuotaParameters `json:"forProvider"`
}

// QuotaStatus defines the observed state of Quota.
type QuotaStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QuotaObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Quota is the Schema for the Quotas API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,kafkajet}
type Quota struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QuotaSpec   `json:"spec"`
	Status            QuotaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QuotaList contains a list of Quotas
type QuotaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Quota `json:"items"`
}

// Repository type metadata.
var (
	Quota_Kind             = "Quota"
	Quota_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Quota_Kind}.String()
	Quota_KindAPIVersion   = Quota_Kind + "." + CRDGroupVersion.String()
	Quota_GroupVersionKind = CRDGroupVersion.WithKind(Quota_Kind)
)

func init() {
	SchemeBuilder.Register(&Quota{}, &QuotaList{})
}
