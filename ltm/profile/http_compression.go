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
	Kind                  string        `json:"kind,omitempty"`
	Name                  string        `json:"name,omitempty"`
	Partition             string        `json:"partition,omitempty"`
	FullPath              string        `json:"fullPath,omitempty"`
	Generation            int           `json:"generation,omitempty"`
	SelfLink              string        `json:"selfLink,omitempty"`
	AllowHTTP10           string        `json:"allowHttp_10,omitempty"`
	AppService            string        `json:"appService,omitempty"`
	BrowserWorkarounds    string        `json:"browserWorkarounds,omitempty"`
	BufferSize            int           `json:"bufferSize,omitempty"`
	ContentTypeExclude    []interface{} `json:"contentTypeExclude,omitempty"`
	ContentTypeInclude    []string      `json:"contentTypeInclude,omitempty"`
	CPUSaver              string        `json:"cpuSaver,omitempty"`
	CPUSaverHigh          int           `json:"cpuSaverHigh,omitempty"`
	CPUSaverLow           int           `json:"cpuSaverLow,omitempty"`
	DefaultsFrom          string        `json:"defaultsFrom,omitempty"`
	DefaultsFromReference struct {
		Link string `json:"link,omitempty"`
	} `json:"defaultsFromReference,omitempty"`
	Description        string        `json:"description,omitempty"`
	GzipLevel          int           `json:"gzipLevel,omitempty"`
	GzipMemoryLevel    int           `json:"gzipMemoryLevel,omitempty"`
	GzipWindowSize     int           `json:"gzipWindowSize,omitempty"`
	KeepAcceptEncoding string        `json:"keepAcceptEncoding,omitempty"`
	MethodPrefer       string        `json:"methodPrefer,omitempty"`
	MinSize            int           `json:"minSize,omitempty"`
	Selective          string        `json:"selective,omitempty"`
	URIExclude         []interface{} `json:"uriExclude,omitempty"`
	URIInclude         []string      `json:"uriInclude,omitempty"`
	VaryHeader         string        `json:"varyHeader,omitempty"`
}

const HTTPCompressionEndpoint = "http-compression"

type HTTPCompressionResource struct {
	b *bigip.BigIP
}

// List retrieves a list of HTTPCompression resources.
func (cr *HTTPCompressionResource) List() (*HTTPCompressionList, error) {
	var items HTTPCompressionList
	// Perform a GET request to retrieve a list of HTTPCompression resource objects
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
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
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
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
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
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
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTPCompressionEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes an HTTPCompression resource by its full path name.
func (cr *HTTPCompressionResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified HTTPCompression resource
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTPCompressionEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
