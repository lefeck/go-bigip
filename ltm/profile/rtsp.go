package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// RTSPList struct contains a list of RTSP resources.
type RTSPList struct {
	Items    []RTSP `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

// RTSP struct contains details about an individual RTSP resource.
type RTSP struct {
	Kind                string `json:"kind,omitempty"`
	Name                string `json:"name,omitempty"`
	Partition           string `json:"partition,omitempty"`
	FullPath            string `json:"fullPath,omitempty"`
	Generation          int    `json:"generation,omitempty"`
	SelfLink            string `json:"selfLink,omitempty"`
	AppService          string `json:"appService,omitempty"`
	CheckSource         string `json:"checkSource,omitempty"`
	DefaultsFrom        string `json:"defaultsFrom,omitempty"`
	Description         string `json:"description,omitempty"`
	IdleTimeout         string `json:"idleTimeout,omitempty"`
	LogProfile          string `json:"logProfile,omitempty"`
	LogPublisher        string `json:"logPublisher,omitempty"`
	MaxHeaderSize       int    `json:"maxHeaderSize,omitempty"`
	MaxQueuedData       int    `json:"maxQueuedData,omitempty"`
	MulticastRedirect   string `json:"multicastRedirect,omitempty"`
	Proxy               string `json:"proxy,omitempty"`
	ProxyHeader         string `json:"proxyHeader,omitempty"`
	RealHTTPPersistence string `json:"realHttpPersistence,omitempty"`
	RtcpPort            int    `json:"rtcpPort,omitempty"`
	RtpPort             int    `json:"rtpPort,omitempty"`
	SessionReconnect    string `json:"sessionReconnect,omitempty"`
	UnicastRedirect     string `json:"unicastRedirect,omitempty"`
}

// RTSPEndpoint is the endpoint constant for RTSP resources.
const RTSPEndpoint = "rtsp"

// RTSPResource struct is used for handling RTSP resources on a bigip.BigIP struct.
type RTSPResource struct {
	b *bigip.BigIP
}

// List retrieves a list of RTSP resources.
func (cr *RTSPResource) List() (*RTSPList, error) {
	var items RTSPList
	// Perform a GET request to retrieve a list of RTSP resource objects
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(RTSPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into RTSPList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves an RTSP resource by its full path name.
func (cr *RTSPResource) Get(fullPathName string) (*RTSP, error) {
	var item RTSP
	// Perform a GET request to retrieve a specific RTSP resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(RTSPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into RTSP struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new RTSP resource using the provided RTSP item.
func (cr *RTSPResource) Create(item RTSP) error {
	// Marshal the RTSP struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new RTSP resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(RTSPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies an RTSP resource identified by its full path name using the provided RTSP item.
func (cr *RTSPResource) Update(fullPathName string, item RTSP) error {
	// Marshal the RTSP struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified RTSP resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(RTSPEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes an RTSP resource by its full path name.
func (cr *RTSPResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified RTSP resource
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(RTSPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
