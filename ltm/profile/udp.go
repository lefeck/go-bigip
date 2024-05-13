package profile

import "github.com/lefeck/go-bigip"

type UDPList struct {
	Items    []UDP  `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type UDP struct {
	Kind                  string `json:"kind"`
	Name                  string `json:"name"`
	Partition             string `json:"partition"`
	FullPath              string `json:"fullPath"`
	Generation            int    `json:"generation"`
	SelfLink              string `json:"selfLink"`
	AllowNoPayload        string `json:"allowNoPayload"`
	AppService            string `json:"appService"`
	BufferMaxBytes        int    `json:"bufferMaxBytes"`
	BufferMaxPackets      int    `json:"bufferMaxPackets"`
	DatagramLoadBalancing string `json:"datagramLoadBalancing"`
	DefaultsFrom          string `json:"defaultsFrom"`
	Description           string `json:"description"`
	IdleTimeout           string `json:"idleTimeout"`
	IPDfMode              string `json:"ipDfMode"`
	IPTosToClient         string `json:"ipTosToClient"`
	IPTTLMode             string `json:"ipTtlMode"`
	IPTTLV4               int    `json:"ipTtlV4"`
	IPTTLV6               int    `json:"ipTtlV6"`
	LinkQosToClient       string `json:"linkQosToClient"`
	NoChecksum            string `json:"noChecksum"`
	ProxyMss              string `json:"proxyMss"`
	SendBufferSize        int    `json:"sendBufferSize"`
	DefaultsFromReference struct {
		Link string `json:"link"`
	} `json:"defaultsFromReference,omitempty"`
}

const UDPEndpoint = "udp"

type UDPResource struct {
	b *bigip.BigIP
}
