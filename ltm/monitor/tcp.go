package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type TCPList struct {
	Items    []TCP  `json:"items"`
	Kind     string `json:"kind"`
	SelfLink string `json:"selflink"`
}

type TCP struct {
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
	IPDscp                   int    `json:"ipDscp,omitempty"`
	Kind                     string `json:"kind,omitempty"`
	ManualResume             string `json:"manualResume,omitempty"`
	Name                     string `json:"name,omitempty"`
	Partition                string `json:"partition,omitempty"`
	Recv                     string `json:"recv,omitempty"`
	RecvDisable              string `json:"recvDisable,omitempty"`
	Reverse                  string `json:"reverse,omitempty"`
	SelfLink                 string `json:"selfLink,omitempty"`
	Send                     string `json:"send,omitempty"`
	TimeUntilUp              int    `json:"timeUntilUp,omitempty"`
	Timeout                  int    `json:"timeout,omitempty"`
	Transparent              string `json:"transparent,omitempty"`
	UpInterval               int    `json:"upInterval,omitempty"`
}

const TCPEndpoint = "tcp"

type TCPResource struct {
	b *bigip.BigIP
}

func (mtr *TCPResource) List() (*TCPList, error) {
	var mtcl TCPList
	res, err := mtr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(TCPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mtcl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mtcl, nil
}

func (mtr *TCPResource) Get(fullPathName string) (*TCP, error) {
	var mtc TCP
	res, err := mtr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(TCPEndpoint).SubStatsResource(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mtc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mtc, nil
}

func (mtr *TCPResource) Create(item TCP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mtr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(TCPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mtr *TCPResource) Update(name string, item TCP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mtr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(TCPEndpoint).SubStatsResource(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mtr *TCPResource) Delete(name string) error {
	_, err := mtr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(TCPEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
