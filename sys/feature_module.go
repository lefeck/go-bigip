package sys

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// FeatureModuleList holds a list of FeatureModule configurations.
type FeatureModuleList struct {
	Items    []FeatureModule `json:"items"`
	Kind     string          `json:"kind"`
	SelfLink string          `json:"selflink"`
}

// FeatureModule holds the configuration of a single FeatureModule.
type FeatureModule struct {
	Disabled   bool   `json:"disabled"`
	FullPath   string `json:"fullPath"`
	Generation int    `json:"generation"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	SelfLink   string `json:"selfLink"`
}

// FeatureModuleEndpoint represents the REST resource for managing FeatureModule.
const FeatureModuleEndpoint = "feature-module"

// FeatureModuleResource provides an API to manage FeatureModule configurations.
type FeatureModuleResource struct {
	b *bigip.BigIP
}

// List retrieves all FeatureModule details.
func (r *FeatureModuleResource) List() (*FeatureModuleList, error) {
	var items FeatureModuleList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(FeatureModuleEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single FeatureModule by node name.
func (r *FeatureModuleResource) Get(name string) (*FeatureModule, error) {
	var item FeatureModule
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(FeatureModuleEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new FeatureModule item.
func (r *FeatureModuleResource) Create(item FeatureModule) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(FeatureModuleEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the FeatureModule item identified by the FeatureModule name.
func (r *FeatureModuleResource) Update(name string, item FeatureModule) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(FeatureModuleEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single FeatureModule identified by the FeatureModule name. If it does not exist, return an error.
func (r *FeatureModuleResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(FeatureModuleEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
