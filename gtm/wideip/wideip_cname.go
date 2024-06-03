package gtm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// CNAMEEndpoint represents the REST resource for managing CNAME.
const CNAMEEndpoint = "cname"

// CNAMEResource provides an API to manage CNAME configurations.
type CNAMEResource struct {
	b *bigip.BigIP
}

// List retrieves all CNAME details.
func (r *CNAMEResource) List() (*WideipList, error) {
	var items WideipList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(CNAMEEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single CNAME by node name.
func (r *CNAMEResource) Get(name string) (*Wideip, error) {
	var item Wideip
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(CNAMEEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new CNAME item.
func (r *CNAMEResource) Create(item Wideip) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(CNAMEEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the CNAME item identified by the CNAME name.
func (r *CNAMEResource) Update(name string, item Wideip) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(CNAMEEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single CNAME identified by the CNAME name. If it does not exist, return an error.
func (r *CNAMEResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(CNAMEEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}

// ShowCNAMEStats retrieves the statistics for a single CNAME record with the given name.
func (r *CNAMEResource) ShowCNAMEStats(name string) (*WideipList, error) {
	var item WideipList

	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(CNAMEEndpoint).ResourceInstance(name).SubStatsResource(StatsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// ShowAllCNAMEStats retrieves the statistics for all CNAME records in the system.
func (r *CNAMEResource) ShowAllCNAMEStats() (*WideipList, error) {
	var item WideipList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(CNAMEEndpoint).SubStatsResource(StatsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}
