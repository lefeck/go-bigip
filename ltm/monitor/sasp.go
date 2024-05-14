package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type SASPList struct {
	Items    []SASP `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}
type SASP struct {
	AppService       string `json:"appService,omitempty"`
	DefaultsFrom     string `json:"defaultsFrom,omitempty"`
	Description      string `json:"description,omitempty"`
	Destination      string `json:"destination,omitempty"`
	FullPath         string `json:"fullPath,omitempty"`
	Generation       int    `json:"generation,omitempty"`
	Interval         string `json:"interval,omitempty"`
	Kind             string `json:"kind,omitempty"`
	Mode             string `json:"mode,omitempty"`
	MonInterval      int    `json:"monInterval,omitempty"`
	Name             string `json:"name,omitempty"`
	Partition        string `json:"partition,omitempty"`
	PrimaryAddress   string `json:"primaryAddress,omitempty"`
	Protocol         string `json:"protocol,omitempty"`
	SecondaryAddress string `json:"secondaryAddress,omitempty"`
	SelfLink         string `json:"selfLink,omitempty"`
	Service          string `json:"service,omitempty"`
	Timeout          int    `json:"timeout,omitempty"`
	TimeUntilUp      int    `json:"timeUntilUp,omitempty"`
}

const SASPEndpoint = "sasp"

type SASPResource struct {
	b *bigip.BigIP
}

func (msr *SASPResource) List() (*SASPList, error) {
	var mscl SASPList
	res, err := msr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SASPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mscl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mscl, nil
}

func (msr *SASPResource) Get(fullPathName string) (*SASP, error) {
	var msc SASP
	res, err := msr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SASPEndpoint).SubStatsResource(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &msc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &msc, nil
}

func (msr *SASPResource) Create(item SASP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SASPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msr *SASPResource) Update(name string, item SASP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SASPEndpoint).SubStatsResource(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msr *SASPResource) Delete(name string) error {
	_, err := msr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SASPEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
