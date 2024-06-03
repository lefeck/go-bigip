package gtm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// ListenerList holds a list of Listener uration.
type ListenerList struct {
	Items    []Listener `json:"items,omitempty"`
	Kind     string     `json:"kind,omitempty"`
	SelfLink string     `json:"selflink,omitempty"`
}

// Listener holds the uration of a single Listener.
type Listener struct {
	ProfilesReference struct {
		IsSubcollection bool   `json:"isSubcollection,omitempty"`
		Link            string `json:"link,omitempty"`
	} `json:"profilesReference,omitempty"`
	Advertise                string `json:"advertise,omitempty"`
	TranslateAddress         string `json:"translateAddress,omitempty"`
	SourceAddressTranslation struct {
		Type string `json:"type,omitempty"`
	} `json:"sourceAddressTranslation,omitempty"`
	SelfLink      string `json:"selfLink,omitempty"`
	VlansDisabled bool   `json:"vlansDisabled,omitempty"`
	Name          string `json:"name,omitempty"`
	IpProtocol    string `json:"ipProtocol,omitempty"`
	FullPath      string `json:"fullPath,omitempty"`
	SourcePort    string `json:"sourcePort,omitempty"`
	Kind          string `json:"kind,omitempty"`
	TranslatePort string `json:"translatePort,omitempty"`
	Address       string `json:"address,omitempty"`
	Generation    int    `json:"generation,omitempty"`
	Port          int    `json:"port,omitempty"`
	Mask          string `json:"mask,omitempty"`
	Enabled       bool   `json:"enabled,omitempty"`
	AutoLasthop   string `json:"autoLasthop,omitempty"`
}

// ListenerEndpoint represents the REST resource for managing Listener.
const ListenerEndpoint = "listener"

// ListenerResource provides an API to manage Listener configurations.
type ListenerResource struct {
	b *bigip.BigIP
}

// List retrieves all Listener details.
func (r *ListenerResource) List() (*ListenerList, error) {
	var items ListenerList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(ListenerEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single Listener by node name.
func (r *ListenerResource) Get(name string) (*Listener, error) {
	var item Listener
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(ListenerEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new Listener item.
func (r *ListenerResource) Create(item Listener) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(ListenerEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the Listener item identified by the Listener name.
func (r *ListenerResource) Update(name string, item Listener) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(ListenerEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single Listener identified by the Listener name. If it does not exist, return an error.
func (r *ListenerResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(ListenerEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
