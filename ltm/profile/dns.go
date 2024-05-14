package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type DNSList struct {
	Items    []DNS  `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type DNS struct {
	Kind                  string `json:"kind"`
	Name                  string `json:"name"`
	Partition             string `json:"partition"`
	FullPath              string `json:"fullPath"`
	Generation            int    `json:"generation"`
	SelfLink              string `json:"selfLink"`
	AppService            string `json:"appService"`
	AvrDnsstatSampleRate  int    `json:"avrDnsstatSampleRate"`
	Cache                 string `json:"cache"`
	DefaultsFrom          string `json:"defaultsFrom"`
	DefaultsFromReference struct {
		Link string `json:"link"`
	} `json:"defaultsFromReference,omitempty"`
	Description                   string `json:"description"`
	DNSSecurity                   string `json:"dnsSecurity"`
	DNS64                         string `json:"dns64"`
	DNS64AdditionalSectionRewrite string `json:"dns64AdditionalSectionRewrite"`
	DNS64Prefix                   string `json:"dns64Prefix"`
	Edns0ClientSubnetInsert       string `json:"edns0ClientSubnetInsert"`
	EnableCache                   string `json:"enableCache"`
	EnableDNSExpress              string `json:"enableDnsExpress"`
	EnableDNSFirewall             string `json:"enableDnsFirewall"`
	EnableDnssec                  string `json:"enableDnssec"`
	EnableGtm                     string `json:"enableGtm"`
	EnableHardwareQueryValidation string `json:"enableHardwareQueryValidation"`
	EnableHardwareResponseCache   string `json:"enableHardwareResponseCache"`
	EnableLogging                 string `json:"enableLogging"`
	EnableRapidResponse           string `json:"enableRapidResponse"`
	LogProfile                    string `json:"logProfile"`
	LogProfileReference           struct {
		Link string `json:"link"`
	} `json:"logProfileReference,omitempty"`
	ProcessRd               string `json:"processRd"`
	ProcessXfr              string `json:"processXfr"`
	RapidResponseLastAction string `json:"rapidResponseLastAction"`
	UnhandledQueryAction    string `json:"unhandledQueryAction"`
	UseLocalBind            string `json:"useLocalBind"`
}

const DNSEndpoint = "dns"

type DNSResource struct {
	b *bigip.BigIP
}

func (cr *DNSResource) List() (*DNSList, error) {
	var items DNSList
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(DNSEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

func (cr *DNSResource) Get(fullPathName string) (*DNS, error) {
	var item DNS
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(DNSEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

func (cr *DNSResource) Create(item DNS) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = cr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(DNSEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (cr *DNSResource) Update(fullPathName string, item DNS) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = cr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(DNSEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (cr *DNSResource) Delete(fullPathName string) error {
	_, err := cr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(DNSEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
