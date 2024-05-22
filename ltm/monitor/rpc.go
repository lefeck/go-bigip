package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type RPCList struct {
	Items    []RPC  `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type RPC struct {
	AppService    string `json:"appService,omitempty"`
	Debug         string `json:"debug,omitempty"`
	DefaultsFrom  string `json:"defaultsFrom,omitempty"`
	Description   string `json:"description,omitempty"`
	Destination   string `json:"destination,omitempty"`
	FullPath      string `json:"fullPath,omitempty"`
	Generation    int    `json:"generation,omitempty"`
	Interval      int    `json:"interval,omitempty"`
	Kind          string `json:"kind,omitempty"`
	ManualResume  string `json:"manualResume,omitempty"`
	Mode          string `json:"mode,omitempty"`
	Name          string `json:"name,omitempty"`
	Partition     string `json:"partition,omitempty"`
	ProgramNumber string `json:"programNumber,omitempty"`
	SelfLink      string `json:"selfLink,omitempty"`
	TimeUntilUp   int    `json:"timeUntilUp,omitempty"`
	Timeout       int    `json:"timeout,omitempty"`
	UpInterval    int    `json:"upInterval,omitempty"`
	VersionNumber string `json:"versionNumber,omitempty"`
}

const RPCEndpoint = "rpc"

type RPCResource struct {
	b *bigip.BigIP
}

func (mrr *RPCResource) List() (*RPCList, error) {
	var mrrcl RPCList
	res, err := mrr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(RPCEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mrrcl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mrrcl, nil
}

func (mrr *RPCResource) Get(fullPathName string) (*RPC, error) {
	var mrrc RPC
	res, err := mrr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(RPCEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mrrc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mrrc, nil
}

func (mrr *RPCResource) Create(item RPC) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mrr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(RPCEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mrr *RPCResource) Update(name string, item RPC) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mrr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(RPCEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mrr *RPCResource) Delete(name string) error {
	_, err := mrr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(RPCEndpoint).SubStatsResource(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
