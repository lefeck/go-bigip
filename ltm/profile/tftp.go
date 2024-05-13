package profile

import "github.com/lefeck/go-bigip"

type TFTPList struct {
	Items    []TFTP `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type TFTP struct {
	Kind         string `json:"kind"`
	Name         string `json:"name"`
	Partition    string `json:"partition"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	SelfLink     string `json:"selfLink"`
	AppService   string `json:"appService"`
	DefaultsFrom string `json:"defaultsFrom"`
	Description  string `json:"description"`
	IdleTimeout  string `json:"idleTimeout"`
	LogProfile   string `json:"logProfile"`
	LogPublisher string `json:"logPublisher"`
}

const TFTPEndpoint = "tcp"

type TFTPResource struct {
	b *bigip.BigIP
}
