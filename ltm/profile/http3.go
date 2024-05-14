package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type HTTP3List struct {
	Items    []HTTP3 `json:"items,omitempty"`
	Kind     string  `json:"kind,omitempty"`
	SelfLink string  `json:"selflink,omitempty"`
}

type HTTP3 struct {
	Kind            string `json:"kind"`
	Name            string `json:"name"`
	Partition       string `json:"partition"`
	FullPath        string `json:"fullPath"`
	Generation      int    `json:"generation"`
	SelfLink        string `json:"selfLink"`
	AppService      string `json:"appService"`
	DefaultsFrom    string `json:"defaultsFrom"`
	Description     string `json:"description"`
	HeaderTableSize int    `json:"headerTableSize"`
}

const HTTP3Endpoint = "http3"

type HTTP3Resource struct {
	b *bigip.BigIP
}

// List retrieves a list of HTTP3 resources.
func (cr *HTTP3Resource) List() (*HTTP3List, error) {
	var items HTTP3List
	// Perform a GET request to retrieve a list of HTTP3 resource objects
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTP3Endpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into HTTP3List struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves an HTTP3 resource by its full path name.
func (cr *HTTP3Resource) Get(fullPathName string) (*HTTP3, error) {
	var item HTTP3
	// Perform a GET request to retrieve a specific HTTP3 resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTP3Endpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into HTTP3 struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new HTTP3 resource using the provided HTTP3 item.
func (cr *HTTP3Resource) Create(item HTTP3) error {
	// Marshal the HTTP3 struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new HTTP3 resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTP3Endpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies an HTTP3 resource identified by its full path name using the provided HTTP3 item.
func (cr *HTTP3Resource) Update(fullPathName string, item HTTP3) error {
	// Marshal the HTTP3 struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified HTTP3 resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTP3Endpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes an HTTP3 resource by its full path name.
func (cr *HTTP3Resource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified HTTP3 resource
	_, err := cr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTP3Endpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
