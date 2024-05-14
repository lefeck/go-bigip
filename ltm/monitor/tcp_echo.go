package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type TCPEchoList struct {
	Items    []TCPEcho `json:"items,omitempty"`
	Kind     string    `json:"kind,omitempty"`
	SelfLink string    `json:"selflink,omitempty"`
}

type TCPEcho struct {
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

const TCPEchoEndpoint = "tcp-echo"

type TCPEchoResource struct {
	b *bigip.BigIP
}

func (mter *TCPEchoResource) List() (*TCPEchoList, error) {
	var mtecl TCPEchoList
	res, err := mter.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(TCPEchoEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mtecl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mtecl, nil
}

func (mter *TCPEchoResource) Get(fullPathName string) (*TCPEcho, error) {
	var mtec TCPEcho
	res, err := mter.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(TCPEchoEndpoint).SubStatsResource(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mtec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mtec, nil
}

func (mter *TCPEchoResource) Create(item TCPEcho) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mter.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(TCPEchoEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mter *TCPEchoResource) Update(name string, item TCPEcho) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mter.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(TCPEchoEndpoint).SubStatsResource(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mter *TCPEchoResource) Delete(name string) error {
	_, err := mter.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(TCPEchoEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
