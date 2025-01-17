// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type KeyInitParameters struct {

	// (String) The name of the key.
	// The name of the key.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The slug of the organization the key should be created for.
	// The slug of the organization the key should be created for.
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// (String) The slug of the project the key should be created for.
	// The slug of the project the key should be created for.
	// +crossplane:generate:reference:type=github.com/justtrack/provider-sentry/apis/sentry/v1alpha1.Project
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Reference to a Project in sentry to populate project.
	// +kubebuilder:validation:Optional
	ProjectRef *v1.Reference `json:"projectRef,omitempty" tf:"-"`

	// Selector for a Project in sentry to populate project.
	// +kubebuilder:validation:Optional
	ProjectSelector *v1.Selector `json:"projectSelector,omitempty" tf:"-"`

	// (Number) Number of events that can be reported within the rate limit window.
	// Number of events that can be reported within the rate limit window.
	RateLimitCount *float64 `json:"rateLimitCount,omitempty" tf:"rate_limit_count,omitempty"`

	// (Number) Length of time that will be considered when checking the rate limit.
	// Length of time that will be considered when checking the rate limit.
	RateLimitWindow *float64 `json:"rateLimitWindow,omitempty" tf:"rate_limit_window,omitempty"`
}

type KeyObservation struct {

	// (String) DSN for the Content Security Policy (CSP) for the key.
	// DSN for the Content Security Policy (CSP) for the key.
	DsnCsp *string `json:"dsnCsp,omitempty" tf:"dsn_csp,omitempty"`

	// (String) DSN for the key.
	// DSN for the key.
	DsnPublic *string `json:"dsnPublic,omitempty" tf:"dsn_public,omitempty"`

	// (String, Deprecated)
	DsnSecret *string `json:"dsnSecret,omitempty" tf:"dsn_secret,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean) Flag indicating the key is active.
	// Flag indicating the key is active.
	IsActive *bool `json:"isActive,omitempty" tf:"is_active,omitempty"`

	// (String) The name of the key.
	// The name of the key.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The slug of the organization the key should be created for.
	// The slug of the organization the key should be created for.
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// (String) The slug of the project the key should be created for.
	// The slug of the project the key should be created for.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// (Number) The ID of the project that the key belongs to.
	// The ID of the project that the key belongs to.
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (String) Public key portion of the client key.
	// Public key portion of the client key.
	Public *string `json:"public,omitempty" tf:"public,omitempty"`

	// (Number) Number of events that can be reported within the rate limit window.
	// Number of events that can be reported within the rate limit window.
	RateLimitCount *float64 `json:"rateLimitCount,omitempty" tf:"rate_limit_count,omitempty"`

	// (Number) Length of time that will be considered when checking the rate limit.
	// Length of time that will be considered when checking the rate limit.
	RateLimitWindow *float64 `json:"rateLimitWindow,omitempty" tf:"rate_limit_window,omitempty"`

	// (String) Secret key portion of the client key.
	// Secret key portion of the client key.
	Secret *string `json:"secret,omitempty" tf:"secret,omitempty"`
}

type KeyParameters struct {

	// (String) The name of the key.
	// The name of the key.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The slug of the organization the key should be created for.
	// The slug of the organization the key should be created for.
	// +kubebuilder:validation:Optional
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// (String) The slug of the project the key should be created for.
	// The slug of the project the key should be created for.
	// +crossplane:generate:reference:type=github.com/justtrack/provider-sentry/apis/sentry/v1alpha1.Project
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Reference to a Project in sentry to populate project.
	// +kubebuilder:validation:Optional
	ProjectRef *v1.Reference `json:"projectRef,omitempty" tf:"-"`

	// Selector for a Project in sentry to populate project.
	// +kubebuilder:validation:Optional
	ProjectSelector *v1.Selector `json:"projectSelector,omitempty" tf:"-"`

	// (Number) Number of events that can be reported within the rate limit window.
	// Number of events that can be reported within the rate limit window.
	// +kubebuilder:validation:Optional
	RateLimitCount *float64 `json:"rateLimitCount,omitempty" tf:"rate_limit_count,omitempty"`

	// (Number) Length of time that will be considered when checking the rate limit.
	// Length of time that will be considered when checking the rate limit.
	// +kubebuilder:validation:Optional
	RateLimitWindow *float64 `json:"rateLimitWindow,omitempty" tf:"rate_limit_window,omitempty"`
}

// KeySpec defines the desired state of Key
type KeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KeyParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider KeyInitParameters `json:"initProvider,omitempty"`
}

// KeyStatus defines the observed state of Key.
type KeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Key is the Schema for the Keys API. Sentry Key resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,sentry}
type Key struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.organization) || (has(self.initProvider) && has(self.initProvider.organization))",message="spec.forProvider.organization is a required parameter"
	Spec   KeySpec   `json:"spec"`
	Status KeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeyList contains a list of Keys
type KeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Key `json:"items"`
}

// Repository type metadata.
var (
	Key_Kind             = "Key"
	Key_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Key_Kind}.String()
	Key_KindAPIVersion   = Key_Kind + "." + CRDGroupVersion.String()
	Key_GroupVersionKind = CRDGroupVersion.WithKind(Key_Kind)
)

func init() {
	SchemeBuilder.Register(&Key{}, &KeyList{})
}
