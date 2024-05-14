package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type TDRList struct {
	Items    []TDR  `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type TDR struct {
	Kind                  string `json:"kind"`
	Name                  string `json:"name"`
	Partition             string `json:"partition"`
	FullPath              string `json:"fullPath"`
	Generation            int    `json:"generation"`
	SelfLink              string `json:"selfLink"`
	AppService            string `json:"appService"`
	DefaultsFrom          string `json:"defaultsFrom"`
	Description           string `json:"description"`
	LogPublisher          string `json:"logPublisher"`
	LogPublisherReference struct {
		Link string `json:"link"`
	} `json:"logPublisherReference"`
	FiltersReference struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"filtersReference"`
	DefaultsFromReference struct {
		Link string `json:"link"`
	} `json:"defaultsFromReference,omitempty"`
}

const TDREndpoint = "tdr"

type TDRResource struct {
	b *bigip.BigIP
}

// List retrieves a list of TDR resources.
func (cr *TDRResource) List() (*TDRList, error) {
	var items TDRList
	// Perform a GET request to retrieve a list of TDR resource objects
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(TDREndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into TDRList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves a TDR resource by its full path name.
func (cr *TDRResource) Get(fullPathName string) (*TDR, error) {
	var item TDR
	// Perform a GET request to retrieve a specific TDR resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(TDREndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into TDR struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new TDR resource using the provided TDR item.
func (cr *TDRResource) Create(item TDR) error {
	// Marshal the TDR struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new TDR resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(TDREndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies a TDR resource identified by its full path name using the provided TDR item.
func (cr *TDRResource) Update(fullPathName string, item TDR) error {
	// Marshal the TDR struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified TDR resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(TDREndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a TDR resource by its full path name.
func (cr *TDRResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified TDR resource
	_, err := cr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(TDREndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
