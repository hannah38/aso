// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210501

type FlexibleServers_FirewallRule_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: The properties of a firewall rule.
	Properties *FirewallRuleProperties_STATUS_ARM `json:"properties,omitempty"`

	// SystemData: The system metadata relating to this resource.
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// The properties of a server firewall rule.
type FirewallRuleProperties_STATUS_ARM struct {
	// EndIpAddress: The end IP address of the server firewall rule. Must be IPv4 format.
	EndIpAddress *string `json:"endIpAddress,omitempty"`

	// StartIpAddress: The start IP address of the server firewall rule. Must be IPv4 format.
	StartIpAddress *string `json:"startIpAddress,omitempty"`
}
