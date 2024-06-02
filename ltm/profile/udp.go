package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type UDPList struct {
	Items    []UDP  `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type UDP struct {
	Kind                  string `json:"kind,omitempty"`
	Name                  string `json:"name,omitempty"`
	Partition             string `json:"partition,omitempty"`
	FullPath              string `json:"fullPath,omitempty"`
	Generation            int    `json:"generation,omitempty"`
	SelfLink              string `json:"selfLink,omitempty"`
	AllowNoPayload        string `json:"allowNoPayload,omitempty"`
	AppService            string `json:"appService,omitempty"`
	BufferMaxBytes        int    `json:"bufferMaxBytes,omitempty"`
	BufferMaxPackets      int    `json:"bufferMaxPackets,omitempty"`
	DatagramLoadBalancing string `json:"datagramLoadBalancing,omitempty"`
	DefaultsFrom          string `json:"defaultsFrom,omitempty"`
	Description           string `json:"description,omitempty"`
	IdleTimeout           string `json:"idleTimeout,omitempty"`
	IPDfMode              string `json:"ipDfMode,omitempty"`
	IPTosToClient         string `json:"ipTosToClient,omitempty"`
	IPTTLMode             string `json:"ipTtlMode,omitempty"`
	IPTTLV4               int    `json:"ipTtlV4,omitempty"`
	IPTTLV6               int    `json:"ipTtlV6,omitempty"`
	LinkQosToClient       string `json:"linkQosToClient,omitempty"`
	NoChecksum            string `json:"noChecksum,omitempty"`
	ProxyMss              string `json:"proxyMss,omitempty"`
	SendBufferSize        int    `json:"sendBufferSize,omitempty"`
	DefaultsFromReference struct {
		Link string `json:"link,omitempty"`
	} `json:"defaultsFromReference,omitempty"`
}

const UDPEndpoint = "udp"

type UDPResource struct {
	b *bigip.BigIP
}

// List retrieves a list of UDP resources.
func (cr *UDPResource) List() (*UDPList, error) {
	var items UDPList
	// Perform a GET request to retrieve a list of UDP resource objects
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(UDPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into UDPList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves a UDP resource by its full path name.
func (cr *UDPResource) Get(fullPathName string) (*UDP, error) {
	var item UDP
	// Perform a GET request to retrieve a specific UDP resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(UDPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into UDP struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new UDP resource using the provided UDP item.
func (cr *UDPResource) Create(item UDP) error {
	// Marshal the UDP struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new UDP resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(UDPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies a UDP resource identified by its full path name using the provided UDP item.
func (cr *UDPResource) Update(fullPathName string, item UDP) error {
	// Marshal the UDP struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified UDP resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(UDPEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a UDP resource by its full path name.
func (cr *UDPResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified UDP resource
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(UDPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
