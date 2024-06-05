package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// GTPList holds a list of GTP uration.
type GTPList struct {
	Items    []GTP  `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

// GTP holds the uration of a single GTP.
type GTP struct {
	Destination        string `json:"destination,omitempty"`
	FullPath           string `json:"fullPath,omitempty"`
	Generation         int    `json:"generation,omitempty"`
	IgnoreDownResponse string `json:"ignoreDownResponse,omitempty"`
	Interval           int    `json:"interval,omitempty"`
	Kind               string `json:"kind,omitempty"`
	Name               string `json:"name,omitempty"`
	Partition          string `json:"partition,omitempty"`
	ProbeAttempts      int    `json:"probeAttempts,omitempty"`
	ProbeInterval      int    `json:"probeInterval,omitempty"`
	ProbeTimeout       int    `json:"probeTimeout,omitempty"`
	ProtocolVersion    int    `json:"protocolVersion,omitempty"`
	SelfLink           string `json:"selfLink,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
}

// GTPEndpoint represents the REST resource for managing GTP.
const GTPEndpoint = "gtp"

// GTPResource provides an API to manage GTP urations.
type GTPResource struct {
	b *bigip.BigIP
}

// List returns a list of all GTPList resources
func (r *GTPResource) List() (*GTPList, error) {
	var mdcl GTPList
	// Makes a GET request from the REST client and parses the returned data
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(GTPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Decode the returned JSON data into the BigIPList type variable
	if err := json.Unmarshal(res, &mdcl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mdcl, nil
}

// Get returns a specific GTP resource identified by its fullPathName
func (r *GTPResource) Get(fullPathName string) (*GTP, error) {
	var mdc GTP
	// Makes a GET request from the REST client and parses the returned data
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(GTPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mdc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mdc, nil
}

// Create adds a new GTP resource provided by the item
func (r *GTPResource) Create(item GTP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	// Makes a POST request from the REST client, specifying the JSON string in the body
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(GTPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies an existing GTP resource identified by name using the provided item
func (r *GTPResource) Update(name string, item GTP) error {
	jsonData, err := json.Marshal(item) // Marshals the item into JSON data
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	// Makes a PUT request from the REST client, specifying the JSON string in the body
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(GTPEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a GTP resource identified by its name
func (r *GTPResource) Delete(name string) error {
	// Makes a DELETE request from the REST client
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(GTPEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
