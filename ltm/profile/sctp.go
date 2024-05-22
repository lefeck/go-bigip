package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type SCTPList struct {
	Items    []SCTP `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type SCTP struct {
	Kind                   string `json:"kind"`
	Name                   string `json:"name"`
	Partition              string `json:"partition"`
	FullPath               string `json:"fullPath"`
	Generation             int    `json:"generation"`
	SelfLink               string `json:"selfLink"`
	AppService             string `json:"appService"`
	ClientSideMultihoming  string `json:"clientSideMultihoming"`
	CookieExpiration       int    `json:"cookieExpiration"`
	DefaultsFrom           string `json:"defaultsFrom"`
	Description            string `json:"description"`
	HeartbeatInterval      int    `json:"heartbeatInterval"`
	HeartbeatMaxBurst      int    `json:"heartbeatMaxBurst"`
	IdleTimeout            int    `json:"idleTimeout"`
	InStreams              int    `json:"inStreams"`
	InitMaxRetries         int    `json:"initMaxRetries"`
	IPTos                  string `json:"ipTos"`
	LinkQos                string `json:"linkQos"`
	MaxBurst               int    `json:"maxBurst"`
	MaxCommunicationPaths  int    `json:"maxCommunicationPaths"`
	MaxPathRetransmitLimit int    `json:"maxPathRetransmitLimit"`
	OutStreams             int    `json:"outStreams"`
	ProxyBufferHigh        int    `json:"proxyBufferHigh"`
	ProxyBufferLow         int    `json:"proxyBufferLow"`
	ReceiveChunks          int    `json:"receiveChunks"`
	ReceiveOrdered         string `json:"receiveOrdered"`
	ReceiveWindowSize      int    `json:"receiveWindowSize"`
	ResetOnTimeout         string `json:"resetOnTimeout"`
	RtoInitial             int    `json:"rtoInitial"`
	RtoMax                 int    `json:"rtoMax"`
	RtoMin                 int    `json:"rtoMin"`
	SackTimeout            int    `json:"sackTimeout"`
	Secret                 string `json:"secret"`
	SendBufferSize         int    `json:"sendBufferSize"`
	SendMaxRetries         int    `json:"sendMaxRetries"`
	SendPartial            string `json:"sendPartial"`
	ServerSideMultihoming  string `json:"serverSideMultihoming"`
	TCPShutdown            string `json:"tcpShutdown"`
	TransmitChunks         int    `json:"transmitChunks"`
}

const SCTPEndpoint = "sctp"

type SCTPResource struct {
	b *bigip.BigIP
}

// List retrieves a list of SCTP resources.
func (cr *SCTPResource) List() (*SCTPList, error) {
	var items SCTPList
	// Perform a GET request to retrieve a list of SCTP resource objects
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(SCTPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into SCTPList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves an SCTP resource by its full path name.
func (cr *SCTPResource) Get(fullPathName string) (*SCTP, error) {
	var item SCTP
	// Perform a GET request to retrieve a specific SCTP resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(SCTPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into SCTP struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new SCTP resource using the provided SCTP item.
func (cr *SCTPResource) Create(item SCTP) error {
	// Marshal the SCTP struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new SCTP resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(SCTPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies an SCTP resource identified by its full path name using the provided SCTP item.
func (cr *SCTPResource) Update(fullPathName string, item SCTP) error {
	// Marshal the SCTP struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified SCTP resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(SCTPEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes an SCTP resource by its full path name.
func (cr *SCTPResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified SCTP resource
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(SCTPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
