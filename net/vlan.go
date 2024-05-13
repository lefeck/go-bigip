package net

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// A VlanList holds a list of Vlan.
type VlanList struct {
	Items    []Vlan `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selfLink,omitempty"`
}

// A Vlan hold the uration for a vlan.
type Vlan struct {
	AutoLasthop         string `json:"autoLasthop,omitempty"`
	CmpHash             string `json:"cmpHash,omitempty"`
	DagRoundRobin       string `json:"dagRoundRobin,omitempty"`
	DagTunnel           string `json:"dagTunnel,omitempty"`
	Failsafe            string `json:"failsafe,omitempty"`
	FailsafeAction      string `json:"failsafeAction,omitempty"`
	FailsafeTimeout     int    `json:"failsafeTimeout,omitempty"`
	FullPath            string `json:"fullPath,omitempty"`
	Generation          int    `json:"generation,omitempty"`
	IfIndex             int    `json:"ifIndex,omitempty"`
	InterfacesReference struct {
		IsSubcollection bool   `json:"isSubcollection,omitempty"`
		Link            string `json:"link,omitempty"`
	} `json:"interfacesReference,omitempty"`
	Kind     string `json:"kind,omitempty"`
	Learning string `json:"learning,omitempty"`
	Mtu      int    `json:"mtu,omitempty"`
	Name     string `json:"name,omitempty"`
	SelfLink string `json:"selfLink,omitempty"`
	Sflow    struct {
		PollInterval       int    `json:"pollInterval,omitempty"`
		PollIntervalGlobal string `json:"pollIntervalGlobal,omitempty"`
		SamplingRate       int    `json:"samplingRate,omitempty"`
		SamplingRateGlobal string `json:"samplingRateGlobal,omitempty"`
	} `json:"sflow,omitempty"`
	SourceChecking string `json:"sourceChecking,omitempty"`
	Tag            int    `json:"tag,omitempty"`
}

type AssignedInterfaceList struct {
	Items    []AssignedInterface `json:"items,omitempty"`
	Kind     string              `json:"kind,omitempty"`
	SelfLink string              `json:"selfLink,omitempty"`
}

type AssignedInterface struct {
	FullPath   string `json:"fullPath,omitempty"`
	Generation int    `json:"generation,omitempty"`
	Kind       string `json:"kind,omitempty"`
	Name       string `json:"name,omitempty"`
	SelfLink   string `json:"selfLink,omitempty"`
	TagMode    string `json:"tagMode,omitempty"`
	Tagged     bool   `json:"tagged,omitempty"`
}

// VlanEndpoint represents the REST resource for managing a vlan.
const VlanEndpoint = "vlan"

const VlanInterfacesEndpoint = "interfaces"

// A VlanResource provides API to manage vlan uration.
type VlanResource struct {
	b *bigip.BigIP
}

// ListAll lists all the vlan urations.
func (vr *VlanResource) List() (*VlanList, error) {
	var vl VlanList
	res, err := vr.b.RestClient.Get().Prefix(bigip.BasePath).ResourceCategory(bigip.TMResource).ManagerName(NetManager).
		Resource(VlanEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &vl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &vl, nil
}

// Get a single vlan uration identified by name.
func (vr *VlanResource) Get(fullPathName string) (*Vlan, error) {
	var vlan Vlan
	res, err := vr.b.RestClient.Get().Prefix(bigip.BasePath).ResourceCategory(bigip.TMResource).ManagerName(NetManager).
		Resource(VlanEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &vlan); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &vlan, nil
}

// Create a new vlan uration.
func (vr *VlanResource) Create(item Vlan) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = vr.b.RestClient.Post().Prefix(bigip.BasePath).ResourceCategory(bigip.TMResource).ManagerName(NetManager).
		Resource(VlanEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Edit a vlan uration identified by id.
func (vr *VlanResource) Update(name string, item Vlan) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = vr.b.RestClient.Put().Prefix(bigip.BasePath).ResourceCategory(bigip.TMResource).ManagerName(NetManager).
		Resource(VlanEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single vlan uration identified by id.
func (vr *VlanResource) Delete(name string) error {
	_, err := vr.b.RestClient.Delete().Prefix(bigip.BasePath).ResourceCategory(bigip.TMResource).ManagerName(NetManager).
		Resource(VlanEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// GetInterfaces gets all interfaces associated to the vlan identified by id.
func (vr *VlanResource) GetVlanAssociatedInterfaces(name string) (*AssignedInterfaceList, error) {
	var ail AssignedInterfaceList
	//if err := vr.c.ReadQuery(BasePath+VlanEndpoint+"/"+id+"/interfaces", &list); err != nil {
	//	return nil, err
	//}
	res, err := vr.b.RestClient.Get().Prefix(bigip.BasePath).ResourceCategory(bigip.TMResource).ManagerName(NetManager).
		Resource(VlanEndpoint).ResourceInstance(name).SubResource(VlanInterfacesEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &ail); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &ail, nil
}

// Edit a vlan uration identified by id.
func (vr *VlanResource) AddInterfaceForVlan(name string, item AssignedInterface) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = vr.b.RestClient.Post().Prefix(bigip.BasePath).ResourceCategory(bigip.TMResource).ManagerName(NetManager).
		Resource(VlanEndpoint).ResourceInstance(name).SubResource(VlanInterfacesEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
