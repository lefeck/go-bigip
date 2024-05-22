package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type ClientLDAPList struct {
	Items    []ClientLDAP `json:"items,omitempty"`
	Kind     string       `json:"kind,omitempty"`
	SelfLink string       `json:"selflink,omitempty"`
}

type ClientLDAP struct {
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

const ClientLDAPEndpoint = "client-ldap"

type ClientLDAPResource struct {
	b *bigip.BigIP
}

// List retrieves a list of ClientLDAP resources.
func (cr *ClientLDAPResource) List() (*ClientLDAPList, error) {
	var items ClientLDAPList
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ClientLDAPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves a ClientLDAP resource by its full path name.
func (cr *ClientLDAPResource) Get(fullPathName string) (*ClientLDAP, error) {
	var item ClientLDAP
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ClientLDAPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new ClientLDAP resource using the provided ClientLDAP item.
func (cr *ClientLDAPResource) Create(item ClientLDAP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ClientLDAPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies a ClientLDAP resource identified by its full path name using the provided ClientLDAP item.
func (cr *ClientLDAPResource) Update(fullPathName string, item ClientLDAP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ClientLDAPEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a ClientLDAP resource by its full path name.
func (cr *ClientLDAPResource) Delete(fullPathName string) error {
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ClientLDAPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
