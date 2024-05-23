package application

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// CustomStatList holds a list of CustomStat configurations.
type CustomStatList struct {
	Items    []CustomStat `json:"items"`
	Kind     string       `json:"kind"`
	SelfLink string       `json:"selflink"`
}

// CustomStat holds the configuration of a single CustomStat.
type CustomStat struct {
	// ... Fields representing the CustomStat configuration
}

// CustomStatEndpoint represents the REST resource for managing CustomStat.
const CustomStatEndpoint = "custom-stat"

// CustomStatResource provides an API to manage CustomStat configurations.
type CustomStatResource struct {
	b *bigip.BigIP
}

// List retrieves all CustomStat details.
func (r *CustomStatResource) List() (*CustomStatList, error) {
	var items CustomStatList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ApplicationEndpoint).SubResource(CustomStatEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single CustomStat by node name.
func (r *CustomStatResource) Get(name string) (*CustomStat, error) {
	var item CustomStat
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ApplicationEndpoint).SubResource(CustomStatEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new CustomStat item.
func (r *CustomStatResource) Create(item CustomStat) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ApplicationEndpoint).SubResource(CustomStatEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the CustomStat item identified by the CustomStat name.
func (r *CustomStatResource) Update(name string, item CustomStat) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ApplicationEndpoint).SubResource(CustomStatEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single CustomStat identified by the CustomStat name. if it is not exist return error
func (r *CustomStatResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ApplicationEndpoint).SubResource(CustomStatEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
