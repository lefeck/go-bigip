package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type SIPList struct {
	Items    []SIP  `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}
type SIP struct {
	AppService    string `json:"appService,omitempty"`
	Cert          string `json:"cert,omitempty"`
	Cipherlist    string `json:"cipherlist,omitempty"`
	Compatibility string `json:"compatibility,omitempty"`
	Debug         string `json:"debug,omitempty"`
	DefaultsFrom  string `json:"defaultsFrom,omitempty"`
	Description   string `json:"description,omitempty"`
	Destination   string `json:"destination,omitempty"`
	FullPath      string `json:"fullPath,omitempty"`
	Filter        string `json:"filter,omitempty"`
	FilterNeg     string `json:"filterNeg,omitempty"`
	Generation    int    `json:"generation,omitempty"`
	Headers       string `json:"headers,omitempty"`
	Interval      int    `json:"interval,omitempty"`
	Key           string `json:"key,omitempty"`
	Kind          string `json:"kind,omitempty"`
	ManualResume  string `json:"manualResume,omitempty"`
	Mode          string `json:"mode,omitempty"`
	Name          string `json:"name,omitempty"`
	Partition     string `json:"partition,omitempty"`
	Request       string `json:"request,omitempty"`
	SelfLink      string `json:"selfLink,omitempty"`
	TimeUntilUp   int    `json:"timeUntilUp,omitempty"`
	Timeout       int    `json:"timeout,omitempty"`
	UpInterval    int    `json:"upInterval,omitempty"`
}

const SIPEndpoint = "sip"

type SIPResource struct {
	b *bigip.BigIP
}

func (msr *SIPResource) List() (*SIPList, error) {
	var mscl SIPList
	res, err := msr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SIPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mscl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mscl, nil
}

func (msr *SIPResource) Get(fullPathName string) (*SIP, error) {
	var msc SIP
	res, err := msr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SIPEndpoint).SubStatsResource(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &msc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &msc, nil
}

func (msr *SIPResource) Create(item SIP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SIPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msr *SIPResource) Update(name string, item SIP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SIPEndpoint).SubStatsResource(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msr *SIPResource) Delete(name string) error {
	_, err := msr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SIPEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
