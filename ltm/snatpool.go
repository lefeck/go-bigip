package ltm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
)

type SnatPoolList struct {
	Kind     string     `json:"kind"`
	SelfLink string     `json:"selfLink"`
	Items    []SnatPool `json:"items"`
}
type SnatPool struct {
	Kind             string   `json:"kind"`
	Name             string   `json:"name"`
	Partition        string   `json:"partition"`
	FullPath         string   `json:"fullPath"`
	Generation       int      `json:"generation"`
	SelfLink         string   `json:"selfLink"`
	Members          []string `json:"members"`
	MembersReference []struct {
		Link string `json:"link"`
	} `json:"membersReference"`
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
