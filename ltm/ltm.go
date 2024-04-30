// Package ltm provides a REST client for the /tm/ltm F5 BigIP API.
package ltm

import (
	"github.com/lefeck/bigip"
)

// LtmManager is a commonly used basepath, providing a large number of api resource types
const LtmManager = "ltm"

// LTM implements a REST client for the F5 BigIP LTM API.
type LTM struct {
	b *bigip.BigIP

	virtual        VirtualResource
	virtualAddress VirtualAddressResource
}

// New creates a new LTM client.
func New(b *bigip.BigIP) LTM {
	return LTM{
		b:              b,
		virtual:        VirtualResource{b: b},
		virtualAddress: VirtualAddressResource{b: b},
	}
}

// Virtual returns a VirtualResource configured to query tm/ltm/virtual API.
func (ltm LTM) Virtual() *VirtualResource {
	return &ltm.virtual
}

func (ltm LTM) VirtualAddress() *VirtualAddressResource {
	return &ltm.virtualAddress
}
