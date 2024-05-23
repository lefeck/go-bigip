package software

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// VolumeList holds a list of UpdateStatus configurations.
type UpdateStatusList struct {
	Items    []UpdateStatus `json:"items"`
	Kind     string         `json:"kind"`
	SelfLink string         `json:"selflink"`
}

// UpdateStatus holds the configuration of a single UpdateStatus.
type UpdateStatus struct {
	Kind         string `json:"kind"`
	Name         string `json:"name"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	SelfLink     string `json:"selfLink"`
	Active       bool   `json:"active"`
	APIRawValues struct {
	} `json:"apiRawValues"`
	Basebuild    string `json:"basebuild"`
	Build        string `json:"build"`
	Product      string `json:"product"`
	UpdateStatus string `json:"updateStatus"`
	Version      string `json:"version"`
	Media        []struct {
		Name                string `json:"name"`
		DefaultBootLocation bool   `json:"defaultBootLocation"`
		Media               string `json:"media"`
		Size                string `json:"size"`
		NameReference       struct {
			Link string `json:"link"`
		} `json:"nameReference"`
	} `json:"media"`
}

// UpdateStatusEndpoint represents the REST resource for managing UpdateStatus.
const UpdateStatusEndpoint = "update-status"

// UpdateStatusResource provides an API to manage UpdateStatus configurations.
type UpdateStatusResource struct {
	b *bigip.BigIP
}

// List retrieves all UpdateStatus details.
func (r *UpdateStatusResource) List() (*UpdateStatusList, error) {
	var items UpdateStatusList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(UpdateStatusEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single UpdateStatus by node name.
func (r *UpdateStatusResource) Get(name string) (*UpdateStatus, error) {
	var item UpdateStatus
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(UpdateStatusEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new UpdateStatus item.
func (r *UpdateStatusResource) Create(item UpdateStatus) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(UpdateStatusEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the UpdateStatus item identified by the UpdateStatus name.
func (r *UpdateStatusResource) Update(name string, item UpdateStatus) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(UpdateStatusEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single UpdateStatus identified by the UpdateStatus name. if it is not exist return error
func (r *UpdateStatusResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(UpdateStatusEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
