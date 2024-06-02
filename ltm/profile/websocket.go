package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type WebSocketList struct {
	Items    []WebSocket `json:"items,omitempty"`
	Kind     string      `json:"kind,omitempty"`
	SelfLink string      `json:"selflink,omitempty"`
}

type WebSocket struct {
	Kind                   string `json:"kind,omitempty"`
	Name                   string `json:"name,omitempty"`
	Partition              string `json:"partition,omitempty"`
	FullPath               string `json:"fullPath,omitempty"`
	Generation             int    `json:"generation,omitempty"`
	SelfLink               string `json:"selfLink,omitempty"`
	AppService             string `json:"appService,omitempty"`
	CompressMode           string `json:"compressMode,omitempty"`
	Compression            string `json:"compression,omitempty"`
	DefaultsFrom           string `json:"defaultsFrom,omitempty"`
	Description            string `json:"description,omitempty"`
	Masking                string `json:"masking,omitempty"`
	NoDelay                string `json:"noDelay,omitempty"`
	PayloadProcessingMode  string `json:"payloadProcessingMode,omitempty"`
	PayloadProtocolProfile string `json:"payloadProtocolProfile,omitempty"`
	WindowBits             int    `json:"windowBits,omitempty"`
}

const WebSocketEndpoint = "websocket"

type WebSocketResource struct {
	b *bigip.BigIP
}

// List retrieves a list of WebSocket resources.
func (cr *WebSocketResource) List() (*WebSocketList, error) {
	var items WebSocketList
	// Perform a GET request to retrieve a list of WebSocket resource objects
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(WebSocketEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into WebSocketList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves a WebSocket resource by its full path name.
func (cr *WebSocketResource) Get(fullPathName string) (*WebSocket, error) {
	var item WebSocket
	// Perform a GET request to retrieve a specific WebSocket resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(WebSocketEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into WebSocket struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new WebSocket resource using the provided WebSocket item.
func (cr *WebSocketResource) Create(item WebSocket) error {
	// Marshal the WebSocket struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new WebSocket resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(WebSocketEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies a WebSocket resource identified by its full path name using the provided WebSocket item.
func (cr *WebSocketResource) Update(fullPathName string, item WebSocket) error {
	// Marshal the WebSocket struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified WebSocket resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(WebSocketEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a WebSocket resource by its full path name.
func (cr *WebSocketResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified WebSocket resource
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(WebSocketEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
