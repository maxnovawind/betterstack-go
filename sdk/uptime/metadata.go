package uptime

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/sudodeo/betterstack-go/models"
)

// ListMetadata returns list of all metadata
func (bs *Betterstack) ListMetadata() (*models.Metadata, error) {
	// TODO: query params
	req, err := http.NewRequest(http.MethodGet, "/api/v3/metadata", nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	data := &models.Metadata{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetMetadataRecord returns a single metadata record.
func (bs *Betterstack) GetMetadataRecord(metadataID string) (*models.MetadataRecord, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v3/metadata/"+metadataID, nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data models.MetadataRecord `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// CreateMetadataRecord creates a new metadata record or returns validation errors.
func (bs *Betterstack) CreateMetadataRecord(bodyParams models.MetadataRecordReqBody) (*models.MetadataRecord, error) {
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "/api/v3/metadata", bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data models.MetadataRecord `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// UpdateMetadataRecord updates an existing metadata record configuration. Send only the parameters you wish to change (e.g., URL).
func (bs *Betterstack) UpdateMetadataRecord(metadataID string, bodyParams models.MetadataRecordReqBody) (*models.MetadataRecord, error) {
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, "/api/v3/metadata/"+metadataID, bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data models.MetadataRecord `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// DeleteMetadataRecord permanently deletes an existing metadata record.
func (bs *Betterstack) DeleteMetadataRecord(metadataID string) error {
	req, err := http.NewRequest(http.MethodDelete, "/api/v3/metadata/"+metadataID, nil)
	if err != nil {
		return err
	}

	_, err = bs.MakeAPIRequest(req)
	if err != nil {
		return err
	}

	return nil
}
