package profile

import "github.com/lefeck/go-bigip"

type HTTP3ConfigList struct {
	Items    []HTTP3 `json:"items,omitempty"`
	Kind     string  `json:"kind,omitempty"`
	SelfLink string  `json:"selflink,omitempty"`
}

type HTTP3 struct {
	Kind            string `json:"kind"`
	Name            string `json:"name"`
	Partition       string `json:"partition"`
	FullPath        string `json:"fullPath"`
	Generation      int    `json:"generation"`
	SelfLink        string `json:"selfLink"`
	AppService      string `json:"appService"`
	DefaultsFrom    string `json:"defaultsFrom"`
	Description     string `json:"description"`
	HeaderTableSize int    `json:"headerTableSize"`
}

const HTTP3Endpoint = "http3"

type HTTP3Resoucre struct {
	b *bigip.BigIP
}
