package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"github.com/lefeck/go-bigip/ltm"
	"strings"
)

type MonitorSNMPDCAConfigList struct {
	Items    []MonitorSNMPDCAConfig `json:"items,omitempty"`
	Kind     string                 `json:"kind,omitempty"`
	SelfLink string                 `json:"selflink,omitempty"`
}

type MonitorSNMPDCAConfig struct {
	AgentType         string `json:"agentType,omitempty"`
	AppService        string `json:"appService,omitempty"`
	Community         string `json:"community,omitempty"`
	CPUCoefficient    string `json:"cpuCoefficient,omitempty"`
	CPUThreshold      string `json:"cpuThreshold,omitempty"`
	DefaultsFrom      string `json:"defaultsFrom,omitempty"`
	Description       string `json:"description,omitempty"`
	Destination       string `json:"destination,omitempty"`
	DiskCoefficient   string `json:"diskCoefficient,omitempty"`
	DiskThreshold     string `json:"diskThreshold,omitempty"`
	FullPath          string `json:"fullPath,omitempty"`
	Generation        int    `json:"generation,omitempty"`
	Interval          int    `json:"interval,omitempty"`
	Kind              string `json:"kind,omitempty"`
	MemoryCoefficient string `json:"memoryCoefficient,omitempty"`
	MemoryThreshold   string `json:"memoryThreshold,omitempty"`
	Name              string `json:"name,omitempty"`
	Partition         string `json:"partition,omitempty"`
	SelfLink          string `json:"selfLink,omitempty"`
	TimeUntilUp       int    `json:"timeUntilUp,omitempty"`
	Timeout           int    `json:"timeout,omitempty"`
	UserDefined       string `json:"userDefined,omitempty"`
	Version           string `json:"version,omitempty"`
}

const MonitorSNMPDCAEndpoint = "/monitor/snmp-dca"

type MonitorSNMPDCAResource struct {
	b *bigip.BigIP
}

func (msdr *MonitorSNMPDCAResource) List() (*MonitorSNMPDCAConfigList, error) {
	var msdcl MonitorSNMPDCAConfigList
	res, err := msdr.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSNMPDCAEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &msdcl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &msdcl, nil
}

func (msdr *MonitorSNMPDCAResource) Get(fullPathName string) (*MonitorSNMPDCAConfig, error) {
	var msdc MonitorSNMPDCAConfig
	res, err := msdr.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSNMPDCAEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &msdc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &msdc, nil
}

func (msdr *MonitorSNMPDCAResource) Create(item MonitorSNMPDCAConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msdr.b.RestClient.Post().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSNMPDCAEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msdr *MonitorSNMPDCAResource) Update(name string, item MonitorSNMPDCAConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msdr.b.RestClient.Put().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSNMPDCAEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msdr *MonitorSNMPDCAResource) Delete(name string) error {
	_, err := msdr.b.RestClient.Delete().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSNMPDCAEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}