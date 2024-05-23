package ecm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// CloudProviderList holds a list of CloudProvider configurations.
type CloudProviderList struct {
	Items    []CloudProvider `json:"items"`
	Kind     string          `json:"kind"`
	SelfLink string          `json:"selflink"`
}

// CloudProvider holds the configuration of a single CloudProvider.
type CloudProvider struct {
	Description      string `json:"description"`
	FullPath         string `json:"fullPath"`
	Generation       int    `json:"generation"`
	Kind             string `json:"kind"`
	Name             string `json:"name"`
	Partition        string `json:"partition"`
	PropertyTemplate []struct {
		Name string `json:"name"`
	} `json:"propertyTemplate"`
	SelfLink string `json:"selfLink"`
}

// ECMCloudProviderEndpoint represents the REST resource for managing ECMCloudProvider.
const CloudProviderEndpoint = "cloud-provider"

// CloudProviderResource provides an API to manage CloudProvider configurations.
type CloudProviderResource struct {
	b *bigip.BigIP
}

// List retrieves all CloudProvider details.
func (r *CloudProviderResource) List() (*CloudProviderList, error) {
	var items CloudProviderList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ECMEndpoint).SubResource(CloudProviderEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single CloudProvider by node name.
func (r *CloudProviderResource) Get(name string) (*CloudProvider, error) {
	var item CloudProvider
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ECMEndpoint).SubResource(CloudProviderEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new CloudProvider item.
func (r *CloudProviderResource) Create(item CloudProvider) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ECMEndpoint).SubResource(CloudProviderEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the CloudProvider item identified by the CloudProvider name.
func (r *CloudProviderResource) Update(name string, item CloudProvider) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ECMEndpoint).SubResource(CloudProviderEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single CloudProvider identified by the CloudProvider name. If it does not exist, return an error.
func (r *CloudProviderResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ECMEndpoint).SubResource(CloudProviderEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
