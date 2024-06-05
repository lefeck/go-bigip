package pool

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// AAAAEndpoint represents the REST resource for managing AAAA.
const AAAAEndpoint = "aaaa"

// AAAAResource provides an API to manage AAAA configurations.
type AAAAResource struct {
	b *bigip.BigIP
}

// List retrieves all AAAA details.
func (r *AAAAResource) List() (*PoolList, error) {
	var items PoolList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(AAAAEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single AAAA by node name.
func (r *AAAAResource) Get(name string) (*Pool, error) {
	var item Pool
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(AAAAEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new AAAA item.
func (r *AAAAResource) Create(item Pool) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(AAAAEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the AAAA item identified by the AAAA name.
func (r *AAAAResource) Update(name string, item Pool) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(AAAAEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single AAAA identified by the AAAA name. If it does not exist, return an error.
func (r *AAAAResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(AAAAEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (r *AAAAResource) ShowAAAAStats(name string) (*PoolStatsList, error) {
	var item PoolStatsList

	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(AAAAEndpoint).SubResourceInstance(name).SubStatsResource(StatsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

func (r *AAAAResource) ShowAllAAAAStats() (*PoolStatsList, error) {
	var item PoolStatsList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(AAAAEndpoint).SubStatsResource(StatsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}
