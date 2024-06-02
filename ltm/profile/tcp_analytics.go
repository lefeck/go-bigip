package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type TCPAnalyticsList struct {
	Items    []TCPAnalytics `json:"items,omitempty"`
	Kind     string         `json:"kind,omitempty"`
	SelfLink string         `json:"selflink,omitempty"`
}

type TCPAnalytics struct {
	Kind                          string `json:"kind,omitempty"`
	Name                          string `json:"name,omitempty"`
	Partition                     string `json:"partition,omitempty"`
	FullPath                      string `json:"fullPath,omitempty"`
	Generation                    int    `json:"generation,omitempty"`
	SelfLink                      string `json:"selfLink,omitempty"`
	AppService                    string `json:"appService,omitempty"`
	CollectCity                   string `json:"collectCity,omitempty"`
	CollectContinent              string `json:"collectContinent,omitempty"`
	CollectCountry                string `json:"collectCountry,omitempty"`
	CollectNexthop                string `json:"collectNexthop,omitempty"`
	CollectPostCode               string `json:"collectPostCode,omitempty"`
	CollectRegion                 string `json:"collectRegion,omitempty"`
	CollectRemoteHostIP           string `json:"collectRemoteHostIp,omitempty"`
	CollectRemoteHostSubnet       string `json:"collectRemoteHostSubnet,omitempty"`
	CollectedByClientSide         string `json:"collectedByClientSide,omitempty"`
	CollectedByServerSide         string `json:"collectedByServerSide,omitempty"`
	CollectedStatsExternalLogging string `json:"collectedStatsExternalLogging,omitempty"`
	CollectedStatsInternalLogging string `json:"collectedStatsInternalLogging,omitempty"`
	DefaultsFrom                  string `json:"defaultsFrom,omitempty"`
	Description                   string `json:"description,omitempty"`
	ExternalLoggingPublisher      string `json:"externalLoggingPublisher,omitempty"`
}

const TCPAnalyticsEndpoint = "tcpanalytics"

type TCPAnalyticsResource struct {
	b *bigip.BigIP
}

// List retrieves a list of TCPAnalytics resources.
func (cr *TCPAnalyticsResource) List() (*TCPAnalyticsList, error) {
	var items TCPAnalyticsList
	// Perform a GET request to retrieve a list of TCPAnalytics resource objects
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(TCPAnalyticsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into TCPAnalyticsList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves a TCPAnalytics resource by its full path name.
func (cr *TCPAnalyticsResource) Get(fullPathName string) (*TCPAnalytics, error) {
	var item TCPAnalytics
	// Perform a GET request to retrieve a specific TCPAnalytics resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(TCPAnalyticsEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into TCPAnalytics struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new TCPAnalytics resource using the provided TCPAnalytics item.
func (cr *TCPAnalyticsResource) Create(item TCPAnalytics) error {
	// Marshal the TCPAnalytics struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new TCPAnalytics resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(TCPAnalyticsEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies a TCPAnalytics resource identified by its full path name using the provided TCPAnalytics item.
func (cr *TCPAnalyticsResource) Update(fullPathName string, item TCPAnalytics) error {
	// Marshal the TCPAnalytics struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified TCPAnalytics resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(TCPAnalyticsEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a TCPAnalytics resource by its full path name.
func (cr *TCPAnalyticsResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified TCPAnalytics resource
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(TCPAnalyticsEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
