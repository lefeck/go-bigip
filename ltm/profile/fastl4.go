package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type FastL4List struct {
	Items    []FastL4 `json:"items,omitempty"`
	Kind     string   `json:"kind,omitempty"`
	SelfLink string   `json:"selflink,omitempty"`
}

type FastL4 struct {
	Kind                  string `json:"kind,omitempty"`
	Name                  string `json:"name,omitempty"`
	Partition             string `json:"partition,omitempty"`
	FullPath              string `json:"fullPath,omitempty"`
	Generation            int    `json:"generation,omitempty"`
	SelfLink              string `json:"selfLink,omitempty"`
	AppService            string `json:"appService,omitempty"`
	ClientTimeout         int    `json:"clientTimeout,omitempty"`
	DefaultsFrom          string `json:"defaultsFrom,omitempty"`
	DefaultsFromReference struct {
		Link string `json:"link,omitempty"`
	} `json:"defaultsFromReference,omitempty"`
	Description                 string `json:"description,omitempty"`
	ExplicitFlowMigration       string `json:"explicitFlowMigration,omitempty"`
	HardwareSynCookie           string `json:"hardwareSynCookie,omitempty"`
	IdleTimeout                 string `json:"idleTimeout,omitempty"`
	IPDfMode                    string `json:"ipDfMode,omitempty"`
	IPTosToClient               string `json:"ipTosToClient,omitempty"`
	IPTosToServer               string `json:"ipTosToServer,omitempty"`
	IPTTLMode                   string `json:"ipTtlMode,omitempty"`
	IPTTLV4                     int    `json:"ipTtlV4,omitempty"`
	IPTTLV6                     int    `json:"ipTtlV6,omitempty"`
	KeepAliveInterval           string `json:"keepAliveInterval,omitempty"`
	LateBinding                 string `json:"lateBinding,omitempty"`
	LinkQosToClient             string `json:"linkQosToClient,omitempty"`
	LinkQosToServer             string `json:"linkQosToServer,omitempty"`
	LooseClose                  string `json:"looseClose,omitempty"`
	LooseInitialization         string `json:"looseInitialization,omitempty"`
	MssOverride                 int    `json:"mssOverride,omitempty"`
	OtherPvaClientpktsThreshold int    `json:"otherPvaClientpktsThreshold,omitempty"`
	OtherPvaOffloadDirection    string `json:"otherPvaOffloadDirection,omitempty"`
	OtherPvaServerpktsThreshold int    `json:"otherPvaServerpktsThreshold,omitempty"`
	OtherPvaWhentoOffload       string `json:"otherPvaWhentoOffload,omitempty"`
	PriorityToClient            string `json:"priorityToClient,omitempty"`
	PriorityToServer            string `json:"priorityToServer,omitempty"`
	PvaAcceleration             string `json:"pvaAcceleration,omitempty"`
	PvaDynamicClientPackets     int    `json:"pvaDynamicClientPackets,omitempty"`
	PvaDynamicServerPackets     int    `json:"pvaDynamicServerPackets,omitempty"`
	PvaFlowAging                string `json:"pvaFlowAging,omitempty"`
	PvaFlowEvict                string `json:"pvaFlowEvict,omitempty"`
	PvaOffloadDynamic           string `json:"pvaOffloadDynamic,omitempty"`
	PvaOffloadDynamicPriority   string `json:"pvaOffloadDynamicPriority,omitempty"`
	PvaOffloadInitialPriority   string `json:"pvaOffloadInitialPriority,omitempty"`
	PvaOffloadState             string `json:"pvaOffloadState,omitempty"`
	ReassembleFragments         string `json:"reassembleFragments,omitempty"`
	ReceiveWindowSize           int    `json:"receiveWindowSize,omitempty"`
	ResetOnTimeout              string `json:"resetOnTimeout,omitempty"`
	RttFromClient               string `json:"rttFromClient,omitempty"`
	RttFromServer               string `json:"rttFromServer,omitempty"`
	ServerSack                  string `json:"serverSack,omitempty"`
	ServerTimestamp             string `json:"serverTimestamp,omitempty"`
	SoftwareSynCookie           string `json:"softwareSynCookie,omitempty"`
	SynCookieDsrFlowResetBy     string `json:"synCookieDsrFlowResetBy,omitempty"`
	SynCookieEnable             string `json:"synCookieEnable,omitempty"`
	SynCookieMss                int    `json:"synCookieMss,omitempty"`
	SynCookieWhitelist          string `json:"synCookieWhitelist,omitempty"`
	TCPCloseTimeout             string `json:"tcpCloseTimeout,omitempty"`
	TCPGenerateIsn              string `json:"tcpGenerateIsn,omitempty"`
	TCPHandshakeTimeout         string `json:"tcpHandshakeTimeout,omitempty"`
	TCPPvaOffloadDirection      string `json:"tcpPvaOffloadDirection,omitempty"`
	TCPPvaWhentoOffload         string `json:"tcpPvaWhentoOffload,omitempty"`
	TCPStripSack                string `json:"tcpStripSack,omitempty"`
	TCPTimeWaitTimeout          int    `json:"tcpTimeWaitTimeout,omitempty"`
	TCPTimestampMode            string `json:"tcpTimestampMode,omitempty"`
	TCPWscaleMode               string `json:"tcpWscaleMode,omitempty"`
	TimeoutRecovery             string `json:"timeoutRecovery,omitempty"`
}

const FastL4Endpoint = "fastl4"

type FastL4Resource struct {
	b *bigip.BigIP
}

func (cr *FastL4Resource) List() (*FastL4List, error) {
	var items FastL4List
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FastL4Endpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

func (cr *FastL4Resource) Get(fullPathName string) (*FastL4, error) {
	var item FastL4
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FastL4Endpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

func (cr *FastL4Resource) Create(item FastL4) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FastL4Endpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (cr *FastL4Resource) Update(fullPathName string, item FastL4) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FastL4Endpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (cr *FastL4Resource) Delete(fullPathName string) error {
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FastL4Endpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
