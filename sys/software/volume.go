package software

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// VolumeList holds a list of Volume configurations.
type VolumeList struct {
	Items    []Volume `json:"items"`
	Kind     string   `json:"kind"`
	SelfLink string   `json:"selflink"`
}

// Volume holds the configuration of a single Volume.
type Volume struct {
	Kind         string `json:"kind"`
	Name         string `json:"name"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	SelfLink     string `json:"selfLink"`
	Active       bool   `json:"active"`
	APIRawValues struct {
	} `json:"apiRawValues"`
	Basebuild string `json:"basebuild"`
	Build     string `json:"build"`
	Product   string `json:"product"`
	Status    string `json:"status"`
	Version   string `json:"version"`
	Media     []struct {
		Name                string `json:"name"`
		DefaultBootLocation bool   `json:"defaultBootLocation"`
		Media               string `json:"media"`
		Size                string `json:"size"`
		NameReference       struct {
			Link string `json:"link"`
		} `json:"nameReference"`
	} `json:"media"`
}

// VolumeEndpoint represents the REST resource for managing Volume.
const VolumeEndpoint = "volume"

// VolumeResource provides an API to manage Volume configurations.
type VolumeResource struct {
	b *bigip.BigIP
}

// List retrieves all Volume details.
func (r *VolumeResource) List() (*VolumeList, error) {
	var items VolumeList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(VolumeEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single Volume by node name.
func (r *VolumeResource) Get(name string) (*Volume, error) {
	var item Volume
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(VolumeEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new Volume item.
func (r *VolumeResource) Create(item Volume) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(VolumeEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the Volume item identified by the Volume name.
func (r *VolumeResource) Update(name string, item Volume) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(VolumeEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single Volume identified by the Volume name. if it is not exist return error
func (r *VolumeResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(VolumeEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
