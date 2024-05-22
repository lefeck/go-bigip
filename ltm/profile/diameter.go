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
	Kind                  string `json:"kind"`
	Name                  string `json:"name"`
	Partition             string `json:"partition"`
	FullPath              string `json:"fullPath"`
	Generation            int    `json:"generation"`
	SelfLink              string `json:"selfLink"`
	AppService            string `json:"appService"`
	ConnectionPrime       string `json:"connectionPrime"`
	DefaultsFrom          string `json:"defaultsFrom"`
	Description           string `json:"description"`
	DestinationRealm      string `json:"destinationRealm"`
	HandshakeTimeout      int    `json:"handshakeTimeout"`
	HostIPRewrite         string `json:"hostIpRewrite"`
	MaxRetransmitAttempts int    `json:"maxRetransmitAttempts"`
	MaxWatchdogFailure    int    `json:"maxWatchdogFailure"`
	OriginHostToClient    string `json:"originHostToClient"`
	OriginHostToServer    string `json:"originHostToServer"`
	OriginRealmToClient   string `json:"originRealmToClient"`
	OriginRealmToServer   string `json:"originRealmToServer"`
	ParentAvp             string `json:"parentAvp"`
	PersistAvp            string `json:"persistAvp"`
	ResetOnTimeout        string `json:"resetOnTimeout"`
	RetransmitTimeout     int    `json:"retransmitTimeout"`
	WatchdogTimeout       int    `json:"watchdogTimeout"`
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
