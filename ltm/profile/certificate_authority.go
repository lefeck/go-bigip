package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type CertificateAuthorityList struct {
	Items    []CertificateAuthority `json:"items,omitempty"`
	Kind     string                 `json:"kind,omitempty"`
	SelfLink string                 `json:"selflink,omitempty"`
}

type CertificateAuthority struct {
	Kind              string `json:"kind"`
	Name              string `json:"name"`
	Partition         string `json:"partition"`
	FullPath          string `json:"fullPath"`
	Generation        int    `json:"generation"`
	SelfLink          string `json:"selfLink"`
	AppService        string `json:"appService"`
	AuthenticateDepth int    `json:"authenticateDepth"`
	CaFile            string `json:"caFile"`
	CrlFile           string `json:"crlFile"`
	DefaultsFrom      string `json:"defaultsFrom"`
	Description       string `json:"description"`
	LocationSpecific  string `json:"locationSpecific"`
	UpdateCrl         string `json:"updateCrl"`
}

const CertificateAuthorityEndpoint = "certificateauthority"

type CertificateAuthorityResource struct {
	b *bigip.BigIP
}

// List retrieves a list of CertificateAuthority resources.
func (cr *CertificateAuthorityResource) List() (*CertificateAuthorityList, error) {
	var items CertificateAuthorityList
	// Perform a GET request to retrieve a list of CertificateAuthority resource objects
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(CertificateAuthorityEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into CertificateAuthorityList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves a CertificateAuthority resource by its full path name.
func (cr *CertificateAuthorityResource) Get(fullPathName string) (*CertificateAuthority, error) {
	var item CertificateAuthority
	// Perform a GET request to retrieve a specific CertificateAuthority resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(CertificateAuthorityEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into CertificateAuthority struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new CertificateAuthority resource using the provided CertificateAuthority item.
func (cr *CertificateAuthorityResource) Create(item CertificateAuthority) error {
	// Marshal the CertificateAuthority struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new CertificateAuthority resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(CertificateAuthorityEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies a CertificateAuthority resource identified by its full path name using the provided CertificateAuthority item.
func (cr *CertificateAuthorityResource) Update(fullPathName string, item CertificateAuthority) error {
	// Marshal the CertificateAuthority struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified CertificateAuthority resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(CertificateAuthorityEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a CertificateAuthority resource by its full path name.
func (cr *CertificateAuthorityResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified CertificateAuthority resource
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(CertificateAuthorityEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
