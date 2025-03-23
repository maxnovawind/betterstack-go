package models

// CreateCatalogRelationRequest represents a request to create a catalog relation
type CreateCatalogRelationRequest struct {
	Name string `json:"name"`
}

// UpdateCatalogRelationRequest represents a request to update a catalog relation
type UpdateCatalogRelationRequest struct {
	Name string `json:"name,omitempty"`
}

// CreateCatalogAttributeRequest represents a request to create a catalog attribute
type CreateCatalogAttributeRequest struct {
	Name string `json:"name"`
	// Additional attribute-specific fields could be added here
}

// UpdateCatalogAttributeRequest represents a request to update a catalog attribute
type UpdateCatalogAttributeRequest struct {
	Name string `json:"name,omitempty"`
	// Additional attribute-specific fields could be added here
}

// CreateCatalogRecordRequest represents a request to create a catalog record
type CreateCatalogRecordRequest struct {
	Attributes []CatalogRecordAttributeRequest `json:"attributes"`
}

// UpdateCatalogRecordRequest represents a request to update a catalog record
type UpdateCatalogRecordRequest struct {
	Attributes []CatalogRecordAttributeRequest `json:"attributes,omitempty"`
}

// CatalogRecordAttributeRequest represents an attribute in a catalog record creation/update request
type CatalogRecordAttributeRequest struct {
	Attribute CatalogRecordAttributeRef   `json:"attribute"`
	Values    []CatalogRecordValueRequest `json:"values"`
}

// CatalogRecordValueRequest represents a value for a catalog record attribute in a request
type CatalogRecordValueRequest struct {
	// Type of the value. Valid options:
	// Scalar types: String
	// Reference types: User, Team, Policy, Schedule, SlackIntegration, LinearIntegration,
	// JiraIntegration, MicrosoftTeamsWebhook, ZapierWebhook, NativeWebhook, PagerDutyWebhook
	Type string `json:"type"`

	// Value is used when Type is String
	Value string `json:"value,omitempty"`

	// ItemID is the ID of a referenced item (for reference types)
	ItemID string `json:"item_id,omitempty"`

	// Name is the name of a referenced item (for reference types)
	Name string `json:"name,omitempty"`

	// Email is the email of a referenced user (for User type)
	Email string `json:"email,omitempty"`
}

// BatchCreateCatalogRecordsRequest represents a request to batch create catalog records
type BatchCreateCatalogRecordsRequest struct {
	Records []CreateCatalogRecordRequest `json:"records"`
}
