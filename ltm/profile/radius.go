package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

//radius

type RADIUSList struct {
	Items    []RADIUS `json:"items,omitempty"`
	Kind     string   `json:"kind,omitempty"`
	SelfLink string   `json:"selflink,omitempty"`
}

type RADIUS struct {
	Kind                  string        `json:"kind"`
	Name                  string        `json:"name"`
	Partition             string        `json:"partition"`
	FullPath              string        `json:"fullPath"`
	Generation            int           `json:"generation"`
	SelfLink              string        `json:"selfLink"`
	AppService            string        `json:"appService"`
	Clients               []interface{} `json:"clients"`
	DefaultsFrom          string        `json:"defaultsFrom"`
	Description           string        `json:"description"`
	PersistAvp            string        `json:"persistAvp"`
	DefaultsFromReference struct {
		Link string `json:"link"`
	} `json:"defaultsFromReference,omitempty"`
}

const RADIUSEndpoint = "radius"

type RADIUSResource struct {
	b *bigip.BigIP
}

// List retrieves a list of RADIUS resources.
func (cr *RADIUSResource) List() (*RADIUSList, error) {
	var items RADIUSList
	// Perform a GET request to retrieve a list of RADIUS resource objects
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(RADIUSEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into RADIUSList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves a RADIUS resource by its full path name.
func (cr *RADIUSResource) Get(fullPathName string) (*RADIUS, error) {
	var item RADIUS
	// Perform a GET request to retrieve a specific RADIUS resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(RADIUSEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into RADIUS struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new RADIUS resource using the provided RADIUS item.
func (cr *RADIUSResource) Create(item RADIUS) error {
	// Marshal the RADIUS struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new RADIUS resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(RADIUSEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies a RADIUS resource identified by its full path name using the provided RADIUS item.
func (cr *RADIUSResource) Update(fullPathName string, item RADIUS) error {
	// Marshal the RADIUS struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified RADIUS resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(RADIUSEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a RADIUS resource by its full path name.
func (cr *RADIUSResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified RADIUS resource
	_, err := cr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(RADIUSEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
