package profile

import "github.com/lefeck/go-bigip"

type FTPList struct {
	Items    []FTP  `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type FTP struct {
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

const FTPEndpoint = "ftp"

type FTPResource struct {
	b *bigip.BigIP
}
