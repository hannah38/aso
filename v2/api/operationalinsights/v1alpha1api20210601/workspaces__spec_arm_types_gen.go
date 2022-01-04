// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type Workspaces_SpecARM struct {
	//ETag: The ETag of the workspace.
	ETag *string `json:"eTag,omitempty"`

	//Location: The geo-location where the resource lives
	Location string `json:"location,omitempty"`

	//Name: The name of the workspace.
	Name string `json:"name"`

	//Properties: Workspace properties.
	Properties WorkspacePropertiesARM `json:"properties"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Workspaces_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (workspacesSpecARM Workspaces_SpecARM) GetAPIVersion() string {
	return "2021-06-01"
}

// GetName returns the Name of the resource
func (workspacesSpecARM Workspaces_SpecARM) GetName() string {
	return workspacesSpecARM.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.OperationalInsights/workspaces"
func (workspacesSpecARM Workspaces_SpecARM) GetType() string {
	return "Microsoft.OperationalInsights/workspaces"
}

//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.OperationalInsights.json#/definitions/WorkspaceProperties
type WorkspacePropertiesARM struct {
	//Features: Workspace features.
	Features *WorkspaceFeaturesARM `json:"features,omitempty"`

	//ForceCmkForQuery: Indicates whether customer managed storage is mandatory for
	//query management.
	ForceCmkForQuery *bool `json:"forceCmkForQuery,omitempty"`

	//ProvisioningState: The provisioning state of the workspace.
	ProvisioningState *WorkspacePropertiesProvisioningState `json:"provisioningState,omitempty"`

	//PublicNetworkAccessForIngestion: The network access type for accessing Log
	//Analytics ingestion.
	PublicNetworkAccessForIngestion *WorkspacePropertiesPublicNetworkAccessForIngestion `json:"publicNetworkAccessForIngestion,omitempty"`

	//PublicNetworkAccessForQuery: The network access type for accessing Log Analytics
	//query.
	PublicNetworkAccessForQuery *WorkspacePropertiesPublicNetworkAccessForQuery `json:"publicNetworkAccessForQuery,omitempty"`

	//RetentionInDays: The workspace data retention in days. Allowed values are per
	//pricing plan. See pricing tiers documentation for details.
	RetentionInDays *int `json:"retentionInDays,omitempty"`

	//Sku: The SKU (tier) of a workspace.
	Sku *WorkspaceSkuARM `json:"sku,omitempty"`

	//WorkspaceCapping: The daily volume cap for ingestion.
	WorkspaceCapping *WorkspaceCappingARM `json:"workspaceCapping,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.OperationalInsights.json#/definitions/WorkspaceCapping
type WorkspaceCappingARM struct {
	//DailyQuotaGb: The workspace daily quota for ingestion.
	DailyQuotaGb *float64 `json:"dailyQuotaGb,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.OperationalInsights.json#/definitions/WorkspaceFeatures
type WorkspaceFeaturesARM struct {
	//AdditionalProperties: Unmatched properties from the message are deserialized
	//this collection
	AdditionalProperties map[string]v1.JSON `json:"additionalProperties,omitempty"`
	ClusterResourceId    *string            `json:"clusterResourceId,omitempty"`

	//DisableLocalAuth: Disable Non-AAD based Auth.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`

	//EnableDataExport: Flag that indicate if data should be exported.
	EnableDataExport *bool `json:"enableDataExport,omitempty"`

	//EnableLogAccessUsingOnlyResourcePermissions: Flag that indicate which permission
	//to use - resource or workspace or both.
	EnableLogAccessUsingOnlyResourcePermissions *bool `json:"enableLogAccessUsingOnlyResourcePermissions,omitempty"`

	//ImmediatePurgeDataOn30Days: Flag that describes if we want to remove the data
	//after 30 days.
	ImmediatePurgeDataOn30Days *bool `json:"immediatePurgeDataOn30Days,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.OperationalInsights.json#/definitions/WorkspaceSku
type WorkspaceSkuARM struct {
	//CapacityReservationLevel: The capacity reservation level in GB for this
	//workspace, when CapacityReservation sku is selected.
	CapacityReservationLevel *int `json:"capacityReservationLevel,omitempty"`

	//Name: The name of the SKU.
	Name WorkspaceSkuName `json:"name"`
}