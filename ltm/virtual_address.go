package ltm

import (
	"context"
	"fmt"
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

const (
	BasePath               = "/mgmt/tm/"
	LTMResource            = "ltm"
	VirtualAddressEndpoint = "virtual-address"
)

type VirtualAddressResource struct {
	c *bigip.BigIP
}

func (vsr *VirtualAddressResource) List() (*VirtualAddressList, error) {
	var vas VirtualAddressList

	//byResult, err := vsr.c.RestClient.Get().Prefix(BasePath).DoRaw(context.Background())
	byResult := vsr.c.RestClient.Get().Prefix(BasePath).Resource(LTMResource).SubResource(VirtualAddressEndpoint).Do(context.Background())
	//if err != nil {
	//	return nil, err
	//}
	fmt.Println(byResult)

	//reader := bytes.NewReader(byResult)
	//dec := json.NewDecoder(byResult)
	//if err := dec.Decode(&vas); err != nil {
	//	return nil, err
	//}

	return &vas, nil
}

//
//func (vsr *VirtualAddressResource) Lists() (*VirtualAddressList, error) {
//	var vsc VirtualAddressList
//
//	resp, err := vsr.doRequest("GET", "", nil)
//	if err != nil {
//		return nil, err
//	}
//	fmt.Println(resp.Header)
//	defer resp.Body.Close()
//	if err := vsr.readError(resp); err != nil {
//		return nil, err
//	}
//	//var vsc VirtualAddressResource
//	dec := json.NewDecoder(resp.Body)
//	if err := dec.Decode(&vsc); err != nil {
//		return nil, err
//	}
//
//	return &vsc, nil
//}

//func (vsr *VirtualAddressResource) GetAddressByVirtualServerName(name string) (string, error) {
//
//	var va VirtualAddress
//
//	resp, err := vsr.c.RestClient.Get().Prefix("/mgmt/tm").Resource("ltm").SubResource("virtual-address").Request(context.Background(), name)
//
//	if err != nil {
//		return "", err
//	}
//	dec := json.NewDecoder(resp.Body)
//	if err := dec.Decode(&va); err != nil {
//		return "", err
//	}
//
//	return va.Address, nil
//}
