package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type MonitorSMTPList struct {
	Items    []MonitorSMTP `json:"items,omitempty"`
	Kind     string        `json:"kind,omitempty"`
	SelfLink string        `json:"selflink,omitempty"`
}

type MonitorSMTP struct {
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

const MonitorSMTPEndpoint = "/monitor/smtp"

type MonitorSMTPResource struct {
	b *bigip.BigIP
}

func (msr *MonitorSMTPResource) List() (*MonitorSMTPList, error) {
	var mscl MonitorSMTPList
	res, err := msr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorSMTPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mscl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mscl, nil
}

func (msr *MonitorSMTPResource) Get(fullPathName string) (*MonitorSMTP, error) {
	var msc MonitorSMTP
	res, err := msr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorSMTPEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &msc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &msc, nil
}

func (msr *MonitorSMTPResource) Create(item MonitorSMTP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorSMTPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msr *MonitorSMTPResource) Update(name string, item MonitorSMTP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorSMTPEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msr *MonitorSMTPResource) Delete(name string) error {
	_, err := msr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorSMTPEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
