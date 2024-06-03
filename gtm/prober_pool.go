package gtm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// ProberPoolList holds a list of ProberPool configuration.
type ProberPoolList struct {
	Items    []ProberPool `json:"items,omitempty"`
	Kind     string       `json:"kind,omitempty"`
	SelfLink string       `json:"selflink,omitempty"`
}

// ProberPool holds the configuration of a single ProberPool.
type ProberPool struct {
	Enabled           bool   `json:"enabled,omitempty"`
	FullPath          string `json:"fullPath,omitempty"`
	Generation        int    `json:"generation,omitempty"`
	Kind              string `json:"kind,omitempty"`
	LoadBalancingMode string `json:"loadBalancingMode,omitempty"`
	MembersReference  struct {
		Members         []ProberPoolMembers `json:"items,omitempty"`
		IsSubcollection bool                `json:"isSubcollection,omitempty"`
		Link            string              `json:"link,omitempty"`
	} `json:"membersReference,omitempty"`
	Name      string `json:"name,omitempty"`
	Partition string `json:"partition,omitempty"`
	SelfLink  string `json:"selfLink,omitempty"`
}

// ProberPoolMembersList holds a list of ProberPoolMembers configuration.
type ProberPoolMembersList struct {
	Items    []ProberPoolMembers `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

// ProberPoolMembers holds the configuration of a single ProberPoolMembers.
type ProberPoolMembers struct {
	Enabled    bool   `json:"enabled"`
	FullPath   string `json:"fullPath"`
	Generation int    `json:"generation"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Order      int    `json:"order"`
	SelfLink   string `json:"selfLink"`
}

// ProberPoolMembersEndpoint represents the REST resource for managing ProberPoolMembers.
const ProberPoolMembersEndpoint = "members"

// ProberPoolEndpoint represents the REST resource for managing ProberPool.
const ProberPoolEndpoint = "prober-pool"

// ProberPoolResource provides an API to manage ProberPool configurations.
type ProberPoolResource struct {
	b *bigip.BigIP
}

// List retrieves all ProberPool details.
func (r *ProberPoolResource) List() (*ProberPoolList, error) {
	var items ProberPoolList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(ProberPoolEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single ProberPool by node name.
func (r *ProberPoolResource) Get(name string) (*ProberPool, error) {
	var item ProberPool
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(ProberPoolEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new ProberPool item.
func (r *ProberPoolResource) Create(item ProberPool) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(ProberPoolEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the ProberPool item identified by the ProberPool name.
func (r *ProberPoolResource) Update(name string, item ProberPool) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(ProberPoolEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single ProberPool identified by the ProberPool name. If it does not exist, return an error.
func (r *ProberPoolResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(ProberPoolEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}

// GetMembers  lists all the ProberPoolMembers configurations.
func (r *ProberPoolResource) GetMembers(name string) (*ProberPoolMembersList, error) {
	var items ProberPoolMembersList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(ProberPoolEndpoint).ResourceInstance(name).SubResourceInstance(ProberPoolMembersEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}
