package ltm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
)

// Get the stats of a specific member under a pool
type MemberStatsList struct {
	Entries  map[string]MemberStatsEntries `json:"entries,omitempty"`
	Kind     string                        `json:"kind,omitempty" pretty:",expanded"`
	SelfLink string                        `json:"selflink,omitempty" pretty:",expanded"`
}

type MemberStatsEntries struct {
	MemberNestedStats MemberStats `json:"nestedStats,omitempty"`
}

type MemberStats struct {
	Kind     string `json:"kind,omitempty" pretty:",expanded"`
	SelfLink string `json:"selflink,omitempty" pretty:",expanded"`
	Entries  struct {
		Addr struct {
			Description string `json:"description"`
		} `json:"addr,omitempty"`
		ConnqAgeEdm struct {
			Value int `json:"value"`
		} `json:"connq.ageEdm,omitempty"`
		ConnqAgeEma struct {
			Value int `json:"value"`
		} `json:"connq.ageEma,omitempty"`
		ConnqAgeHead struct {
			Value int `json:"value"`
		} `json:"connq.ageHead,omitempty"`
		ConnqAgeMax struct {
			Value int `json:"value"`
		} `json:"connq.ageMax,omitempty"`
		ConnqDepth struct {
			Value int `json:"value"`
		} `json:"connq.depth,omitempty"`
		ConnqServiced struct {
			Value int `json:"value"`
		} `json:"connq.serviced,omitempty"`
		CurSessions struct {
			Value int `json:"value"`
		} `json:"curSessions,omitempty"`
		Rule struct {
			Description string `json:"description,omitempty"`
		} `json:"Rule,omitempty"`
		Status struct {
			Description string `json:"description,omitempty"`
		} `json:"Status,omitempty"`
		NodeName struct {
			Description string `json:"description,omitempty"`
		} `json:"nodeName,omitempty"`
		PoolName struct {
			Description string `json:"description,omitempty"`
		} `json:"poolName,omitempty"`
		Port struct {
			Value int `json:"value"`
		} `json:"port,omitempty"`
		ServersideBitsIn struct {
			Value int `json:"value"`
		} `json:"serverside.bitsIn,omitempty"`
		ServersideBitsOut struct {
			Value int `json:"value"`
		} `json:"serverside.bitsOut,omitempty"`
		ServersideCurConns struct {
			Value int `json:"value"`
		} `json:"serverside.curConns,omitempty"`
		ServersideMaxConns struct {
			Value int `json:"value"`
		} `json:"serverside.maxConns,omitempty"`
		ServersidePktsIn struct {
			Value int `json:"value"`
		} `json:"serverside.pktsIn,omitempty"`
		ServersidePktsOut struct {
			Value int `json:"value"`
		} `json:"serverside.pktsOut,omitempty"`
		ServersideTotConns struct {
			Value int `json:"value"`
		} `json:"serverside.totConns,omitempty"`
		StatusAvailabilityState struct {
			Description string `json:"description,omitempty"`
		} `json:"status.availabilityState,omitempty"`
		StatusEnabledState struct {
			Description string `json:"description,omitempty"`
		} `json:"status.enabledState,omitempty"`
		StatusStatusReason struct {
			Description string `json:"description,omitempty"`
		} `json:"status.statusReason,omitempty"`
		TotRequests struct {
			Value int `json:"value"`
		} `json:"totRequests,omitempty"`
	} `json:"entries,omitempty"`
}

// Specify pool and member, get the specified member stats.
func (psr *PoolStatsResource) GetMemberStats(poolFullPathName, memberFullPathName string) (*MemberStatsList, error) {
	var msl MemberStatsList
	res, err := psr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(PoolEndpoint).ResourceInstance(poolFullPathName).SubResource(poolMembersEndpoint).
		SubResourceInstance(memberFullPathName).SubStatsResource(StatsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &msl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &msl, nil
}

type PoolAllMemberStatsList struct {
	Entries  map[string]PoolAllMemberStatsEntries `json:"entries,omitempty"`
	Kind     string                               `json:"kind,omitempty" pretty:",expanded"`
	SelfLink string                               `json:"selflink,omitempty" pretty:",expanded"`
}

type PoolAllMemberStatsEntries struct {
	NestedPoolAllMemberStats MemberStats `json:"nestedStats,omitempty"`
}

// Get the stats of all members in a pool.
func (psr *PoolStatsResource) GetPoolAllMemberStats(poolFullPathName string) (*PoolAllMemberStatsList, error) {
	var pams PoolAllMemberStatsList
	res, err := psr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(PoolEndpoint).ResourceInstance(poolFullPathName).SubResource(poolMembersEndpoint).
		SubStatsResource(StatsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &pams); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &pams, nil
}
