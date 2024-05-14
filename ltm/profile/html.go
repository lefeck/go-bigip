package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type HTMLList struct {
	Items    []HTML `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type HTML struct {
	Kind             string   `json:"kind"`
	Name             string   `json:"name"`
	Partition        string   `json:"partition"`
	FullPath         string   `json:"fullPath"`
	Generation       int      `json:"generation"`
	SelfLink         string   `json:"selfLink"`
	AppService       string   `json:"appService"`
	ContentDetection string   `json:"contentDetection"`
	ContentSelection []string `json:"contentSelection"`
	DefaultsFrom     string   `json:"defaultsFrom"`
	Description      string   `json:"description"`
}

const HTMLEndpoint = "html"

type HTMLResource struct {
	b *bigip.BigIP
}

// List retrieves a list of HTML resources.
func (cr *HTMLResource) List() (*HTMLList, error) {
	var items HTMLList
	// Perform a GET request to retrieve a list of HTML resource objects
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTMLEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into HTMLList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves an HTML resource by its full path name.
func (cr *HTMLResource) Get(fullPathName string) (*HTML, error) {
	var item HTML
	// Perform a GET request to retrieve a specific HTML resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTMLEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into HTML struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new HTML resource using the provided HTML item.
func (cr *HTMLResource) Create(item HTML) error {
	// Marshal the HTML struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new HTML resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTMLEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies an HTML resource identified by its full path name using the provided HTML item.
func (cr *HTMLResource) Update(fullPathName string, item HTML) error {
	// Marshal the HTML struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified HTML resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTMLEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes an HTML resource by its full path name.
func (cr *HTMLResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified HTML resource
	_, err := cr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTMLEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
