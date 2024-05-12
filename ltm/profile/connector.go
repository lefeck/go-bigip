package profile

import "github.com/lefeck/go-bigip"

type ConnectorConfigList struct {
	Items    []ConnectorConfig `json:"items,omitempty"`
	Kind     string            `json:"kind,omitempty"`
	SelfLink string            `json:"selflink,omitempty"`
}

type ConnectorConfig struct {
	Kind               string `json:"kind"`
	Name               string `json:"name"`
	Partition          string `json:"partition"`
	FullPath           string `json:"fullPath"`
	Generation         int    `json:"generation"`
	SelfLink           string `json:"selfLink"`
	AppService         string `json:"appService"`
	ConnectOnData      string `json:"connectOnData"`
	ConnectionTimeout  int    `json:"connectionTimeout"`
	EntryVirtualServer string `json:"entryVirtualServer"`
	ServiceDownAction  string `json:"serviceDownAction"`
}

const ConnectorEndpoint = "connector"

type ConnectorResource struct {
	b *bigip.BigIP
}
