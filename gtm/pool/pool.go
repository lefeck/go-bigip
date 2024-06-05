package pool

import "github.com/lefeck/go-bigip"

// PoolEndpoint is the REST resource for managing pool in BigIP
const PoolEndpoint = "pool"
const GTMManager = "gtm"

// PoolResource struct is a container for different pool types
type PoolResource struct {
	a     AResource
	aaaa  AAAAResource
	cname CNAMEResource
	mx    MXResource
	naptr NAPTRResource
	srv   SRVResource
}

// NewPoolResource constructs a new instance of PoolResource with a given bigip.BigIP instance
func NewPoolResource(b *bigip.BigIP) PoolResource {
	return PoolResource{
		a:     AResource{b: b},
		aaaa:  AAAAResource{b: b},
		cname: CNAMEResource{b: b},
		mx:    MXResource{b: b},
		naptr: NAPTRResource{b: b},
		srv:   SRVResource{b: b},
	}
}

// A returns a reference to the AResource instance
func (p *PoolResource) A() *AResource {
	return &p.a
}

// AAAA returns a reference to the AAAAResource instance
func (p *PoolResource) AAAA() *AAAAResource {
	return &p.aaaa
}

// CNAME returns a reference to the CNAMEResource instance
func (p *PoolResource) CNAME() *CNAMEResource {
	return &p.cname
}

// MX returns a reference to the MXResource instance
func (p *PoolResource) MX() *MXResource {
	return &p.mx
}

// NAPTR returns a reference to the NAPTRResource instance
func (p *PoolResource) NAPTR() *NAPTRResource {
	return &p.naptr
}

// SRV returns a reference to the SRVResource instance
func (p *PoolResource) SRV() *SRVResource {
	return &p.srv
}
