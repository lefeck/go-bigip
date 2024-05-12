package profile

import "github.com/lefeck/go-bigip"

type Http2ConfigList struct {
	Items    []Http2Config `json:"items,omitempty"`
	Kind     string        `json:"kind,omitempty"`
	SelfLink string        `json:"selflink,omitempty"`
}

type Http2Config struct {
	Kind                           string   `json:"kind"`
	Name                           string   `json:"name"`
	Partition                      string   `json:"partition"`
	FullPath                       string   `json:"fullPath"`
	Generation                     int      `json:"generation"`
	SelfLink                       string   `json:"selfLink"`
	ActivationModes                []string `json:"activationModes"`
	AppService                     string   `json:"appService"`
	ConcurrentStreamsPerConnection int      `json:"concurrentStreamsPerConnection"`
	ConnectionIdleTimeout          int      `json:"connectionIdleTimeout"`
	DefaultsFrom                   string   `json:"defaultsFrom"`
	Description                    string   `json:"description"`
	EnforceTLSRequirements         string   `json:"enforceTlsRequirements"`
	FrameSize                      int      `json:"frameSize"`
	HeaderTableSize                int      `json:"headerTableSize"`
	IncludeContentLength           string   `json:"includeContentLength"`
	InsertHeader                   string   `json:"insertHeader"`
	InsertHeaderName               string   `json:"insertHeaderName"`
	ReceiveWindow                  int      `json:"receiveWindow"`
	WriteSize                      int      `json:"writeSize"`
}

const Http2Endpoint = "http2"

type Http2Resoucre struct {
	b *bigip.BigIP
}
