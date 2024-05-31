package ltm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
)

type SnatPoolList struct {
	Kind     string     `json:"kind,omitempty"`
	SelfLink string     `json:"selfLink,omitempty"`
	Items    []SnatPool `json:"items,omitempty"`
}

type SnatPool struct {
	Kind             string   `json:"kind,omitempty"`
	Name             string   `json:"name,omitempty"`
	Partition        string   `json:"partition,omitempty"`
	FullPath         string   `json:"fullPath,omitempty"`
	Generation       int      `json:"generation,omitempty"`
	SelfLink         string   `json:"selfLink,omitempty"`
	Members          []string `json:"members,omitempty"`
	MembersReference []struct {
		Link string `json:"link,omitempty"`
	} `json:"membersReference,omitempty"`
}

type SnatPoolResource struct {
	b *bigip.BigIP
}

// VirtualEndpoint is the base path of the ltm API.
const SnatPoolEndpoint = "snatpool"

func (spr *SnatPoolResource) List() (*SnatPoolList, error) {
	res, err := spr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(SnatPoolEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	var spl SnatPoolList
	if err := json.Unmarshal(res, &spl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &spl, nil
}

// https://192.168.13.91/mgmt/tm/ltm/snatpool/~Common~snat_pool_yw?expandSubcollections=true
func (spr *SnatPoolResource) Get(name string) (*SnatPool, error) {
	res, err := spr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(SnatPoolEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	var sp SnatPool
	if err := json.Unmarshal(res, &sp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &sp, nil
}
