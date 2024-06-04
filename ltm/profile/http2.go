package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type HTTP2List struct {
	Items    []HTTP2 `json:"items,omitempty"`
	Kind     string  `json:"kind,omitempty"`
	SelfLink string  `json:"selflink,omitempty"`
}

type HTTP2 struct {
	Kind                           string   `json:"kind,omitempty"`
	Name                           string   `json:"name,omitempty"`
	Partition                      string   `json:"partition,omitempty"`
	FullPath                       string   `json:"fullPath,omitempty"`
	Generation                     int      `json:"generation,omitempty"`
	SelfLink                       string   `json:"selfLink,omitempty"`
	ActivationModes                []string `json:"activationModes,omitempty"`
	AppService                     string   `json:"appService,omitempty"`
	ConcurrentStreamsPerConnection int      `json:"concurrentStreamsPerConnection,omitempty"`
	ConnectionIdleTimeout          int      `json:"connectionIdleTimeout,omitempty"`
	DefaultsFrom                   string   `json:"defaultsFrom,omitempty"`
	Description                    string   `json:"description,omitempty"`
	EnforceTLSRequirements         string   `json:"enforceTlsRequirements,omitempty"`
	FrameSize                      int      `json:"frameSize,omitempty"`
	HeaderTableSize                int      `json:"headerTableSize,omitempty"`
	IncludeContentLength           string   `json:"includeContentLength,omitempty"`
	InsertHeader                   string   `json:"insertHeader,omitempty"`
	InsertHeaderName               string   `json:"insertHeaderName,omitempty"`
	ReceiveWindow                  int      `json:"receiveWindow,omitempty"`
	WriteSize                      int      `json:"writeSize,omitempty"`
}

const HTTP2Endpoint = "http2"

type HTTP2Resource struct {
	b *bigip.BigIP
}

// List retrieves a list of HTTP2 resources.
func (cr *HTTP2Resource) List() (*HTTP2List, error) {
	var items HTTP2List
	// Perform a GET request to retrieve a list of HTTP2 resource objects
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTP2Endpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into HTTP2List struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves an HTTP2 resource by its full path name.
func (cr *HTTP2Resource) Get(fullPathName string) (*HTTP2, error) {
	var item HTTP2
	// Perform a GET request to retrieve a specific HTTP2 resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTP2Endpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into HTTP2 struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new HTTP2 resource using the provided HTTP2 item.
func (cr *HTTP2Resource) Create(item HTTP2) error {
	// Marshal the HTTP2 struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new HTTP2 resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTP2Endpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies an HTTP2 resource identified by its full path name using the provided HTTP2 item.
func (cr *HTTP2Resource) Update(fullPathName string, item HTTP2) error {
	// Marshal the HTTP2 struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified HTTP2 resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTP2Endpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes an HTTP2 resource by its full path name.
func (cr *HTTP2Resource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified HTTP2 resource
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTP2Endpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
