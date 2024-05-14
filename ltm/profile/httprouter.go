package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type HTTPRouterList struct {
	Items    []HTTPRouter `json:"items,omitempty"`
	Kind     string       `json:"kind,omitempty"`
	SelfLink string       `json:"selflink,omitempty"`
}

type HTTPRouter struct {
	Kind         string `json:"kind"`
	Name         string `json:"name"`
	Partition    string `json:"partition"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	SelfLink     string `json:"selfLink"`
	AppService   string `json:"appService"`
	DefaultsFrom string `json:"defaultsFrom"`
	Description  string `json:"description"`
}

const HTTPRouterEndpoint = "httprouter"

type HTTPRouterResource struct {
	b *bigip.BigIP
}

// List retrieves a list of HTTPRouter resources.
func (cr *HTTPRouterResource) List() (*HTTPRouterList, error) {
	var items HTTPRouterList
	// Perform a GET request to retrieve a list of HTTPRouter resource objects
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTPRouterEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into HTTPRouterList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves an HTTPRouter resource by its full path name.
func (cr *HTTPRouterResource) Get(fullPathName string) (*HTTPRouter, error) {
	var item HTTPRouter
	// Perform a GET request to retrieve a specific HTTPRouter resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTPRouterEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into HTTPRouter struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new HTTPRouter resource using the provided HTTPRouter item.
func (cr *HTTPRouterResource) Create(item HTTPRouter) error {
	// Marshal the HTTPRouter struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new HTTPRouter resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTPRouterEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies an HTTPRouter resource identified by its full path name using the provided HTTPRouter item.
func (cr *HTTPRouterResource) Update(fullPathName string, item HTTPRouter) error {
	// Marshal the HTTPRouter struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified HTTPRouter resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTPRouterEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes an HTTPRouter resource by its full path name.
func (cr *HTTPRouterResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified HTTPRouter resource
	_, err := cr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTPRouterEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
