package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type XMLList struct {
	Items    []XML  `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type XML struct {
	Kind                 string        `json:"kind"`
	Name                 string        `json:"name"`
	Partition            string        `json:"partition"`
	FullPath             string        `json:"fullPath"`
	Generation           int           `json:"generation"`
	SelfLink             string        `json:"selfLink"`
	AppService           string        `json:"appService"`
	DefaultsFrom         string        `json:"defaultsFrom"`
	Description          string        `json:"description"`
	MultipleQueryMatches string        `json:"multipleQueryMatches"`
	NamespaceMappings    []interface{} `json:"namespaceMappings"`
	XpathQueries         []interface{} `json:"xpathQueries"`
}

const XMLEndpoint = "xml"

type XMLResource struct {
	b *bigip.BigIP
}

// List retrieves a list of XML resources.
func (cr *XMLResource) List() (*XMLList, error) {
	var items XMLList
	// Perform a GET request to retrieve a list of XML resource objects
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(XMLEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into XMLList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves an XML resource by its full path name.
func (cr *XMLResource) Get(fullPathName string) (*XML, error) {
	var item XML
	// Perform a GET request to retrieve a specific XML resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(XMLEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into XML struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new XML resource using the provided XML item.
func (cr *XMLResource) Create(item XML) error {
	// Marshal the XML struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new XML resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(XMLEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies an XML resource identified by its full path name using the provided XML item.
func (cr *XMLResource) Update(fullPathName string, item XML) error {
	// Marshal the XML struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified XML resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(XMLEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes an XML resource by its full path name.
func (cr *XMLResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified XML resource
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(XMLEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
