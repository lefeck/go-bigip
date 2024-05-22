package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type WAPList struct {
	Items    []WAP  `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type WAP struct {
	AccountingNode string `json:"accountingNode,omitempty"`
	AccountingPort string `json:"accountingPort,omitempty"`
	AppService     string `json:"appService,omitempty"`
	CallId         string `json:"callId,omitempty"`
	Debug          string `json:"debug,omitempty"`
	DefaultsFrom   string `json:"defaultsFrom,omitempty"`
	Description    string `json:"description,omitempty"`
	Destination    string `json:"destination,omitempty"`
	FramedAddress  string `json:"framedAddress,omitempty"`
	FullPath       string `json:"fullPath,omitempty"`
	Generation     int    `json:"generation,omitempty"`
	Interval       int    `json:"interval,omitempty"`
	Kind           string `json:"kind,omitempty"`
	ManualResume   string `json:"manualResume,omitempty"`
	Name           string `json:"name,omitempty"`
	Partition      string `json:"partition,omitempty"`
	Recv           string `json:"recv,omitempty"`
	SelfLink       string `json:"selfLink,omitempty"`
	Send           string `json:"send,omitempty"`
	ServerId       string `json:"serverId,omitempty"`
	SessionId      string `json:"sessionId,omitempty"`
	TimeUntilUp    int    `json:"timeUntilUp,omitempty"`
	Timeout        int    `json:"timeout,omitempty"`
	UpInterval     int    `json:"upInterval,omitempty"`
}

const WAPEndpoint = "wap"

type WAPResource struct {
	b *bigip.BigIP
}

func (mwr *WAPResource) List() (*WAPList, error) {
	var mwcl WAPList
	res, err := mwr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(WAPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mwcl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mwcl, nil
}

func (mwr *WAPResource) Get(fullPathName string) (*WAP, error) {
	var mwc WAP
	res, err := mwr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(WAPEndpoint).SubStatsResource(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mwc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mwc, nil
}

func (mwr *WAPResource) Create(item WAP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mwr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(WAPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mwr *WAPResource) Update(name string, item WAP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mwr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(WAPEndpoint).SubStatsResource(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mwr *WAPResource) Delete(name string) error {
	_, err := mwr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(WAPEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
