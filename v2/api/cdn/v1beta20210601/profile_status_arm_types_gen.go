// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210601

type Profile_STATUSARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Kind: Kind of the profile. Used by portal to differentiate traditional CDN profile and new AFD profile.
	Kind *string `json:"kind,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: Resource name.
	Name       *string                      `json:"name,omitempty"`
	Properties *ProfileProperties_STATUSARM `json:"properties,omitempty"`

	// Sku: The pricing tier (defines Azure Front Door Standard or Premium or a CDN provider, feature list and rate) of the
	// profile.
	Sku        *Sku_STATUSARM        `json:"sku,omitempty"`
	SystemData *SystemData_STATUSARM `json:"systemData,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

type ProfileProperties_STATUSARM struct {
	// FrontDoorId: The Id of the frontdoor.
	FrontDoorId *string `json:"frontDoorId,omitempty"`

	// OriginResponseTimeoutSeconds: Send and receive timeout on forwarding request to the origin. When timeout is reached, the
	// request fails and returns.
	OriginResponseTimeoutSeconds *int `json:"originResponseTimeoutSeconds,omitempty"`

	// ProvisioningState: Provisioning status of the profile.
	ProvisioningState *ProfileProperties_STATUS_ProvisioningState `json:"provisioningState,omitempty"`

	// ResourceState: Resource status of the profile.
	ResourceState *ProfileProperties_STATUS_ResourceState `json:"resourceState,omitempty"`
}

type Sku_STATUSARM struct {
	// Name: Name of the pricing tier.
	Name *Sku_STATUS_Name `json:"name,omitempty"`
}

type Sku_STATUS_Name string

const (
	Sku_STATUS_Name_Custom_Verizon                     = Sku_STATUS_Name("Custom_Verizon")
	Sku_STATUS_Name_Premium_AzureFrontDoor             = Sku_STATUS_Name("Premium_AzureFrontDoor")
	Sku_STATUS_Name_Premium_Verizon                    = Sku_STATUS_Name("Premium_Verizon")
	Sku_STATUS_Name_StandardPlus_955BandWidth_ChinaCdn = Sku_STATUS_Name("StandardPlus_955BandWidth_ChinaCdn")
	Sku_STATUS_Name_StandardPlus_AvgBandWidth_ChinaCdn = Sku_STATUS_Name("StandardPlus_AvgBandWidth_ChinaCdn")
	Sku_STATUS_Name_StandardPlus_ChinaCdn              = Sku_STATUS_Name("StandardPlus_ChinaCdn")
	Sku_STATUS_Name_Standard_955BandWidth_ChinaCdn     = Sku_STATUS_Name("Standard_955BandWidth_ChinaCdn")
	Sku_STATUS_Name_Standard_Akamai                    = Sku_STATUS_Name("Standard_Akamai")
	Sku_STATUS_Name_Standard_AvgBandWidth_ChinaCdn     = Sku_STATUS_Name("Standard_AvgBandWidth_ChinaCdn")
	Sku_STATUS_Name_Standard_AzureFrontDoor            = Sku_STATUS_Name("Standard_AzureFrontDoor")
	Sku_STATUS_Name_Standard_ChinaCdn                  = Sku_STATUS_Name("Standard_ChinaCdn")
	Sku_STATUS_Name_Standard_Microsoft                 = Sku_STATUS_Name("Standard_Microsoft")
	Sku_STATUS_Name_Standard_Verizon                   = Sku_STATUS_Name("Standard_Verizon")
)