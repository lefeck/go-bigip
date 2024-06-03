package gtm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
)

// PersistList holds a list of Persist configurations.
type PersistList struct {
	Items    []Persist `json:"items,omitempty"`
	Kind     string    `json:"kind,omitempty"`
	SelfLink string    `json:"selflink,omitempty"`
}

// Persist holds the configuration of a single Persist.
type Persist struct {
	APIRawValues struct {
		APIAnonymous string `json:"apiAnonymous,omitempty"`
	} `json:"apiRawValues,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selfLink,omitempty"`
}

// PersistEndpoint represents the REST resource for managing Persist.
const PersistEndpoint = "persist"

// PersistResource provides an API to manage Persist configurations.
type PersistResource struct {
	b *bigip.BigIP
}

// List retrieves all Persist details.
func (r *PersistResource) List() (*PersistList, error) {
	var items PersistList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(PersistEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}
