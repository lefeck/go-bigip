package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type SNMPDCAList struct {
	Items    []SNMPDCA `json:"items,omitempty"`
	Kind     string    `json:"kind,omitempty"`
	SelfLink string    `json:"selflink,omitempty"`
}

type SNMPDCA struct {
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

const SNMPDCAEndpoint = "snmp-dca"

type SNMPDCAResource struct {
	b *bigip.BigIP
}

func (msdr *SNMPDCAResource) List() (*SNMPDCAList, error) {
	var msdcl SNMPDCAList
	res, err := msdr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SNMPDCAEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &msdcl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &msdcl, nil
}

func (msdr *SNMPDCAResource) Get(fullPathName string) (*SNMPDCA, error) {
	var msdc SNMPDCA
	res, err := msdr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SNMPDCAEndpoint).SubStatsResource(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &msdc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &msdc, nil
}

func (msdr *SNMPDCAResource) Create(item SNMPDCA) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msdr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SNMPDCAEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msdr *SNMPDCAResource) Update(name string, item SNMPDCA) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msdr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SNMPDCAEndpoint).SubStatsResource(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msdr *SNMPDCAResource) Delete(name string) error {
	_, err := msdr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SNMPDCAEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
