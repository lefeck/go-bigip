package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type NetflowList struct {
	Items    []Netflow `json:"items,omitempty"`
	Kind     string    `json:"kind,omitempty"`
	SelfLink string    `json:"selflink,omitempty"`
}

type Netflow struct {
	Kind           string `json:"kind,omitempty"`
	Name           string `json:"name,omitempty"`
	Partition      string `json:"partition,omitempty"`
	FullPath       string `json:"fullPath,omitempty"`
	Generation     int    `json:"generation,omitempty"`
	SelfLink       string `json:"selfLink,omitempty"`
	AppService     string `json:"appService,omitempty"`
	DefaultsFrom   string `json:"defaultsFrom,omitempty"`
	Description    string `json:"description,omitempty"`
	NetflowVersion string `json:"netflowVersion,omitempty"`
	SamplingRate   int    `json:"samplingRate,omitempty"`
}

const NetflowEndpoint = "netflow"

type NetflowResource struct {
	b *bigip.BigIP
}

// List retrieves a list of Netflow resources.
func (cr *NetflowResource) List() (*NetflowList, error) {
	var items NetflowList
	// Perform a GET request to retrieve a list of Netflow resource objects
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(NetflowEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into NetflowList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves a Netflow resource by its full path name.
func (cr *NetflowResource) Get(fullPathName string) (*Netflow, error) {
	var item Netflow
	// Perform a GET request to retrieve a specific Netflow resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(NetflowEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into Netflow struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new Netflow resource using the provided Netflow item.
func (cr *NetflowResource) Create(item Netflow) error {
	// Marshal the Netflow struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new Netflow resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(NetflowEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies a Netflow resource identified by its full path name using the provided Netflow item.
func (cr *NetflowResource) Update(fullPathName string, item Netflow) error {
	// Marshal the Netflow struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified Netflow resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(NetflowEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a Netflow resource by its full path name.
func (cr *NetflowResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified Netflow resource
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(NetflowEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
