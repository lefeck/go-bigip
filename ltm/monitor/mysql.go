package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type MySQLList struct {
	Items    []MySQL `json:"items,omitempty"`
	Kind     string  `json:"kind,omitempty"`
	SelfLink string  `json:"selflink,omitempty"`
}
type MySQL struct {
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

const MySQLEndpoint = "mysql"

type MySQLResource struct {
	b *bigip.BigIP
}

func (mmr *MySQLResource) List() (*MySQLList, error) {
	var mmcl MySQLList
	res, err := mmr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(MySQLEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mmcl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mmcl, nil
}

func (mmr *MySQLResource) Get(fullPathName string) (*MySQL, error) {
	var mmc MySQL
	res, err := mmr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(MySQLEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mmc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mmc, nil
}

func (mmr *MySQLResource) Create(item MySQL) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mmr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(MySQLEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mmr *MySQLResource) Update(name string, item MySQL) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mmr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(MySQLEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mmr *MySQLResource) Delete(name string) error {
	_, err := mmr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(MySQLEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
