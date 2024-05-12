package profile

import "github.com/lefeck/go-bigip"

type Fastl4ConfigList struct {
	Items    []Fastl4Config `json:"items,omitempty"`
	Kind     string         `json:"kind,omitempty"`
	SelfLink string         `json:"selflink,omitempty"`
}

type Fastl4Config struct {
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
