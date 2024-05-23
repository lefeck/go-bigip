package software

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// BlockDeviceHotfixList holds a list of BlockDeviceHotfix configurations.
type BlockDeviceHotfixList struct {
	Items    []BlockDeviceHotfix `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

// BlockDeviceHotfix holds the configuration of a single BlockDeviceHotfix.
type BlockDeviceHotfix struct {
}

// BlockDeviceHotfixEndpoint represents the REST resource for managing BlockDeviceHotfix.
const BlockDeviceHotfixEndpoint = "block-device-hotfix"

// BlockDeviceHotfixResource provides an API to manage BlockDeviceHotfix configurations.
type BlockDeviceHotfixResource struct {
	b *bigip.BigIP
}

// List retrieves all BlockDeviceHotfix details.
func (r *BlockDeviceHotfixResource) List() (*BlockDeviceHotfixList, error) {
	var items BlockDeviceHotfixList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(BlockDeviceHotfixEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single BlockDeviceHotfix by node name.
func (r *BlockDeviceHotfixResource) Get(name string) (*BlockDeviceHotfix, error) {
	var item BlockDeviceHotfix
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(BlockDeviceHotfixEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new BlockDeviceHotfix item.
func (r *BlockDeviceHotfixResource) Create(item BlockDeviceHotfix) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(BlockDeviceHotfixEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the BlockDeviceHotfix item identified by the BlockDeviceHotfix name.
func (r *BlockDeviceHotfixResource) Update(name string, item BlockDeviceHotfix) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(BlockDeviceHotfixEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single BlockDeviceHotfix identified by the BlockDeviceHotfix name. if it is not exist return error
func (r *BlockDeviceHotfixResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(BlockDeviceHotfixEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
