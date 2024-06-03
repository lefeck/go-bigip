package gtm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// PoolAAAAEndpoint represents the REST resource for managing PoolAAAA.
const PoolAAAAEndpoint = "aaaa"

// PoolAAAAResource provides an API to manage PoolAAAA configurations.
type PoolAAAAResource struct {
	b *bigip.BigIP
}

// List retrieves all PoolAAAA details.
func (r *PoolAAAAResource) List() (*PoolList, error) {
	var items PoolList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(PoolAAAAEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single PoolAAAA by node name.
func (r *PoolAAAAResource) Get(name string) (*Pool, error) {
	var item Pool
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(PoolAAAAEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new PoolAAAA item.
func (r *PoolAAAAResource) Create(item Pool) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(PoolAAAAEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the PoolAAAA item identified by the PoolAAAA name.
func (r *PoolAAAAResource) Update(name string, item Pool) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(PoolAAAAEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single PoolAAAA identified by the PoolAAAA name. If it does not exist, return an error.
func (r *PoolAAAAResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(PoolAAAAEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (r *PoolAAAAResource) ShowAAAAStats(name string) (*PoolStatsList, error) {
	var item PoolStatsList

	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(PoolAAAAEndpoint).SubResourceInstance(name).SubStatsResource(StatsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

func (r *PoolAAAAResource) ShowAllAAAAStats() (*PoolStatsList, error) {
	var item PoolStatsList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(PoolAAAAEndpoint).SubStatsResource(StatsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}
