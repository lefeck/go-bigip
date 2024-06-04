package pool

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lefeck/go-bigip"
	"strings"
)

// NAPTREndpoint represents the REST resource for managing NAPTR.
const NAPTREndpoint = "naptr"

// PoolNAPTRResource provides an API to manage PoolNAPTR configurations.
type PoolNAPTRResource struct {
	b *bigip.BigIP
}

// List retrieves all PoolNAPTR details.
func (r *PoolNAPTRResource) List() (*PoolList, error) {
	var items PoolList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(NAPTREndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single PoolNAPTR by node name.
func (r *PoolNAPTRResource) Get(name string) (*Pool, error) {
	var item Pool
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(NAPTREndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new PoolNAPTR item.
func (r *PoolNAPTRResource) Create(item Pool) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(NAPTREndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the PoolNAPTR item identified by the PoolNAPTR name.
func (r *PoolNAPTRResource) Update(name string, item Pool) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(NAPTREndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single PoolNAPTR identified by the PoolNAPTR name. If it does not exist, return an error.
func (r *PoolNAPTRResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(NAPTREndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (r *PoolNAPTRResource) ShowNAPTRStats(name string) (*PoolStatsList, error) {
	var item PoolStatsList

	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(NAPTREndpoint).SubResourceInstance(name).SubStatsResource(StatsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

func (r *PoolNAPTRResource) ShowAllNAPTRStats() (*PoolStatsList, error) {
	var item PoolStatsList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(NAPTREndpoint).SubStatsResource(StatsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}
