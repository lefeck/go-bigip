package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"github.com/lefeck/go-bigip/ltm"
	"strings"
)

type MonitorTCPHalfOpenConfigList struct {
	Items    []MonitorTCPHalfOpenConfig `json:"items,omitempty"`
	Kind     string                     `json:"kind,omitempty"`
	SelfLink string                     `json:"selflink,omitempty"`
}

type MonitorTCPHalfOpenConfig struct {
	AppService   string `json:"appService,omitempty"`
	DefaultsFrom string `json:"defaultsFrom,omitempty"`
	Description  string `json:"description,omitempty"`
	Destination  string `json:"destination,omitempty"`
	FullPath     string `json:"fullPath,omitempty"`
	Generation   int    `json:"generation,omitempty"`
	Interval     int    `json:"interval,omitempty"`
	Kind         string `json:"kind,omitempty"`
	ManualResume string `json:"manualResume,omitempty"`
	Name         string `json:"name,omitempty"`
	Partition    string `json:"partition,omitempty"`
	SelfLink     string `json:"selfLink,omitempty"`
	TimeUntilUp  int    `json:"timeUntilUp,omitempty"`
	Timeout      int    `json:"timeout,omitempty"`
	Transparent  string `json:"transparent,omitempty"`
	UpInterval   int    `json:"upInterval,omitempty"`
}

const MonitorTCPHalfOpenEndpoint = "/monitor/tcp-half-open"

type MonitorTCPHalfOpenResource struct {
	b *bigip.BigIP
}

func (mthor *MonitorTCPHalfOpenResource) List() (*MonitorTCPHalfOpenConfigList, error) {
	var mthocl MonitorTCPHalfOpenConfigList
	res, err := mthor.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorTCPHalfOpenEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mthocl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mthocl, nil
}

func (mthor *MonitorTCPHalfOpenResource) Get(fullPathName string) (*MonitorTCPHalfOpenConfig, error) {
	var mthoc MonitorTCPHalfOpenConfig
	res, err := mthor.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorTCPHalfOpenEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mthoc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mthoc, nil
}

func (mthor *MonitorTCPHalfOpenResource) Create(item MonitorTCPHalfOpenConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mthor.b.RestClient.Post().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorTCPHalfOpenEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mthor *MonitorTCPHalfOpenResource) Update(name string, item MonitorTCPHalfOpenConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mthor.b.RestClient.Put().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorTCPHalfOpenEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mthor *MonitorTCPHalfOpenResource) Delete(name string) error {
	_, err := mthor.b.RestClient.Delete().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorTCPHalfOpenEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
