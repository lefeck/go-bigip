package profile

type CertificateAuthorityConfigList struct {
	Items    []CertificateAuthorityConfig `json:"items,omitempty"`
	Kind     string                       `json:"kind,omitempty"`
	SelfLink string                       `json:"selflink,omitempty"`
}

type CertificateAuthorityConfig struct {
}
