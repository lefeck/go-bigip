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
	Kind                        string        `json:"kind"`
	Name                        string        `json:"name"`
	Partition                   string        `json:"partition"`
	FullPath                    string        `json:"fullPath"`
	Generation                  int           `json:"generation"`
	SelfLink                    string        `json:"selfLink"`
	AppService                  string        `json:"appService"`
	Applications                []interface{} `json:"applications"`
	ApplicationsReference       []interface{} `json:"applicationsReference"`
	CacheAgingRate              int           `json:"cacheAgingRate"`
	CacheClientCacheControlMode string        `json:"cacheClientCacheControlMode"`
	CacheInsertAgeHeader        string        `json:"cacheInsertAgeHeader"`
	CacheMaxAge                 int           `json:"cacheMaxAge"`
	CacheMaxEntries             int           `json:"cacheMaxEntries"`
	CacheObjectMaxSize          int           `json:"cacheObjectMaxSize"`
	CacheObjectMinSize          int           `json:"cacheObjectMinSize"`
	CacheSize                   int           `json:"cacheSize"`
	CacheURIExclude             []interface{} `json:"cacheUriExclude"`
	CacheURIInclude             []string      `json:"cacheUriInclude"`
	CacheURIIncludeOverride     []interface{} `json:"cacheUriIncludeOverride"`
	CacheURIPinned              []interface{} `json:"cacheUriPinned"`
	DefaultsFrom                string        `json:"defaultsFrom"`
	DefaultsFromReference       struct {
		Link string `json:"link"`
	} `json:"defaultsFromReference,omitempty"`
	MetadataCacheMaxSize int `json:"metadataCacheMaxSize"`
}

const WebAccelerationEndpoint = "webacceleration"

type WebAccelerationResource struct {
	b *bigip.BigIP
}

// List retrieves a list of WebAcceleration resources.
func (cr *WebAccelerationResource) List() (*WebAccelerationList, error) {
	var items WebAccelerationList
	// Perform a GET request to retrieve a list of WebAcceleration resource objects
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
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
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
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
	_, err = cr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
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
	_, err = cr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(WebAccelerationEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a WebAcceleration resource by its full path name.
func (cr *WebAccelerationResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified WebAcceleration resource
	_, err := cr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(WebAccelerationEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
