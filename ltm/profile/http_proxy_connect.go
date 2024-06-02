package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type HTTPProxyConnectList struct {
	Items    []HTTPProxyConnect `json:"items,omitempty"`
	Kind     string             `json:"kind,omitempty"`
	SelfLink string             `json:"selflink,omitempty"`
}

type HTTPProxyConnect struct {
	Kind         string `json:"kind,omitempty"`
	Name         string `json:"name,omitempty"`
	Partition    string `json:"partition,omitempty"`
	FullPath     string `json:"fullPath,omitempty"`
	Generation   int    `json:"generation,omitempty"`
	SelfLink     string `json:"selfLink,omitempty"`
	AppService   string `json:"appService,omitempty"`
	DefaultState string `json:"defaultState,omitempty"`
	DefaultsFrom string `json:"defaultsFrom,omitempty"`
	Description  string `json:"description,omitempty"`
}

const HTTPProxyConnectEndpoint = "http-proxy-connect"

type HTTPProxyConnectResource struct {
	b *bigip.BigIP
}

// List retrieves a list of HTTPProxyConnect resources.
func (cr *HTTPProxyConnectResource) List() (*HTTPProxyConnectList, error) {
	var items HTTPProxyConnectList
	// Perform a GET request to retrieve a list of HTTPProxyConnect resource objects
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTPProxyConnectEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into HTTPProxyConnectList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves an HTTPProxyConnect resource by its full path name.
func (cr *HTTPProxyConnectResource) Get(fullPathName string) (*HTTPProxyConnect, error) {
	var item HTTPProxyConnect
	// Perform a GET request to retrieve a specific HTTPProxyConnect resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTPProxyConnectEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into HTTPProxyConnect struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new HTTPProxyConnect resource using the provided HTTPProxyConnect item.
func (cr *HTTPProxyConnectResource) Create(item HTTPProxyConnect) error {
	// Marshal the HTTPProxyConnect struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new HTTPProxyConnect resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTPProxyConnectEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies an HTTPProxyConnect resource identified by its full path name using the provided HTTPProxyConnect item.
func (cr *HTTPProxyConnectResource) Update(fullPathName string, item HTTPProxyConnect) error {
	// Marshal the HTTPProxyConnect struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified HTTPProxyConnect resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTPProxyConnectEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes an HTTPProxyConnect resource by its full path name.
func (cr *HTTPProxyConnectResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified HTTPProxyConnect resource
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTPProxyConnectEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
