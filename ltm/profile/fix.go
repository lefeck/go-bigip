package profile

import "github.com/lefeck/go-bigip"

type FixConfigList struct {
	Items    []Fastl4Config `json:"items,omitempty"`
	Kind     string         `json:"kind,omitempty"`
	SelfLink string         `json:"selflink,omitempty"`
}

type FixConfig struct {
	Kind                     string        `json:"kind"`
	Name                     string        `json:"name"`
	Partition                string        `json:"partition"`
	FullPath                 string        `json:"fullPath"`
	Generation               int           `json:"generation"`
	SelfLink                 string        `json:"selfLink"`
	AppService               string        `json:"appService"`
	DefaultsFrom             string        `json:"defaultsFrom"`
	Description              string        `json:"description"`
	ErrorAction              string        `json:"errorAction"`
	FullLogonParsing         string        `json:"fullLogonParsing"`
	MessageLogPublisher      string        `json:"messageLogPublisher"`
	QuickParsing             string        `json:"quickParsing"`
	ReportLogPublisher       string        `json:"reportLogPublisher"`
	ResponseParsing          string        `json:"responseParsing"`
	SenderTagClass           []interface{} `json:"senderTagClass"`
	StatisticsSampleInterval int           `json:"statisticsSampleInterval"`
}

const FixEndpoint = "fix"

type FixResource struct {
	b *bigip.BigIP
}
