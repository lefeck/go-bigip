package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type FixList struct {
	Items    []Fix  `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type Fix struct {
	Kind                     string        `json:"kind"`
	Name                     string        `json:"name"`
	Partition                string        `json:"partition"`
	FullPath                 string        `json:"fullPath"`
	Generation               int           `json:"generation"`
	SelfLink                 string        `json:"selfLink"`
	AppService               string        `json:"appService"`
	DefaultsFrom             string        `json:"defaultsFrom"`
	Description              string        `json:"description"`
	ErrorAction              string        `json:"errorAction"`
	FullLogonParsing         string        `json:"fullLogonParsing"`
	MessageLogPublisher      string        `json:"messageLogPublisher"`
	QuickParsing             string        `json:"quickParsing"`
	ReportLogPublisher       string        `json:"reportLogPublisher"`
	ResponseParsing          string        `json:"responseParsing"`
	SenderTagClass           []interface{} `json:"senderTagClass"`
	StatisticsSampleInterval int           `json:"statisticsSampleInterval"`
}

const FixEndpoint = "fix"

type FixResource struct {
	b *bigip.BigIP
}

func (cr *FixResource) List() (*FixList, error) {
	var items FixList
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FixEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

func (cr *FixResource) Get(fullPathName string) (*Fix, error) {
	var item Fix
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FixEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

func (cr *FixResource) Create(item Fix) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = cr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FixEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (cr *FixResource) Update(fullPathName string, item Fix) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = cr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FixEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (cr *FixResource) Delete(fullPathName string) error {
	_, err := cr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FixEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
