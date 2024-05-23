// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package disk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// LogicalDiskList holds a list of LogicalDisk configuration.
type LogicalDiskList struct {
	Items    []LogicalDisk `json:"items"`
	Kind     string        `json:"kind"`
	SelfLink string        `json:"selflink"`
}

// LogicalDisk holds the configuration of a single LogicalDisk.
type LogicalDisk struct {
	FullPath   string `json:"fullPath,omitempty"`
	Generation int    `json:"generation,omitempty"`
	Kind       string `json:"kind,omitempty"`
	Mode       string `json:"mode,omitempty"`
	Name       string `json:"name,omitempty"`
	SelfLink   string `json:"selfLink,omitempty"`
	Size       int    `json:"size,omitempty"`
	VgFree     int    `json:"vgFree,omitempty"`
	VgInUse    int    `json:"vgInUse,omitempty"`
	VgReserved int    `json:"vgReserved,omitempty"`
}

// LogicalDiskEndpoint represents the REST resource for managing LogicalDisk.
const LogicalDiskEndpoint = "logical-disk"

// LogicalDiskResource provides an API to manage LogicalDisk configurations.
type LogicalDiskResource struct {
	b *bigip.BigIP
}

// List all logical disk details
func (r *LogicalDiskResource) List() (*LogicalDiskList, error) {
	var items LogicalDiskList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(LogicalDiskEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get a single logical disk details by the node name
func (r *LogicalDiskResource) Get(name string) (*LogicalDisk, error) {
	var item LogicalDisk
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(LogicalDiskEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create a new logical disk item
func (r *LogicalDiskResource) Create(item LogicalDisk) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(LogicalDiskEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update the logical disk item identified by the logical disk name, otherwise an error will be reported.
func (r *LogicalDiskResource) Update(name string, item LogicalDisk) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(LogicalDiskEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single logical disk identified by the logical disk name. if it is not exist return error
func (r *LogicalDiskResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(LogicalDiskEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
