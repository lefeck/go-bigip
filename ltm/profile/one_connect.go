package profile

import "github.com/lefeck/go-bigip"

type OneConnectConfigList struct {
	Items    []OneConnectConfig `json:"items,omitempty"`
	Kind     string             `json:"kind,omitempty"`
	SelfLink string             `json:"selflink,omitempty"`
}

type OneConnectConfig struct {
	Kind                string `json:"kind"`
	Name                string `json:"name"`
	Partition           string `json:"partition"`
	FullPath            string `json:"fullPath"`
	Generation          int    `json:"generation"`
	SelfLink            string `json:"selfLink"`
	AppService          string `json:"appService"`
	DefaultsFrom        string `json:"defaultsFrom"`
	Description         string `json:"description"`
	IdleTimeoutOverride string `json:"idleTimeoutOverride"`
	LimitType           string `json:"limitType"`
	MaxAge              int    `json:"maxAge"`
	MaxReuse            int    `json:"maxReuse"`
	MaxSize             int    `json:"maxSize"`
	SharePools          string `json:"sharePools"`
	SourceMask          string `json:"sourceMask"`
}

const OneConnectEndpoint = "oneconnect"

type OneConnectResoucre struct {
	b *bigip.BigIP
}
