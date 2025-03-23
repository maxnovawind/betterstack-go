package models

// Metadata represents the metadata information in the response.
type Metadata struct {
	Data       []MetadataRecord `json:"data,omitempty"`       // Data contains the metadata records.
	Pagination Pagination       `json:"pagination,omitempty"` // Pagination contains the pagination information.
}

// MetadataRecord represents a metadata record in the response.
type MetadataRecord struct {
	ID            string             `json:"id,omitempty"`
	Type          string             `json:"type,omitempty"`
	Attributes    MetadataAttributes `json:"attributes,omitempty"`
	Relationships Relationships      `json:"relationships,omitempty"`
}

type MetadataAttributes struct {
	Key       string          `json:"key,omitempty"`
	Values    []MetadataValue `json:"values,omitempty"`
	TeamName  string          `json:"team_name,omitempty"`
	OwnerID   string          `json:"owner_id,omitempty"`
	OwnerType string          `json:"owner_type,omitempty"`
}

// MetadataValue represents a value in the metadata attributes
type MetadataValue struct {
	Type   string `json:"type,omitempty"`
	Value  string `json:"value,omitempty"`   // Used for String type values
	ItemID int    `json:"item_id,omitempty"` // Used for Schedule type values
	Name   string `json:"name,omitempty"`    // Used for Schedule type values
}
