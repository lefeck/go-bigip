package software

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// HotfixList holds a list of Hotfix configurations.
type HotfixList struct {
	Items    []Hotfix `json:"items"`
	Kind     string   `json:"kind"`
	SelfLink string   `json:"selflink"`
}

// Hotfix holds the configuration of a single Hotfix.
type Hotfix struct {
	Build      string `json:"build"`
	Checksum   string `json:"checksum"`
	FullPath   string `json:"fullPath"`
	Generation int    `json:"generation"`
	ID         string `json:"id"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Product    string `json:"product"`
	SelfLink   string `json:"selfLink"`
	Title      string `json:"title"`
	Verified   string `json:"verified"`
	Version    string `json:"version"`
}

// HotfixEndpoint represents the REST resource for managing Hotfix.
const HotfixEndpoint = "hotfix"

// HotfixResource provides an API to manage Hotfix configurations.
type HotfixResource struct {
	b *bigip.BigIP
}

// List retrieves all Hotfix details.
func (r *HotfixResource) List() (*HotfixList, error) {
	var items HotfixList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(HotfixEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single Hotfix by node name.
func (r *HotfixResource) Get(name string) (*Hotfix, error) {
	var item Hotfix
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(HotfixEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new Hotfix item.
func (r *HotfixResource) Create(item Hotfix) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(HotfixEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the Hotfix item identified by the Hotfix name.
func (r *HotfixResource) Update(name string, item Hotfix) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(HotfixEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single Hotfix identified by the Hotfix name. if it is not exist return error
func (r *HotfixResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(HotfixEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
