package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type Fastl4List struct {
	Items    []Fastl4 `json:"items,omitempty"`
	Kind     string   `json:"kind,omitempty"`
	SelfLink string   `json:"selflink,omitempty"`
}

type Fastl4 struct {
	Kind                  string `json:"kind"`
	Name                  string `json:"name"`
	Partition             string `json:"partition"`
	FullPath              string `json:"fullPath"`
	Generation            int    `json:"generation"`
	SelfLink              string `json:"selfLink"`
	AppService            string `json:"appService"`
	ClientTimeout         int    `json:"clientTimeout"`
	DefaultsFrom          string `json:"defaultsFrom"`
	DefaultsFromReference struct {
		Link string `json:"link"`
	} `json:"defaultsFromReference"`
	Description                 string `json:"description"`
	ExplicitFlowMigration       string `json:"explicitFlowMigration"`
	HardwareSynCookie           string `json:"hardwareSynCookie"`
	IdleTimeout                 string `json:"idleTimeout"`
	IPDfMode                    string `json:"ipDfMode"`
	IPTosToClient               string `json:"ipTosToClient"`
	IPTosToServer               string `json:"ipTosToServer"`
	IPTTLMode                   string `json:"ipTtlMode"`
	IPTTLV4                     int    `json:"ipTtlV4"`
	IPTTLV6                     int    `json:"ipTtlV6"`
	KeepAliveInterval           string `json:"keepAliveInterval"`
	LateBinding                 string `json:"lateBinding"`
	LinkQosToClient             string `json:"linkQosToClient"`
	LinkQosToServer             string `json:"linkQosToServer"`
	LooseClose                  string `json:"looseClose"`
	LooseInitialization         string `json:"looseInitialization"`
	MssOverride                 int    `json:"mssOverride"`
	OtherPvaClientpktsThreshold int    `json:"otherPvaClientpktsThreshold"`
	OtherPvaOffloadDirection    string `json:"otherPvaOffloadDirection"`
	OtherPvaServerpktsThreshold int    `json:"otherPvaServerpktsThreshold"`
	OtherPvaWhentoOffload       string `json:"otherPvaWhentoOffload"`
	PriorityToClient            string `json:"priorityToClient"`
	PriorityToServer            string `json:"priorityToServer"`
	PvaAcceleration             string `json:"pvaAcceleration"`
	PvaDynamicClientPackets     int    `json:"pvaDynamicClientPackets"`
	PvaDynamicServerPackets     int    `json:"pvaDynamicServerPackets"`
	PvaFlowAging                string `json:"pvaFlowAging"`
	PvaFlowEvict                string `json:"pvaFlowEvict"`
	PvaOffloadDynamic           string `json:"pvaOffloadDynamic"`
	PvaOffloadDynamicPriority   string `json:"pvaOffloadDynamicPriority"`
	PvaOffloadInitialPriority   string `json:"pvaOffloadInitialPriority"`
	PvaOffloadState             string `json:"pvaOffloadState"`
	ReassembleFragments         string `json:"reassembleFragments"`
	ReceiveWindowSize           int    `json:"receiveWindowSize"`
	ResetOnTimeout              string `json:"resetOnTimeout"`
	RttFromClient               string `json:"rttFromClient"`
	RttFromServer               string `json:"rttFromServer"`
	ServerSack                  string `json:"serverSack"`
	ServerTimestamp             string `json:"serverTimestamp"`
	SoftwareSynCookie           string `json:"softwareSynCookie"`
	SynCookieDsrFlowResetBy     string `json:"synCookieDsrFlowResetBy"`
	SynCookieEnable             string `json:"synCookieEnable"`
	SynCookieMss                int    `json:"synCookieMss"`
	SynCookieWhitelist          string `json:"synCookieWhitelist"`
	TCPCloseTimeout             string `json:"tcpCloseTimeout"`
	TCPGenerateIsn              string `json:"tcpGenerateIsn"`
	TCPHandshakeTimeout         string `json:"tcpHandshakeTimeout"`
	TCPPvaOffloadDirection      string `json:"tcpPvaOffloadDirection"`
	TCPPvaWhentoOffload         string `json:"tcpPvaWhentoOffload"`
	TCPStripSack                string `json:"tcpStripSack"`
	TCPTimeWaitTimeout          int    `json:"tcpTimeWaitTimeout"`
	TCPTimestampMode            string `json:"tcpTimestampMode"`
	TCPWscaleMode               string `json:"tcpWscaleMode"`
	TimeoutRecovery             string `json:"timeoutRecovery"`
}

const Fastl4Endpoint = "fastl4"

type Fastl4Resource struct {
	b *bigip.BigIP
}

func (cr *Fastl4Resource) List() (*Fastl4List, error) {
	var items Fastl4List
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(Fastl4Endpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

func (cr *Fastl4Resource) Get(fullPathName string) (*Fastl4, error) {
	var item Fastl4
	res, err := cr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(Fastl4Endpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

func (cr *Fastl4Resource) Create(item Fastl4) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = cr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(Fastl4Endpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (cr *Fastl4Resource) Update(fullPathName string, item Fastl4) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = cr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(Fastl4Endpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (cr *Fastl4Resource) Delete(fullPathName string) error {
	_, err := cr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(Fastl4Endpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
