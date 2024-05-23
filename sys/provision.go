package sys

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
)

// ProvisionList holds a list of Provision configurations.
type ProvisionList struct {
	Items    []Provision `json:"items"`
	Kind     string      `json:"kind"`
	SelfLink string      `json:"selflink"`
}

// Provision holds the configuration of a single Provision.
type Provision struct {
	Kind        string `json:"kind"`
	Name        string `json:"name"`
	FullPath    string `json:"fullPath"`
	Generation  int    `json:"generation"`
	SelfLink    string `json:"selfLink"`
	CPURatio    int    `json:"cpuRatio"`
	DiskRatio   int    `json:"diskRatio"`
	Level       string `json:"level"`
	MemoryRatio int    `json:"memoryRatio"`
}

// ProvisionEndpoint represents the REST resource for managing Provision.
const ProvisionEndpoint = "provision"

// ProvisionResource provides an API to manage Provision configurations.
type ProvisionResource struct {
	b *bigip.BigIP
}

// List retrieves all Provision details.
func (r *ProvisionResource) List() (*ProvisionList, error) {
	var items ProvisionList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ProvisionEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single Provision by node name.
func (r *ProvisionResource) Get(name string) (*Provision, error) {
	var item Provision
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ProvisionEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}
