package profile

import "github.com/lefeck/go-bigip"

type ImapConfigList struct {
	Items    []ImapConfig `json:"items,omitempty"`
	Kind     string       `json:"kind,omitempty"`
	SelfLink string       `json:"selflink,omitempty"`
}

type ImapConfig struct {
	Kind           string `json:"kind"`
	Name           string `json:"name"`
	Partition      string `json:"partition"`
	FullPath       string `json:"fullPath"`
	Generation     int    `json:"generation"`
	SelfLink       string `json:"selfLink"`
	ActivationMode string `json:"activationMode"`
	AppService     string `json:"appService"`
	DefaultsFrom   string `json:"defaultsFrom"`
	Description    string `json:"description"`
}

const ImapEndpoint = "iamp"

type ImapResoucre struct {
	b *bigip.BigIP
}
