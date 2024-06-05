package wideip

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
func (r *AAAAResource) List() (*WideipList, error) {
	var items WideipList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(AAAAEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single AAAA by node name.
func (r *AAAAResource) Get(name string) (*Wideip, error) {
	var item Wideip
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(AAAAEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new AAAA item.
func (r *AAAAResource) Create(item Wideip) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(AAAAEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the AAAA item identified by the AAAA name.
func (r *AAAAResource) Update(name string, item Wideip) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(AAAAEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single AAAA identified by the AAAA name. If it does not exist, return an error.
func (r *AAAAResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(AAAAEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}

// ShowAAAAStats retrieves the statistics for a single AAAA record with the given name.
func (r *AAAAResource) ShowAAAAStats(name string) (*WideipList, error) {
	var item WideipList

	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(AAAAEndpoint).ResourceInstance(name).SubStatsResource(StatsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// ShowAllAAAAStats retrieves the statistics for all AAAA records in the system.
func (r *AAAAResource) ShowAllAAAAStats() (*WideipList, error) {
	var item WideipList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(AAAAEndpoint).SubStatsResource(StatsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}
