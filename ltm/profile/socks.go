package profile

type socks struct {
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
