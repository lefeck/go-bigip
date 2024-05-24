package crypto

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// KeyList holds a list of CryptoKey configuration.
type KeyList struct {
	Items    []Key  `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

// Key holds the configuration of a single CryptoKey.
type Key struct {
	FullPath     string `json:"fullPath,omitempty"`
	Generation   int    `json:"generation,omitempty"`
	KeySize      string `json:"keySize,omitempty"`
	KeyType      string `json:"keyType,omitempty"`
	Kind         string `json:"kind,omitempty"`
	Name         string `json:"name,omitempty"`
	SecurityType string `json:"securityType,omitempty"`
	SelfLink     string `json:"selfLink,omitempty"`
}

// KeyEndpoint represents the REST resource for managing CryptoKey.
const KeyEndpoint = "key"

// KeyResource provides an API to manage Key configurations.
type KeyResource struct {
	b *bigip.BigIP
}

// List retrieves all Key details.
func (r *KeyResource) List() (*KeyList, error) {
	var items KeyList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(CryptoEndpoint).SubResource(KeyEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single Key by node name.
func (r *KeyResource) Get(name string) (*Key, error) {
	var item Key
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(CryptoEndpoint).SubResource(KeyEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new Key item.
func (r *KeyResource) Create(item Key) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(CryptoEndpoint).SubResource(KeyEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the Key item identified by the Key name.
func (r *KeyResource) Update(name string, item Key) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(CryptoEndpoint).SubResource(KeyEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single Key identified by the Key name. If it does not exist, return an error.
func (r *KeyResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(CryptoEndpoint).SubResource(KeyEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
