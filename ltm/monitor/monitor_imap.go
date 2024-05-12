package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"github.com/lefeck/go-bigip/ltm"
	"strings"
)

type MonitorIMAPConfigList struct {
	Items    []MonitorIMAPConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

type MonitorIMAPConfig struct {
	Debug        string `json:"debug"`
	DefaultsFrom string `json:"defaultsFrom,omitempty"`
	Description  string `json:"description,omitempty"`
	Destination  string `json:"destination"`
	Folder       string `json:"folder"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	Interval     int    `json:"interval"`
	Kind         string `json:"kind"`
	ManualResume string `json:"manualResume"`
	Name         string `json:"name"`
	Partition    string `json:"partition"`
	SelfLink     string `json:"selfLink"`
	TimeUntilUp  int    `json:"timeUntilUp"`
	Timeout      int    `json:"timeout"`
	UpInterval   int    `json:"upInterval"`
}

const MonitorIMAPEndpoint = "/monitor/imap"

type MonitorIMAPResource struct {
	b *bigip.BigIP
}

func (mir *MonitorIMAPResource) List() (*MonitorIMAPConfigList, error) {
	var micl MonitorIMAPConfigList
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

func (mir *MonitorIMAPResource) Get(fullPathName string) (*MonitorIMAPConfig, error) {
	var mic MonitorIMAPConfig
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

func (mir *MonitorIMAPResource) Create(item MonitorIMAPConfig) error {
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

func (mir *MonitorIMAPResource) Update(name string, item MonitorIMAPConfig) error {
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

func (mir *MonitorIMAPResource) Delete(name string) error {
	_, err := mir.b.RestClient.Delete().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorICMPEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
