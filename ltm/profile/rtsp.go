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
	Kind                string `json:"kind"`
	Name                string `json:"name"`
	Partition           string `json:"partition"`
	FullPath            string `json:"fullPath"`
	Generation          int    `json:"generation"`
	SelfLink            string `json:"selfLink"`
	AppService          string `json:"appService"`
	CheckSource         string `json:"checkSource"`
	DefaultsFrom        string `json:"defaultsFrom"`
	Description         string `json:"description"`
	IdleTimeout         string `json:"idleTimeout"`
	LogProfile          string `json:"logProfile"`
	LogPublisher        string `json:"logPublisher"`
	MaxHeaderSize       int    `json:"maxHeaderSize"`
	MaxQueuedData       int    `json:"maxQueuedData"`
	MulticastRedirect   string `json:"multicastRedirect"`
	Proxy               string `json:"proxy"`
	ProxyHeader         string `json:"proxyHeader"`
	RealHTTPPersistence string `json:"realHttpPersistence"`
	RtcpPort            int    `json:"rtcpPort"`
	RtpPort             int    `json:"rtpPort"`
	SessionReconnect    string `json:"sessionReconnect"`
	UnicastRedirect     string `json:"unicastRedirect"`
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
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
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
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
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
	_, err = cr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
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
	_, err = cr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(RTSPEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes an RTSP resource by its full path name.
func (cr *RTSPResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified RTSP resource
	_, err := cr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(RTSPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
