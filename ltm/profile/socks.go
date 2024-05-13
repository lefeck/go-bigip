package profile

import "github.com/lefeck/go-bigip"

type SocksList struct {
	Items    []WebSocket `json:"items,omitempty"`
	Kind     string      `json:"kind,omitempty"`
	SelfLink string      `json:"selflink,omitempty"`
}

type Socks struct {
	Kind                   string   `json:"kind"`
	Name                   string   `json:"name"`
	Partition              string   `json:"partition"`
	FullPath               string   `json:"fullPath"`
	Generation             int      `json:"generation"`
	SelfLink               string   `json:"selfLink"`
	AppService             string   `json:"appService"`
	DefaultConnectHandling string   `json:"defaultConnectHandling"`
	DefaultsFrom           string   `json:"defaultsFrom"`
	Description            string   `json:"description"`
	DNSResolver            string   `json:"dnsResolver"`
	Ipv6                   string   `json:"ipv6"`
	ProtocolVersions       []string `json:"protocolVersions"`
	RouteDomain            string   `json:"routeDomain"`
	RouteDomainReference   struct {
		Link string `json:"link"`
	} `json:"routeDomainReference"`
	TunnelName          string `json:"tunnelName"`
	TunnelNameReference struct {
		Link string `json:"link"`
	} `json:"tunnelNameReference"`
}

const SocksEndpoint = "socks"

type SocksResource struct {
	b *bigip.BigIP
}
