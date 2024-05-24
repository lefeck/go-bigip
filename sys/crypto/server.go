package crypto

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// ServerList holds a list of CryptoServer configuration.
type ServerList struct {
	Items    []Server `json:"items,omitempty"`
	Kind     string   `json:"kind,omitempty"`
	SelfLink string   `json:"selflink,omitempty"`
}

// Server holds the configuration of a single CryptoServer.
type Server struct {
	FullPath     string `json:"fullPath,omitempty"`
	Generation   int    `json:"generation,omitempty"`
	ServerSize   string `json:"ServerSize,omitempty"`
	ServerType   string `json:"ServerType,omitempty"`
	Kind         string `json:"kind,omitempty"`
	Name         string `json:"name,omitempty"`
	SecurityType string `json:"securityType,omitempty"`
	SelfLink     string `json:"selfLink,omitempty"`
}

// ServerEndpoint represents the REST resource for managing CryptoServer.
const ServerEndpoint = "server"

// ServerResource provides an API to manage Server configurations.
type ServerResource struct {
	b *bigip.BigIP
}

// List retrieves all Server details.
func (r *ServerResource) List() (*ServerList, error) {
	var items ServerList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(CryptoEndpoint).SubResource(ServerEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single Server by node name.
func (r *ServerResource) Get(name string) (*Server, error) {
	var item Server
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(CryptoEndpoint).SubResource(ServerEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new Server item.
func (r *ServerResource) Create(item Server) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(CryptoEndpoint).SubResource(ServerEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the Server item identified by the Server name.
func (r *ServerResource) Update(name string, item Server) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(CryptoEndpoint).SubResource(ServerEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single Server identified by the Server name. If it does not exist, return an error.
func (r *ServerResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(CryptoEndpoint).SubResource(ServerEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
