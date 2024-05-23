package ipfix

import "github.com/lefeck/go-bigip"

// IPFixEndpoint represents the REST resource for managing IPFix.
const IPFixEndpoint = "ipfix"
const SysManager = "sys"

type IPFixResource struct {
	iPFixElement IPFixElementResource
}

func NewIPFix(b *bigip.BigIP) IPFixResource {
	return IPFixResource{
		iPFixElement: IPFixElementResource{b: b},
	}
}

func (ipfix IPFixResource) Element() *IPFixElementResource {
	return &ipfix.iPFixElement
}
