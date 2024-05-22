package net

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// A TrunkList holds a list of Trunks.
type TrunkList struct {
	Items    []Trunk `json:"items,omitempty"`
	Kind     string  `json:"kind,omitempty"`
	SelfLink string  `json:"selfLink,omitempty"`
}

// A Trunk hold the uration for a trunk.
type Trunk struct {
	Bandwidth           int      `json:"bandwidth,omitempty"`
	CfgMbrCount         int      `json:"cfgMbrCount,omitempty"`
	DistributionHash    string   `json:"distributionHash,omitempty"`
	FullPath            string   `json:"fullPath,omitempty"`
	Generation          int      `json:"generation,omitempty"`
	ID                  int      `json:"id,omitempty"`
	Interfaces          []string `json:"interfaces,omitempty"`
	InterfacesReference []struct {
		Link string `json:"link,omitempty"`
	} `json:"interfacesReference,omitempty"`
	Kind             string `json:"kind,omitempty"`
	Lacp             string `json:"lacp,omitempty"`
	LacpMode         string `json:"lacpMode,omitempty"`
	LacpTimeout      string `json:"lacpTimeout,omitempty"`
	LinkSelectPolicy string `json:"linkSelectPolicy,omitempty"`
	MacAddress       string `json:"macAddress,omitempty"`
	Media            string `json:"media,omitempty"`
	Name             string `json:"name,omitempty"`
	QinqEthertype    string `json:"qinqEthertype,omitempty"`
	SelfLink         string `json:"selfLink,omitempty"`
	Stp              string `json:"stp,omitempty"`
	Type             string `json:"type,omitempty"`
	WorkingMbrCount  int    `json:"workingMbrCount,omitempty"`
}

// TrunkEndpoint represents the REST resource for managing a trunk.
const TrunkEndpoint = "trunk"

// A TrunkResource provides API to manage trunks uration.
type TrunkResource struct {
	b *bigip.BigIP
}

// ListAll lists all the trunk urations.
func (tr *TrunkResource) List() (*TrunkList, error) {
	var tl TrunkList
	res, err := tr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).
		Resource(TrunkEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &tl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &tl, nil
}

// Get a single trunk uration identified by id.
func (tr *TrunkResource) Get(fullPathName string) (*Trunk, error) {
	var trunk Trunk
	res, err := tr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).
		Resource(TrunkEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &trunk); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &trunk, nil
}

// Create a new trunk uration.
func (tr *TrunkResource) Create(item Trunk) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = tr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).
		Resource(TrunkEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Edit a trunk uration identified by id.
func (tr *TrunkResource) Edit(name string, item Trunk) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = tr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).
		Resource(TrunkEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single trunk uration identified by id.
func (tr *TrunkResource) Delete(name string) error {
	_, err := tr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).
		Resource(TrunkEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
