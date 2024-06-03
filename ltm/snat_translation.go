package ltm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// snat-translation - Configures an explicit secure network address translation (SNAT) translation address.
// SnatTranslationList is a list contains multiple SnatTranslation objects.
type SnatTranslationList struct {
	Kind     string            `json:"kind,omitempty"`
	SelfLink string            `json:"selfLink,omitempty"`
	Items    []SnatTranslation `json:"items,omitempty"`
}

// SnatTranslation represents an F5 BIG-IP LTM SnatTranslation configuration.
type SnatTranslation struct {
	Kind                  string `json:"kind,omitempty"`
	Name                  string `json:"name,omitempty"`
	Partition             string `json:"partition,omitempty"`
	FullPath              string `json:"fullPath,omitempty"`
	Generation            int    `json:"generation,omitempty"`
	SelfLink              string `json:"selfLink,omitempty"`
	Address               string `json:"address,omitempty"`
	Arp                   string `json:"arp,omitempty"`
	ConnectionLimit       int    `json:"connectionLimit,omitempty"`
	Enabled               bool   `json:"enabled,omitempty"`
	Disabled              bool   `json:"disabled,omitempty"`
	InheritedTrafficGroup string `json:"inheritedTrafficGroup,omitempty"`
	IPIdleTimeout         string `json:"ipIdleTimeout,omitempty"`
	TCPIdleTimeout        string `json:"tcpIdleTimeout,omitempty"`
	TrafficGroup          string `json:"trafficGroup,omitempty"`
	TrafficGroupReference struct {
		Link string `json:"link,omitempty"`
	} `json:"trafficGroupReference,omitempty"`
	UDPIdleTimeout string `json:"udpIdleTimeout,omitempty"`
	Unit           int    `json:"unit,omitempty"`
}

// SnatTranslationResource provides an API to manage SnatTranslation object.
type SnatTranslationResource struct {
	b *bigip.BigIP
}

// SnatTranslationEndpoint represents the REST resource for managing SnatTranslation.
const SnatTranslationEndpoint = "snat-translation"

// List all the SnatTranslationstate instances.
func (str *SnatTranslationResource) List() (*SnatTranslationList, error) {
	res, err := str.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(SnatTranslationEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	var items SnatTranslationList
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get a single SnatTranslation identified by name.
func (str *SnatTranslationResource) Get(fullPathName string) (*SnatTranslation, error) {
	res, err := str.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(SnatTranslationEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	var item SnatTranslation
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create a new SnatTranslation instance.
func (str *SnatTranslationResource) Create(item SnatTranslation) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = str.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(SnatTranslationEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update a SnatTranslation instance identified by name.
func (str *SnatTranslationResource) Update(fullPathName string, item SnatTranslation) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = str.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(SnatTranslationEndpoint).ResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single SnatTranslation instance identified by name.
func (str *SnatTranslationResource) Delete(fullPathName string) error {
	_, err := str.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(SnatTranslationEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Enabling a SnatTranslation item identified by the SnatTranslation name.
func (str *SnatTranslationResource) Enable(name string) error {
	item := SnatTranslation{Enabled: true}
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = str.b.RestClient.Patch().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(SnatTranslationEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Disabling a SnatTranslation item identified by the SnatTranslationname.
func (str *SnatTranslationResource) Disable(name string) error {
	item := SnatTranslation{Disabled: true}
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = str.b.RestClient.Patch().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(SnatTranslationEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
