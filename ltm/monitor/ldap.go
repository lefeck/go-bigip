package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type LDAPList struct {
	Items    []LDAP `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}
type LDAP struct {
	AppService          string `json:"appService,omitempty"`
	Base                string `json:"base,omitempty"`
	ChaseReferrals      string `json:"chaseReferrals,omitempty"`
	Debug               string `json:"debug,omitempty"`
	DefaultsFrom        string `json:"defaultsFrom,omitempty"`
	Description         string `json:"description,omitempty"`
	Destination         string `json:"destination,omitempty"`
	Filter              string `json:"filter,omitempty"`
	FullPath            string `json:"fullPath,omitempty"`
	Generation          int    `json:"generation,omitempty"`
	Interval            int    `json:"interval,omitempty"`
	Kind                string `json:"kind,omitempty"`
	MandatoryAttributes string `json:"mandatoryAttributes,omitempty"`
	ManualResume        string `json:"manualResume,omitempty"`
	Name                string `json:"name,omitempty"`
	Partition           string `json:"partition,omitempty"`
	Security            string `json:"security,omitempty"`
	SelfLink            string `json:"selfLink,omitempty"`
	TimeUntilUp         int    `json:"timeUntilUp,omitempty"`
	Timeout             int    `json:"timeout,omitempty"`
	UpInterval          int    `json:"upInterval,omitempty"`
}

const LDAPEndpoint = "//ldap"

type LDAPResource struct {
	b *bigip.BigIP
}

func (mlr *LDAPResource) List() (*LDAPList, error) {
	var mlcl LDAPList
	res, err := mlr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(LDAPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mlcl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mlcl, nil
}

func (mlr *LDAPResource) Get(fullPathName string) (*LDAP, error) {
	var mlc LDAP
	res, err := mlr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(LDAPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mlc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mlc, nil
}

func (mlr *LDAPResource) Create(item LDAP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mlr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(LDAPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mlr *LDAPResource) Update(name string, item LDAP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mlr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(LDAPEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mlr *LDAPResource) Delete(name string) error {
	_, err := mlr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(LDAPEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
