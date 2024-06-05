package wideip

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// WideipList holds a list of WideipA configuration.
type WideipList struct {
	Items    []Wideip `json:"items,omitempty"`
	Kind     string   `json:"kind,omitempty"`
	SelfLink string   `json:"selflink,omitempty"`
}

// Wideip holds the configuration of a single WideipA.
type Wideip struct {
	Enabled              bool   `json:"enabled,omitempty"`
	FailureRcode         string `json:"failureRcode,omitempty"`
	FailureRcodeResponse string `json:"failureRcodeResponse,omitempty"`
	FailureRcodeTTL      int    `json:"failureRcodeTtl,omitempty"`
	FullPath             string `json:"fullPath,omitempty"`
	Generation           int    `json:"generation,omitempty"`
	Kind                 string `json:"kind,omitempty"`
	LastResortPool       string `json:"lastResortPool,omitempty"`
	MinimalResponse      string `json:"minimalResponse,omitempty"`
	Name                 string `json:"name,omitempty"`
	Partition            string `json:"partition,omitempty"`
	PersistCidrIpv4      int    `json:"persistCidrIpv4,omitempty"`
	PersistCidrIpv6      int    `json:"persistCidrIpv6,omitempty"`
	Persistence          string `json:"persistence,omitempty"`
	PoolLbMode           string `json:"poolLbMode,omitempty"`
	Pools                []struct {
		Name          string `json:"name,omitempty"`
		NameReference struct {
			Link string `json:"link,omitempty"`
		} `json:"nameReference,omitempty"`
		Order     int    `json:"order,omitempty"`
		Partition string `json:"partition,omitempty"`
		Ratio     int    `json:"ratio,omitempty"`
	} `json:"pools,omitempty"`
	SelfLink       string `json:"selfLink,omitempty"`
	TTLPersistence int    `json:"ttlPersistence,omitempty"`
}

// StatsEndpoint represents the REST resource for managing stats.
const StatsEndpoint = "stats"

// AEndpoint represents the REST resource for managing A records.
const AEndpoint = "a"

// AResource provides an API to manage A record configurations.
type AResource struct {
	b *bigip.BigIP
}

// List retrieves all A record details.
func (r *AResource) List() (*WideipList, error) {
	var items WideipList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(AEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single A record by node name.
func (r *AResource) Get(name string) (*Wideip, error) {
	var item Wideip
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(AEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new A record item.
func (r *AResource) Create(item Wideip) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(AEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the A record item identified by the A record name.
func (r *AResource) Update(name string, item Wideip) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(AEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single A record identified by the A record name. If it does not exist, return an error.
func (r *AResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(AEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// ShowAStats retrieves the statistics for a single A record with the given name.
//
//	name: The name of the A record for which you want to retrieve statistics.
//
// Returns:
//   - *WideipList: Pointer to a structure containing the list of wide IP A records and their statistic details.
//   - error: If an error occurs during the operation, it will be returned.
func (r *AResource) ShowAStats(name string) (*WideipList, error) {
	var item WideipList

	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(AEndpoint).ResourceInstance(name).SubStatsResource(StatsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// ShowAllAStats retrieves the statistics for all A records in the system.
// Returns:
//   - *WideipList: Pointer to a structure containing the list of all wide IP A records and their statistic details.
//   - error: If an error occurs during the operation, it will be returned.
func (r *AResource) ShowAllAStats() (*WideipList, error) {
	var item WideipList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(WideipEndpoint).SubResource(AEndpoint).SubStatsResource(StatsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}
