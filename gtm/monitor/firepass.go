package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// FirepassList holds a list of Firepass uration.
type FirepassList struct {
	Items    []Firepass `json:"items,omitempty"`
	Kind     string     `json:"kind,omitempty"`
	SelfLink string     `json:"selflink,omitempty"`
}

// Firepass holds the uration of a single Firepass.
type Firepass struct {
	Cipherlist         string `json:"cipherlist,omitempty"`
	ConcurrencyLimit   int    `json:"concurrencyLimit,omitempty"`
	Destination        string `json:"destination,omitempty"`
	FullPath           string `json:"fullPath,omitempty"`
	Generation         int    `json:"generation,omitempty"`
	IgnoreDownResponse string `json:"ignoreDownResponse,omitempty"`
	Interval           int    `json:"interval,omitempty"`
	Kind               string `json:"kind,omitempty"`
	MaxLoadAverage     int    `json:"maxLoadAverage,omitempty"`
	Name               string `json:"name,omitempty"`
	Partition          string `json:"partition,omitempty"`
	ProbeTimeout       int    `json:"probeTimeout,omitempty"`
	SelfLink           string `json:"selfLink,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
	Username           string `json:"username,omitempty"`
}

// FirepassEndpoint represents the REST resource for managing Firepass.
const FirepassEndpoint = "firepass"

// FirepassResource provides an API to manage Firepass urations.
type FirepassResource struct {
	b *bigip.BigIP
}

// List returns a list of all FirepassList resources
func (r *FirepassResource) List() (*FirepassList, error) {
	var mdcl FirepassList
	// Makes a GET request from the REST client and parses the returned data
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(FirepassEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Decode the returned JSON data into the BigIPList type variable
	if err := json.Unmarshal(res, &mdcl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mdcl, nil
}

// Get returns a specific Firepass resource identified by its fullPathName
func (r *FirepassResource) Get(fullPathName string) (*Firepass, error) {
	var mdc Firepass
	// Makes a GET request from the REST client and parses the returned data
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(FirepassEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mdc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mdc, nil
}

// Create adds a new Firepass resource provided by the item
func (r *FirepassResource) Create(item Firepass) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	// Makes a POST request from the REST client, specifying the JSON string in the body
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(FirepassEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies an existing Firepass resource identified by name using the provided item
func (r *FirepassResource) Update(name string, item Firepass) error {
	jsonData, err := json.Marshal(item) // Marshals the item into JSON data
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	// Makes a PUT request from the REST client, specifying the JSON string in the body
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(FirepassEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a Firepass resource identified by its name
func (r *FirepassResource) Delete(name string) error {
	// Makes a DELETE request from the REST client
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(FirepassEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
