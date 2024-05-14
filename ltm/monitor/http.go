package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type HTTPList struct {
	Items    []HTTP `json:"items"`
	Kind     string `json:"kind"`
	SelfLink string `json:"selflink"`
}
type HTTP struct {
	Name                     string `json:"name,omitempty"`
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

const HTTPEndpoint = "http"

type HTTPResource struct {
	b *bigip.BigIP
}

func (mhr *HTTPResource) List() (*HTTPList, error) {
	var mhcl HTTPList
	res, err := mhr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(HTTPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mhcl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mhcl, nil
}

func (mhr *HTTPResource) Get(fullPathName string) (*HTTP, error) {
	var mhc HTTP
	res, err := mhr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(HTTPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mhc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mhc, nil
}

func (mhr *HTTPResource) Create(item HTTP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mhr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(HTTPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mhr *HTTPResource) Update(name string, item HTTP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mhr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(HTTPEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mhr *HTTPResource) Delete(name string) error {
	_, err := mhr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(HTTPEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
