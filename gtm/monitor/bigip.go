package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// BigIPList contains a list of BigIP uration.
type BigIPList struct {
	Items    []BigIP `json:"items,omitempty"`
	Kind     string  `json:"kind,omitempty"`
	SelfLink string  `json:"selflink,omitempty"`
}

// BigIP contains a single BigIP.
type BigIP struct {
	AggregateDynamicRatios string `json:"aggregateDynamicRatios,omitempty"`
	Destination            string `json:"destination,omitempty"`
	FullPath               string `json:"fullPath,omitempty"`
	Generation             int    `json:"generation,omitempty"`
	IgnoreDownResponse     string `json:"ignoreDownResponse,omitempty"`
	Interval               int    `json:"interval,omitempty"`
	Kind                   string `json:"kind,omitempty"`
	Name                   string `json:"name,omitempty"`
	Partition              string `json:"partition,omitempty"`
	SelfLink               string `json:"selfLink,omitempty"`
	Timeout                int    `json:"timeout,omitempty"`
}

// BigIPEndpoint represents the REST resource for managing BigIP.
const BigIPEndpoint = "bigip"

// BigIPResource provides an API to manage BigIP urations.
type BigIPResource struct {
	b *bigip.BigIP
}

// List returns a list of all BigIP resources
func (r *BigIPResource) List() (*BigIPList, error) {
	var mdcl BigIPList // Defines a variable of type BigIPList
	// Makes a GET request from the REST client and parses the returned data
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(BigIPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Decode the returned JSON data into the BigIPList type variable
	if err := json.Unmarshal(res, &mdcl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mdcl, nil
}

// Get returns a specific BigIP resource identified by its fullPathName
func (r *BigIPResource) Get(fullPathName string) (*BigIP, error) {
	var mdc BigIP // Defines a variable of type BigIP
	// Makes a GET request from the REST client and parses the returned data
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(BigIPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mdc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mdc, nil
}

// Create adds a new BigIP resource provided by the item
func (r *BigIPResource) Create(item BigIP) error {
	jsonData, err := json.Marshal(item) // Marshals the item into JSON data
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	// Makes a POST request from the REST client, specifying the JSON string in the body
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(BigIPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies an existing BigIP resource identified by name using the provided item
func (r *BigIPResource) Update(name string, item BigIP) error {
	jsonData, err := json.Marshal(item) // Marshals the item into JSON data
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	// Makes a PUT request from the REST client, specifying the JSON string in the body
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(BigIPEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a BigIP resource identified by its name
func (r *BigIPResource) Delete(name string) error {
	// Makes a DELETE request from the REST client
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(MonitorEndpoint).SubResource(BigIPEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
