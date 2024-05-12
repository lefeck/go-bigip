package profile

import "github.com/lefeck/go-bigip"

type Http3ConfigList struct {
	Items    []Http3Config `json:"items,omitempty"`
	Kind     string        `json:"kind,omitempty"`
	SelfLink string        `json:"selflink,omitempty"`
}

type Http3Config struct {
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

const Http3Endpoint = "http3"

type Http3Resoucre struct {
	b *bigip.BigIP
}
