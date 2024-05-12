package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"github.com/lefeck/go-bigip/ltm"
	"strings"
)

type MonitorFirepassConfigList struct {
	Items    []MonitorFirepassConfig `json:"items,omitempty"`
	Kind     string                  `json:"kind,omitempty"`
	SelfLink string                  `json:"selflink,omitempty"`
}

type MonitorFirepassConfig struct {
	AppService       string `json:"appService,omitempty"`
	Cipherlist       string `json:"cipherlist,omitempty"`
	ConcurrencyLimit int    `json:"concurrencyLimit,omitempty"`
	DefaultsFrom     string `json:"defaultsFrom,omitempty"`
	Description      string `json:"description,omitempty"`
	Destination      string `json:"destination,omitempty"`
	FullPath         string `json:"fullPath,omitempty"`
	Generation       int    `json:"generation,omitempty"`
	Interval         int    `json:"interval,omitempty"`
	Kind             string `json:"kind,omitempty"`
	MaxLoadAverage   int    `json:"maxLoadAverage,omitempty"`
	Name             string `json:"name,omitempty"`
	Partition        string `json:"partition,omitempty"`
	SelfLink         string `json:"selfLink,omitempty"`
	TimeUntilUp      int    `json:"timeUntilUp,omitempty"`
	Timeout          int    `json:"timeout,omitempty"`
	UpInterval       int    `json:"upInterval,omitempty"`
	Username         string `json:"username,omitempty"`
}

const MonitorFirepassEndpoint = "/monitor/firepass"

type MonitorFirepassResource struct {
	b *bigip.BigIP
}

func (mfr *MonitorFirepassResource) List() (*MonitorFirepassConfigList, error) {
	var mfcl MonitorFirepassConfigList
	res, err := mfr.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorFirepassEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mfcl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mfcl, nil
}

func (mfr *MonitorFirepassResource) Get(fullPathName string) (*MonitorFirepassConfig, error) {
	var mfc MonitorFirepassConfig
	res, err := mfr.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorFirepassEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mfc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mfc, nil
}

func (mfr *MonitorFirepassResource) Create(item MonitorFirepassConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mfr.b.RestClient.Post().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorFirepassEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mfr *MonitorFirepassResource) Update(name string, item MonitorFirepassConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mfr.b.RestClient.Put().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorFirepassEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mfr *MonitorFirepassResource) Delete(name string) error {
	_, err := mfr.b.RestClient.Delete().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorFirepassEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
