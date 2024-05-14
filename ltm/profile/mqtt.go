package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type MQTTList struct {
	Items    []MQTT `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type MQTT struct {
	Kind         string `json:"kind"`
	Name         string `json:"name"`
	Partition    string `json:"partition"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	SelfLink     string `json:"selfLink"`
	AppService   string `json:"appService"`
	DefaultsFrom string `json:"defaultsFrom"`
	Description  string `json:"description"`
}

const MQTTEndpoint = "mqtt"

type MQTTResource struct {
	b *bigip.BigIP
}

// List retrieves a list of MQTT resources.
func (cr *MQTTResource) List() (*MQTTList, error) {
	var items MQTTList
	// Perform a GET request to retrieve a list of MQTT resource objects
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(MQTTEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into MQTTList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves an MQTT resource by its full path name.
func (cr *MQTTResource) Get(fullPathName string) (*MQTT, error) {
	var item MQTT
	// Perform a GET request to retrieve a specific MQTT resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(MQTTEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into MQTT struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new MQTT resource using the provided MQTT item.
func (cr *MQTTResource) Create(item MQTT) error {
	// Marshal the MQTT struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new MQTT resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(MQTTEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies an MQTT resource identified by its full path name using the provided MQTT item.
func (cr *MQTTResource) Update(fullPathName string, item MQTT) error {
	// Marshal the MQTT struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified MQTT resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(MQTTEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes an MQTT resource by its full path name.
func (cr *MQTTResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified MQTT resource
	_, err := cr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(MQTTEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
