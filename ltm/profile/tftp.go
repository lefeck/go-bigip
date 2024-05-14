package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type TFTPList struct {
	Items    []TFTP `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type TFTP struct {
	Kind         string `json:"kind"`
	Name         string `json:"name"`
	Partition    string `json:"partition"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	SelfLink     string `json:"selfLink"`
	AppService   string `json:"appService"`
	DefaultsFrom string `json:"defaultsFrom"`
	Description  string `json:"description"`
	IdleTimeout  string `json:"idleTimeout"`
	LogProfile   string `json:"logProfile"`
	LogPublisher string `json:"logPublisher"`
}

const TFTPEndpoint = "tcp"

type TFTPResource struct {
	b *bigip.BigIP
}

// List retrieves a list of TFTP resources.
func (cr *TFTPResource) List() (*TFTPList, error) {
	var items TFTPList
	// Perform a GET request to retrieve a list of TFTP resource objects
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(TFTPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into TFTPList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves a TFTP resource by its full path name.
func (cr *TFTPResource) Get(fullPathName string) (*TFTP, error) {
	var item TFTP
	// Perform a GET request to retrieve a specific TFTP resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(TFTPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into TFTP struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new TFTP resource using the provided TFTP item.
func (cr *TFTPResource) Create(item TFTP) error {
	// Marshal the TFTP struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new TFTP resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(TFTPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies a TFTP resource identified by its full path name using the provided TFTP item.
func (cr *TFTPResource) Update(fullPathName string, item TFTP) error {
	// Marshal the TFTP struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified TFTP resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(TFTPEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a TFTP resource by its full path name.
func (cr *TFTPResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified TFTP resource
	_, err := cr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(TFTPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
