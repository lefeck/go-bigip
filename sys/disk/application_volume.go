package disk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// ApplicationVolumeList holds a list of ApplicationVolume configuration.
type ApplicationVolumeList struct {
	Items    []ApplicationVolume `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

// ApplicationVolume holds the configuration of a single ApplicationVolume.
type ApplicationVolume struct {
}

// ApplicationVolumeEndpoint represents the REST resource for managing ApplicationVolume.
const ApplicationVolumeEndpoint = "application-volume"

// ApplicationVolumeResource provides an API to manage ApplicationVolume configurations.
type ApplicationVolumeResource struct {
	b *bigip.BigIP
}

// List all application volume details
func (r *ApplicationVolumeResource) List() (*ApplicationVolumeList, error) {
	var items ApplicationVolumeList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ApplicationVolumeEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get a single application volume details by the node name
func (r *ApplicationVolumeResource) Get(name string) (*ApplicationVolume, error) {
	var item ApplicationVolume
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ApplicationVolumeEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create a new application volume item
func (r *ApplicationVolumeResource) Create(item ApplicationVolume) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ApplicationVolumeEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update the application volume item identified by the application volume name, otherwise an error will be reported.
func (r *ApplicationVolumeResource) Update(name string, item ApplicationVolume) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ApplicationVolumeEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single application volume identified by the application volume name. if it is not exist return error
func (r *ApplicationVolumeResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ApplicationVolumeEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
