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
	Kind                          string `json:"kind"`
	Name                          string `json:"name"`
	Partition                     string `json:"partition"`
	FullPath                      string `json:"fullPath"`
	Generation                    int    `json:"generation"`
	SelfLink                      string `json:"selfLink"`
	AppService                    string `json:"appService"`
	CollectCity                   string `json:"collectCity"`
	CollectContinent              string `json:"collectContinent"`
	CollectCountry                string `json:"collectCountry"`
	CollectNexthop                string `json:"collectNexthop"`
	CollectPostCode               string `json:"collectPostCode"`
	CollectRegion                 string `json:"collectRegion"`
	CollectRemoteHostIP           string `json:"collectRemoteHostIp"`
	CollectRemoteHostSubnet       string `json:"collectRemoteHostSubnet"`
	CollectedByClientSide         string `json:"collectedByClientSide"`
	CollectedByServerSide         string `json:"collectedByServerSide"`
	CollectedStatsExternalLogging string `json:"collectedStatsExternalLogging"`
	CollectedStatsInternalLogging string `json:"collectedStatsInternalLogging"`
	DefaultsFrom                  string `json:"defaultsFrom"`
	Description                   string `json:"description"`
	ExternalLoggingPublisher      string `json:"externalLoggingPublisher"`
}

const TCPAnalyticsEndpoint = "tcpanalytics"

type TCPAnalyticsResource struct {
	b *bigip.BigIP
}

// List retrieves a list of TCPAnalytics resources.
func (cr *TCPAnalyticsResource) List() (*TCPAnalyticsList, error) {
	var items TCPAnalyticsList
	// Perform a GET request to retrieve a list of TCPAnalytics resource objects
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
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
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
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
	_, err = cr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
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
	_, err = cr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(TCPAnalyticsEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a TCPAnalytics resource by its full path name.
func (cr *TCPAnalyticsResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified TCPAnalytics resource
	_, err := cr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(TCPAnalyticsEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
