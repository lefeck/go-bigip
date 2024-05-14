package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type FirepassList struct {
	Items    []Firepass `json:"items,omitempty"`
	Kind     string     `json:"kind,omitempty"`
	SelfLink string     `json:"selflink,omitempty"`
}

type Firepass struct {
	AppService       string `json:"appService,omitempty"`
	Cipherlist       string `json:"cipherlist,omitempty"`
	ConcurrencyLimit int    `json:"concurrencyLimit,omitempty"`
	DefaultsFrom     string `json:"defaultsFrom,omitempty"`
	Description      string `json:"description,omitempty"`
	Destination      string `json:"destination,omitempty"`
	FullPath         string `json:"fullPath,omitempty"`
	Generation       int    `json:"generation,omitempty"`
	Interval         int    `json:"interval,omitempty"`
	Kind             string `json:"kind,omitempty"`
	MaxLoadAverage   int    `json:"maxLoadAverage,omitempty"`
	Name             string `json:"name,omitempty"`
	Partition        string `json:"partition,omitempty"`
	SelfLink         string `json:"selfLink,omitempty"`
	TimeUntilUp      int    `json:"timeUntilUp,omitempty"`
	Timeout          int    `json:"timeout,omitempty"`
	UpInterval       int    `json:"upInterval,omitempty"`
	Username         string `json:"username,omitempty"`
}

const FirepassEndpoint = "firepass"

type FirepassResource struct {
	b *bigip.BigIP
}

func (mfr *FirepassResource) List() (*FirepassList, error) {
	var mfcl FirepassList
	res, err := mfr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(FirepassEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mfcl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mfcl, nil
}

func (mfr *FirepassResource) Get(fullPathName string) (*Firepass, error) {
	var mfc Firepass
	res, err := mfr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(FirepassEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mfc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mfc, nil
}

func (mfr *FirepassResource) Create(item Firepass) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mfr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(FirepassEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mfr *FirepassResource) Update(name string, item Firepass) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mfr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(FirepassEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mfr *FirepassResource) Delete(name string) error {
	_, err := mfr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(FirepassEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
