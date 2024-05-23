package sys

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
	"time"
)

// ClockConfigList holds a list of Clock configuration.
//type ClockConfigList struct {
//	Entries  map[string]ClockEntry `json:"entries"`
//	Kind     string                `json:"kind"`
//	SelfLink string                `json:"selflink"`
//}

type FullDate struct {
	Description time.Time `json:"description"`
}

type ClockEntry struct {
	Entries struct {
		FullDate FullDate `json:"fullDate"`
	} `json:"entries"`
}

type NestedStats struct {
	NestedStats ClockEntry `json:"nestedStats"`
}

type MainEntry struct {
	SysClock NestedStats `json:"https://localhost/mgmt/tm/sys/clock/0"`
}

type Clock struct {
	Kind     string    `json:"kind"`
	SelfLink string    `json:"selfLink"`
	Entries  MainEntry `json:"entries"`
}

// ClockEndpoint represents the REST resource for managing Clock.
const ClockEndpoint = "clock"

// ClockResource provides an API to manage Clock configurations.
type ClockResource struct {
	b *bigip.BigIP
}

// Get retrieves the details of a single Clock by node name.
func (r *ClockResource) Show() (*Clock, error) {
	var item Clock
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ClockEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new Clock item.
func (r *ClockResource) Create(item Clock) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ClockEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the Clock item identified by the Clock name.
func (r *ClockResource) Update(name string, item Clock) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ClockEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single Clock identified by the Clock name. if it is not exist return error
func (r *ClockResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ClockEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
