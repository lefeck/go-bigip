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
	Kind           string        `json:"kind"`
	Name           string        `json:"name"`
	Partition      string        `json:"partition"`
	FullPath       string        `json:"fullPath"`
	Generation     int           `json:"generation"`
	SelfLink       string        `json:"selfLink"`
	AcceptXff      string        `json:"acceptXff"`
	AppService     string        `json:"appService"`
	BasicAuthRealm string        `json:"basicAuthRealm"`
	DefaultsFrom   string        `json:"defaultsFrom"`
	Description    string        `json:"description"`
	EncryptCookies []interface{} `json:"encryptCookies"`
	Enforcement    struct {
		AllowWsHeaderName     string   `json:"allowWsHeaderName"`
		ExcessClientHeaders   string   `json:"excessClientHeaders"`
		ExcessServerHeaders   string   `json:"excessServerHeaders"`
		KnownMethods          []string `json:"knownMethods"`
		MaxHeaderCount        int      `json:"maxHeaderCount"`
		MaxHeaderSize         int      `json:"maxHeaderSize"`
		MaxRequests           int      `json:"maxRequests"`
		OversizeClientHeaders string   `json:"oversizeClientHeaders"`
		OversizeServerHeaders string   `json:"oversizeServerHeaders"`
		Pipeline              string   `json:"pipeline"`
		RfcCompliance         string   `json:"rfcCompliance"`
		TruncatedRedirects    string   `json:"truncatedRedirects"`
		UnknownMethod         string   `json:"unknownMethod"`
	} `json:"enforcement"`
	ExplicitProxy struct {
		BadRequestMessage      string        `json:"badRequestMessage"`
		BadResponseMessage     string        `json:"badResponseMessage"`
		ConnectErrorMessage    string        `json:"connectErrorMessage"`
		DefaultConnectHandling string        `json:"defaultConnectHandling"`
		DNSErrorMessage        string        `json:"dnsErrorMessage"`
		DNSResolver            string        `json:"dnsResolver"`
		HostNames              []interface{} `json:"hostNames"`
		Ipv6                   string        `json:"ipv6"`
		RouteDomain            string        `json:"routeDomain"`
		TunnelName             string        `json:"tunnelName"`
		TunnelOnAnyRequest     string        `json:"tunnelOnAnyRequest"`
	} `json:"explicitProxy"`
	FallbackHost        string        `json:"fallbackHost"`
	FallbackStatusCodes []interface{} `json:"fallbackStatusCodes"`
	HeaderErase         string        `json:"headerErase"`
	HeaderInsert        string        `json:"headerInsert"`
	Hsts                struct {
		IncludeSubdomains string `json:"includeSubdomains"`
		MaximumAge        int    `json:"maximumAge"`
		Mode              string `json:"mode"`
		Preload           string `json:"preload"`
	} `json:"hsts"`
	InsertXforwardedFor       string        `json:"insertXforwardedFor"`
	LwsSeparator              string        `json:"lwsSeparator"`
	LwsWidth                  int           `json:"lwsWidth"`
	OneconnectStatusReuse     string        `json:"oneconnectStatusReuse"`
	OneconnectTransformations string        `json:"oneconnectTransformations"`
	ProxyType                 string        `json:"proxyType"`
	RedirectRewrite           string        `json:"redirectRewrite"`
	RequestChunking           string        `json:"requestChunking"`
	ResponseChunking          string        `json:"responseChunking"`
	ResponseHeadersPermitted  []interface{} `json:"responseHeadersPermitted"`
	ServerAgentName           string        `json:"serverAgentName"`
	Sflow                     struct {
		PollInterval       int    `json:"pollInterval"`
		PollIntervalGlobal string `json:"pollIntervalGlobal"`
		SamplingRate       int    `json:"samplingRate"`
		SamplingRateGlobal string `json:"samplingRateGlobal"`
	} `json:"sflow"`
	ViaHostName           string        `json:"viaHostName"`
	ViaRequest            string        `json:"viaRequest"`
	ViaResponse           string        `json:"viaResponse"`
	XffAlternativeNames   []interface{} `json:"xffAlternativeNames"`
	DefaultsFromReference struct {
		Link string `json:"link"`
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
