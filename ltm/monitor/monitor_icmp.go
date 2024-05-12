package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"github.com/lefeck/go-bigip/ltm"
	"strings"
)

type MonitorICMPConfigList struct {
	Items    []MonitorICMPConfig `json:"items,omitempty"`
	Kind     string              `json:"kind,omitempty"`
	SelfLink string              `json:"selflink,omitempty"`
}

type MonitorICMPConfig struct {
	Adaptive                 string `json:"adaptive,omitempty"`
	AdaptiveDivergenceType   string `json:"adaptiveDivergenceType,omitempty"`
	AdaptiveDivergenceValue  int    `json:"adaptiveDivergenceValue,omitempty"`
	AdaptiveLimit            int    `json:"adaptiveLimit,omitempty"`
	AdaptiveSamplingTimespan int    `json:"adaptiveSamplingTimespan,omitempty"`
	AppService               string `json:"appService,omitempty"`
	DefaultsFrom             string `json:"defaultsFrom,omitempty"`
	Description              string `json:"description,omitempty"`
	Destination              string `json:"destination,omitempty"`
	FullPath                 string `json:"fullPath,omitempty"`
	Generation               int    `json:"generation,omitempty"`
	Interval                 int    `json:"interval,omitempty"`
	Kind                     string `json:"kind,omitempty"`
	ManualResume             string `json:"manualResume,omitempty"`
	Name                     string `json:"name,omitempty"`
	Partition                string `json:"partition,omitempty"`
	SelfLink                 string `json:"selfLink,omitempty"`
	TimeUntilUp              int    `json:"timeUntilUp,omitempty"`
	Timeout                  int    `json:"timeout,omitempty"`
	Transparent              string `json:"transparent,omitempty"`
	UpInterval               int    `json:"upInterval,omitempty"`
}

const MonitorICMPEndpoint = "/monitor/icmp"

type MonitorICMPResource struct {
	b *bigip.BigIP
}

func (mir *MonitorICMPResource) List() (*MonitorICMPConfigList, error) {
	var micl MonitorICMPConfigList
	res, err := mir.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorICMPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &micl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &micl, nil
}

func (mir *MonitorICMPResource) Get(fullPathName string) (*MonitorICMPConfig, error) {
	var mic MonitorICMPConfig
	res, err := mir.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorICMPEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mic); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mic, nil
}

func (mir *MonitorICMPResource) Create(item MonitorICMPConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mir.b.RestClient.Post().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorICMPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mir *MonitorICMPResource) Update(name string, item MonitorICMPConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mir.b.RestClient.Put().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorICMPEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mir *MonitorICMPResource) Delete(name string) error {
	_, err := mir.b.RestClient.Delete().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorICMPEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
