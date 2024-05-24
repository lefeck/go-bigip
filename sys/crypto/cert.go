package crypto

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// CertList holds a list of Cert configurations.
type CertList struct {
	Items    []Cert `json:"items"`
	Kind     string `json:"kind"`
	SelfLink string `json:"selflink"`
}

// Cert holds the configuration of a single Cert.
type Cert struct {
	APIRawValues struct {
		CertificateKeySize string `json:"certificateKeySize,omitempty"`
		Expiration         string `json:"expiration,omitempty"`
		PublicKeyType      string `json:"publicKeyType,omitempty"`
	} `json:"apiRawValues,omitempty"`
	Country      string `json:"country,omitempty"`
	CommonName   string `json:"commonName,omitempty"`
	Fingerprint  string `json:"fingerprint,omitempty"`
	FullPath     string `json:"fullPath,omitempty"`
	Generation   int    `json:"generation,omitempty"`
	Kind         string `json:"kind,omitempty"`
	Name         string `json:"name,omitempty"`
	Organization string `json:"organization,omitempty"`
	Ou           string `json:"ou,omitempty"`
	SelfLink     string `json:"selfLink,omitempty"`
}

// CertEndpoint represents the REST resource for managing Cert.
const CertEndpoint = "cert"

// CertResource provides an API to manage Cert configurations.
type CertResource struct {
	b *bigip.BigIP
}

// List retrieves all Cert details.
func (r *CertResource) List() (*CertList, error) {
	var items CertList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(CryptoEndpoint).SubResource(CertEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single Cert by node name.
func (r *CertResource) Get(name string) (*Cert, error) {
	var item Cert
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(CryptoEndpoint).SubResource(CertEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new Cert item.
func (r *CertResource) Create(item Cert) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(CryptoEndpoint).SubResource(CertEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the Cert item identified by the Cert name.
func (r *CertResource) Update(name string, item Cert) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(CryptoEndpoint).SubResource(CertEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single Cert identified by the Cert name. If it does not exist, return an error.
func (r *CertResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(CryptoEndpoint).SubResource(CertEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
