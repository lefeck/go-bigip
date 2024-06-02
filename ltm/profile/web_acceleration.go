package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type WebAccelerationList struct {
	Items    []WebAcceleration `json:"items,omitempty"`
	Kind     string            `json:"kind,omitempty"`
	SelfLink string            `json:"selflink,omitempty"`
}

type WebAcceleration struct {
	Kind                        string        `json:"kind,omitempty"`
	Name                        string        `json:"name,omitempty"`
	Partition                   string        `json:"partition,omitempty"`
	FullPath                    string        `json:"fullPath,omitempty"`
	Generation                  int           `json:"generation,omitempty"`
	SelfLink                    string        `json:"selfLink,omitempty"`
	AppService                  string        `json:"appService,omitempty"`
	Applications                []interface{} `json:"applications,omitempty"`
	ApplicationsReference       []interface{} `json:"applicationsReference,omitempty"`
	CacheAgingRate              int           `json:"cacheAgingRate,omitempty"`
	CacheClientCacheControlMode string        `json:"cacheClientCacheControlMode,omitempty"`
	CacheInsertAgeHeader        string        `json:"cacheInsertAgeHeader,omitempty"`
	CacheMaxAge                 int           `json:"cacheMaxAge,omitempty"`
	CacheMaxEntries             int           `json:"cacheMaxEntries,omitempty"`
	CacheObjectMaxSize          int           `json:"cacheObjectMaxSize,omitempty"`
	CacheObjectMinSize          int           `json:"cacheObjectMinSize,omitempty"`
	CacheSize                   int           `json:"cacheSize,omitempty"`
	CacheURIExclude             []interface{} `json:"cacheUriExclude,omitempty"`
	CacheURIInclude             []string      `json:"cacheUriInclude,omitempty"`
	CacheURIIncludeOverride     []interface{} `json:"cacheUriIncludeOverride,omitempty"`
	CacheURIPinned              []interface{} `json:"cacheUriPinned,omitempty"`
	DefaultsFrom                string        `json:"defaultsFrom,omitempty"`
	DefaultsFromReference       struct {
		Link string `json:"link,omitempty"`
	} `json:"defaultsFromReference,omitempty"`
	MetadataCacheMaxSize int `json:"metadataCacheMaxSize,omitempty"`
}

const WebAccelerationEndpoint = "webacceleration"

type WebAccelerationResource struct {
	b *bigip.BigIP
}

// List retrieves a list of WebAcceleration resources.
func (cr *WebAccelerationResource) List() (*WebAccelerationList, error) {
	var items WebAccelerationList
	// Perform a GET request to retrieve a list of WebAcceleration resource objects
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(WebAccelerationEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into WebAccelerationList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves a WebAcceleration resource by its full path name.
func (cr *WebAccelerationResource) Get(fullPathName string) (*WebAcceleration, error) {
	var item WebAcceleration
	// Perform a GET request to retrieve a specific WebAcceleration resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(WebAccelerationEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into WebAcceleration struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new WebAcceleration resource using the provided WebAcceleration item.
func (cr *WebAccelerationResource) Create(item WebAcceleration) error {
	// Marshal the WebAcceleration struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new WebAcceleration resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(WebAccelerationEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies a WebAcceleration resource identified by its full path name using the provided WebAcceleration item.
func (cr *WebAccelerationResource) Update(fullPathName string, item WebAcceleration) error {
	// Marshal the WebAcceleration struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified WebAcceleration resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(WebAccelerationEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a WebAcceleration resource by its full path name.
func (cr *WebAccelerationResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified WebAcceleration resource
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(WebAccelerationEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
