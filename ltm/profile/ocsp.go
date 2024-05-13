package profile

import "github.com/lefeck/go-bigip"

type OcspList struct {
	Items    []Ocsp `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type Ocsp struct {
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

const OcspEndpoint = "ocsp"

type OcspResoucre struct {
	b *bigip.BigIP
}
