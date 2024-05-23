package ltm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// VirtualAddressList is a list containing multiple VirtualAddress objects.
type VirtualAddressList struct {
	Items    []VirtualAddress `json:"items,omitempty"`
	Kind     string           `json:"kind,omitempty" pretty:",expanded"`
	SelfLink string           `json:"selfLink,omitempty" pretty:",expanded"`
}

// VirtualAddress represents an F5 BIG-IP LTM virtual address configuration.
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

// VirtualAddressResource is used to interact with the virtual address API.
type VirtualAddressResource struct {
	b *bigip.BigIP
}

// List retrieves the list of all virtual addresses configured on the BIG-IP.
func (vars *VirtualAddressResource) List() (*VirtualAddressList, error) {
	var val VirtualAddressList
	res, err := vars.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(VirtualAddressEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &val); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &val, nil
}

// Get retrieves a single virtual address configuration identified by fullPathName.
func (vars *VirtualAddressResource) Get(fullPathName string) (*VirtualAddress, error) {
	var va VirtualAddress
	res, err := vars.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(VirtualAddressEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &va); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &va, nil
}

// GetAddressByVirtualServerName retrieves the IP address for a given virtual server identified by fullPathName.
func (vars *VirtualAddressResource) GetAddressByVirtualServerName(fullPathName string) (string, error) {
	var va VirtualAddress
	res, err := vars.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(VirtualAddressEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(res, &va); err != nil {
		return "", fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return va.Address, nil
}

// Update modifies a virtual address configuration identified by name with the given VirtualAddress object.
func (vr *VirtualAddressResource) Update(name string, item VirtualAddress) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = vr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(VirtualAddressEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single virtual address identified by the virtual address name. if it is not exist return error
func (vr *VirtualAddressResource) Delete(name string) error {
	_, err := vr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(VirtualAddressEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Enabling a virtual address item identified by the virtual address.
func (vr *VirtualAddressResource) Enable(name string) error {
	item := VirtualAddress{Enabled: "yes"}
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = vr.b.RestClient.Patch().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(VirtualAddressEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Disabling a virtual address item identified by the virtual address.
func (vr *VirtualAddressResource) Disable(name string) error {
	item := VirtualAddress{Enabled: "no"}

	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = vr.b.RestClient.Patch().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(VirtualAddressEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
