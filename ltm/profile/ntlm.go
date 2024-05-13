package profile

import "github.com/lefeck/go-bigip"

type NtlmList struct {
	Items    []Ntlm `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type Ntlm struct {
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

const NtlmEndpoint = "ntlm"

type NtlmResource struct {
	b *bigip.BigIP
}
