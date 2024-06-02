package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type NTLMList struct {
	Items    []NTLM `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type NTLM struct {
	Kind                   string `json:"kind,omitempty"`
	Name                   string `json:"name,omitempty"`
	Partition              string `json:"partition,omitempty"`
	FullPath               string `json:"fullPath,omitempty"`
	Generation             int    `json:"generation,omitempty"`
	SelfLink               string `json:"selfLink,omitempty"`
	AppService             string `json:"appService,omitempty"`
	DefaultsFrom           string `json:"defaultsFrom,omitempty"`
	Description            string `json:"description,omitempty"`
	InsertCookieDomain     string `json:"insertCookieDomain,omitempty"`
	InsertCookieName       string `json:"insertCookieName,omitempty"`
	InsertCookiePassphrase string `json:"insertCookiePassphrase,omitempty"`
	KeyByCookie            string `json:"keyByCookie,omitempty"`
	KeyByCookieName        string `json:"keyByCookieName,omitempty"`
	KeyByDomain            string `json:"keyByDomain,omitempty"`
	KeyByIPAddress         string `json:"keyByIpAddress,omitempty"`
	KeyByTarget            string `json:"keyByTarget,omitempty"`
	KeyByUser              string `json:"keyByUser,omitempty"`
	KeyByWorkstation       string `json:"keyByWorkstation,omitempty"`
}

const NTLMEndpoint = "ntlm"

type NTLMResource struct {
	b *bigip.BigIP
}

// List retrieves a list of NTLM resources.
func (cr *NTLMResource) List() (*NTLMList, error) {
	var items NTLMList
	// Perform a GET request to retrieve a list of NTLM resource objects
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(NTLMEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into NTLMList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves an NTLM resource by its full path name.
func (cr *NTLMResource) Get(fullPathName string) (*NTLM, error) {
	var item NTLM
	// Perform a GET request to retrieve a specific NTLM resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(NTLMEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into NTLM struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new NTLM resource using the provided NTLM item.
func (cr *NTLMResource) Create(item NTLM) error {
	// Marshal the NTLM struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new NTLM resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(NTLMEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies an NTLM resource identified by its full path name using the provided NTLM item.
func (cr *NTLMResource) Update(fullPathName string, item NTLM) error {
	// Marshal the NTLM struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified NTLM resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(NTLMEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes an NTLM resource by its full path name.
func (cr *NTLMResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified NTLM resource
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(NTLMEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
