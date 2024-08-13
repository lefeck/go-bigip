package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type PartitionList struct {
	Item     []Partition `json:"items,omitempty"`
	Kind     string      `json:"kind"`
	SelfLink string      `json:"selfLink"`
}

type Partition struct {
	Kind               string `json:"kind,omitempty"`
	Name               string `json:"name,omitempty"`
	FullPath           string `json:"fullPath,omitempty"`
	Generation         int    `json:"generation,omitempty"`
	SelfLink           string `json:"selfLink,omitempty"`
	DefaultRouteDomain int    `json:"defaultRouteDomain,omitempty"`
	Description        string `json:"description,omitempty"`
}

type PartitionResource struct {
	b *bigip.BigIP
}

// PartitionEndpoint is the base path of the auth API.
const PartitionEndpoint = "partition"

func (r *PartitionResource) List() (*PartitionList, error) {
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).
		ManagerName(AuthManager).Resource(PartitionEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	var pl PartitionList
	if err := json.Unmarshal(res, &pl); err != nil {
		panic(err)
	}
	return &pl, nil
}

// Get a single Partition configuration identified by name.
func (r *PartitionResource) Get(name string) (*Partition, error) {
	var item Partition
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(AuthManager).
		Resource(PartitionEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create a new Partition configuration.
func (r *PartitionResource) Create(item Partition) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(AuthManager).
		Resource(PartitionEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Edit a Partition configuration identified by name.
func (r *PartitionResource) Update(fullPathName string, item Partition) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(AuthManager).
		Resource(PartitionEndpoint).ResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single Partition configuration identified by name.
func (r *PartitionResource) Delete(fullPathName string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(AuthManager).
		Resource(PartitionEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
