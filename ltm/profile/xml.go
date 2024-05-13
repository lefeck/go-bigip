package profile

import "github.com/lefeck/go-bigip"

type XMLList struct {
	Items    []XML  `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type XML struct {
	Kind                 string        `json:"kind"`
	Name                 string        `json:"name"`
	Partition            string        `json:"partition"`
	FullPath             string        `json:"fullPath"`
	Generation           int           `json:"generation"`
	SelfLink             string        `json:"selfLink"`
	AppService           string        `json:"appService"`
	DefaultsFrom         string        `json:"defaultsFrom"`
	Description          string        `json:"description"`
	MultipleQueryMatches string        `json:"multipleQueryMatches"`
	NamespaceMappings    []interface{} `json:"namespaceMappings"`
	XpathQueries         []interface{} `json:"xpathQueries"`
}

const XMLEndpoint = "xml"

type XMLResource struct {
	b *bigip.BigIP
}
