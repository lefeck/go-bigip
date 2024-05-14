package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type SIPList struct {
	Items    []SIP  `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type SIP struct {
	Kind                       string `json:"kind"`
	Name                       string `json:"name"`
	Partition                  string `json:"partition"`
	FullPath                   string `json:"fullPath"`
	Generation                 int    `json:"generation"`
	SelfLink                   string `json:"selfLink"`
	AlgEnable                  string `json:"algEnable"`
	AppService                 string `json:"appService"`
	Community                  string `json:"community"`
	DefaultsFrom               string `json:"defaultsFrom"`
	Description                string `json:"description"`
	DialogAware                string `json:"dialogAware"`
	DialogEstablishmentTimeout int    `json:"dialogEstablishmentTimeout"`
	EnableSipFirewall          string `json:"enableSipFirewall"`
	InsertRecordRouteHeader    string `json:"insertRecordRouteHeader"`
	InsertViaHeader            string `json:"insertViaHeader"`
	LogProfile                 string `json:"logProfile"`
	LogPublisher               string `json:"logPublisher"`
	MaxMediaSessions           int    `json:"maxMediaSessions"`
	MaxRegistrations           int    `json:"maxRegistrations"`
	MaxSessionsPerRegistration int    `json:"maxSessionsPerRegistration"`
	MaxSize                    int    `json:"maxSize"`
	RegistrationTimeout        int    `json:"registrationTimeout"`
	RtpProxyStyle              string `json:"rtpProxyStyle"`
	SecureViaHeader            string `json:"secureViaHeader"`
	Security                   string `json:"security"`
	SipSessionTimeout          int    `json:"sipSessionTimeout"`
	TerminateOnBye             string `json:"terminateOnBye"`
	UserViaHeader              string `json:"userViaHeader"`
}

const SIPEndpoint = "sip"

type SIPResource struct {
	b *bigip.BigIP
}

// List retrieves a list of SIP resources.
func (cr *SIPResource) List() (*SIPList, error) {
	var items SIPList
	// Perform a GET request to retrieve a list of SIP resource objects
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(SIPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into SIPList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves a SIP resource by its full path name.
func (cr *SIPResource) Get(fullPathName string) (*SIP, error) {
	var item SIP
	// Perform a GET request to retrieve a specific SIP resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(SIPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into SIP struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new SIP resource using the provided SIP item.
func (cr *SIPResource) Create(item SIP) error {
	// Marshal the SIP struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new SIP resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(SIPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies a SIP resource identified by its full path name using the provided SIP item.
func (cr *SIPResource) Update(fullPathName string, item SIP) error {
	// Marshal the SIP struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified SIP resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(SIPEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a SIP resource by its full path name.
func (cr *SIPResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified SIP resource
	_, err := cr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(SIPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
