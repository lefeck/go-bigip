package profile

import "github.com/lefeck/go-bigip"

type NtlmConfigList struct {
	Items    []NtlmConfig `json:"items,omitempty"`
	Kind     string       `json:"kind,omitempty"`
	SelfLink string       `json:"selflink,omitempty"`
}

type NtlmConfig struct {
	Kind                   string `json:"kind"`
	Name                   string `json:"name"`
	Partition              string `json:"partition"`
	FullPath               string `json:"fullPath"`
	Generation             int    `json:"generation"`
	SelfLink               string `json:"selfLink"`
	AppService             string `json:"appService"`
	DefaultsFrom           string `json:"defaultsFrom"`
	Description            string `json:"description"`
	InsertCookieDomain     string `json:"insertCookieDomain"`
	InsertCookieName       string `json:"insertCookieName"`
	InsertCookiePassphrase string `json:"insertCookiePassphrase"`
	KeyByCookie            string `json:"keyByCookie"`
	KeyByCookieName        string `json:"keyByCookieName"`
	KeyByDomain            string `json:"keyByDomain"`
	KeyByIPAddress         string `json:"keyByIpAddress"`
	KeyByTarget            string `json:"keyByTarget"`
	KeyByUser              string `json:"keyByUser"`
	KeyByWorkstation       string `json:"keyByWorkstation"`
}

const NtlmEndpoint = "iamp"

type NtlmResoucre struct {
	b *bigip.BigIP
}
