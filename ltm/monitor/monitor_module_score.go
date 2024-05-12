package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"github.com/lefeck/go-bigip/ltm"
	"strings"
)

type MonitorModuleScoreConfigList struct {
	Items    []MonitorModuleScoreConfig `json:"items,omitempty"`
	Kind     string                     `json:"kind,omitempty"`
	SelfLink string                     `json:"selflink,omitempty"`
}

type MonitorModuleScoreConfig struct {
	AppService    string `json:"appService,omitempty"`
	Debug         string `json:"debug,omitempty"`
	DefaultsFrom  string `json:"defaultsFrom,omitempty"`
	Description   string `json:"description,omitempty"`
	Destination   string `json:"destination,omitempty"`
	FullPath      string `json:"fullPath,omitempty"`
	Generation    int    `json:"generation,omitempty"`
	Interval      int    `json:"interval,omitempty"`
	Kind          string `json:"kind,omitempty"`
	Name          string `json:"name,omitempty"`
	Partition     string `json:"partition,omitempty"`
	Pool          string `json:"pool,omitempty"`
	SelfLink      string `json:"selfLink,omitempty"`
	SnmpCommunity string `json:"snmpCommunity,omitempty"`
	SnmpIpAddress string `json:"snmpIpAddress,omitempty"`
	SnmpPort      int    `json:"snmpPort,omitempty"`
	SnmpVersion   string `json:"snmpVersion,omitempty"`
	TimeUntilUp   int    `json:"timeUntilUp,omitempty"`
	Timeout       int    `json:"timeout,omitempty"`
	UpInterval    int    `json:"upInterval,omitempty"`
}

const MonitorModuleScoreEndpoint = "/monitor/module-score"

type MonitorModuleScoreResource struct {
	b *bigip.BigIP
}

func (mmsr *MonitorModuleScoreResource) List() (*MonitorModuleScoreConfigList, error) {
	var mmscl MonitorModuleScoreConfigList
	res, err := mmsr.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorModuleScoreEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mmscl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mmscl, nil
}

func (mmsr *MonitorModuleScoreResource) Get(fullPathName string) (*MonitorModuleScoreConfig, error) {
	var mmsc MonitorModuleScoreConfig
	res, err := mmsr.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorModuleScoreEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mmsc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mmsc, nil
}

func (mmsr *MonitorModuleScoreResource) Create(item MonitorModuleScoreConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mmsr.b.RestClient.Post().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorModuleScoreEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mmsr *MonitorModuleScoreResource) Update(name string, item MonitorModuleScoreConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mmsr.b.RestClient.Put().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorModuleScoreEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mmsr *MonitorModuleScoreResource) Delete(name string) error {
	_, err := mmsr.b.RestClient.Delete().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorModuleScoreEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
