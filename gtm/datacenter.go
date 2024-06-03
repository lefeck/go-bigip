package gtm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// DatacenterList holds a list of Datacenter configuration.
type DatacenterList struct {
	Items    []Datacenter `json:"items,omitempty"`
	Kind     string       `json:"kind,omitempty"`
	SelfLink string       `json:"selflink,omitempty"`
}

// Datacenter holds the configuration of a single Datacenter.
type Datacenter struct {
	Kind             string `json:"kind,omitempty"`
	Name             string `json:"name,omitempty"`
	Partition        string `json:"partition,omitempty"`
	FullPath         string `json:"fullPath,omitempty"`
	Generation       int    `json:"generation,omitempty"`
	SelfLink         string `json:"selfLink,omitempty"`
	Contact          string `json:"contact,omitempty"`
	Enabled          bool   `json:"enabled,omitempty"`
	Location         string `json:"location,omitempty"`
	ProberFallback   string `json:"proberFallback,omitempty"`
	ProberPreference string `json:"proberPreference,omitempty"`
}

// DatacenterEndpoint represents the REST resource for managing Datacenter.
const DatacenterEndpoint = "datacenter"

// DatacenterResource provides an API to manage Datacenter configurations.
type DatacenterResource struct {
	b *bigip.BigIP
}

// List retrieves all Datacenter details.
func (r *DatacenterResource) List() (*DatacenterList, error) {
	var items DatacenterList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(DatacenterEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single Datacenter by node name.
func (r *DatacenterResource) Get(name string) (*Datacenter, error) {
	var item Datacenter
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(DatacenterEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new Datacenter item.
func (r *DatacenterResource) Create(item Datacenter) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(DatacenterEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the Datacenter item identified by the Datacenter name.
func (r *DatacenterResource) Update(name string, item Datacenter) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(DatacenterEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single Datacenter identified by the Datacenter name. If it does not exist, return an error.
func (r *DatacenterResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(DatacenterEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
