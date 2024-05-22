package net

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// A AddressList holds a list of Address.
type AddressList struct {
	Items    []Address `json:"items,omitempty"`
	Kind     string    `json:"kind,omitempty"`
	SelfLink string    `json:"selfLink,omitempty"`
}

// A Address hold the uration for a Address.
type Address struct {
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Partition  string `json:"partition"`
	FullPath   string `json:"fullPath"`
	Generation int    `json:"generation"`
	SelfLink   string `json:"selfLink"`
	Addresses  []struct {
		Name string `json:"name"`
	} `json:"addresses"`
}

const AddressEndpoint = "address-list"

type AddressResource struct {
	b *bigip.BigIP
}

// List lists all the address configurations.
func (ar *AddressResource) List() (*AddressList, error) {
	var al AddressList
	res, err := ar.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).
		Resource(AddressEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &al); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &al, nil
}

// Get a single address configuration identified by fullPathName.
func (ar *AddressResource) Get(fullPathName string) (*Address, error) {
	var address Address
	res, err := ar.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).
		Resource(AddressEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &address); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &address, nil
}

// Create a new address configuration.
func (ar *AddressResource) Create(item Address) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = ar.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).
		Resource(AddressEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Edit an address configuration identified by name.
func (ar *AddressResource) Update(name string, item Address) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = ar.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).
		Resource(AddressEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single address configuration identified by name.
func (ar *AddressResource) Delete(name string) error {
	_, err := ar.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).
		Resource(AddressEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
