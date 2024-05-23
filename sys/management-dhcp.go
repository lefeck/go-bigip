package sys

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// ManagementDHCPList holds a list of ManagementDHCP configuration.
type ManagementDHCPList struct {
	Items    []ManagementDHCP `json:"items"`
	Kind     string           `json:"kind"`
	SelfLink string           `json:"selflink"`
}

// ManagementDHCP holds the configuration of a single ManagementDHCP.
type ManagementDHCP struct {
	Kind           string   `json:"kind"`
	Name           string   `json:"name"`
	Partition      string   `json:"partition"`
	FullPath       string   `json:"fullPath"`
	Generation     int      `json:"generation"`
	SelfLink       string   `json:"selfLink"`
	RequestOptions []string `json:"requestOptions"`
}

// ManagementDHCPEndpoint represents the REST resource for managing ManagementDHCP.
const ManagementDHCPEndpoint = "management-dhcp"

// ManagementDHCPResource provides an API to manage ManagementDHCP configurations.
type ManagementDHCPResource struct {
	b *bigip.BigIP
}

// List all management DHCP details
func (r *ManagementDHCPResource) List() (*ManagementDHCPList, error) {
	var items ManagementDHCPList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ManagementDHCPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get a single management DHCP details by the node name
func (r *ManagementDHCPResource) Get(name string) (*ManagementDHCP, error) {
	var item ManagementDHCP
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ManagementDHCPEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create a new management DHCP item
func (r *ManagementDHCPResource) Create(item ManagementDHCP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ManagementDHCPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update the management DHCP item identified by the management DHCP name, otherwise an error will be reported.
func (r *ManagementDHCPResource) Update(name string, item ManagementDHCP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ManagementDHCPEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single management DHCP identified by the management DHCP name. if it is not exist return error
func (r *ManagementDHCPResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ManagementDHCPEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
