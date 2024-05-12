package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"github.com/lefeck/go-bigip/ltm"
	"strings"
)

type MonitorSMTPConfigList struct {
	Items    []MonitorSMTPConfig `json:"items,omitempty"`
	Kind     string              `json:"kind,omitempty"`
	SelfLink string              `json:"selflink,omitempty"`
}

type MonitorSMTPConfig struct {
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

func (msr *MonitorSMTPResource) List() (*MonitorSMTPConfigList, error) {
	var mscl MonitorSMTPConfigList
	res, err := msr.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSMTPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mscl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mscl, nil
}

func (msr *MonitorSMTPResource) Get(fullPathName string) (*MonitorSMTPConfig, error) {
	var msc MonitorSMTPConfig
	res, err := msr.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSMTPEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &msc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &msc, nil
}

func (msr *MonitorSMTPResource) Create(item MonitorSMTPConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msr.b.RestClient.Post().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSMTPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msr *MonitorSMTPResource) Update(name string, item MonitorSMTPConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msr.b.RestClient.Put().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSMTPEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msr *MonitorSMTPResource) Delete(name string) error {
	_, err := msr.b.RestClient.Delete().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSMTPEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
