package profile

import "github.com/lefeck/go-bigip"

type SmtpsConfigList struct {
	Items    []SmtpsConfig `json:"items,omitempty"`
	Kind     string        `json:"kind,omitempty"`
	SelfLink string        `json:"selflink,omitempty"`
}

type SmtpsConfig struct {
	Kind           string `json:"kind"`
	Name           string `json:"name"`
	Partition      string `json:"partition"`
	FullPath       string `json:"fullPath"`
	Generation     int    `json:"generation"`
	SelfLink       string `json:"selfLink"`
	ActivationMode string `json:"activationMode"`
	AppService     string `json:"appService"`
	DefaultsFrom   string `json:"defaultsFrom"`
	Description    string `json:"description"`
}

type ProfileSmtpsResource struct {
	b *bigip.BigIP
}

func (p *ProfileSmtpsResource) List() (*SmtpsConfigList, error) {

}
