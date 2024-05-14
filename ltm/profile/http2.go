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
	Kind                           string   `json:"kind"`
	Name                           string   `json:"name"`
	Partition                      string   `json:"partition"`
	FullPath                       string   `json:"fullPath"`
	Generation                     int      `json:"generation"`
	SelfLink                       string   `json:"selfLink"`
	ActivationModes                []string `json:"activationModes"`
	AppService                     string   `json:"appService"`
	ConcurrentStreamsPerConnection int      `json:"concurrentStreamsPerConnection"`
	ConnectionIdleTimeout          int      `json:"connectionIdleTimeout"`
	DefaultsFrom                   string   `json:"defaultsFrom"`
	Description                    string   `json:"description"`
	EnforceTLSRequirements         string   `json:"enforceTlsRequirements"`
	FrameSize                      int      `json:"frameSize"`
	HeaderTableSize                int      `json:"headerTableSize"`
	IncludeContentLength           string   `json:"includeContentLength"`
	InsertHeader                   string   `json:"insertHeader"`
	InsertHeaderName               string   `json:"insertHeaderName"`
	ReceiveWindow                  int      `json:"receiveWindow"`
	WriteSize                      int      `json:"writeSize"`
}

const HTTP2Endpoint = "http2"

type HTTP2Resource struct {
	b *bigip.BigIP
}

// 把下面这段代码中HTTP都改成HTTP2， 代码如下：
// List retrieves a list of HTTP2 resources.
func (cr *HTTP2Resource) List() (*HTTP2List, error) {
	var items HTTP2List
	// Perform a GET request to retrieve a list of HTTP2 resource objects
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
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
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
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
	_, err = cr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
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
	_, err = cr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTP2Endpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes an HTTP2 resource by its full path name.
func (cr *HTTP2Resource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified HTTP2 resource
	_, err := cr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTP2Endpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
