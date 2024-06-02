package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type DiameterList struct {
	Items    []Diameter `json:"items,omitempty"`
	Kind     string     `json:"kind,omitempty"`
	SelfLink string     `json:"selflink,omitempty"`
}

type Diameter struct {
	Kind                  string `json:"kind,omitempty"`
	Name                  string `json:"name,omitempty"`
	Partition             string `json:"partition,omitempty"`
	FullPath              string `json:"fullPath,omitempty"`
	Generation            int    `json:"generation,omitempty"`
	SelfLink              string `json:"selfLink,omitempty"`
	AppService            string `json:"appService,omitempty"`
	ConnectionPrime       string `json:"connectionPrime,omitempty"`
	DefaultsFrom          string `json:"defaultsFrom,omitempty"`
	Description           string `json:"description,omitempty"`
	DestinationRealm      string `json:"destinationRealm,omitempty"`
	HandshakeTimeout      int    `json:"handshakeTimeout,omitempty"`
	HostIPRewrite         string `json:"hostIpRewrite,omitempty"`
	MaxRetransmitAttempts int    `json:"maxRetransmitAttempts,omitempty"`
	MaxWatchdogFailure    int    `json:"maxWatchdogFailure,omitempty"`
	OriginHostToClient    string `json:"originHostToClient,omitempty"`
	OriginHostToServer    string `json:"originHostToServer,omitempty"`
	OriginRealmToClient   string `json:"originRealmToClient,omitempty"`
	OriginRealmToServer   string `json:"originRealmToServer,omitempty"`
	ParentAvp             string `json:"parentAvp,omitempty"`
	PersistAvp            string `json:"persistAvp,omitempty"`
	ResetOnTimeout        string `json:"resetOnTimeout,omitempty"`
	RetransmitTimeout     int    `json:"retransmitTimeout,omitempty"`
	WatchdogTimeout       int    `json:"watchdogTimeout,omitempty"`
}

const DiameterEndpoint = "diameter"

type DiameterResource struct {
	b *bigip.BigIP
}

func (cr *DiameterResource) List() (*DiameterList, error) {
	var items DiameterList
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(DiameterEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

func (cr *DiameterResource) Get(fullPathName string) (*Diameter, error) {
	var item Diameter
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(DiameterEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

func (cr *DiameterResource) Create(item Diameter) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(DiameterEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (cr *DiameterResource) Update(fullPathName string, item Diameter) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(DiameterEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (cr *DiameterResource) Delete(fullPathName string) error {
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(DiameterEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
