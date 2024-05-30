package net

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// A PortList contains a list of Port.
type PortList struct {
	Items    []Port `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selfLink,omitempty"`
}

// Port contains for port attributes.
type Port struct {
	Kind       string       `json:"kind"`
	Name       string       `json:"name"`
	Partition  string       `json:"partition"`
	FullPath   string       `json:"fullPath"`
	Generation int          `json:"generation"`
	SelfLink   string       `json:"selfLink"`
	Ports      []PortMember `json:"ports"`
}

type PortMember struct {
	Name string `json:"name"`
}

// PortEndpoint represents the REST resource for managing port list.
const PortEndpoint = "port-list"

// PortResource provides an API to manage port list object.
type PortResource struct {
	b *bigip.BigIP
}

// List all the port configurations.
func (ar *PortResource) List() (*PortList, error) {
	var al PortList
	res, err := ar.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).
		Resource(PortEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &al); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &al, nil
}

// Get a single Port configuration identified by fullPathName.
func (ar *PortResource) Get(fullPathName string) (*Port, error) {
	var Port Port
	res, err := ar.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).
		Resource(PortEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &Port); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &Port, nil
}

// Create a new Port configuration.
func (ar *PortResource) Create(item Port) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = ar.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).
		Resource(PortEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Edit an Port configuration identified by name.
func (ar *PortResource) Update(name string, item Port) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = ar.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).
		Resource(PortEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single Port configuration identified by name.
func (ar *PortResource) Delete(name string) error {
	_, err := ar.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).
		Resource(PortEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
