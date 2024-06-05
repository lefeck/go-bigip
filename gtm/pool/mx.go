package pool

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lefeck/go-bigip"
	"strings"
)

// MXEndpoint represents the REST resource for managing MX.
const MXEndpoint = "mx"

// MXResource provides an API to manage MX configurations.
type MXResource struct {
	b *bigip.BigIP
}

// List retrieves all MX details.
func (r *MXResource) List() (*PoolList, error) {
	var items PoolList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(MXEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single MX by node name.
func (r *MXResource) Get(name string) (*Pool, error) {
	var item Pool
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(MXEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new MX item.
func (r *MXResource) Create(item Pool) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(MXEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the MX item identified by the MX name.
func (r *MXResource) Update(name string, item Pool) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(MXEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single MX identified by the MX name. If it does not exist, return an error.
func (r *MXResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(MXEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (r *MXResource) ShowMXStats(name string) (*PoolStatsList, error) {
	var item PoolStatsList

	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(MXEndpoint).SubResourceInstance(name).SubStatsResource(StatsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

func (r *MXResource) ShowAllMXStats() (*PoolStatsList, error) {
	var item PoolStatsList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(MXEndpoint).SubStatsResource(StatsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}
