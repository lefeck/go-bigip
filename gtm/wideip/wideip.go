package wideip

import "github.com/lefeck/go-bigip"

const WideipEndpoint = "wideip"
const GTMManager = "gtm"

type WideipResource struct {
	a     AResource
	aaaa  AAAAResource
	cname CNAMEResource
	mx    MXResource
	naptr NAPTRResource
	srv   SRVResource
}

func NewWideip(b *bigip.BigIP) WideipResource {
	return WideipResource{
		a:     AResource{b: b},
		aaaa:  AAAAResource{b: b},
		cname: CNAMEResource{b: b},
		mx:    MXResource{b: b},
		naptr: NAPTRResource{b: b},
		srv:   SRVResource{b: b},
	}
}

func (w WideipResource) A() *AResource {
	return &w.a
}

func (w WideipResource) AAAA() *AAAAResource {
	return &w.aaaa
}

func (w WideipResource) CNAME() *CNAMEResource {
	return &w.cname
}

func (w WideipResource) MX() *MXResource {
	return &w.mx
}

func (w WideipResource) NAPTR() *NAPTRResource {
	return &w.naptr
}

func (w WideipResource) SRV() *SRVResource {
	return &w.srv
}
