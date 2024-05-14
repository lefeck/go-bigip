package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type RealServerList struct {
	Items    []RealServer `json:"items,omitempty"`
	Kind     string       `json:"kind,omitempty"`
	SelfLink string       `json:"selflink,omitempty"`
}
type RealServer struct {
	Agent        string `json:"agent,omitempty"`
	AppService   string `json:"appService,omitempty"`
	DefaultsFrom string `json:"defaultsFrom,omitempty"`
	Description  string `json:"description,omitempty"`
	Destination  string `json:"destination,omitempty"`
	FullPath     string `json:"fullPath,omitempty"`
	Generation   int    `json:"generation,omitempty"`
	Interval     int    `json:"interval,omitempty"`
	Kind         string `json:"kind,omitempty"`
	Method       string `json:"method,omitempty"`
	Metrics      string `json:"metrics,omitempty"`
	Name         string `json:"name,omitempty"`
	Partition    string `json:"partition,omitempty"`
	SelfLink     string `json:"selfLink,omitempty"`
	TimeUntilUp  int    `json:"timeUntilUp,omitempty"`
	Timeout      int    `json:"timeout,omitempty"`
	TmCommand    string `json:"tmCommand,omitempty"`
}

const RealServerEndpoint = "real-server"

type RealServerResource struct {
	b *bigip.BigIP
}

func (mrsr *RealServerResource) List() (*RealServerList, error) {
	var mrscl RealServerList
	res, err := mrsr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(RealServerEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mrscl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mrscl, nil
}

func (mrsr *RealServerResource) Get(fullPathName string) (*RealServer, error) {
	var mrsc RealServer
	res, err := mrsr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(RealServerEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mrsc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mrsc, nil
}

func (mrsr *RealServerResource) Create(item RealServer) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mrsr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(RealServerEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mrsr *RealServerResource) Update(name string, item RealServer) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mrsr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(RealServerEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mrsr *RealServerResource) Delete(name string) error {
	_, err := mrsr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(RealServerEndpoint).SubStatsResource(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
