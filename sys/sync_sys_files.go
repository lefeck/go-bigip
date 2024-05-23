package sys

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// SyncSysFilesList holds a list of SyncSysFiles configuration.
type SyncSysFilesList struct {
	Items    []SyncSysFiles `json:"items"`
	Kind     string         `json:"kind"`
	SelfLink string         `json:"selflink"`
}

// SyncSysFiles holds the configuration of a single SyncSysFiles.
type SyncSysFiles struct {
}

// SyncSysFilesEndpoint represents the REST resource for managing SyncSysFiles.
const SyncSysFilesEndpoint = "sync-sys-files"

// SyncSysFilesResource provides an API to manage SyncSysFiles configurations.
type SyncSysFilesResource struct {
	b *bigip.BigIP
}

// List all sync sys files details
func (r *SyncSysFilesResource) List() (*SyncSysFilesList, error) {
	var items SyncSysFilesList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SyncSysFilesEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get a single sync sys files details by the node name
func (r *SyncSysFilesResource) Get(name string) (*SyncSysFiles, error) {
	var item SyncSysFiles
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SyncSysFilesEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create a new sync sys files item
func (r *SyncSysFilesResource) Create(item SyncSysFiles) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SyncSysFilesEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update the sync sys files item identified by the sync sys files name, otherwise an error will be reported.
func (r *SyncSysFilesResource) Update(name string, item SyncSysFiles) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SyncSysFilesEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single sync sys files identified by the sync sys files name. if it is not exist return error
func (r *SyncSysFilesResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SyncSysFilesEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
