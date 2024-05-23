package sys

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// ManagementRouteList holds a list of ManagementRoute configuration.
type ManagementRouteList struct {
	Items    []ManagementRoute `json:"items"`
	Kind     string            `json:"kind"`
	SelfLink string            `json:"selflink"`
}

// ManagementRoute holds the configuration of a single ManagementRoute.
type ManagementRoute struct {
	Kind        string `json:"kind"`
	Name        string `json:"name"`
	Partition   string `json:"partition"`
	FullPath    string `json:"fullPath"`
	Generation  int    `json:"generation"`
	SelfLink    string `json:"selfLink"`
	Description string `json:"description"`
	Gateway     string `json:"gateway"`
	Mtu         int    `json:"mtu"`
	Network     string `json:"network"`
}

// ManagementRouteEndpoint represents the REST resource for managing ManagementRoute.
const ManagementRouteEndpoint = "management-route"

// ManagementRouteResource provides an API to manage ManagementRoute configurations.
type ManagementRouteResource struct {
	b *bigip.BigIP
}

// List all management route details
func (r *ManagementRouteResource) List() (*ManagementRouteList, error) {
	var items ManagementRouteList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ManagementRouteEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get a single management route details by the node name
func (r *ManagementRouteResource) Get(name string) (*ManagementRoute, error) {
	var item ManagementRoute
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ManagementRouteEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create a new management route item
func (r *ManagementRouteResource) Create(item ManagementRoute) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ManagementRouteEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update the management route item identified by the management route name, otherwise an error will be reported.
func (r *ManagementRouteResource) Update(name string, item ManagementRoute) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ManagementRouteEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single management route identified by the management route name. if it is not exist return error
func (r *ManagementRouteResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ManagementRouteEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
