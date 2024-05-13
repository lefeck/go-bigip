package ltm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
)

type VirtualAddressStatsList struct {
	Kind       string                              `json:"kind"`
	Generation int                                 `json:"generation"`
	SelfLink   string                              `json:"selfLink"`
	Entries    map[string]VirtualAddressStatsEntry `json:"entries"`
}

type VirtualAddressStatsEntry struct {
	NestedVirtualAddressStats NestedVirtualAddressStats `json:"nestedStats"`
}

type NestedVirtualAddressStats struct {
	Entries NestedVirtualAddressEntries `json:"entries"`
}

type NestedVirtualAddressEntries struct {
	Addr                    DescriptionOrValue `json:"addr"`
	ClientsideBitsIn        DescriptionOrValue `json:"clientside.bitsIn"`
	ClientsideBitsOut       DescriptionOrValue `json:"clientside.bitsOut"`
	ClientsideCurConns      DescriptionOrValue `json:"clientside.curConns"`
	ClientsideMaxConns      DescriptionOrValue `json:"clientside.maxConns"`
	ClientsidePktsIn        DescriptionOrValue `json:"clientside.pktsIn"`
	ClientsidePktsOut       DescriptionOrValue `json:"clientside.pktsOut"`
	ClientsideTotConns      DescriptionOrValue `json:"clientside.totConns"`
	TmName                  DescriptionOrValue `json:"tmName"`
	StatusAvailabilityState DescriptionOrValue `json:"status.availabilityState"`
	StatusEnabledState      DescriptionOrValue `json:"status.enabledState"`
	StatusStatusReason      DescriptionOrValue `json:"status.statusReason"`
}

type DescriptionOrValue struct {
	Description string `json:"description,omitempty"`
	Value       int    `json:"value,omitempty"`
}

// VirtualStatsResource provides an API to manage VirtualStats urations.
type VirtualAddressStatsResource struct {
	b *bigip.BigIP
}

func (vasr *VirtualAddressStatsResource) List() (*VirtualAddressStatsList, error) {
	res, err := vasr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(VirtualAddressEndpoint).SubStatsResource(StatsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	var vasl VirtualAddressStatsList
	if err := json.Unmarshal(res, &vasl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &vasl, nil
}

func (vasr *VirtualAddressStatsResource) Get(name string) (*VirtualAddressStatsList, error) {
	res, err := vasr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(VirtualAddressEndpoint).SubResourceInstance(name).SubStatsResource(StatsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	var vasl VirtualAddressStatsList
	if err := json.Unmarshal(res, &vasl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &vasl, nil
}
