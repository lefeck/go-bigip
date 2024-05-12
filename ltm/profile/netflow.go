package profile

import "github.com/lefeck/go-bigip"

type NetflowConfigList struct {
	Items    []NetflowConfig `json:"items,omitempty"`
	Kind     string          `json:"kind,omitempty"`
	SelfLink string          `json:"selflink,omitempty"`
}

type NetflowConfig struct {
	Kind           string `json:"kind"`
	Name           string `json:"name"`
	Partition      string `json:"partition"`
	FullPath       string `json:"fullPath"`
	Generation     int    `json:"generation"`
	SelfLink       string `json:"selfLink"`
	AppService     string `json:"appService"`
	DefaultsFrom   string `json:"defaultsFrom"`
	Description    string `json:"description"`
	NetflowVersion string `json:"netflowVersion"`
	SamplingRate   int    `json:"samplingRate"`
}

const NetflowEndpoint = "iamp"

type NetflowResoucre struct {
	b *bigip.BigIP
}
