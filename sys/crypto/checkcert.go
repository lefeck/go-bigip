package crypto

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// CheckCertList holds a list of CheckCert configurations.
type CheckCertList struct {
	Items    []CheckCert `json:"items"`
	Kind     string      `json:"kind"`
	SelfLink string      `json:"selflink"`
}

// CheckCert holds the configuration of a single CheckCert.
type CheckCert struct {
}

// CheckCertEndpoint represents the REST resource for managing CheckCert.
const CheckCertEndpoint = "check-cert"

// CheckCertResource provides an API to manage CheckCert configurations.
type CheckCertResource struct {
	b *bigip.BigIP
}

// List retrieves all CheckCert details.
func (r *CheckCertResource) List() (*CheckCertList, error) {
	var items CheckCertList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(CryptoEndpoint).SubResource(CheckCertEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single CheckCert by node name.
func (r *CheckCertResource) Get(name string) (*CheckCert, error) {
	var item CheckCert
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(CryptoEndpoint).SubResource(CheckCertEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new CheckCert item.
func (r *CheckCertResource) Create(item CheckCert) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(CryptoEndpoint).SubResource(CheckCertEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the CheckCert item identified by the CheckCert name.
func (r *CheckCertResource) Update(name string, item CheckCert) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(CryptoEndpoint).SubResource(CheckCertEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single CheckCert identified by the CheckCert name. If it does not exist, return an error.
func (r *CheckCertResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(CryptoEndpoint).SubResource(CheckCertEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
