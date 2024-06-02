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
	Kind                  string `json:"kind,omitempty"`
	Name                  string `json:"name,omitempty"`
	Partition             string `json:"partition,omitempty"`
	FullPath              string `json:"fullPath,omitempty"`
	Generation            int    `json:"generation,omitempty"`
	SelfLink              string `json:"selfLink,omitempty"`
	AppService            string `json:"appService,omitempty"`
	AvrDnsstatSampleRate  int    `json:"avrDnsstatSampleRate,omitempty"`
	Cache                 string `json:"cache,omitempty"`
	DefaultsFrom          string `json:"defaultsFrom,omitempty"`
	DefaultsFromReference struct {
		Link string `json:"link"`
	} `json:"defaultsFromReference,omitempty"`
	Description                   string `json:"description,omitempty"`
	DNSSecurity                   string `json:"dnsSecurity,omitempty"`
	DNS64                         string `json:"dns64,omitempty"`
	DNS64AdditionalSectionRewrite string `json:"dns64AdditionalSectionRewrite,omitempty"`
	DNS64Prefix                   string `json:"dns64Prefix,omitempty"`
	Edns0ClientSubnetInsert       string `json:"edns0ClientSubnetInsert,omitempty"`
	EnableCache                   string `json:"enableCache,omitempty"`
	EnableDNSExpress              string `json:"enableDnsExpress,omitempty"`
	EnableDNSFirewall             string `json:"enableDnsFirewall,omitempty"`
	EnableDnssec                  string `json:"enableDnssec,omitempty"`
	EnableGtm                     string `json:"enableGtm,omitempty"`
	EnableHardwareQueryValidation string `json:"enableHardwareQueryValidation,omitempty"`
	EnableHardwareResponseCache   string `json:"enableHardwareResponseCache,omitempty"`
	EnableLogging                 string `json:"enableLogging,omitempty"`
	EnableRapidResponse           string `json:"enableRapidResponse,omitempty"`
	LogProfile                    string `json:"logProfile,omitempty"`
	LogProfileReference           struct {
		Link string `json:"link,omitempty"`
	} `json:"logProfileReference,omitempty"`
	ProcessRd               string `json:"processRd,omitempty"`
	ProcessXfr              string `json:"processXfr,omitempty"`
	RapidResponseLastAction string `json:"rapidResponseLastAction,omitempty"`
	UnhandledQueryAction    string `json:"unhandledQueryAction,omitempty"`
	UseLocalBind            string `json:"useLocalBind,omitempty"`
}

const DNSEndpoint = "dns"

type DNSResource struct {
	b *bigip.BigIP
}

func (cr *DNSResource) List() (*DNSList, error) {
	var items DNSList
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
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
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
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
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
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
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(DNSEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (cr *DNSResource) Delete(fullPathName string) error {
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(DNSEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
