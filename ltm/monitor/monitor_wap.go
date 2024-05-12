package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"github.com/lefeck/go-bigip/ltm"
	"strings"
)

type MonitorWAPConfigList struct {
	Items    []MonitorWAPConfig `json:"items,omitempty"`
	Kind     string             `json:"kind,omitempty"`
	SelfLink string             `json:"selflink,omitempty"`
}

type MonitorWAPConfig struct {
	AccountingNode string `json:"accountingNode,omitempty"`
	AccountingPort string `json:"accountingPort,omitempty"`
	AppService     string `json:"appService,omitempty"`
	CallId         string `json:"callId,omitempty"`
	Debug          string `json:"debug,omitempty"`
	DefaultsFrom   string `json:"defaultsFrom,omitempty"`
	Description    string `json:"description,omitempty"`
	Destination    string `json:"destination,omitempty"`
	FramedAddress  string `json:"framedAddress,omitempty"`
	FullPath       string `json:"fullPath,omitempty"`
	Generation     int    `json:"generation,omitempty"`
	Interval       int    `json:"interval,omitempty"`
	Kind           string `json:"kind,omitempty"`
	ManualResume   string `json:"manualResume,omitempty"`
	Name           string `json:"name,omitempty"`
	Partition      string `json:"partition,omitempty"`
	Recv           string `json:"recv,omitempty"`
	SelfLink       string `json:"selfLink,omitempty"`
	Send           string `json:"send,omitempty"`
	ServerId       string `json:"serverId,omitempty"`
	SessionId      string `json:"sessionId,omitempty"`
	TimeUntilUp    int    `json:"timeUntilUp,omitempty"`
	Timeout        int    `json:"timeout,omitempty"`
	UpInterval     int    `json:"upInterval,omitempty"`
}

const MonitorWAPEndpoint = "/monitor/wap"

type MonitorWAPResource struct {
	b *bigip.BigIP
}

func (mwr *MonitorWAPResource) List() (*MonitorWAPConfigList, error) {
	var mwcl MonitorWAPConfigList
	res, err := mwr.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorWAPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mwcl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mwcl, nil
}

func (mwr *MonitorWAPResource) Get(fullPathName string) (*MonitorWAPConfig, error) {
	var mwc MonitorWAPConfig
	res, err := mwr.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorWAPEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mwc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mwc, nil
}

func (mwr *MonitorWAPResource) Create(item MonitorWAPConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mwr.b.RestClient.Post().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorWAPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mwr *MonitorWAPResource) Update(name string, item MonitorWAPConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mwr.b.RestClient.Put().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorWAPEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mwr *MonitorWAPResource) Delete(name string) error {
	_, err := mwr.b.RestClient.Delete().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorWAPEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
