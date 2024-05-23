package sys

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// ScriptdList holds a list of Scriptd configuration.
type ScriptdList struct {
	Items    []Scriptd `json:"items"`
	Kind     string    `json:"kind"`
	SelfLink string    `json:"selflink"`
}

// Scriptd holds the configuration of a single Scriptd.
type Scriptd struct {
}

// ScriptdEndpoint represents the REST resource for managing Scriptd.
const ScriptdEndpoint = "scriptd"

// ScriptdResource provides an API to manage Scriptd configurations.
type ScriptdResource struct {
	b *bigip.BigIP
}

// List all scriptd details
func (r *ScriptdResource) List() (*ScriptdList, error) {
	var items ScriptdList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ScriptdEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get a single scriptd details by the node name
func (r *ScriptdResource) Get(name string) (*Scriptd, error) {
	var item Scriptd
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ScriptdEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create a new scriptd item
func (r *ScriptdResource) Create(item Scriptd) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ScriptdEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update the scriptd item identified by the scriptd name, otherwise an error will be reported.
func (r *ScriptdResource) Update(name string, item Scriptd) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ScriptdEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single scriptd identified by the scriptd name. if it is not exist return error
func (r *ScriptdResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ScriptdEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
