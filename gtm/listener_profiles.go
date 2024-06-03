package gtm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// ListenerProfilesList holds a list of ListenerProfiles configurations.
type ListenerProfilesList struct {
	Items    []ListenerProfiles `json:"items,omitempty"`
	Kind     string             `json:"kind,omitempty"`
	SelfLink string             `json:"selflink,omitempty"`
}

// ListenerProfiles holds the configuration of a single ListenerProfiles.
type ListenerProfiles struct {
	Kind       string `json:"kind,omitempty"`
	Name       string `json:"name,omitempty"`
	Partition  string `json:"partition,omitempty"`
	FullPath   string `json:"fullPath,omitempty"`
	Generation int    `json:"generation,omitempty"`
	SelfLink   string `json:"selfLink,omitempty"`
	// ... Add other ListenerProfiles specific fields
}

// ListenerProfilesEndpoint represents the REST resource for managing ListenerProfiles.
const ListenerProfilesEndpoint = "listenerprofiles"

// ListenerProfilesResource provides an API to manage ListenerProfiles configurations.
type ListenerProfilesResource struct {
	b *bigip.BigIP
}

// List retrieves all ListenerProfiles details.
func (r *ListenerProfilesResource) List() (*ListenerProfilesList, error) {
	var items ListenerProfilesList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(ListenerProfilesEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single ListenerProfiles by node name.
func (r *ListenerProfilesResource) Get(name string) (*ListenerProfiles, error) {
	var item ListenerProfiles
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(ListenerProfilesEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new ListenerProfiles item.
func (r *ListenerProfilesResource) Create(item ListenerProfiles) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(ListenerProfilesEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the ListenerProfiles item identified by the ListenerProfiles name.
func (r *ListenerProfilesResource) Update(name string, item ListenerProfiles) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(ListenerProfilesEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single ListenerProfiles identified by the ListenerProfiles name. If it does not exist, return an error.
func (r *ListenerProfilesResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(ListenerProfilesEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
