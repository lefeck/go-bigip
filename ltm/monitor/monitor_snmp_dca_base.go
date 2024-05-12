package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"github.com/lefeck/go-bigip/ltm"
	"strings"
)

type MonitorSNMPDCABaseConfigList struct {
	Items    []MonitorSNMPDCABaseConfig `json:"items,omitempty"`
	Kind     string                     `json:"kind,omitempty"`
	SelfLink string                     `json:"selflink,omitempty"`
}

type MonitorSNMPDCABaseConfig struct {
	AppService   string `json:"appService,omitempty"`
	Community    string `json:"community,omitempty"`
	DefaultsFrom string `json:"defaultsFrom,omitempty"`
	Description  string `json:"description,omitempty"`
	Destination  string `json:"destination,omitempty"`
	FullPath     string `json:"fullPath,omitempty"`
	Generation   int    `json:"generation,omitempty"`
	Interval     int    `json:"interval,omitempty"`
	Kind         string `json:"kind,omitempty"`
	Name         string `json:"name,omitempty"`
	Partition    string `json:"partition,omitempty"`
	SelfLink     string `json:"selfLink,omitempty"`
	TimeUntilUp  int    `json:"timeUntilUp,omitempty"`
	Timeout      int    `json:"timeout,omitempty"`
	UserDefined  string `json:"userDefined,omitempty"`
	Version      string `json:"version,omitempty"`
}

const MonitorSNMPDCABaseEndpoint = "/monitor/snmp-dca-base"

type MonitorSNMPDCABaseResource struct {
	b *bigip.BigIP
}

func (msdbr *MonitorSNMPDCABaseResource) List() (*MonitorSNMPDCABaseConfigList, error) {
	var msdbcl MonitorSNMPDCABaseConfigList
	res, err := msdbr.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSOAPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &msdbcl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &msdbcl, nil
}

func (msdbr *MonitorSNMPDCABaseResource) Get(fullPathName string) (*MonitorSNMPDCABaseConfig, error) {
	var msdbc MonitorSNMPDCABaseConfig
	res, err := msdbr.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSOAPEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &msdbc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &msdbc, nil
}
func (msdbr *MonitorSNMPDCABaseResource) Create(item MonitorSNMPDCABaseConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msdbr.b.RestClient.Post().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSOAPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msdbr *MonitorSNMPDCABaseResource) Update(name string, item MonitorSNMPDCABaseConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msdbr.b.RestClient.Put().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSOAPEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msdbr *MonitorSNMPDCABaseResource) Delete(name string) error {
	_, err := msdbr.b.RestClient.Delete().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSOAPEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
