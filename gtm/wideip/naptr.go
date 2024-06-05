package wideip

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lefeck/go-bigip"
	"strings"
)

// NAPTREndpoint represents the REST resource for managing NAPTR.
const NAPTREndpoint = "naptr"

// NAPTRResource provides an API to manage NAPTR configurations.
type NAPTRResource struct {
	b *bigip.BigIP
}

// List retrieves all NAPTR details.
func (r *NAPTRResource) List() (*WideipList, error) {
	var items WideipList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(NAPTREndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single NAPTR by node name.
func (r *NAPTRResource) Get(name string) (*Wideip, error) {
	var item Wideip
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(NAPTREndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new NAPTR item.
func (r *NAPTRResource) Create(item Wideip) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(NAPTREndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the NAPTR item identified by the NAPTR name.
func (r *NAPTRResource) Update(name string, item Wideip) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(NAPTREndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single NAPTR identified by the NAPTR name. If it does not exist, return an error.
func (r *NAPTRResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(NAPTREndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}

// ShowNAPTRStats retrieves the statistics for a single NAPTR record with the given name.
func (r *NAPTRResource) ShowNAPTRStats(name string) (*WideipList, error) {
	var item WideipList

	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(NAPTREndpoint).ResourceInstance(name).SubStatsResource(StatsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// ShowAllNAPTRAStats retrieves the statistics for all NAPTR records in the system.
func (r *NAPTRResource) ShowAllNAPTRStats() (*WideipList, error) {
	var item WideipList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(NAPTREndpoint).SubStatsResource(StatsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}
