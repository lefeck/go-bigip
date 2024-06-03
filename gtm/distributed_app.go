package gtm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// DistributedAppList contains a list of DistributedApp.
type DistributedAppList struct {
	Items    []DistributedApp `json:"items,omitempty"`
	Kind     string           `json:"kind,omitempty"`
	SelfLink string           `json:"selflink,omitempty"`
}

// DistributedApp holds the uration of a single DistributedApp.
type DistributedApp struct {
	DependencyLevel string `json:"dependencyLevel,omitempty"`
	FullPath        string `json:"fullPath,omitempty"`
	Generation      int    `json:"generation,omitempty"`
	Kind            string `json:"kind,omitempty"`
	Name            string `json:"name,omitempty"`
	Partition       string `json:"partition,omitempty"`
	PersistCidrIpv4 int    `json:"persistCidrIpv4,omitempty"`
	PersistCidrIpv6 int    `json:"persistCidrIpv6,omitempty"`
	Persistence     string `json:"persistence,omitempty"`
	SelfLink        string `json:"selfLink,omitempty"`
	TTLPersistence  int    `json:"ttlPersistence,omitempty"`
	Wideips         []struct {
		Name string `json:"name,omitempty"`
	} `json:"wideips,omitempty"`
}

// DistributedAppEndpoint represents the REST resource for managing DistributedApp.
const DistributedAppEndpoint = "distributed-app"

// DistributedAppResource provides an API to manage DistributedApp configurations.
type DistributedAppResource struct {
	b *bigip.BigIP
}

// List retrieves all DistributedApp details.
func (r *DistributedAppResource) List() (*DistributedAppList, error) {
	var items DistributedAppList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(DistributedAppEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single DistributedApp by node name.
func (r *DistributedAppResource) Get(name string) (*DistributedApp, error) {
	var item DistributedApp
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(DistributedAppEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new DistributedApp item.
func (r *DistributedAppResource) Create(item DistributedApp) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(DistributedAppEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the DistributedApp item identified by the DistributedApp name.
func (r *DistributedAppResource) Update(name string, item DistributedApp) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(DistributedAppEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single DistributedApp identified by the DistributedApp name. If it does not exist, return an error.
func (r *DistributedAppResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(DistributedAppEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
