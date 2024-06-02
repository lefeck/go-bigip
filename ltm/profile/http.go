package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type HTTPList struct {
	Items    []HTTP `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type HTTP struct {
	Kind           string        `json:"kind,omitempty"`
	Name           string        `json:"name,omitempty"`
	Partition      string        `json:"partition,omitempty"`
	FullPath       string        `json:"fullPath,omitempty"`
	Generation     int           `json:"generation,omitempty"`
	SelfLink       string        `json:"selfLink,omitempty"`
	AcceptXff      string        `json:"acceptXff,omitempty"`
	AppService     string        `json:"appService,omitempty"`
	BasicAuthRealm string        `json:"basicAuthRealm,omitempty"`
	DefaultsFrom   string        `json:"defaultsFrom,omitempty"`
	Description    string        `json:"description,omitempty"`
	EncryptCookies []interface{} `json:"encryptCookies,omitempty"`
	Enforcement    struct {
		AllowWsHeaderName     string   `json:"allowWsHeaderName,omitempty"`
		ExcessClientHeaders   string   `json:"excessClientHeaders,omitempty"`
		ExcessServerHeaders   string   `json:"excessServerHeaders,omitempty"`
		KnownMethods          []string `json:"knownMethods,omitempty"`
		MaxHeaderCount        int      `json:"maxHeaderCount,omitempty"`
		MaxHeaderSize         int      `json:"maxHeaderSize,omitempty"`
		MaxRequests           int      `json:"maxRequests,omitempty"`
		OversizeClientHeaders string   `json:"oversizeClientHeaders,omitempty"`
		OversizeServerHeaders string   `json:"oversizeServerHeaders,omitempty"`
		Pipeline              string   `json:"pipeline,omitempty"`
		RfcCompliance         string   `json:"rfcCompliance,omitempty"`
		TruncatedRedirects    string   `json:"truncatedRedirects,omitempty"`
		UnknownMethod         string   `json:"unknownMethod,omitempty"`
	} `json:"enforcement,omitempty"`
	ExplicitProxy struct {
		BadRequestMessage      string        `json:"badRequestMessage,omitempty"`
		BadResponseMessage     string        `json:"badResponseMessage,omitempty"`
		ConnectErrorMessage    string        `json:"connectErrorMessage,omitempty"`
		DefaultConnectHandling string        `json:"defaultConnectHandling,omitempty"`
		DNSErrorMessage        string        `json:"dnsErrorMessage,omitempty"`
		DNSResolver            string        `json:"dnsResolver,omitempty"`
		HostNames              []interface{} `json:"hostNames,omitempty"`
		Ipv6                   string        `json:"ipv6,omitempty"`
		RouteDomain            string        `json:"routeDomain,omitempty"`
		TunnelName             string        `json:"tunnelName,omitempty"`
		TunnelOnAnyRequest     string        `json:"tunnelOnAnyRequest,omitempty"`
	} `json:"explicitProxy,omitempty"`
	FallbackHost        string        `json:"fallbackHost,omitempty"`
	FallbackStatusCodes []interface{} `json:"fallbackStatusCodes,omitempty"`
	HeaderErase         string        `json:"headerErase,omitempty"`
	HeaderInsert        string        `json:"headerInsert,omitempty"`
	Hsts                struct {
		IncludeSubdomains string `json:"includeSubdomains,omitempty"`
		MaximumAge        int    `json:"maximumAge,omitempty"`
		Mode              string `json:"mode,omitempty"`
		Preload           string `json:"preload,omitempty"`
	} `json:"hsts,omitempty"`
	InsertXforwardedFor       string        `json:"insertXforwardedFor,omitempty"`
	LwsSeparator              string        `json:"lwsSeparator,omitempty"`
	LwsWidth                  int           `json:"lwsWidth,omitempty"`
	OneconnectStatusReuse     string        `json:"oneconnectStatusReuse,omitempty"`
	OneconnectTransformations string        `json:"oneconnectTransformations,omitempty"`
	ProxyType                 string        `json:"proxyType,omitempty"`
	RedirectRewrite           string        `json:"redirectRewrite,omitempty"`
	RequestChunking           string        `json:"requestChunking,omitempty"`
	ResponseChunking          string        `json:"responseChunking,omitempty"`
	ResponseHeadersPermitted  []interface{} `json:"responseHeadersPermitted,omitempty"`
	ServerAgentName           string        `json:"serverAgentName,omitempty"`
	Sflow                     struct {
		PollInterval       int    `json:"pollInterval,omitempty"`
		PollIntervalGlobal string `json:"pollIntervalGlobal,omitempty"`
		SamplingRate       int    `json:"samplingRate,omitempty"`
		SamplingRateGlobal string `json:"samplingRateGlobal,omitempty"`
	} `json:"sflow,omitempty"`
	ViaHostName           string        `json:"viaHostName,omitempty"`
	ViaRequest            string        `json:"viaRequest,omitempty"`
	ViaResponse           string        `json:"viaResponse,omitempty"`
	XffAlternativeNames   []interface{} `json:"xffAlternativeNames,omitempty"`
	DefaultsFromReference struct {
		Link string `json:"link,omitempty"`
	} `json:"defaultsFromReference,omitempty"`
}

const HTTPEndpoint = "http"

type HTTPResource struct {
	b *bigip.BigIP
}

// List retrieves a list of HTTP resources.
func (cr *HTTPResource) List() (*HTTPList, error) {
	var items HTTPList
	// Perform a GET request to retrieve a list of HTTP resource objects
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into HTTPList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves an HTTP resource by its full path name.
func (cr *HTTPResource) Get(fullPathName string) (*HTTP, error) {
	var item HTTP
	// Perform a GET request to retrieve a specific HTTP resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into HTTP struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new HTTP resource using the provided HTML item.
func (cr *HTTPResource) Create(item HTTP) error {
	// Marshal the HTML struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new HTML resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies an HTTP resource identified by its full path name using the provided HTTP item.
func (cr *HTTPResource) Update(fullPathName string, item HTTP) error {
	// Marshal the HTTP struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified HTTP resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTPEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes an HTTP resource by its full path name.
func (cr *HTTPResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified HTTP resource
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(HTTPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
