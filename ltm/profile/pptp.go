package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type PPTPList struct {
	Items    []PPTP `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type PPTP struct {
	Kind                 string `json:"kind"`
	Name                 string `json:"name"`
	Partition            string `json:"partition"`
	FullPath             string `json:"fullPath"`
	Generation           int    `json:"generation"`
	SelfLink             string `json:"selfLink"`
	AppService           string `json:"appService"`
	CsvFormat            string `json:"csvFormat"`
	DefaultsFrom         string `json:"defaultsFrom"`
	Description          string `json:"description"`
	IncludeDestinationIP string `json:"includeDestinationIp"`
	PublisherName        string `json:"publisherName"`
}

const PPTPEndpoint = "pptp"

type PPTPResource struct {
	b *bigip.BigIP
}

// List retrieves a list of PPTP resources.
func (cr *PPTPResource) List() (*PPTPList, error) {
	var items PPTPList
	// Perform a GET request to retrieve a list of PPTP resource objects
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(PPTPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into PPTPList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves a PPTP resource by its full path name.
func (cr *PPTPResource) Get(fullPathName string) (*PPTP, error) {
	var item PPTP
	// Perform a GET request to retrieve a specific PPTP resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(PPTPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into PPTP struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new PPTP resource using the provided PPTP item.
func (cr *PPTPResource) Create(item PPTP) error {
	// Marshal the PPTP struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new PPTP resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(PPTPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies a PPTP resource identified by its full path name using the provided PPTP item.
func (cr *PPTPResource) Update(fullPathName string, item PPTP) error {
	// Marshal the PPTP struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified PPTP resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(PPTPEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a PPTP resource by its full path name.
func (cr *PPTPResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified PPTP resource
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(PPTPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
