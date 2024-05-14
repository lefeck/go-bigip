package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type POP3List struct {
	Items    []POP3 `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type POP3 struct {
	Kind           string `json:"kind"`
	Name           string `json:"name"`
	Partition      string `json:"partition"`
	FullPath       string `json:"fullPath"`
	Generation     int    `json:"generation"`
	SelfLink       string `json:"selfLink"`
	ActivationMode string `json:"activationMode"`
	AppService     string `json:"appService"`
	DefaultsFrom   string `json:"defaultsFrom"`
	Description    string `json:"description"`
}

const POP3Endpoint = "pop3"

type POP3Resource struct {
	b *bigip.BigIP
}

// List retrieves a list of POP3 resources.
func (cr *POP3Resource) List() (*POP3List, error) {
	var items POP3List
	// Perform a GET request to retrieve a list of POP3 resource objects
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(POP3Endpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into POP3List struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves a POP3 resource by its full path name.
func (cr *POP3Resource) Get(fullPathName string) (*POP3, error) {
	var item POP3
	// Perform a GET request to retrieve a specific POP3 resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(POP3Endpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into POP3 struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new POP3 resource using the provided POP3 item.
func (cr *POP3Resource) Create(item POP3) error {
	// Marshal the POP3 struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new POP3 resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(POP3Endpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies a POP3 resource identified by its full path name using the provided POP3 item.
func (cr *POP3Resource) Update(fullPathName string, item POP3) error {
	// Marshal the POP3 struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified POP3 resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(POP3Endpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a POP3 resource by its full path name.
func (cr *POP3Resource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified POP3 resource
	_, err := cr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(POP3Endpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
