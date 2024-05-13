package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type MonitorInbandList struct {
	Items    []MonitorInband `json:"items,omitempty"`
	Kind     string          `json:"kind,omitempty"`
	SelfLink string          `json:"selflink,omitempty"`
}

type MonitorInband struct {
	AppService      string `json:"appService,omitempty"`
	DefaultsFrom    string `json:"defaultsFrom,omitempty"`
	Description     string `json:"description,omitempty"`
	FailureInterval int    `json:"failureInterval,omitempty"`
	Failures        int    `json:"failures,omitempty"`
	FullPath        string `json:"fullPath,omitempty"`
	Generation      int    `json:"generation,omitempty"`
	Kind            string `json:"kind,omitempty"`
	Name            string `json:"name,omitempty"`
	Partition       string `json:"partition,omitempty"`
	ResponseTime    int    `json:"responseTime,omitempty"`
	RetryTime       int    `json:"retryTime,omitempty"`
	SelfLink        string `json:"selfLink,omitempty"`
}

const MonitorInbandEndpoint = "/monitor/inband"

type MonitorInbandResource struct {
	b *bigip.BigIP
}

func (mir *MonitorInbandResource) List() (*MonitorInbandList, error) {
	var micl MonitorInbandList
	res, err := mir.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorInbandEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &micl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &micl, nil
}

func (mir *MonitorInbandResource) Get(fullPathName string) (*MonitorInband, error) {
	var mic MonitorInband
	res, err := mir.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorInbandEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mic); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mic, nil
}

func (mir *MonitorInbandResource) Create(item MonitorInband) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mir.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorInbandEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mir *MonitorInbandResource) Update(name string, item MonitorInband) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mir.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorInbandEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mir *MonitorInbandResource) Delete(name string) error {
	_, err := mir.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorInbandEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
