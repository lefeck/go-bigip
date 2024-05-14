package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type SMBList struct {
	Items    []SMB  `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}
type SMB struct {
	AppService   string `json:"appService,omitempty"`
	Debug        string `json:"debug,omitempty"`
	DefaultsFrom string `json:"defaultsFrom,omitempty"`
	Description  string `json:"description,omitempty"`
	Destination  string `json:"destination,omitempty"`
	FullPath     string `json:"fullPath,omitempty"`
	Generation   int    `json:"generation,omitempty"`
	Get          string `json:"get,omitempty"`
	Interval     int    `json:"interval,omitempty"`
	Kind         string `json:"kind,omitempty"`
	ManualResume string `json:"manualResume,omitempty"`
	Name         string `json:"name,omitempty"`
	Partition    string `json:"partition,omitempty"`
	SelfLink     string `json:"selfLink,omitempty"`
	Server       string `json:"server,omitempty"`
	Service      string `json:"service,omitempty"`
	TimeUntilUp  int    `json:"timeUntilUp,omitempty"`
	Timeout      int    `json:"timeout,omitempty"`
	UpInterval   int    `json:"upInterval,omitempty"`
}

const SMBEndpoint = "smb"

type SMBResource struct {
	b *bigip.BigIP
}

func (msr *SMBResource) List() (*SMBList, error) {
	var mscl SMBList
	res, err := msr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SMBEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mscl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mscl, nil
}

func (msr *SMBResource) Get(fullPathName string) (*SMB, error) {
	var msc SMB
	res, err := msr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SMBEndpoint).SubStatsResource(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &msc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &msc, nil
}

func (msr *SMBResource) Create(item SMB) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SMBEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func (msr *SMBResource) Update(name string, item SMB) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SMBEndpoint).SubStatsResource(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msr *SMBResource) Delete(name string) error {
	_, err := msr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SMBEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
