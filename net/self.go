package net

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// A SelfList holds a list of Self.
type SelfList struct {
	Items    []Self `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selfLink,omitempty"`
}

// A Self hold the uration for a Self IP.
type Self struct {
	Address               string `json:"address,omitempty"`
	AddressSource         string `json:"addressSource,omitempty"`
	Floating              string `json:"floating,omitempty"`
	FullPath              string `json:"fullPath,omitempty"`
	Generation            int    `json:"generation,omitempty"`
	InheritedTrafficGroup string `json:"inheritedTrafficGroup,omitempty"`
	Kind                  string `json:"kind,omitempty"`
	Name                  string `json:"name,omitempty"`
	SelfLink              string `json:"selfLink,omitempty"`
	TrafficGroup          string `json:"trafficGroup,omitempty"`
	TrafficGroupReference struct {
		Link string `json:"link,omitempty"`
	} `json:"trafficGroupReference,omitempty"`
	Unit          int    `json:"unit,omitempty"`
	Vlan          string `json:"vlan,omitempty"`
	VlanReference struct {
		Link string `json:"link,omitempty"`
	} `json:"vlanReference,omitempty"`
}

// SelfEndpoint represents the REST resource for managing a self IP.
const SelfEndpoint = "self"

// A SelfResource provides API to manage self ip uration.
type SelfResource struct {
	b *bigip.BigIP
}

// ListAll lists all the self ip urations.
func (sr *SelfResource) List() (*SelfList, error) {
	var sl SelfList
	res, err := sr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).
		Resource(SelfEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &sl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &sl, nil
}

// Get a single self ip uration identified by id.
func (sr *SelfResource) Get(fullPathName string) (*Self, error) {
	var self Self
	res, err := sr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).
		Resource(SelfEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &self); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &self, nil
}

// Create a new self ip uration.
func (sr *SelfResource) Create(item Self) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = sr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).
		Resource(SelfEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Edit a self ip uration identified by id.
func (sr *SelfResource) Edit(name string, item Self) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = sr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).
		Resource(SelfEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single self ip uration identified by id.
func (sr *SelfResource) Delete(name string) error {
	_, err := sr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).
		Resource(SelfEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
