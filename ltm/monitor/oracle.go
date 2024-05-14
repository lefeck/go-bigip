package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type OracleList struct {
	Items    []Oracle `json:"items,omitempty"`
	Kind     string   `json:"kind,omitempty"`
	SelfLink string   `json:"selflink,omitempty"`
}
type Oracle struct {
	AppService   string `json:"appService,omitempty"`
	Count        string `json:"count,omitempty"`
	Database     string `json:"database,omitempty"`
	Debug        string `json:"debug,omitempty"`
	DefaultsFrom string `json:"defaultsFrom,omitempty"`
	Description  string `json:"description,omitempty"`
	Destination  string `json:"destination,omitempty"`
	FullPath     string `json:"fullPath,omitempty"`
	Generation   int    `json:"generation,omitempty"`
	Interval     int    `json:"interval,omitempty"`
	Kind         string `json:"kind,omitempty"`
	ManualResume string `json:"manualResume,omitempty"`
	Name         string `json:"name,omitempty"`
	Partition    string `json:"partition,omitempty"`
	Recv         string `json:"recv,omitempty"`
	RecvColumn   string `json:"recvColumn,omitempty"`
	RecvRow      string `json:"recvRow,omitempty"`
	SelfLink     string `json:"selfLink,omitempty"`
	Send         string `json:"send,omitempty"`
	TimeUntilUp  int    `json:"timeUntilUp,omitempty"`
	Timeout      int    `json:"timeout,omitempty"`
	UpInterval   int    `json:"upInterval,omitempty"`
}

const OracleEndpoint = "oracle"

type OracleResource struct {
	b *bigip.BigIP
}

func (mor *OracleResource) List() (*OracleList, error) {
	var mocl OracleList
	res, err := mor.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(OracleEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mocl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mocl, nil
}

func (mor *OracleResource) Get(fullPathName string) (*Oracle, error) {
	var moc Oracle
	res, err := mor.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(OracleEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &moc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &moc, nil
}

func (mor *OracleResource) Create(item Oracle) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mor.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(OracleEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mor *OracleResource) Update(name string, item Oracle) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mor.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(OracleEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mor *OracleResource) Delete(name string) error {
	_, err := mor.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(OracleEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
