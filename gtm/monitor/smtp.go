package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// SMTPList holds a list of SMTP configuration.
type SMTPList struct {
	Items    []SMTP `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

// SMTP holds the configuration of a single SMTP.
type SMTP struct {
	Debug              string `json:"debug,omitempty"`
	Destination        string `json:"destination,omitempty"`
	FullPath           string `json:"fullPath,omitempty"`
	Generation         int    `json:"generation,omitempty"`
	IgnoreDownResponse string `json:"ignoreDownResponse,omitempty"`
	Interval           int    `json:"interval,omitempty"`
	Kind               string `json:"kind,omitempty"`
	Name               string `json:"name,omitempty"`
	Partition          string `json:"partition,omitempty"`
	ProbeTimeout       int    `json:"probeTimeout,omitempty"`
	SelfLink           string `json:"selfLink,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
}

// SMTPEndpoint represents the REST resource for managing SMTP.
const SMTPEndpoint = "smtp"

// SMTPResource provides an API to manage SMTP configurations.
type SMTPResource struct {
	b *bigip.BigIP
}

// List returns a list of all SMTPList resources
func (r *SMTPResource) List() (*SMTPList, error) {
	var mdcl SMTPList
	// Makes a GET request from the REST client and parses the returned data
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(SMTPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Decode the returned JSON data into the BigIPList type variable
	if err := json.Unmarshal(res, &mdcl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mdcl, nil
}

// Get returns a specific SMTP resource identified by its fullPathName
func (r *SMTPResource) Get(fullPathName string) (*SMTP, error) {
	var mdc SMTP
	// Makes a GET request from the REST client and parses the returned data
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(SMTPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mdc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mdc, nil
}

// Create adds a new SMTP resource provided by the item
func (r *SMTPResource) Create(item SMTP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	// Makes a POST request from the REST client, specifying the JSON string in the body
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(SMTPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies an existing SMTP resource identified by name using the provided item
func (r *SMTPResource) Update(name string, item SMTP) error {
	jsonData, err := json.Marshal(item) // Marshals the item into JSON data
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	// Makes a PUT request from the REST client, specifying the JSON string in the body
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(SMTPEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a SMTP resource identified by its name
func (r *SMTPResource) Delete(name string) error {
	// Makes a DELETE request from the REST client
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(SMTPEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
