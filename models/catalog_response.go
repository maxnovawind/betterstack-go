package models

import "time"

// CatalogRelationResponse represents a response containing catalog relation data
type CatalogRelationResponse struct {
	Data       CatalogRelation `json:"data,omitempty"`
	Pagination *Pagination     `json:"pagination,omitempty"`
}

// CatalogRelationsResponse represents a response containing a list of catalog relations
type CatalogRelationsResponse struct {
	Data       []CatalogRelation `json:"data,omitempty"`
	Pagination *Pagination       `json:"pagination,omitempty"`
}

// CatalogRelation represents a catalog relation
type CatalogRelation struct {
	ID         string               `json:"id,omitempty"`
	Type       string               `json:"type,omitempty"`
	Attributes CatalogRelationAttrs `json:"attributes,omitempty"`
}

// CatalogRelationAttrs contains catalog relation attributes
type CatalogRelationAttrs struct {
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// CatalogAttributeResponse represents a response containing catalog attribute data
type CatalogAttributeResponse struct {
	Data       CatalogAttribute `json:"data,omitempty"`
	Pagination *Pagination      `json:"pagination,omitempty"`
}

// CatalogAttributesResponse represents a response containing a list of catalog attributes
type CatalogAttributesResponse struct {
	Data       []CatalogAttribute `json:"data,omitempty"`
	Pagination *Pagination        `json:"pagination,omitempty"`
}

// CatalogAttribute represents a catalog attribute
type CatalogAttribute struct {
	ID         string                `json:"id,omitempty"`
	Type       string                `json:"type,omitempty"`
	Attributes CatalogAttributeAttrs `json:"attributes,omitempty"`
}

// CatalogAttributeAttrs contains catalog attribute attributes
type CatalogAttributeAttrs struct {
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// CatalogRecordResponse represents a response containing catalog record data
type CatalogRecordResponse struct {
	Data       CatalogRecord `json:"data,omitempty"`
	Pagination *Pagination   `json:"pagination,omitempty"`
}

// CatalogRecordsResponse represents a response containing a list of catalog records
type CatalogRecordsResponse struct {
	Data       []CatalogRecord `json:"data,omitempty"`
	Pagination *Pagination     `json:"pagination,omitempty"`
}

// CatalogRecord represents a catalog record
type CatalogRecord struct {
	ID         string             `json:"id,omitempty"`
	Type       string             `json:"type,omitempty"`
	Attributes CatalogRecordAttrs `json:"attributes,omitempty"`
}

// CatalogRecordAttrs contains catalog record attributes
type CatalogRecordAttrs struct {
	Attributes []CatalogRecordAttribute `json:"attributes,omitempty"`
	CreatedAt  time.Time                `json:"created_at,omitempty"`
	UpdatedAt  time.Time                `json:"updated_at,omitempty"`
}

// CatalogRecordAttribute represents an attribute in a catalog record
type CatalogRecordAttribute struct {
	Attribute CatalogRecordAttributeRef `json:"attribute,omitempty"`
	Values    []CatalogRecordValue      `json:"values,omitempty"`
}

// CatalogRecordAttributeRef references a catalog attribute by ID or name
type CatalogRecordAttributeRef struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// CatalogRecordValue represents a value for a catalog record attribute
type CatalogRecordValue struct {
	Type   string `json:"type,omitempty"`
	Value  string `json:"value,omitempty"`
	ItemID string `json:"item_id,omitempty"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
}
