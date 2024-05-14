package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type SOAPList struct {
	Items    []SOAP `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}
type SOAP struct {
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

const SOAPEndpoint = "soap"

type SOAPResource struct {
	b *bigip.BigIP
}

func (msr *SOAPResource) List() (*SOAPList, error) {
	var mscl SOAPList
	res, err := msr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SOAPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mscl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mscl, nil
}
func (msr *SOAPResource) Get(fullPathName string) (*SOAP, error) {
	var msc SOAP
	res, err := msr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SOAPEndpoint).SubStatsResource(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &msc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &msc, nil
}

func (msr *SOAPResource) Create(item SOAP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SOAPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msr *SOAPResource) Update(name string, item SOAP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = msr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SOAPEndpoint).SubStatsResource(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (msr *SOAPResource) Delete(name string) error {
	_, err := msr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(SOAPEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
