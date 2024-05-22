package ltm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
)

type NodeStatsList struct {
	Kind     string               `json:"kind"`
	SelfLink string               `json:"selfLink"`
	Entries  map[string]NodeEntry `json:"entries"`
}

type NodeEntry struct {
	NestedStats NestedNodeStats `json:"nestedStats"`
}

type NestedNodeStats struct {
	Kind     string         `json:"kind"`
	SelfLink string         `json:"selfLink"`
	Entries  NodeStatsEntry `json:"entries"`
}

type NodeStatsEntry struct {
	Addr                    DescriptionOrValue `json:"addr"`
	CurSessions             DescriptionOrValue `json:"curSessions"`
	Rule                    DescriptionOrValue `json:"Rule"`
	Status                  DescriptionOrValue `json:"Status"`
	TmName                  DescriptionOrValue `json:"tmName"`
	ServersideBitsIn        DescriptionOrValue `json:"serverside.bitsIn"`
	ServersideBitsOut       DescriptionOrValue `json:"serverside.bitsOut"`
	ServersideCurConns      DescriptionOrValue `json:"serverside.curConns"`
	ServersideMaxConns      DescriptionOrValue `json:"serverside.maxConns"`
	ServersidePktsIn        DescriptionOrValue `json:"serverside.pktsIn"`
	ServersidePktsOut       DescriptionOrValue `json:"serverside.pktsOut"`
	ServersideTotConns      DescriptionOrValue `json:"serverside.totConns"`
	SessionStatus           DescriptionOrValue `json:"sessionStatus"`
	StatusAvailabilityState DescriptionOrValue `json:"status.availabilityState"`
	StatusEnabledState      DescriptionOrValue `json:"status.enabledState"`
	StatusStatusReason      DescriptionOrValue `json:"status.statusReason"`
	TotRequests             DescriptionOrValue `json:"totRequests"`
}

// NodeStatsResource provides an API to manage NodeStats entries.
type NodeStatsResource struct {
	b *bigip.BigIP
}

func (nsr *NodeStatsResource) List() (*NodeStatsList, error) {
	var nsl NodeStatsList
	res, err := nsr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(NodeEndpoint).SubStatsResource(StatsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &nsl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &nsl, nil
}
