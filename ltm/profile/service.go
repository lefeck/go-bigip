package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type ServiceList struct {
	Items    []Service `json:"items,omitempty"`
	Kind     string    `json:"kind,omitempty"`
	SelfLink string    `json:"selflink,omitempty"`
}
type Service struct {
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Partition  string `json:"partition"`
	FullPath   string `json:"fullPath"`
	Generation int    `json:"generation"`
	SelfLink   string `json:"selfLink"`
	AppService string `json:"appService"`
	Type       string `json:"type"`
}

const ServiceEndpoint = "service"

type ServiceResource struct {
	b *bigip.BigIP
}

// List retrieves a list of Service resources.
func (cr *ServiceResource) List() (*ServiceList, error) {
	var items ServiceList
	// Perform a GET request to retrieve a list of Service resource objects
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ServiceEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into ServiceList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves a Service resource by its full path name.
func (cr *ServiceResource) Get(fullPathName string) (*Service, error) {
	var item Service
	// Perform a GET request to retrieve a specific Service resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ServiceEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into Service struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new Service resource using the provided Service item.
func (cr *ServiceResource) Create(item Service) error {
	// Marshal the Service struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new Service resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ServiceEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies a Service resource identified by its full path name using the provided Service item.
func (cr *ServiceResource) Update(fullPathName string, item Service) error {
	// Marshal the Service struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified Service resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ServiceEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a Service resource by its full path name.
func (cr *ServiceResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified Service resource
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ServiceEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
