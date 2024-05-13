package profile

import "github.com/lefeck/go-bigip"

type HTTPList struct {
	Items    []HTTP `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type HTTP struct {
	Kind           string        `json:"kind"`
	Name           string        `json:"name"`
	Partition      string        `json:"partition"`
	FullPath       string        `json:"fullPath"`
	Generation     int           `json:"generation"`
	SelfLink       string        `json:"selfLink"`
	AcceptXff      string        `json:"acceptXff"`
	AppService     string        `json:"appService"`
	BasicAuthRealm string        `json:"basicAuthRealm"`
	DefaultsFrom   string        `json:"defaultsFrom"`
	Description    string        `json:"description"`
	EncryptCookies []interface{} `json:"encryptCookies"`
	Enforcement    struct {
		AllowWsHeaderName     string   `json:"allowWsHeaderName"`
		ExcessClientHeaders   string   `json:"excessClientHeaders"`
		ExcessServerHeaders   string   `json:"excessServerHeaders"`
		KnownMethods          []string `json:"knownMethods"`
		MaxHeaderCount        int      `json:"maxHeaderCount"`
		MaxHeaderSize         int      `json:"maxHeaderSize"`
		MaxRequests           int      `json:"maxRequests"`
		OversizeClientHeaders string   `json:"oversizeClientHeaders"`
		OversizeServerHeaders string   `json:"oversizeServerHeaders"`
		Pipeline              string   `json:"pipeline"`
		RfcCompliance         string   `json:"rfcCompliance"`
		TruncatedRedirects    string   `json:"truncatedRedirects"`
		UnknownMethod         string   `json:"unknownMethod"`
	} `json:"enforcement"`
	ExplicitProxy struct {
		BadRequestMessage      string        `json:"badRequestMessage"`
		BadResponseMessage     string        `json:"badResponseMessage"`
		ConnectErrorMessage    string        `json:"connectErrorMessage"`
		DefaultConnectHandling string        `json:"defaultConnectHandling"`
		DNSErrorMessage        string        `json:"dnsErrorMessage"`
		DNSResolver            string        `json:"dnsResolver"`
		HostNames              []interface{} `json:"hostNames"`
		Ipv6                   string        `json:"ipv6"`
		RouteDomain            string        `json:"routeDomain"`
		TunnelName             string        `json:"tunnelName"`
		TunnelOnAnyRequest     string        `json:"tunnelOnAnyRequest"`
	} `json:"explicitProxy"`
	FallbackHost        string        `json:"fallbackHost"`
	FallbackStatusCodes []interface{} `json:"fallbackStatusCodes"`
	HeaderErase         string        `json:"headerErase"`
	HeaderInsert        string        `json:"headerInsert"`
	Hsts                struct {
		IncludeSubdomains string `json:"includeSubdomains"`
		MaximumAge        int    `json:"maximumAge"`
		Mode              string `json:"mode"`
		Preload           string `json:"preload"`
	} `json:"hsts"`
	InsertXforwardedFor       string        `json:"insertXforwardedFor"`
	LwsSeparator              string        `json:"lwsSeparator"`
	LwsWidth                  int           `json:"lwsWidth"`
	OneconnectStatusReuse     string        `json:"oneconnectStatusReuse"`
	OneconnectTransformations string        `json:"oneconnectTransformations"`
	ProxyType                 string        `json:"proxyType"`
	RedirectRewrite           string        `json:"redirectRewrite"`
	RequestChunking           string        `json:"requestChunking"`
	ResponseChunking          string        `json:"responseChunking"`
	ResponseHeadersPermitted  []interface{} `json:"responseHeadersPermitted"`
	ServerAgentName           string        `json:"serverAgentName"`
	Sflow                     struct {
		PollInterval       int    `json:"pollInterval"`
		PollIntervalGlobal string `json:"pollIntervalGlobal"`
		SamplingRate       int    `json:"samplingRate"`
		SamplingRateGlobal string `json:"samplingRateGlobal"`
	} `json:"sflow"`
	ViaHostName           string        `json:"viaHostName"`
	ViaRequest            string        `json:"viaRequest"`
	ViaResponse           string        `json:"viaResponse"`
	XffAlternativeNames   []interface{} `json:"xffAlternativeNames"`
	DefaultsFromReference struct {
		Link string `json:"link"`
	} `json:"defaultsFromReference,omitempty"`
}

const HTTPEndpoint = "http"

type HTTPResource struct {
	b *bigip.BigIP
}
