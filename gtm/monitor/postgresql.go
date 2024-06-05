package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// PostgreSQLList holds a list of PostgreSQL configuration.
type PostgreSQLList struct {
	Items    []PostgreSQL `json:"items,omitempty"`
	Kind     string       `json:"kind,omitempty"`
	SelfLink string       `json:"selflink,omitempty"`
}

// PostgreSQL holds the configuration of a single PostgreSQL.
type PostgreSQL struct {
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

// PostgreSQLEndpoint represents the REST resource for managing PostgreSQL.
const PostgreSQLEndpoint = "postgresql"

// PostgreSQLResource provides an API to manage PostgreSQL configurations.
type PostgreSQLResource struct {
	b *bigip.BigIP
}

// List returns a list of all PostgreSQLList resources
func (r *PostgreSQLResource) List() (*PostgreSQLList, error) {
	var mdcl PostgreSQLList
	// Makes a GET request from the REST client and parses the returned data
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(PostgreSQLEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Decode the returned JSON data into the BigIPList type variable
	if err := json.Unmarshal(res, &mdcl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mdcl, nil
}

// Get returns a specific PostgreSQL resource identified by its fullPathName
func (r *PostgreSQLResource) Get(fullPathName string) (*PostgreSQL, error) {
	var mdc PostgreSQL
	// Makes a GET request from the REST client and parses the returned data
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(PostgreSQLEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mdc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mdc, nil
}

// Create adds a new PostgreSQL resource provided by the item
func (r *PostgreSQLResource) Create(item PostgreSQL) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	// Makes a POST request from the REST client, specifying the JSON string in the body
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(PostgreSQLEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies an existing PostgreSQL resource identified by name using the provided item
func (r *PostgreSQLResource) Update(name string, item PostgreSQL) error {
	jsonData, err := json.Marshal(item) // Marshals the item into JSON data
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	// Makes a PUT request from the REST client, specifying the JSON string in the body
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(PostgreSQLEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a PostgreSQL resource identified by its name
func (r *PostgreSQLResource) Delete(name string) error {
	// Makes a DELETE request from the REST client
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(PostgreSQLEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
