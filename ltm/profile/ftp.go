package profile

import "github.com/lefeck/go-bigip"

type FtpConfigList struct {
	Items    []FtpConfig `json:"items,omitempty"`
	Kind     string      `json:"kind,omitempty"`
	SelfLink string      `json:"selflink,omitempty"`
}

type FtpConfig struct {
	Kind                   string `json:"kind"`
	Name                   string `json:"name"`
	Partition              string `json:"partition"`
	FullPath               string `json:"fullPath"`
	Generation             int    `json:"generation"`
	SelfLink               string `json:"selfLink"`
	AllowActiveMode        string `json:"allowActiveMode"`
	AllowFtps              string `json:"allowFtps"`
	AppService             string `json:"appService"`
	DefaultsFrom           string `json:"defaultsFrom"`
	Description            string `json:"description"`
	EnforceTLSSessionReuse string `json:"enforceTlsSessionReuse"`
	FtpsMode               string `json:"ftpsMode"`
	InheritParentProfile   string `json:"inheritParentProfile"`
	InheritVlanList        string `json:"inheritVlanList"`
	LogProfile             string `json:"logProfile"`
	LogPublisher           string `json:"logPublisher"`
	Port                   int    `json:"port"`
	Security               string `json:"security"`
	TranslateExtended      string `json:"translateExtended"`
}

const FtpEndpoint = "ftp"

type FtpResource struct {
	b *bigip.BigIP
}
