package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"github.com/lefeck/go-bigip/ltm"
	"strings"
)

type MonitorSOAPConfigList struct {
	Items    []MonitorSOAPConfig `json:"items,omitempty"`
	Kind     string              `json:"kind,omitempty"`
	SelfLink string              `json:"selflink,omitempty"`
}

type MonitorSOAPConfig struct {
	AppService     string `json:"appService,omitempty"`
	Debug          string `json:"debug,omitempty"`
	DefaultsFrom   string `json:"defaultsFrom,omitempty"`
	Description    string `json:"description,omitempty"`
	Destination    string `json:"destination,omitempty"`
	ExpectFault    string `json:"expectFault,omitempty"`
	FullPath       string `json:"fullPath,omitempty"`
	Generation     int    `json:"generation,omitempty"`
	Interval       int    `json:"interval,omitempty"`
	Kind           string `json:"kind,omitempty"`
	ManualResume   string `json:"manualResume,omitempty"`
	Method         string `json:"method,omitempty"`
	Name           string `json:"name,omitempty"`
	Namespace      string `json:"namespace,omitempty"`
	ParameterName  string `json:"parameterName,omitempty"`
	ParameterType  string `json:"parameterType,omitempty"`
	ParameterValue string `json:"parameterValue,omitempty"`
	Partition      string `json:"partition,omitempty"`
	Protocol       string `json:"protocol,omitempty"`
	ReturnType     string `json:"returnType,omitempty"`
	ReturnValue    string `json:"returnValue,omitempty"`
	SelfLink       string `json:"selfLink,omitempty"`
	SoapAction     string `json:"soapAction,omitempty"`
	TimeUntilUp    int    `json:"timeUntilUp,omitempty"`
	Timeout        int    `json:"timeout,omitempty"`
	UpInterval     int    `json:"upInterval,omitempty"`
	UrlPath        string `json:"urlPath,omitempty"`
}

const MonitorSOAPEndpoint = "/monitor/soap"

type MonitorSOAPResource struct {
	b *bigip.BigIP
}

func (msr *MonitorSOAPResource) List() (*MonitorSOAPConfigList, error) {
	var mscl MonitorSOAPConfigList
	res, err := msr.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSOAPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mscl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mscl, nil
}
func (msr *MonitorSOAPResource) Get(fullPathName string) (*MonitorSOAPConfig, error) {
	var msc MonitorSOAPConfig
	res, err := msr.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSOAPEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &msc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &msc, nil
}

func (msr *MonitorSOAPResource) Create(item MonitorSOAPConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msr.b.RestClient.Post().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSOAPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msr *MonitorSOAPResource) Update(name string, item MonitorSOAPConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msr.b.RestClient.Put().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSOAPEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msr *MonitorSOAPResource) Delete(name string) error {
	_, err := msr.b.RestClient.Delete().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorSOAPEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
