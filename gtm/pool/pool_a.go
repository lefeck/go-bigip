package gtm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// PoolList holds a list of Pool configuration.
type PoolList struct {
	Items    []Pool `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

// Pool holds the configuration of a single Pool.
type Pool struct {
	AlternateMode             string `json:"alternateMode,omitempty"`
	DynamicRatio              string `json:"dynamicRatio,omitempty"`
	Enabled                   bool   `json:"enabled,omitempty"`
	FallbackIP                string `json:"fallbackIp,omitempty"`
	FallbackMode              string `json:"fallbackMode,omitempty"`
	FullPath                  string `json:"fullPath,omitempty"`
	Generation                int    `json:"generation,omitempty"`
	Kind                      string `json:"kind,omitempty"`
	LimitMaxBps               int    `json:"limitMaxBps,omitempty"`
	LimitMaxBpsStatus         string `json:"limitMaxBpsStatus,omitempty"`
	LimitMaxConnections       int    `json:"limitMaxConnections,omitempty"`
	LimitMaxConnectionsStatus string `json:"limitMaxConnectionsStatus,omitempty"`
	LimitMaxPps               int    `json:"limitMaxPps,omitempty"`
	LimitMaxPpsStatus         string `json:"limitMaxPpsStatus,omitempty"`
	LoadBalancingMode         string `json:"loadBalancingMode,omitempty"`
	ManualResume              string `json:"manualResume,omitempty"`
	MaxAnswersReturned        int    `json:"maxAnswersReturned,omitempty"`
	MembersReference          struct {
		Members         []PoolMembers `json:"items,omitempty"`
		IsSubcollection bool          `json:"isSubcollection,omitempty"`
		Link            string        `json:"link,omitempty"`
	} `json:"membersReference,omitempty"`
	Monitor                  string `json:"monitor,omitempty"`
	Name                     string `json:"name,omitempty"`
	Partition                string `json:"partition,omitempty"`
	QosHitRatio              int    `json:"qosHitRatio,omitempty"`
	QosHops                  int    `json:"qosHops,omitempty"`
	QosKilobytesSecond       int    `json:"qosKilobytesSecond,omitempty"`
	QosLcs                   int    `json:"qosLcs,omitempty"`
	QosPacketRate            int    `json:"qosPacketRate,omitempty"`
	QosRtt                   int    `json:"qosRtt,omitempty"`
	QosTopology              int    `json:"qosTopology,omitempty"`
	QosVsCapacity            int    `json:"qosVsCapacity,omitempty"`
	QosVsScore               int    `json:"qosVsScore,omitempty"`
	SelfLink                 string `json:"selfLink,omitempty"`
	TTL                      int    `json:"ttl,omitempty"`
	VerifyMemberAvailability string `json:"verifyMemberAvailability,omitempty"`
}

type PoolMembersList struct {
	Items    []PoolMembers `json:"items,omitempty"`
	Kind     string        `json:"kind,omitempty"`
	SelfLink string        `json:"selflink,omitempty"`
}

// PoolMembers holds the configuration of a single PoolMembers.
type PoolMembers struct {
	Enabled                   bool   `json:"enabled,omitempty"`
	FullPath                  string `json:"fullPath,omitempty"`
	Generation                int    `json:"generation,omitempty"`
	Kind                      string `json:"kind,omitempty"`
	LimitMaxBps               int    `json:"limitMaxBps,omitempty"`
	LimitMaxBpsStatus         string `json:"limitMaxBpsStatus,omitempty"`
	LimitMaxConnections       int    `json:"limitMaxConnections,omitempty"`
	LimitMaxConnectionsStatus string `json:"limitMaxConnectionsStatus,omitempty"`
	LimitMaxPps               int    `json:"limitMaxPps,omitempty"`
	LimitMaxPpsStatus         string `json:"limitMaxPpsStatus,omitempty"`
	MemberOrder               int    `json:"memberOrder,omitempty"`
	Monitor                   string `json:"monitor,omitempty"`
	Name                      string `json:"name,omitempty"`
	Partition                 string `json:"partition,omitempty"`
	Ratio                     int    `json:"ratio,omitempty"`
	SelfLink                  string `json:"selfLink,omitempty"`
}

type PoolStatsList struct {
	Entries  map[string]PoolStatsEntries `json:"entries,omitempty"`
	Kind     string                      `json:"kind,omitempty" pretty:",expanded"`
	SelfLink string                      `json:"selflink,omitempty" pretty:",expanded"`
}

type PoolStatsEntries struct {
	NestedPoolStats PoolStats `json:"nestedStats,omitempty"`
}

type PoolStats struct {
	Entries struct {
		Alternate struct {
			Value int `json:"value"`
		} `json:"alternate"`
		Dropped struct {
			Value int `json:"value"`
		} `json:"dropped"`
		Fallback struct {
			Value int `json:"value"`
		} `json:"fallback"`
		PoolType struct {
			Description string `json:"description"`
		} `json:"poolType"`
		Preferred struct {
			Value int `json:"value"`
		} `json:"preferred"`
		ReturnFromDNS struct {
			Value int `json:"value"`
		} `json:"returnFromDns"`
		ReturnToDNS struct {
			Value int `json:"value"`
		} `json:"returnToDns"`
		Status_availabilityState struct {
			Description string `json:"description"`
		} `json:"status.availabilityState"`
		Status_enabledState struct {
			Description string `json:"description"`
		} `json:"status.enabledState"`
		Status_statusReason struct {
			Description string `json:"description"`
		} `json:"status.statusReason"`
		TmName struct {
			Description string `json:"description"`
		} `json:"tmName"`
	} `json:"entries"`
}

// PoolEndpoint represents the REST resource for managing Pool.
const PoolAEndpoint = "a"

// StatsEndpoint represents the REST resource for managing stats.
const StatsEndpoint = "stats"

// PoolAResource provides an API to manage PoolA configurations.
type PoolAResource struct {
	b *bigip.BigIP
}

// List retrieves all PoolA details.
func (r *PoolAResource) List() (*PoolList, error) {
	var items PoolList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(PoolAEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single PoolA by node name.
func (r *PoolAResource) Get(name string) (*Pool, error) {
	var item Pool
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(PoolAEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new PoolA item.
func (r *PoolAResource) Create(item Pool) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(PoolAEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the PoolA item identified by the PoolA name.
func (r *PoolAResource) Update(name string, item Pool) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(PoolAEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single PoolA identified by the PoolA name. If it does not exist, return an error.
func (r *PoolAResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(PoolAEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (r *PoolAResource) ShowAStats(name string) (*PoolStatsList, error) {
	var item PoolStatsList

	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(PoolAEndpoint).SubResourceInstance(name).SubStatsResource(StatsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

func (r *PoolAResource) ShowAllAStats() (*PoolStatsList, error) {
	var item PoolStatsList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PoolEndpoint).SubResource(PoolAEndpoint).SubStatsResource(StatsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}
