package profile

import "github.com/lefeck/go-bigip"

type OcspConfigList struct {
	Items    []OcspConfig `json:"items,omitempty"`
	Kind     string       `json:"kind,omitempty"`
	SelfLink string       `json:"selflink,omitempty"`
}

type OcspConfig struct {
	Kind         string `json:"kind"`
	Name         string `json:"name"`
	Partition    string `json:"partition"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	SelfLink     string `json:"selfLink"`
	AppService   string `json:"appService"`
	DefaultsFrom string `json:"defaultsFrom"`
	MaxAge       int    `json:"maxAge"`
	Nonce        string `json:"nonce"`
}

const OcspEndpoint = "iamp"

type OcspResoucre struct {
	b *bigip.BigIP
}
