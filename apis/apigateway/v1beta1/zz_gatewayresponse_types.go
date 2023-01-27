/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type GatewayResponseObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type GatewayResponseParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Map of parameters (paths, query strings and headers) of the Gateway Response.
	// +kubebuilder:validation:Optional
	ResponseParameters map[string]*string `json:"responseParameters,omitempty" tf:"response_parameters,omitempty"`

	// Map of templates used to transform the response body.
	// +kubebuilder:validation:Optional
	ResponseTemplates map[string]*string `json:"responseTemplates,omitempty" tf:"response_templates,omitempty"`

	// Response type of the associated GatewayResponse.
	// +kubebuilder:validation:Required
	ResponseType *string `json:"responseType" tf:"response_type,omitempty"`

	// String identifier of the associated REST API.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.RestAPI
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`

	// Reference to a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDRef *v1.Reference `json:"restApiIdRef,omitempty" tf:"-"`

	// Selector for a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDSelector *v1.Selector `json:"restApiIdSelector,omitempty" tf:"-"`

	// HTTP status code of the Gateway Response.
	// +kubebuilder:validation:Optional
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

// GatewayResponseSpec defines the desired state of GatewayResponse
type GatewayResponseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayResponseParameters `json:"forProvider"`
}

// GatewayResponseStatus defines the observed state of GatewayResponse.
type GatewayResponseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayResponseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayResponse is the Schema for the GatewayResponses API. Provides an API Gateway Gateway Response for a REST API Gateway.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type GatewayResponse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewayResponseSpec   `json:"spec"`
	Status            GatewayResponseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayResponseList contains a list of GatewayResponses
type GatewayResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GatewayResponse `json:"items"`
}

// Repository type metadata.
var (
	GatewayResponse_Kind             = "GatewayResponse"
	GatewayResponse_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GatewayResponse_Kind}.String()
	GatewayResponse_KindAPIVersion   = GatewayResponse_Kind + "." + CRDGroupVersion.String()
	GatewayResponse_GroupVersionKind = CRDGroupVersion.WithKind(GatewayResponse_Kind)
)

func init() {
	SchemeBuilder.Register(&GatewayResponse{}, &GatewayResponseList{})
}
