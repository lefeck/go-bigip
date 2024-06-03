package gtm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// ServerList holds a list of Server configuration.
type ServerList struct {
	Items    []Server `json:"items,omitempty"`
	Kind     string   `json:"kind,omitempty"`
	SelfLink string   `json:"selflink,omitempty"`
}

// Server holds the configuration of a single Server.
type Server struct {
	Addresses []struct {
		DeviceName  string `json:"deviceName,omitempty"`
		Name        string `json:"name,omitempty"`
		Translation string `json:"translation,omitempty"`
	} `json:"addresses,omitempty"`
	Datacenter          string `json:"datacenter,omitempty"`
	DatacenterReference struct {
		Link string `json:"link,omitempty"`
	} `json:"datacenterReference,omitempty"`
	DevicesReference struct {
		IsSubcollection bool   `json:"isSubcollection,omitempty"`
		Link            string `json:"link,omitempty"`
	} `json:"devicesReference,omitempty"`
	Enabled                   bool   `json:"enabled,omitempty"`
	ExposeRouteDomains        string `json:"exposeRouteDomains,omitempty"`
	FullPath                  string `json:"fullPath,omitempty"`
	Generation                int    `json:"generation,omitempty"`
	IqAllowPath               string `json:"iqAllowPath,omitempty"`
	IqAllowServiceCheck       string `json:"iqAllowServiceCheck,omitempty"`
	IqAllowSnmp               string `json:"iqAllowSnmp,omitempty"`
	Kind                      string `json:"kind,omitempty"`
	LimitCPUUsage             int    `json:"limitCpuUsage,omitempty"`
	LimitCPUUsageStatus       string `json:"limitCpuUsageStatus,omitempty"`
	LimitMaxBps               int    `json:"limitMaxBps,omitempty"`
	LimitMaxBpsStatus         string `json:"limitMaxBpsStatus,omitempty"`
	LimitMaxConnections       int    `json:"limitMaxConnections,omitempty"`
	LimitMaxConnectionsStatus string `json:"limitMaxConnectionsStatus,omitempty"`
	LimitMaxPps               int    `json:"limitMaxPps,omitempty"`
	LimitMaxPpsStatus         string `json:"limitMaxPpsStatus,omitempty"`
	LimitMemAvail             int    `json:"limitMemAvail,omitempty"`
	LimitMemAvailStatus       string `json:"limitMemAvailStatus,omitempty"`
	LinkDiscovery             string `json:"linkDiscovery,omitempty"`
	Monitor                   string `json:"monitor,omitempty"`
	Name                      string `json:"name,omitempty"`
	Partition                 string `json:"partition,omitempty"`
	ProberFallback            string `json:"proberFallback,omitempty"`
	ProberPreference          string `json:"proberPreference,omitempty"`
	Product                   string `json:"product,omitempty"`
	SelfLink                  string `json:"selfLink,omitempty"`
	VirtualServerDiscovery    string `json:"virtualServerDiscovery,omitempty"`
	VirtualServersReference   struct {
		VirtualServers  []ServerVirtualServers `json:"items,omitempty"`
		IsSubcollection bool                   `json:"isSubcollection,omitempty"`
		Link            string                 `json:"link,omitempty"`
	} `json:"virtualServersReference,omitempty"`
}

// ServerVirtualServersList holds a list of ServerVirtualServers configuration.
type ServerVirtualServersList struct {
	Items    []ServerVirtualServers `json:"items"`
	Kind     string                 `json:"kind"`
	SelfLink string                 `json:"selflink"`
}

// ServerVirtualServers holds the configuration of a single ServerVirtualServers.
type ServerVirtualServers struct {
	LimitMaxPpsStatus        string `json:"limitMaxPpsStatus,omitempty"`
	Kind                     string `json:"kind,omitempty"`
	LimitMaxBps              int    `json:"limitMaxBps,omitempty"`
	Destination              string `json:"destination,omitempty"`
	LimitMaxConnections      string `json:"limitMaxConnections,omitempty"`
	Enabled                  bool   `json:"enabled,omitempty"`
	SelfLink                 string `json:"selfLink,omitempty"`
	LimitMaxConnectionStatus string `json:"limitMaxConnectionStatus,omitempty"`
	TranslationPort          int    `json:"translationPort,omitempty"`
	FullPath                 string `json:"fullPath,omitempty"`
	LimitMaxPps              int    `json:"limitMaxPps,omitempty"`
	Generation               int    `json:"generation,omitempty"`
	LimitMaxBpsStatus        string `json:"limitMaxBpsStatus,omitempty"`
	TranslationAddress       string `json:"translationAddress,omitempty"`
	Name                     string `json:"name,omitempty"`
}

// ServerVirtualServersEndpoint represents the REST resource for managing ServerVirtualServers.
const ServerVirtualServersEndpoint = "virtual-servers"

// ServerEndpoint represents the REST resource for managing Server.
const ServerEndpoint = "server"

// ServerResource provides an API to manage Server configurations.
type ServerResource struct {
	b *bigip.BigIP
}

// ListAll  lists all the Server configurations.
func (r *ServerResource) List() (*ServerList, error) {
	var items ServerList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(ServerEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get a single Server configuration identified by name.
func (r *ServerResource) Get(fullPathName string) (*Server, error) {
	var item Server
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(ServerEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// GetVirtualServers lists all the ServerVirtualServers configurations.
func (r *ServerResource) GetVirtualServers(fullPathName string) (*ServerVirtualServersList, error) {
	var items ServerVirtualServersList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(ServerEndpoint).ResourceInstance(fullPathName).SubStatsResource(ServerVirtualServersEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Create a new Server configuration.
func (r *ServerResource) Create(item Server) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(ServerEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Edit a Server configuration identified by name.
func (r *ServerResource) Update(fullPathName string, item Server) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(DatacenterEndpoint).ResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single Server configuration identified by name.
func (r *ServerResource) Delete(fullPathName string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(DatacenterEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
