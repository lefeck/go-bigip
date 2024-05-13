package profile

import "github.com/lefeck/go-bigip"

type HTTPCompressionList struct {
	Items    []HTTPCompression `json:"items,omitempty"`
	Kind     string            `json:"kind,omitempty"`
	SelfLink string            `json:"selflink,omitempty"`
}

type HTTPCompression struct {
	Kind                  string        `json:"kind"`
	Name                  string        `json:"name"`
	Partition             string        `json:"partition"`
	FullPath              string        `json:"fullPath"`
	Generation            int           `json:"generation"`
	SelfLink              string        `json:"selfLink"`
	AllowHTTP10           string        `json:"allowHttp_10"`
	AppService            string        `json:"appService"`
	BrowserWorkarounds    string        `json:"browserWorkarounds"`
	BufferSize            int           `json:"bufferSize"`
	ContentTypeExclude    []interface{} `json:"contentTypeExclude"`
	ContentTypeInclude    []string      `json:"contentTypeInclude"`
	CPUSaver              string        `json:"cpuSaver"`
	CPUSaverHigh          int           `json:"cpuSaverHigh"`
	CPUSaverLow           int           `json:"cpuSaverLow"`
	DefaultsFrom          string        `json:"defaultsFrom"`
	DefaultsFromReference struct {
		Link string `json:"link"`
	} `json:"defaultsFromReference"`
	Description        string        `json:"description"`
	GzipLevel          int           `json:"gzipLevel"`
	GzipMemoryLevel    int           `json:"gzipMemoryLevel"`
	GzipWindowSize     int           `json:"gzipWindowSize"`
	KeepAcceptEncoding string        `json:"keepAcceptEncoding"`
	MethodPrefer       string        `json:"methodPrefer"`
	MinSize            int           `json:"minSize"`
	Selective          string        `json:"selective"`
	URIExclude         []interface{} `json:"uriExclude"`
	URIInclude         []string      `json:"uriInclude"`
	VaryHeader         string        `json:"varyHeader"`
}

const HTTPCompressionEndpoint = "http-compression"

type HTTPCompressionResoucre struct {
	b *bigip.BigIP
}
