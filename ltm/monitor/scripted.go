package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type ScriptedList struct {
	Items    []Scripted `json:"items,omitempty"`
	Kind     string     `json:"kind,omitempty"`
	SelfLink string     `json:"selflink,omitempty"`
}

type Scripted struct {
	AppService   string `json:"appService,omitempty"`
	Debug        string `json:"debug,omitempty"`
	DefaultsFrom string `json:"defaultsFrom,omitempty"`
	Description  string `json:"description,omitempty"`
	Destination  string `json:"destination,omitempty"`
	Filename     string `json:"filename,omitempty"`
	FullPath     string `json:"fullPath,omitempty"`
	Generation   int    `json:"generation,omitempty"`
	Interval     int    `json:"interval,omitempty"`
	Kind         string `json:"kind,omitempty"`
	ManualResume string `json:"manualResume,omitempty"`
	Name         string `json:"name,omitempty"`
	Partition    string `json:"partition,omitempty"`
	SelfLink     string `json:"selfLink,omitempty"`
	TimeUntilUp  int    `json:"timeUntilUp,omitempty"`
	Timeout      int    `json:"timeout,omitempty"`
	UpInterval   int    `json:"upInterval,omitempty"`
}

const ScriptedEndpoint = "scripted"

type ScriptedResource struct {
	b *bigip.BigIP
}

func (msr *ScriptedResource) List() (*ScriptedList, error) {
	var mscl ScriptedList
	res, err := msr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(ScriptedEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mscl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mscl, nil
}

func (msr *ScriptedResource) Get(fullPathName string) (*Scripted, error) {
	var msc Scripted
	res, err := msr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(ScriptedEndpoint).SubStatsResource(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &msc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &msc, nil
}

func (msr *ScriptedResource) Create(item Scripted) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(ScriptedEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msr *ScriptedResource) Update(name string, item Scripted) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(ScriptedEndpoint).SubStatsResource(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msr *ScriptedResource) Delete(name string) error {
	_, err := msr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(ScriptedEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
