package ltm

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/lefeck/bigip"
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

type VirtualAddressResource struct {
	c *bigip.BigIP
}

func (vasr *VirtualAddressResource) List() (*VirtualAddressList, error) {
	var val VirtualAddressList
	result, err := vasr.c.RestClient.Get().Prefix(BasePath).ManagerName(LTMManager).Resource(VirtualAddressEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(result)
	dec := json.NewDecoder(reader)
	if err := dec.Decode(&val); err != nil {
		return nil, err
	}

	return &val, nil
}

func (vasr *VirtualAddressResource) GetAddressByVirtualServerName(name string) (string, error) {
	var va VirtualAddress

	result, err := vasr.c.RestClient.Get().Prefix(BasePath).ManagerName(LTMManager).Resource(VirtualAddressEndpoint).Suffix(Suffix).ResourceNameFullPath(name).DoRaw(context.Background())
	if err != nil {
		return "", err
	}
	//fmt.Println(result)
	reader := bytes.NewReader(result)
	dec := json.NewDecoder(reader)
	if err := dec.Decode(&va); err != nil {
		return "", err
	}
	return va.Address, nil
}
