package gtm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// LinkList holds a list of Link configuration.
type LinkList struct {
	Items    []Link `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

// Link holds the configuration of a single Link.
type Link struct {
	Datacenter          string `json:"datacenter,omitempty"`
	DatacenterReference struct {
		Link string `json:"link,omitempty"`
	} `json:"datacenterReference,omitempty"`
	DuplexBilling             string `json:"duplexBilling,omitempty"`
	Enabled                   bool   `json:"enabled,omitempty"`
	FullPath                  string `json:"fullPath,omitempty"`
	Generation                int    `json:"generation,omitempty"`
	Kind                      string `json:"kind,omitempty"`
	LimitMaxInboundBps        int    `json:"limitMaxInboundBps,omitempty"`
	LimitMaxInboundBpsStatus  string `json:"limitMaxInboundBpsStatus,omitempty"`
	LimitMaxOutboundBps       int    `json:"limitMaxOutboundBps,omitempty"`
	LimitMaxOutboundBpsStatus string `json:"limitMaxOutboundBpsStatus,omitempty"`
	LimitMaxTotalBps          int    `json:"limitMaxTotalBps,omitempty"`
	LimitMaxTotalBpsStatus    string `json:"limitMaxTotalBpsStatus,omitempty"`
	LinkRatio                 int    `json:"linkRatio,omitempty"`
	Monitor                   string `json:"monitor,omitempty"`
	Name                      string `json:"name,omitempty"`
	Partition                 string `json:"partition,omitempty"`
	PrepaidSegment            int    `json:"prepaidSegment,omitempty"`
	RouterAddresses           []struct {
		Name        string `json:"name,omitempty"`
		Translation string `json:"translation,omitempty"`
	} `json:"routerAddresses,omitempty"`
	SelfLink      string `json:"selfLink,omitempty"`
	UplinkAddress string `json:"uplinkAddress,omitempty"`
	Weighting     string `json:"weighting,omitempty"`
}

// LinkEndpoint represents the REST resource for managing Link.
const LinkEndpoint = "link"

// LinkResource provides an API to manage Link configurations.
type LinkResource struct {
	b *bigip.BigIP
}

// List retrieves all Link details.
func (r *LinkResource) List() (*LinkList, error) {
	var items LinkList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(LinkEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single Link by node name.
func (r *LinkResource) Get(name string) (*Link, error) {
	var item Link
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(LinkEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new Link item.
func (r *LinkResource) Create(item Link) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(LinkEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the Link item identified by the Link name.
func (r *LinkResource) Update(name string, item Link) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(LinkEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single Link identified by the Link name. If it does not exist, return an error.
func (r *LinkResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(LinkEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
