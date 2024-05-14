package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type SocksList struct {
	Items    []Socks `json:"items,omitempty"`
	Kind     string  `json:"kind,omitempty"`
	SelfLink string  `json:"selflink,omitempty"`
}

type Socks struct {
	Kind                   string   `json:"kind"`
	Name                   string   `json:"name"`
	Partition              string   `json:"partition"`
	FullPath               string   `json:"fullPath"`
	Generation             int      `json:"generation"`
	SelfLink               string   `json:"selfLink"`
	AppService             string   `json:"appService"`
	DefaultConnectHandling string   `json:"defaultConnectHandling"`
	DefaultsFrom           string   `json:"defaultsFrom"`
	Description            string   `json:"description"`
	DNSResolver            string   `json:"dnsResolver"`
	Ipv6                   string   `json:"ipv6"`
	ProtocolVersions       []string `json:"protocolVersions"`
	RouteDomain            string   `json:"routeDomain"`
	RouteDomainReference   struct {
		Link string `json:"link"`
	} `json:"routeDomainReference"`
	TunnelName          string `json:"tunnelName"`
	TunnelNameReference struct {
		Link string `json:"link"`
	} `json:"tunnelNameReference"`
}

const SocksEndpoint = "socks"

type SocksResource struct {
	b *bigip.BigIP
}

// List retrieves a list of Socks resources.
func (cr *SocksResource) List() (*SocksList, error) {
	var items SocksList
	// Perform a GET request to retrieve a list of Socks resource objects
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(SocksEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into SocksList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves a Socks resource by its full path name.
func (cr *SocksResource) Get(fullPathName string) (*Socks, error) {
	var item Socks
	// Perform a GET request to retrieve a specific Socks resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(SocksEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into Socks struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new Socks resource using the provided Socks item.
func (cr *SocksResource) Create(item Socks) error {
	// Marshal the Socks struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new Socks resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(SocksEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies a Socks resource identified by its full path name using the provided Socks item.
func (cr *SocksResource) Update(fullPathName string, item Socks) error {
	// Marshal the Socks struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified Socks resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(SocksEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a Socks resource by its full path name.
func (cr *SocksResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified Socks resource
	_, err := cr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(SocksEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
