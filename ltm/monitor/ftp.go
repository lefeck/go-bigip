package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type FTPList struct {
	Items    []FTP  `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}
type FTP struct {
	Adaptive                 string `json:"adaptive,omitempty"`
	AdaptiveDivergenceType   string `json:"adaptiveDivergenceType,omitempty"`
	AdaptiveDivergenceValue  int    `json:"adaptiveDivergenceValue,omitempty"`
	AdaptiveLimit            int    `json:"adaptiveLimit,omitempty"`
	AdaptiveSamplingTimespan int    `json:"adaptiveSamplingTimespan,omitempty"`
	AppService               string `json:"appService,omitempty"`
	Debug                    string `json:"debug,omitempty,omitempty"`
	DefaultsFrom             string `json:"defaultsFrom,omitempty"`
	Description              string `json:"description,omitempty"`
	Destination              string `json:"destination,omitempty,omitempty"`
	Filename                 string `json:"filename,omitempty,omitempty"`
	FullPath                 string `json:"fullPath,omitempty,omitempty"`
	Generation               int    `json:"generation,omitempty,omitempty"`
	Interval                 int    `json:"interval,omitempty,omitempty"`
	Kind                     string `json:"kind,omitempty,omitempty"`
	ManualResume             string `json:"manualResume,omitempty,omitempty"`
	Mode                     string `json:"mode,omitempty,omitempty"`
	Name                     string `json:"name,omitempty,omitempty"`
	Partition                string `json:"partition,omitempty,omitempty"`
	SelfLink                 string `json:"selfLink,omitempty,omitempty"`
	TimeUntilUp              int    `json:"timeUntilUp,omitempty,omitempty"`
	Timeout                  int    `json:"timeout,omitempty,omitempty"`
	UpInterval               int    `json:"upInterval,omitempty,omitempty"`
}

const FTPEndpoint = "ftp"

type FTPResource struct {
	b *bigip.BigIP
}

func (mfr *FTPResource) List() (*FTPList, error) {
	var mfcl FTPList
	res, err := mfr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(FTPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mfcl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mfcl, nil
}

func (mfr *FTPResource) Get(fullPathName string) (*FTP, error) {
	var mfc FTP
	res, err := mfr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(FTPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mfc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mfc, nil
}

func (mfr *FTPResource) Create(item FTP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mfr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(FTPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mfr *FTPResource) Update(name string, item FTP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mfr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(FTPEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mfr *FTPResource) Delete(name string) error {
	_, err := mfr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(FTPEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
