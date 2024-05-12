package profile

import "github.com/lefeck/go-bigip"

type FasthttpConfigList struct {
	Items    []FasthttpConfig `json:"items,omitempty"`
	Kind     string           `json:"kind,omitempty"`
	SelfLink string           `json:"selflink,omitempty"`
}

type FasthttpConfig struct {
	Kind                        string `json:"kind"`
	Name                        string `json:"name"`
	Partition                   string `json:"partition"`
	FullPath                    string `json:"fullPath"`
	Generation                  int    `json:"generation"`
	SelfLink                    string `json:"selfLink"`
	AppService                  string `json:"appService"`
	ClientCloseTimeout          int    `json:"clientCloseTimeout"`
	ConnpoolIdleTimeoutOverride int    `json:"connpoolIdleTimeoutOverride"`
	ConnpoolMaxReuse            int    `json:"connpoolMaxReuse"`
	ConnpoolMaxSize             int    `json:"connpoolMaxSize"`
	ConnpoolMinSize             int    `json:"connpoolMinSize"`
	ConnpoolReplenish           string `json:"connpoolReplenish"`
	ConnpoolStep                int    `json:"connpoolStep"`
	DefaultsFrom                string `json:"defaultsFrom"`
	Description                 string `json:"description"`
	ForceHTTP10Response         string `json:"forceHttp_10Response"`
	HardwareSynCookie           string `json:"hardwareSynCookie"`
	HeaderInsert                string `json:"headerInsert"`
	HTTP11CloseWorkarounds      string `json:"http_11CloseWorkarounds"`
	IdleTimeout                 int    `json:"idleTimeout"`
	InsertXforwardedFor         string `json:"insertXforwardedFor"`
	Layer7                      string `json:"layer_7"`
	MaxHeaderSize               int    `json:"maxHeaderSize"`
	MaxRequests                 int    `json:"maxRequests"`
	MssOverride                 int    `json:"mssOverride"`
	ReceiveWindowSize           int    `json:"receiveWindowSize"`
	ResetOnTimeout              string `json:"resetOnTimeout"`
	ServerCloseTimeout          int    `json:"serverCloseTimeout"`
	ServerSack                  string `json:"serverSack"`
	ServerTimestamp             string `json:"serverTimestamp"`
	UncleanShutdown             string `json:"uncleanShutdown"`
}

const FasthttpEndpoint = "fasthttp"

type FasthttpResource struct {
	b *bigip.BigIP
}
