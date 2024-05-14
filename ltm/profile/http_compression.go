package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type HTTPCompressionList struct {
	Items    []HTTPCompression `json:"items,omitempty"`
	Kind     string            `json:"kind,omitempty"`
	SelfLink string            `json:"selflink,omitempty"`
}

type HTTPCompression struct {
	Kind                  string        `json:"kind"`
	Name                  string        `json:"name"`
	Partition             string        `json:"partition"`
	FullPath              string        `json:"fullPath"`
	Generation            int           `json:"generation"`
	SelfLink              string        `json:"selfLink"`
	AllowHTTP10           string        `json:"allowHttp_10"`
	AppService            string        `json:"appService"`
	BrowserWorkarounds    string        `json:"browserWorkarounds"`
	BufferSize            int           `json:"bufferSize"`
	ContentTypeExclude    []interface{} `json:"contentTypeExclude"`
	ContentTypeInclude    []string      `json:"contentTypeInclude"`
	CPUSaver              string        `json:"cpuSaver"`
	CPUSaverHigh          int           `json:"cpuSaverHigh"`
	CPUSaverLow           int           `json:"cpuSaverLow"`
	DefaultsFrom          string        `json:"defaultsFrom"`
	DefaultsFromReference struct {
		Link string `json:"link"`
	} `json:"defaultsFromReference"`
	Description        string        `json:"description"`
	GzipLevel          int           `json:"gzipLevel"`
	GzipMemoryLevel    int           `json:"gzipMemoryLevel"`
	GzipWindowSize     int           `json:"gzipWindowSize"`
	KeepAcceptEncoding string        `json:"keepAcceptEncoding"`
	MethodPrefer       string        `json:"methodPrefer"`
	MinSize            int           `json:"minSize"`
	Selective          string        `json:"selective"`
	URIExclude         []interface{} `json:"uriExclude"`
	URIInclude         []string      `json:"uriInclude"`
	VaryHeader         string        `json:"varyHeader"`
}

const HTTPCompressionEndpoint = "http-compression"

type HTTPCompressionResource struct {
	b *bigip.BigIP
}

// List retrieves a list of HTTPCompression resources.
func (cr *HTTPCompressionResource) List() (*HTTPCompressionList, error) {
	var items HTTPCompressionList
	// Perform a GET request to retrieve a list of HTTPCompression resource objects
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTPCompressionEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into HTTPCompressionList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves an HTTPCompression resource by its full path name.
func (cr *HTTPCompressionResource) Get(fullPathName string) (*HTTPCompression, error) {
	var item HTTPCompression
	// Perform a GET request to retrieve a specific HTTPCompression resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTPCompressionEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into HTTPCompression struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new HTTPCompression resource using the provided HTTPCompression item.
func (cr *HTTPCompressionResource) Create(item HTTPCompression) error {
	// Marshal the HTTPCompression struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new HTTPCompression resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTPCompressionEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies an HTTPCompression resource identified by its full path name using the provided HTTPCompression item.
func (cr *HTTPCompressionResource) Update(fullPathName string, item HTTPCompression) error {
	// Marshal the HTTPCompression struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified HTTPCompression resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTPCompressionEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes an HTTPCompression resource by its full path name.
func (cr *HTTPCompressionResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified HTTPCompression resource
	_, err := cr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTPCompressionEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
