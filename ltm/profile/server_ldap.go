package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type ServerLDAPList struct {
	Items    []ServerLDAP `json:"items,omitempty"`
	Kind     string       `json:"kind,omitempty"`
	SelfLink string       `json:"selflink,omitempty"`
}

type ServerLDAP struct {
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

const ServerLDAPEndpoint = "server-ldap"

type ServerLDAPResource struct {
	b *bigip.BigIP
}

// List retrieves a list of ServerLDAP resources.
func (sr *ServerLDAPResource) List() (*ServerLDAPList, error) {
	var items ServerLDAPList
	res, err := sr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ServerLDAPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves a ServerLDAP resource by its full path name.
func (sr *ServerLDAPResource) Get(fullPathName string) (*ServerLDAP, error) {
	var item ServerLDAP
	res, err := sr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ServerLDAPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new ServerLDAP resource using the provided ServerLDAP item.
func (sr *ServerLDAPResource) Create(item ServerLDAP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	_, err = sr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ServerLDAPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies a ServerLDAP resource identified by its full path name using the provided ServerLDAP item.
func (sr *ServerLDAPResource) Update(fullPathName string, item ServerLDAP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	_, err = sr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ServerLDAPEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a ServerLDAP resource by its full path name.
func (sr *ServerLDAPResource) Delete(fullPathName string) error {
	_, err := sr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ServerLDAPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
