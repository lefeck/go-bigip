package software

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// BlockDeviceImageList holds a list of BlockDeviceImage configurations.
type BlockDeviceImageList struct {
	Items    []BlockDeviceImage `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

// BlockDeviceImage holds the configuration of a single BlockDeviceImage.
type BlockDeviceImage struct {
}

// BlockDeviceImageEndpoint represents the REST resource for managing BlockDeviceImage.
const BlockDeviceImageEndpoint = "block-device-image"

// BlockDeviceImageResource provides an API to manage BlockDeviceImage configurations.
type BlockDeviceImageResource struct {
	b *bigip.BigIP
}

// List retrieves all BlockDeviceImage details.
func (r *BlockDeviceImageResource) List() (*BlockDeviceImageList, error) {
	var items BlockDeviceImageList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(BlockDeviceImageEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single BlockDeviceImage by node name.
func (r *BlockDeviceImageResource) Get(name string) (*BlockDeviceImage, error) {
	var item BlockDeviceImage
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(BlockDeviceImageEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new BlockDeviceImage item.
func (r *BlockDeviceImageResource) Create(item BlockDeviceImage) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(BlockDeviceImageEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the BlockDeviceImage item identified by the BlockDeviceImage name.
func (r *BlockDeviceImageResource) Update(name string, item BlockDeviceImage) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(BlockDeviceImageEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single BlockDeviceImage identified by the BlockDeviceImage name. if it is not exist return error
func (r *BlockDeviceImageResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(BlockDeviceImageEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
