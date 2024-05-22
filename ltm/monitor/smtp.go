package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type SMTPList struct {
	Items    []SMTP `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type SMTP struct {
	AppService   string `json:"appService,omitempty"`
	Debug        string `json:"debug,omitempty"`
	DefaultsFrom string `json:"defaultsFrom,omitempty"`
	Description  string `json:"description,omitempty"`
	Destination  string `json:"destination,omitempty,omitempty"`
	Domain       string `json:"domain,omitempty,omitempty"`
	FullPath     string `json:"fullPath,omitempty,omitempty"`
	Generation   int    `json:"generation,omitempty,omitempty"`
	Interval     int    `json:"interval,omitempty,omitempty"`
	Kind         string `json:"kind,omitempty,omitempty"`
	ManualResume string `json:"manualResume,omitempty,omitempty"`
	Name         string `json:"name,omitempty,omitempty"`
	Partition    string `json:"partition,omitempty,omitempty"`
	SelfLink     string `json:"selfLink,omitempty,omitempty"`
	TimeUntilUp  int    `json:"timeUntilUp,omitempty,omitempty"`
	Timeout      int    `json:"timeout,omitempty,omitempty"`
	UpInterval   int    `json:"upInterval,omitempty,omitempty"`
}

const SMTPEndpoint = "smtp"

type SMTPResource struct {
	b *bigip.BigIP
}

func (msr *SMTPResource) List() (*SMTPList, error) {
	var mscl SMTPList
	res, err := msr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SMTPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mscl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mscl, nil
}

func (msr *SMTPResource) Get(fullPathName string) (*SMTP, error) {
	var msc SMTP
	res, err := msr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SMTPEndpoint).SubStatsResource(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &msc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &msc, nil
}

func (msr *SMTPResource) Create(item SMTP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SMTPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msr *SMTPResource) Update(name string, item SMTP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SMTPEndpoint).SubStatsResource(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msr *SMTPResource) Delete(name string) error {
	_, err := msr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SMTPEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
