package ltm

import (
	"context"
	"encoding/json"
	"github.com/lefeck/go-bigip"
)

type VirtualAddressList struct {
	Items    []VirtualAddress `json:"items,omitempty"`
	Kind     string           `json:"kind,omitempty" pretty:",expanded"`
	SelfLink string           `json:"selfLink,omitempty" pretty:",expanded"`
}

type VirtualAddress struct {
	Kind                  string `json:"kind,omitempty"`
	Name                  string `json:"name,omitempty"`
	Partition             string `json:"partition,omitempty"`
	FullPath              string `json:"fullPath,omitempty"`
	Generation            int    `json:"generation,omitempty"`
	SelfLink              string `json:"selfLink,omitempty"`
	Address               string `json:"address,omitempty"`
	Arp                   string `json:"arp,omitempty"`
	AutoDelete            string `json:"autoDelete,omitempty"`
	ConnectionLimit       int    `json:"connectionLimit,omitempty"`
	Enabled               string `json:"enabled,omitempty"`
	Floating              string `json:"floating,omitempty"`
	IcmpEcho              string `json:"icmpEcho,omitempty"`
	InheritedTrafficGroup string `json:"inheritedTrafficG,omitemptyroup"`
	Mask                  string `json:"mask,omitempty"`
	RouteAdvertisement    string `json:"routeAdvertisemen,omitemptyt"`
	ServerScope           string `json:"serverScope,omitempty"`
	Spanning              string `json:"spanning,omitempty"`
	TrafficGroup          string `json:"trafficGroup,omitempty"`
	TrafficGroupReference struct {
		Link string `json:"link,omitempty"`
	} `json:"trafficGroupReference,omitempty"`
	Unit int `json:"unit,omitempty"`
}

// VirtualAddressEndpoint is the base path of the ltm API.
const VirtualAddressEndpoint = "virtual-address"

type VirtualAddressResource struct {
	b *bigip.BigIP
}

func (vars *VirtualAddressResource) List() (*VirtualAddressList, error) {
	var val VirtualAddressList
	res, err := vars.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(VirtualAddressEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &val); err != nil {
		panic(err)
	}
	return &val, nil
}

func (vars *VirtualAddressResource) GetAddressByVirtualServerName(fullPathName string) (string, error) {
	var va VirtualAddress
	res, err := vars.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(VirtualAddressEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(res, &va); err != nil {
		panic(err)
	}
	return va.Address, nil
}
