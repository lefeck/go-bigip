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
	Kind                   string `json:"kind,omitempty"`
	Name                   string `json:"name,omitempty"`
	Partition              string `json:"partition,omitempty"`
	FullPath               string `json:"fullPath,omitempty"`
	Generation             int    `json:"generation,omitempty"`
	SelfLink               string `json:"selfLink,omitempty"`
	AppService             string `json:"appService,omitempty"`
	ClientSideMultihoming  string `json:"clientSideMultihoming,omitempty"`
	CookieExpiration       int    `json:"cookieExpiration,omitempty"`
	DefaultsFrom           string `json:"defaultsFrom,omitempty"`
	Description            string `json:"description,omitempty"`
	HeartbeatInterval      int    `json:"heartbeatInterval,omitempty"`
	HeartbeatMaxBurst      int    `json:"heartbeatMaxBurst,omitempty"`
	IdleTimeout            int    `json:"idleTimeout,omitempty"`
	InStreams              int    `json:"inStreams,omitempty"`
	InitMaxRetries         int    `json:"initMaxRetries,omitempty"`
	IPTos                  string `json:"ipTos,omitempty"`
	LinkQos                string `json:"linkQos,omitempty"`
	MaxBurst               int    `json:"maxBurst,omitempty"`
	MaxCommunicationPaths  int    `json:"maxCommunicationPaths,omitempty"`
	MaxPathRetransmitLimit int    `json:"maxPathRetransmitLimit,omitempty"`
	OutStreams             int    `json:"outStreams,omitempty"`
	ProxyBufferHigh        int    `json:"proxyBufferHigh,omitempty"`
	ProxyBufferLow         int    `json:"proxyBufferLow,omitempty"`
	ReceiveChunks          int    `json:"receiveChunks,omitempty"`
	ReceiveOrdered         string `json:"receiveOrdered,omitempty"`
	ReceiveWindowSize      int    `json:"receiveWindowSize,omitempty"`
	ResetOnTimeout         string `json:"resetOnTimeout,omitempty"`
	RtoInitial             int    `json:"rtoInitial,omitempty"`
	RtoMax                 int    `json:"rtoMax,omitempty"`
	RtoMin                 int    `json:"rtoMin,omitempty"`
	SackTimeout            int    `json:"sackTimeout,omitempty"`
	Secret                 string `json:"secret,omitempty"`
	SendBufferSize         int    `json:"sendBufferSize,omitempty"`
	SendMaxRetries         int    `json:"sendMaxRetries,omitempty"`
	SendPartial            string `json:"sendPartial,omitempty"`
	ServerSideMultihoming  string `json:"serverSideMultihoming,omitempty"`
	TCPShutdown            string `json:"tcpShutdown,omitempty"`
	TransmitChunks         int    `json:"transmitChunks,omitempty"`
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
