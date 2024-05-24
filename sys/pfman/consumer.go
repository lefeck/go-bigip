package pfman

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// ConsumerList holds a list of Consumer configurations.
type ConsumerList struct {
	Items    []Consumer `json:"items"`
	Kind     string     `json:"kind"`
	SelfLink string     `json:"selflink"`
}

// Consumer holds the configuration of a single Consumer.
type Consumer struct {
}

// ConsumerEndpoint represents the REST resource for managing Consumer.
const ConsumerEndpoint = "consumer"

// ConsumerResource provides an API to manage Consumer configurations.
type ConsumerResource struct {
	b *bigip.BigIP
}

// List retrieves all Consumer details.
func (r *ConsumerResource) List() (*ConsumerList, error) {
	var items ConsumerList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(PFManEndpoint).SubResource(ConsumerEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single Consumer by node name.
func (r *ConsumerResource) Get(name string) (*Consumer, error) {
	var item Consumer
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(PFManEndpoint).SubResource(ConsumerEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new Consumer item.
func (r *ConsumerResource) Create(item Consumer) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(PFManEndpoint).SubResource(ConsumerEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the Consumer item identified by the Consumer name.
func (r *ConsumerResource) Update(name string, item Consumer) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(PFManEndpoint).SubResource(ConsumerEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single Consumer identified by the Consumer name. If it does not exist, return an error.
func (r *ConsumerResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(PFManEndpoint).SubResource(ConsumerEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
