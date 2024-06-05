package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// OracleList holds a list of Oracle configuration.
type OracleList struct {
	Items    []Oracle `json:"items,omitempty"`
	Kind     string   `json:"kind,omitempty"`
	SelfLink string   `json:"selflink,omitempty"`
}

// Oracle holds the configuration of a single Oracle.
type Oracle struct {
	Count              string `json:"count,omitempty"`
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

// OracleEndpoint represents the REST resource for managing Oracle.
const OracleEndpoint = "/monitor/oracle"

// OracleResource provides an API to manage Oracle configurations.
type OracleResource struct {
	b *bigip.BigIP
}

// List returns a list of all OracleList resources
func (r *OracleResource) List() (*OracleList, error) {
	var mdcl OracleList
	// Makes a GET request from the REST client and parses the returned data
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(OracleEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Decode the returned JSON data into the BigIPList type variable
	if err := json.Unmarshal(res, &mdcl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mdcl, nil
}

// Get returns a specific Oracle resource identified by its fullPathName
func (r *OracleResource) Get(fullPathName string) (*Oracle, error) {
	var mdc Oracle
	// Makes a GET request from the REST client and parses the returned data
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(OracleEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mdc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mdc, nil
}

// Create adds a new Oracle resource provided by the item
func (r *OracleResource) Create(item Oracle) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	// Makes a POST request from the REST client, specifying the JSON string in the body
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(OracleEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies an existing Oracle resource identified by name using the provided item
func (r *OracleResource) Update(name string, item Oracle) error {
	jsonData, err := json.Marshal(item) // Marshals the item into JSON data
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	// Makes a PUT request from the REST client, specifying the JSON string in the body
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(OracleEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a Oracle resource identified by its name
func (r *OracleResource) Delete(name string) error {
	// Makes a DELETE request from the REST client
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(OracleEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
