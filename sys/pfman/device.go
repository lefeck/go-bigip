package pfman

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// DeviceList holds a list of Device configurations.
type DeviceList struct {
	Items    []Device `json:"items"`
	Kind     string   `json:"kind"`
	SelfLink string   `json:"selflink"`
}

// Device holds the configuration of a single Device.
type Device struct {
}

// DeviceEndpoint represents the REST resource for managing Device.
const DeviceEndpoint = "device"

// DeviceResource provides an API to manage Device configurations.
type DeviceResource struct {
	b *bigip.BigIP
}

// List retrieves all Device details.
func (r *DeviceResource) List() (*DeviceList, error) {
	var items DeviceList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(PFManEndpoint).SubResource(DeviceEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single Device by node name.
func (r *DeviceResource) Get(name string) (*Device, error) {
	var item Device
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(PFManEndpoint).SubResource(DeviceEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new Device item.
func (r *DeviceResource) Create(item Device) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(PFManEndpoint).SubResource(DeviceEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the Device item identified by the Device name.
func (r *DeviceResource) Update(name string, item Device) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(PFManEndpoint).SubResource(DeviceEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single Device identified by the Device name. If it does not exist, return an error.
func (r *DeviceResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(PFManEndpoint).SubResource(DeviceEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
