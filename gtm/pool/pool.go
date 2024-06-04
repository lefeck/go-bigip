package pool

import "github.com/lefeck/go-bigip"

const PoolEndpoint = "pool"
const GTMManager = "gtm"

type PoolResource struct {
	poolA    PoolAResource
	poolAAAA PoolAAAAResource
	cname    PoolCNAMEResource
	mx       PoolMXResource
	naptr    PoolNAPTRResource
	srv      SRVResource
}

func NewPoolResource(b *bigip.BigIP) PoolResource {
	return PoolResource{
		poolA:    PoolAResource{b: b},
		poolAAAA: PoolAAAAResource{b: b},
		cname:    PoolCNAMEResource{b: b},
		mx:       PoolMXResource{b: b},
		naptr:    PoolNAPTRResource{b: b},
		srv:      SRVResource{b: b},
	}
}

func (p *PoolResource) A() *PoolAResource {
	return &p.poolA
}

func (p *PoolResource) AAAA() *PoolAAAAResource {
	return &p.poolAAAA
}

func (p *PoolResource) CNAME() *PoolCNAMEResource {
	return &p.cname
}

func (p *PoolResource) MX() *PoolMXResource {
	return &p.mx
}

func (p *PoolResource) NAPTR() *PoolNAPTRResource {
	return &p.naptr
}
func (p *PoolResource) SRV() *SRVResource {
	return &p.srv
}
