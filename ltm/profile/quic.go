package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type QUICList struct {
	Items    []QUIC `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type QUIC struct {
	Kind                               string `json:"kind,omitempty"`
	Name                               string `json:"name,omitempty"`
	Partition                          string `json:"partition,omitempty"`
	FullPath                           string `json:"fullPath,omitempty"`
	Generation                         int    `json:"generation,omitempty"`
	SelfLink                           string `json:"selfLink,omitempty"`
	AppService                         string `json:"appService,omitempty"`
	BidiConcurrentStreamsPerConnection int    `json:"bidiConcurrentStreamsPerConnection,omitempty"`
	DefaultsFrom                       string `json:"defaultsFrom,omitempty"`
	Description                        string `json:"description,omitempty"`
	SpinBit                            string `json:"spinBit,omitempty"`
	UniConcurrentStreamsPerConnection  int    `json:"uniConcurrentStreamsPerConnection,omitempty"`
}

const QUICEndpoint = "quic"

type QUICResource struct {
	b *bigip.BigIP
}

// List retrieves a list of QUIC resources.
func (cr *QUICResource) List() (*QUICList, error) {
	var items QUICList
	// Perform a GET request to retrieve a list of QUIC resource objects
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(QUICEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into QUICList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves a QUIC resource by its full path name.
func (cr *QUICResource) Get(fullPathName string) (*QUIC, error) {
	var item QUIC
	// Perform a GET request to retrieve a specific QUIC resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(QUICEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into QUIC struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new QUIC resource using the provided QUIC item.
func (cr *QUICResource) Create(item QUIC) error {
	// Marshal the QUIC struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new QUIC resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(QUICEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies a QUIC resource identified by its full path name using the provided QUIC item.
func (cr *QUICResource) Update(fullPathName string, item QUIC) error {
	// Marshal the QUIC struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified QUIC resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(QUICEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a QUIC resource by its full path name.
func (cr *QUICResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified QUIC resource
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(QUICEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
