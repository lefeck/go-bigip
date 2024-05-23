package sys

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lefeck/go-bigip"
	"strings"
)

// ServiceList holds a list of Service configuration.
type ServiceList struct {
	Items    []Service `json:"items"`
	Kind     string    `json:"kind"`
	SelfLink string    `json:"selflink"`
}

// Service holds the configuration of a single Service.
type Service struct {
	Kind         string `json:"kind"`
	Name         string `json:"name"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	SelfLink     string `json:"selfLink"`
	APIRawValues struct {
		Boot    string `json:"Boot"`
		Default string `json:"Default"`
	} `json:"apiRawValues"`
}

// ServiceEndpoint represents the REST resource for managing Service.
const ServiceEndpoint = "service"

// ServiceResource provides an API to manage Service configurations.
type ServiceResource struct {
	b *bigip.BigIP
}

// List all service details
func (r *ServiceResource) List() (*ServiceList, error) {
	var items ServiceList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ServiceEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get a single service details by the node name
func (r *ServiceResource) Get(name string) (*Service, error) {
	var item Service
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ServiceEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create a new service item
func (r *ServiceResource) Create(item Service) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ServiceEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update the service item identified by the service name, otherwise an error will be reported.
func (r *ServiceResource) Update(name string, item Service) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ServiceEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single service identified by the service name. if it is not exist return error
func (r *ServiceResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ServiceEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
