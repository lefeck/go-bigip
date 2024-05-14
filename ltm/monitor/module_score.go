package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type ModuleScoreList struct {
	Items    []ModuleScore `json:"items,omitempty"`
	Kind     string        `json:"kind,omitempty"`
	SelfLink string        `json:"selflink,omitempty"`
}
type ModuleScore struct {
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

const ModuleScoreEndpoint = "module-score"

type ModuleScoreResource struct {
	b *bigip.BigIP
}

func (mmsr *ModuleScoreResource) List() (*ModuleScoreList, error) {
	var mmscl ModuleScoreList
	res, err := mmsr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(ModuleScoreEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mmscl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mmscl, nil
}

func (mmsr *ModuleScoreResource) Get(fullPathName string) (*ModuleScore, error) {
	var mmsc ModuleScore
	res, err := mmsr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(ModuleScoreEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mmsc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mmsc, nil
}

func (mmsr *ModuleScoreResource) Create(item ModuleScore) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mmsr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(ModuleScoreEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mmsr *ModuleScoreResource) Update(name string, item ModuleScore) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mmsr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(ModuleScoreEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mmsr *ModuleScoreResource) Delete(name string) error {
	_, err := mmsr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(ModuleScoreEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
