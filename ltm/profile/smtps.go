package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type SMTPSList struct {
	Items    []SMTPS `json:"items,omitempty"`
	Kind     string  `json:"kind,omitempty"`
	SelfLink string  `json:"selflink,omitempty"`
}

type SMTPS struct {
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

const SMTPSEndpoint = "smtps"

type SMTPSResource struct {
	b *bigip.BigIP
}

// List retrieves a list of SMTPS resources.
func (cr *SMTPSResource) List() (*SMTPSList, error) {
	var items SMTPSList
	// Perform a GET request to retrieve a list of SMTPS resource objects
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(SMTPSEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into SMTPSList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves an SMTPS resource by its full path name.
func (cr *SMTPSResource) Get(fullPathName string) (*SMTPS, error) {
	var item SMTPS
	// Perform a GET request to retrieve a specific SMTPS resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(SMTPSEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into SMTPS struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new SMTPS resource using the provided SMTPS item.
func (cr *SMTPSResource) Create(item SMTPS) error {
	// Marshal the SMTPS struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new SMTPS resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(SMTPSEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies an SMTPS resource identified by its full path name using the provided SMTPS item.
func (cr *SMTPSResource) Update(fullPathName string, item SMTPS) error {
	// Marshal the SMTPS struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified SMTPS resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(SMTPSEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes an SMTPS resource by its full path name.
func (cr *SMTPSResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified SMTPS resource
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(SMTPSEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
