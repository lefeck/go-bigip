package net

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
)

// A InterfaceList holds a list of Interface.
type InterfaceList struct {
	Items    []Interface `json:"items,omitempty"`
	Kind     string      `json:"kind,omitempty"`
	SelfLink string      `json:"selfLink,omitempty"`
}

type Interface struct {
	Bundle                 string `json:"bundle,omitempty"`
	BundleSpeed            string `json:"bundleSpeed,omitempty"`
	Enabled                bool   `json:"enabled,omitempty"`
	FlowControl            string `json:"flowControl,omitempty"`
	ForceGigabitFiber      string `json:"forceGigabitFiber,omitempty"`
	ForwardErrorCorrection string `json:"forwardErrorCorrection,omitempty"`
	FullPath               string `json:"fullPath,omitempty"`
	Generation             int    `json:"generation,omitempty"`
	IfIndex                int    `json:"ifIndex,omitempty"`
	Kind                   string `json:"kind,omitempty"`
	LldpAdmin              string `json:"lldpAdmin,omitempty"`
	LldpTlvmap             int    `json:"lldpTlvmap,omitempty"`
	MacAddress             string `json:"macAddress,omitempty"`
	MediaActive            string `json:"mediaActive,omitempty"`
	MediaFixed             string `json:"mediaFixed,omitempty"`
	MediaSfp               string `json:"mediaSfp,omitempty"`
	Mtu                    int    `json:"mtu,omitempty"`
	Name                   string `json:"name,omitempty"`
	PreferPort             string `json:"preferPort,omitempty"`
	QinqEthertype          string `json:"qinqEthertype,omitempty"`
	SelfLink               string `json:"selfLink,omitempty"`
	Sflow                  struct {
		PollInterval       int    `json:"pollInterval,omitempty"`
		PollIntervalGlobal string `json:"pollIntervalGlobal,omitempty"`
	} `json:"sflow,omitempty"`
	Stp             string `json:"stp,omitempty"`
	StpAutoEdgePort string `json:"stpAutoEdgePort,omitempty"`
	StpEdgePort     string `json:"stpEdgePort,omitempty"`
	StpLinkType     string `json:"stpLinkType,omitempty"`
}

// InterfaceEndpoint represents the REST resource for managing interfaces.
const InterfaceEndpoint = "interface"

// A InetResource provides an API to manage Interface urations.
type InetResource struct {
	b *bigip.BigIP
}

// ListAll lists all interfaces uration.
func (ir *InetResource) List() (*InterfaceList, error) {
	var items InterfaceList
	res, err := ir.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).Resource(InterfaceEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get a single interface uration identified by id.
func (ir *InetResource) Get(fullPathName string) (*Interface, error) {
	var item Interface
	res, err := ir.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).
		Resource(InterfaceEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

func (ir *InetResource) ShowStats(fullPathName string) (*InterfaceStatsList, error) {
	var item InterfaceStatsList
	res, err := ir.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).
		Resource(InterfaceEndpoint).SubResourceInstance(fullPathName).SubStatsResource(StatsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}
