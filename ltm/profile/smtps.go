package profile

import "github.com/lefeck/go-bigip"

type SMTPSList struct {
	Items    []SMTPS `json:"items,omitempty"`
	Kind     string  `json:"kind,omitempty"`
	SelfLink string  `json:"selflink,omitempty"`
}

type SMTPS struct {
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

const SMTPSEndpoint = "smtps"

type SMTPSResource struct {
	b *bigip.BigIP
}

func (p *SMTPS) List() (*SMTPSList, error) {
	var items SMTPSList

	return &items, nil
}
