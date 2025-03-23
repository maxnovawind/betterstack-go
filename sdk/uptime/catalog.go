package uptime

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sudodeo/betterstack-go/models"
)

// Catalog Relations

// ListCatalogRelations lists all catalog relations
func (bs *Betterstack) ListCatalogRelations(ctx context.Context, page, perPage int) (*models.CatalogRelationsResponse, error) {
	endpoint := "/api/v2/catalog/relations"

	// Add query parameters if provided
	if page > 0 || perPage > 0 {
		endpoint += "?"
		if page > 0 {
			endpoint += fmt.Sprintf("page=%d", page)
		}
		if perPage > 0 {
			if page > 0 {
				endpoint += "&"
			}
			endpoint += fmt.Sprintf("per_page=%d", perPage)
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	response := &models.CatalogRelationsResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetCatalogRelation gets a specific catalog relation by ID
func (bs *Betterstack) GetCatalogRelation(ctx context.Context, relationID string) (*models.CatalogRelationResponse, error) {
	endpoint := fmt.Sprintf("/api/v2/catalog/relations/%s", relationID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	response := &models.CatalogRelationResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// CreateCatalogRelation creates a new catalog relation
func (bs *Betterstack) CreateCatalogRelation(ctx context.Context, params *models.CreateCatalogRelationRequest) (*models.CatalogRelationResponse, error) {
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/api/v2/catalog/relations", bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	response := &models.CatalogRelationResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateCatalogRelation updates an existing catalog relation
func (bs *Betterstack) UpdateCatalogRelation(ctx context.Context, relationID string, params *models.UpdateCatalogRelationRequest) (*models.CatalogRelationResponse, error) {
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("/api/v2/catalog/relations/%s", relationID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, endpoint, bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	response := &models.CatalogRelationResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// DeleteCatalogRelation deletes a catalog relation
func (bs *Betterstack) DeleteCatalogRelation(ctx context.Context, relationID string) error {
	endpoint := fmt.Sprintf("/api/v2/catalog/relations/%s", relationID)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}

	_, err = bs.MakeAPIRequest(req)
	if err != nil {
		return err
	}

	return nil
}

// Catalog Attributes

// ListCatalogAttributes lists all attributes for a catalog relation
func (bs *Betterstack) ListCatalogAttributes(ctx context.Context, relationID string, page, perPage int) (*models.CatalogAttributesResponse, error) {
	endpoint := fmt.Sprintf("/api/v2/catalog/relations/%s/attributes", relationID)

	// Add query parameters if provided
	if page > 0 || perPage > 0 {
		endpoint += "?"
		if page > 0 {
			endpoint += fmt.Sprintf("page=%d", page)
		}
		if perPage > 0 {
			if page > 0 {
				endpoint += "&"
			}
			endpoint += fmt.Sprintf("per_page=%d", perPage)
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	response := &models.CatalogAttributesResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetCatalogAttribute gets a specific catalog attribute
func (bs *Betterstack) GetCatalogAttribute(ctx context.Context, relationID, attributeID string) (*models.CatalogAttributeResponse, error) {
	endpoint := fmt.Sprintf("/api/v2/catalog/relations/%s/attributes/%s", relationID, attributeID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	response := &models.CatalogAttributeResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// CreateCatalogAttribute creates a new catalog attribute
func (bs *Betterstack) CreateCatalogAttribute(ctx context.Context, relationID string, params *models.CreateCatalogAttributeRequest) (*models.CatalogAttributeResponse, error) {
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("/api/v2/catalog/relations/%s/attributes", relationID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	response := &models.CatalogAttributeResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateCatalogAttribute updates an existing catalog attribute
func (bs *Betterstack) UpdateCatalogAttribute(ctx context.Context, relationID, attributeID string, params *models.UpdateCatalogAttributeRequest) (*models.CatalogAttributeResponse, error) {
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("/api/v2/catalog/relations/%s/attributes/%s", relationID, attributeID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, endpoint, bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	response := &models.CatalogAttributeResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// DeleteCatalogAttribute deletes a catalog attribute
func (bs *Betterstack) DeleteCatalogAttribute(ctx context.Context, relationID, attributeID string) error {
	endpoint := fmt.Sprintf("/api/v2/catalog/relations/%s/attributes/%s", relationID, attributeID)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}

	_, err = bs.MakeAPIRequest(req)
	if err != nil {
		return err
	}

	return nil
}

// Catalog Records

// ListCatalogRecords lists all records for a catalog relation
func (bs *Betterstack) ListCatalogRecords(ctx context.Context, relationID string, page, perPage int) (*models.CatalogRecordsResponse, error) {
	endpoint := fmt.Sprintf("/api/v2/catalog/relations/%s/records", relationID)

	// Add query parameters if provided
	if page > 0 || perPage > 0 {
		endpoint += "?"
		if page > 0 {
			endpoint += fmt.Sprintf("page=%d", page)
		}
		if perPage > 0 {
			if page > 0 {
				endpoint += "&"
			}
			endpoint += fmt.Sprintf("per_page=%d", perPage)
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	response := &models.CatalogRecordsResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetCatalogRecord gets a specific catalog record
func (bs *Betterstack) GetCatalogRecord(ctx context.Context, relationID, recordID string) (*models.CatalogRecordResponse, error) {
	endpoint := fmt.Sprintf("/api/v2/catalog/relations/%s/records/%s", relationID, recordID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	response := &models.CatalogRecordResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// CreateCatalogRecord creates a new catalog record
func (bs *Betterstack) CreateCatalogRecord(ctx context.Context, relationID string, params *models.CreateCatalogRecordRequest) (*models.CatalogRecordResponse, error) {
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("/api/v2/catalog/relations/%s/records", relationID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	response := &models.CatalogRecordResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// BatchCreateCatalogRecords batch creates catalog records
func (bs *Betterstack) BatchCreateCatalogRecords(ctx context.Context, relationID string, params []models.CreateCatalogRecordRequest) (*models.CatalogRecordsResponse, error) {
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("/api/v2/catalog/relations/%s/records/batch", relationID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	response := &models.CatalogRecordsResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateCatalogRecord updates an existing catalog record
func (bs *Betterstack) UpdateCatalogRecord(ctx context.Context, relationID, recordID string, params *models.UpdateCatalogRecordRequest) (*models.CatalogRecordResponse, error) {
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("/api/v2/catalog/relations/%s/records/%s", relationID, recordID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, endpoint, bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	response := &models.CatalogRecordResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// DeleteCatalogRecord deletes a catalog record
func (bs *Betterstack) DeleteCatalogRecord(ctx context.Context, relationID, recordID string) error {
	endpoint := fmt.Sprintf("/api/v2/catalog/relations/%s/records/%s", relationID, recordID)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}

	_, err = bs.MakeAPIRequest(req)
	if err != nil {
		return err
	}

	return nil
}
