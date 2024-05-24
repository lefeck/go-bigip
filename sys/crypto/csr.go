package crypto

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// CsrList holds a list of Csr configurations.
type CsrList struct {
	Items    []Csr  `json:"items"`
	Kind     string `json:"kind"`
	SelfLink string `json:"selflink"`
}

// Csr holds the configuration of a single Csr.
type Csr struct {
}

// CsrEndpoint represents the REST resource for managing Csr.
const CsrEndpoint = "csr"

// CsrResource provides an API to manage Csr configurations.
type CsrResource struct {
	b *bigip.BigIP
}

// List retrieves all Csr details.
func (r *CsrResource) List() (*CsrList, error) {
	var items CsrList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(CryptoEndpoint).SubResource(CsrEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single Csr by node name.
func (r *CsrResource) Get(name string) (*Csr, error) {
	var item Csr
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(CryptoEndpoint).SubResource(CsrEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new Csr item.
func (r *CsrResource) Create(item Csr) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(CryptoEndpoint).SubResource(CsrEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the Csr item identified by the Csr name.
func (r *CsrResource) Update(name string, item Csr) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(CryptoEndpoint).SubResource(CsrEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single Csr identified by the Csr name. If it does not exist, return an error.
func (r *CsrResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(CryptoEndpoint).SubResource(CsrEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
