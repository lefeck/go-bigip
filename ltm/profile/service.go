package profile

import "github.com/lefeck/go-bigip"

type ServiceList struct {
	Items    []Service `json:"items,omitempty"`
	Kind     string    `json:"kind,omitempty"`
	SelfLink string    `json:"selflink,omitempty"`
}
type Service struct {
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Partition  string `json:"partition"`
	FullPath   string `json:"fullPath"`
	Generation int    `json:"generation"`
	SelfLink   string `json:"selfLink"`
	AppService string `json:"appService"`
	Type       string `json:"type"`
}

const ServiceEndpoint = "service"

type ServiceResource struct {
	b *bigip.BigIP
}
