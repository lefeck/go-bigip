package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"github.com/lefeck/go-bigip/ltm"
	"strings"
)

type MonitorSIPConfigList struct {
	Items    []MonitorSIPConfig `json:"items,omitempty"`
	Kind     string             `json:"kind,omitempty"`
	SelfLink string             `json:"selflink,omitempty"`
}

type MonitorSIPConfig struct {
	AppService    string `json:"appService,omitempty"`
	Cert          string `json:"cert,omitempty"`
	Cipherlist    string `json:"cipherlist,omitempty"`
	Compatibility string `json:"compatibility,omitempty"`
	Debug         string `json:"debug,omitempty"`
	DefaultsFrom  string `json:"defaultsFrom,omitempty"`
	Description   string `json:"description,omitempty"`
	Destination   string `json:"destination,omitempty"`
	FullPath      string `json:"fullPath,omitempty"`
	Filter        string `json:"filter,omitempty"`
	FilterNeg     string `json:"filterNeg,omitempty"`
	Generation    int    `json:"generation,omitempty"`
	Headers       string `json:"headers,omitempty"`
	Interval      int    `json:"interval,omitempty"`
	Key           string `json:"key,omitempty"`
	Kind          string `json:"kind,omitempty"`
	ManualResume  string `json:"manualResume,omitempty"`
	Mode          string `json:"mode,omitempty"`
	Name          string `json:"name,omitempty"`
	Partition     string `json:"partition,omitempty"`
	Request       string `json:"request,omitempty"`
	SelfLink      string `json:"selfLink,omitempty"`
	TimeUntilUp   int    `json:"timeUntilUp,omitempty"`
	Timeout       int    `json:"timeout,omitempty"`
	UpInterval    int    `json:"upInterval,omitempty"`
}

const MonitorSIPEndpoint = "/monitor/sip"

type MonitorSIPResource struct {
	b *bigip.BigIP
}

func (msr *MonitorSIPResource) List() (*MonitorSIPConfigList, error) {
	var mscl MonitorSIPConfigList
	res, err := msr.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSIPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mscl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mscl, nil
}

func (msr *MonitorSIPResource) Get(fullPathName string) (*MonitorSIPConfig, error) {
	var msc MonitorSIPConfig
	res, err := msr.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSIPEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &msc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &msc, nil
}

func (msr *MonitorSIPResource) Create(item MonitorSIPConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msr.b.RestClient.Post().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSIPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msr *MonitorSIPResource) Update(name string, item MonitorSIPConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msr.b.RestClient.Put().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSIPEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msr *MonitorSIPResource) Delete(name string) error {
	_, err := msr.b.RestClient.Delete().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSIPEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
