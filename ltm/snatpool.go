package ltm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// SnatPoolList is a list contains multiple SnatPool objects.
type SnatPoolList struct {
	Kind     string     `json:"kind,omitempty"`
	SelfLink string     `json:"selfLink,omitempty"`
	Items    []SnatPool `json:"items,omitempty"`
}

// SnatPool represents an F5 BIG-IP LTM SnatPool configuration.
type SnatPool struct {
	Kind             string   `json:"kind,omitempty"`
	Name             string   `json:"name,omitempty"`
	Partition        string   `json:"partition,omitempty"`
	FullPath         string   `json:"fullPath,omitempty"`
	Generation       int      `json:"generation,omitempty"`
	SelfLink         string   `json:"selfLink,omitempty"`
	Members          []string `json:"members,omitempty"`
	MembersReference []struct {
		Link string `json:"link,omitempty"`
	} `json:"membersReference,omitempty"`
}

// SnatPoolResource provides an API to manage SnatPool object.
type SnatPoolResource struct {
	b *bigip.BigIP
}

// SnatPoolEndpoint represents the REST resource for managing SnatPool.
const SnatPoolEndpoint = "snatpool"

// List all the snatpool instances.
func (spr *SnatPoolResource) List() (*SnatPoolList, error) {
	res, err := spr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(SnatPoolEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	var spl SnatPoolList
	if err := json.Unmarshal(res, &spl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &spl, nil
}

// Get a single snatpool identified by name.
func (spr *SnatPoolResource) Get(fullPathName string) (*SnatPool, error) {
	res, err := spr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(SnatPoolEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	var sp SnatPool
	if err := json.Unmarshal(res, &sp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &sp, nil
}

// Create a new snatpool instance.
func (pr *SnatPoolResource) Create(item SnatPool) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = pr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(SnatPoolEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update a snatpool instance identified by name.
func (pr *SnatPoolResource) Update(fullPathName string, item SnatPool) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = pr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(SnatPoolEndpoint).ResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single pool instance identified by name.
func (pr *SnatPoolResource) Delete(fullPathName string) error {
	_, err := pr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(SnatPoolEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
