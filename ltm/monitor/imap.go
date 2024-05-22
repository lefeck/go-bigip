package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type IMAPList struct {
	Items    []IMAP `json:"items"`
	Kind     string `json:"kind"`
	SelfLink string `json:"selflink"`
}

type IMAP struct {
	Debug        string `json:"debug"`
	DefaultsFrom string `json:"defaultsFrom,omitempty"`
	Description  string `json:"description,omitempty"`
	Destination  string `json:"destination"`
	Folder       string `json:"folder"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	Interval     int    `json:"interval"`
	Kind         string `json:"kind"`
	ManualResume string `json:"manualResume"`
	Name         string `json:"name"`
	Partition    string `json:"partition"`
	SelfLink     string `json:"selfLink"`
	TimeUntilUp  int    `json:"timeUntilUp"`
	Timeout      int    `json:"timeout"`
	UpInterval   int    `json:"upInterval"`
}

const IMAPEndpoint = "imap"

type IMAPResource struct {
	b *bigip.BigIP
}

func (mir *IMAPResource) List() (*IMAPList, error) {
	var micl IMAPList
	res, err := mir.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(IMAPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &micl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &micl, nil
}

func (mir *IMAPResource) Get(fullPathName string) (*IMAP, error) {
	var mic IMAP
	res, err := mir.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(IMAPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mic); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mic, nil
}

func (mir *IMAPResource) Create(item IMAP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mir.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(IMAPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mir *IMAPResource) Update(name string, item IMAP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mir.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(IMAPEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mir *IMAPResource) Delete(name string) error {
	_, err := mir.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(IMAPEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
