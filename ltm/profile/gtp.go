package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type GTPList struct {
	Items    []GTP  `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type GTP struct {
	Kind         string `json:"kind"`
	Name         string `json:"name"`
	Partition    string `json:"partition"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	SelfLink     string `json:"selfLink"`
	AppService   string `json:"appService"`
	DefaultsFrom string `json:"defaultsFrom"`
	Description  string `json:"description"`
	IngressMax   int    `json:"ingressMax"`
}

const GTPEndpoint = "gtp"

type GTPResource struct {
	b *bigip.BigIP
}

// List retrieves a list of GTP resources.
func (cr *GTPResource) List() (*GTPList, error) {
	var items GTPList
	// Perform a GET request to retrieve a list of GTP resource objects
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(GTPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into GTPList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves a GTP resource by its full path name.
func (cr *GTPResource) Get(fullPathName string) (*GTP, error) {
	var item GTP
	// Perform a GET request to retrieve a specific GTP resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(GTPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into GTP struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new GTP resource using the provided GTP item.
func (cr *GTPResource) Create(item GTP) error {
	// Marshal the GTP struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new GTP resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(GTPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies a GTP resource identified by its full path name using the provided GTP item.
func (cr *GTPResource) Update(fullPathName string, item GTP) error {
	// Marshal the GTP struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified GTP resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(GTPEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a GTP resource by its full path name.
func (cr *GTPResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified GTP resource
	_, err := cr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(GTPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
