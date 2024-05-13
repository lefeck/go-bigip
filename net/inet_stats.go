package net

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
)

type InterfaceStatsList struct {
	Entries  map[string]InterfaceStatsEntries `json:"entries,omitempty"`
	Kind     string                           `json:"kind,omitempty"`
	SelfLink string                           `json:"selflink,omitempty"`
}

type InterfaceStatsEntries struct {
	NestedInterfaceStats InterfaceStats `json:"nestedStats,omitempty"`
}

type InterfaceStats struct {
	Entries struct {
		CountersBitsIn struct {
			Value int `json:"value"`
		} `json:"counters.bitsIn,omitempty"`
		CountersBitsOut struct {
			Value int `json:"value"`
		} `json:"counters.bitsOut,omitempty"`
		CountersDropsAll struct {
			Value int `json:"value"`
		} `json:"counters.dropsAll,omitempty"`
		CountersErrorsAll struct {
			Value int `json:"value"`
		} `json:"counters.errorsAll,omitempty"`
		CountersPktsIn struct {
			Value int `json:"value"`
		} `json:"counters.pktsIn,omitempty"`
		CountersPktsOut struct {
			Value int `json:"value"`
		} `json:"counters.pktsOut,omitempty"`
		MediaActive struct {
			Description string `json:"description,omitempty"`
		} `json:"mediaActive,omitempty"`
		Status struct {
			Description string `json:"description,omitempty"`
		} `json:"status,omitempty"`
		TmName struct {
			Description string `json:"description,omitempty"`
		} `json:"tmName,omitempty"`
	} `json:"entries,omitempty"`
}

// InterfaceStatsEndpoint represents the REST resource for managing InterfaceStats.
const InterfaceStatsEndpoint = "interface"
const StatsEndpoint = "stats"

// InterfaceStatsResource provides an API to manage InterfaceStats urations.
type InetStatsResource struct {
	b *bigip.BigIP
}

func (isr *InetStatsResource) List() (*InterfaceStatsList, error) {
	var item InterfaceStatsList
	res, err := isr.b.RestClient.Get().Prefix(bigip.BasePath).ResourceCategory(bigip.TMResource).ManagerName(NetManager).Resource(InterfaceStatsEndpoint).SubResource(StatsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}
