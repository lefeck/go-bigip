package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type FIXList struct {
	Items    []FIX  `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type FIX struct {
	Kind                     string        `json:"kind,omitempty"`
	Name                     string        `json:"name,omitempty"`
	Partition                string        `json:"partition,omitempty"`
	FullPath                 string        `json:"fullPath,omitempty"`
	Generation               int           `json:"generation,omitempty"`
	SelfLink                 string        `json:"selfLink,omitempty"`
	AppService               string        `json:"appService,omitempty"`
	DefaultsFrom             string        `json:"defaultsFrom,omitempty"`
	Description              string        `json:"description,omitempty"`
	ErrorAction              string        `json:"errorAction,omitempty"`
	FullLogonParsing         string        `json:"fullLogonParsing,omitempty"`
	MessageLogPublisher      string        `json:"messageLogPublisher,omitempty"`
	QuickParsing             string        `json:"quickParsing,omitempty"`
	ReportLogPublisher       string        `json:"reportLogPublisher,omitempty"`
	ResponseParsing          string        `json:"responseParsing,omitempty"`
	SenderTagClass           []interface{} `json:"senderTagClass,omitempty"`
	StatisticsSampleInterval int           `json:"statisticsSampleInterval,omitempty"`
}

const FIXEndpoint = "fix"

type FIXResource struct {
	b *bigip.BigIP
}

func (cr *FIXResource) List() (*FIXList, error) {
	var items FIXList
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FIXEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

func (cr *FIXResource) Get(fullPathName string) (*FIX, error) {
	var item FIX
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FIXEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

func (cr *FIXResource) Create(item FIX) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FIXEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (cr *FIXResource) Update(fullPathName string, item FIX) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FIXEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (cr *FIXResource) Delete(fullPathName string) error {
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FIXEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
