package global_settings

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// General holds the configuration of a single General.
type General struct {
	AutoDiscovery                     string `json:"autoDiscovery,omitempty"`
	AutoDiscoveryInterval             int    `json:"autoDiscoveryInterval,omitempty"`
	AutomaticConfigurationSaveTimeout int    `json:"automaticConfigurationSaveTimeout,omitempty"`
	CacheLdnsServers                  string `json:"cacheLdnsServers,omitempty"`
	DomainNameCheck                   string `json:"domainNameCheck,omitempty"`
	DrainPersistentRequests           string `json:"drainPersistentRequests,omitempty"`
	ForwardStatus                     string `json:"forwardStatus,omitempty"`
	GtmSetsRecursion                  string `json:"gtmSetsRecursion,omitempty"`
	HeartbeatInterval                 int    `json:"heartbeatInterval,omitempty"`
	Kind                              string `json:"kind,omitempty"`
	MonitorDisabledObjects            string `json:"monitorDisabledObjects,omitempty"`
	NethsmTimeout                     int    `json:"nethsmTimeout,omitempty"`
	SelfLink                          string `json:"selfLink,omitempty"`
	SendWildcardRrs                   string `json:"sendWildcardRrs,omitempty"`
	StaticPersistCidrIpv4             int    `json:"staticPersistCidrIpv4,omitempty"`
	StaticPersistCidrIpv6             int    `json:"staticPersistCidrIpv6,omitempty"`
	Synchronization                   string `json:"synchronization,omitempty"`
	SynchronizationGroupName          string `json:"synchronizationGroupName,omitempty"`
	SynchronizationTimeTolerance      int    `json:"synchronizationTimeTolerance,omitempty"`
	SynchronizationTimeout            int    `json:"synchronizationTimeout,omitempty"`
	SynchronizeZoneFiles              string `json:"synchronizeZoneFiles,omitempty"`
	SynchronizeZoneFilesTimeout       int    `json:"synchronizeZoneFilesTimeout,omitempty"`
	VirtualsDependOnServerState       string `json:"virtualsDependOnServerState,omitempty"`
}

// GlobalSettingsGeneralEndpoint represents the REST resource for managing GlobalSettingsGeneral.
const GeneralEndpoint = "general"

// GlobalSettingsGeneralResource provides an API to manage GlobalSettingsGeneral configurations.
type GeneralResource struct {
	b *bigip.BigIP
}

// List  lists all the General configurations.
func (r *GeneralResource) List() (*General, error) {
	var item General
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(GlobalSettingsEndpoint).SubResource(GeneralEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Update a General configuration.
func (r *GeneralResource) Update(item General) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(GlobalSettingsEndpoint).SubResource(GeneralEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
