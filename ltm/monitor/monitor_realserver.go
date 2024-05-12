package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"github.com/lefeck/go-bigip/ltm"
	"strings"
)

type MonitorRealServerConfigList struct {
	Items    []MonitorRealServerConfig `json:"items,omitempty"`
	Kind     string                    `json:"kind,omitempty"`
	SelfLink string                    `json:"selflink,omitempty"`
}

type MonitorRealServerConfig struct {
	Agent        string `json:"agent,omitempty"`
	AppService   string `json:"appService,omitempty"`
	DefaultsFrom string `json:"defaultsFrom,omitempty"`
	Description  string `json:"description,omitempty"`
	Destination  string `json:"destination,omitempty"`
	FullPath     string `json:"fullPath,omitempty"`
	Generation   int    `json:"generation,omitempty"`
	Interval     int    `json:"interval,omitempty"`
	Kind         string `json:"kind,omitempty"`
	Method       string `json:"method,omitempty"`
	Metrics      string `json:"metrics,omitempty"`
	Name         string `json:"name,omitempty"`
	Partition    string `json:"partition,omitempty"`
	SelfLink     string `json:"selfLink,omitempty"`
	TimeUntilUp  int    `json:"timeUntilUp,omitempty"`
	Timeout      int    `json:"timeout,omitempty"`
	TmCommand    string `json:"tmCommand,omitempty"`
}

const MonitorRealServerEndpoint = "/monitor/real-server"

type MonitorRealServerResource struct {
	b *bigip.BigIP
}

func (mrsr *MonitorRealServerResource) List() (*MonitorRealServerConfigList, error) {
	var mrscl MonitorRealServerConfigList
	res, err := mrsr.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorRealServerEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mrscl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mrscl, nil
}

func (mrsr *MonitorRealServerResource) Get(fullPathName string) (*MonitorRealServerConfig, error) {
	var mrsc MonitorRealServerConfig
	res, err := mrsr.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorRealServerEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mrsc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mrsc, nil
}

func (mrsr *MonitorRealServerResource) Create(item MonitorRealServerConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mrsr.b.RestClient.Post().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorRealServerEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mrsr *MonitorRealServerResource) Update(name string, item MonitorRealServerConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mrsr.b.RestClient.Put().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorRealServerEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mrsr *MonitorRealServerResource) Delete(name string) error {
	_, err := mrsr.b.RestClient.Delete().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorRealServerEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
