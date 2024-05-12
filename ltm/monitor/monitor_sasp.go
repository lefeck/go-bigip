package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"github.com/lefeck/go-bigip/ltm"
	"strings"
)

type MonitorSASPConfigList struct {
	Items    []MonitorSASPConfig `json:"items,omitempty"`
	Kind     string              `json:"kind,omitempty"`
	SelfLink string              `json:"selflink,omitempty"`
}

type MonitorSASPConfig struct {
	AppService       string `json:"appService,omitempty"`
	DefaultsFrom     string `json:"defaultsFrom,omitempty"`
	Description      string `json:"description,omitempty"`
	Destination      string `json:"destination,omitempty"`
	FullPath         string `json:"fullPath,omitempty"`
	Generation       int    `json:"generation,omitempty"`
	Interval         string `json:"interval,omitempty"`
	Kind             string `json:"kind,omitempty"`
	Mode             string `json:"mode,omitempty"`
	MonInterval      int    `json:"monInterval,omitempty"`
	Name             string `json:"name,omitempty"`
	Partition        string `json:"partition,omitempty"`
	PrimaryAddress   string `json:"primaryAddress,omitempty"`
	Protocol         string `json:"protocol,omitempty"`
	SecondaryAddress string `json:"secondaryAddress,omitempty"`
	SelfLink         string `json:"selfLink,omitempty"`
	Service          string `json:"service,omitempty"`
	Timeout          int    `json:"timeout,omitempty"`
	TimeUntilUp      int    `json:"timeUntilUp,omitempty"`
}

const MonitorSASPEndpoint = "/monitor/sasp"

type MonitorSASPResource struct {
	b *bigip.BigIP
}

func (msr *MonitorSASPResource) List() (*MonitorSASPConfigList, error) {
	var mscl MonitorSASPConfigList
	res, err := msr.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSASPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mscl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mscl, nil
}

func (msr *MonitorSASPResource) Get(fullPathName string) (*MonitorSASPConfig, error) {
	var msc MonitorSASPConfig
	res, err := msr.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSASPEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &msc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &msc, nil
}

func (msr *MonitorSASPResource) Create(item MonitorSASPConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msr.b.RestClient.Post().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSASPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msr *MonitorSASPResource) Update(name string, item MonitorSASPConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msr.b.RestClient.Put().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSASPEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msr *MonitorSASPResource) Delete(name string) error {
	_, err := msr.b.RestClient.Delete().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSASPEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
