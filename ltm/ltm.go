// Package ltm provides a REST client for the /tm/ltm F5 BigIP API.
package ltm

import (
	"github.com/lefeck/go-bigip"
)

// LtmManager is a commonly used basepath, providing a large number of api resource types
const LtmManager = "ltm"

// LTM implements a REST client for the F5 BigIP LTM API.
type LTM struct {
	b *bigip.BigIP

	virtual             VirtualResource
	virtualAddress      VirtualAddressResource
	virtualAddressStats VirtualAddressStatsResource
	virtualStats        VirtualStatsResource
	pool                PoolResource
	rule                RuleResource
	poolMembers         PoolMembersResource
	poolStats           PoolStatsResource
	snatPool            SnatPoolResource
	nodeStats           NodeStatsResource
}

// New creates a new LTM client.
func New(b *bigip.BigIP) LTM {
	return LTM{
		b:                   b,
		virtual:             VirtualResource{b: b},
		virtualAddress:      VirtualAddressResource{b: b},
		virtualAddressStats: VirtualAddressStatsResource{b: b},
		virtualStats:        VirtualStatsResource{b: b},
		pool:                PoolResource{b: b},
		snatPool:            SnatPoolResource{b: b},
		poolStats:           PoolStatsResource{b: b},
		poolMembers:         PoolMembersResource{b: b},
		rule:                RuleResource{b: b},
		nodeStats:           NodeStatsResource{b: b},
	}
}

// Virtual returns a VirtualResource configured to query tm/ltm/virtual API.
func (ltm LTM) Virtual() *VirtualResource {
	return &ltm.virtual
}

func (ltm LTM) VirtualAddress() *VirtualAddressResource {
	return &ltm.virtualAddress
}

func (ltm LTM) VirtualAddressStats() *VirtualAddressStatsResource {
	return &ltm.virtualAddressStats
}

func (ltm LTM) VirtualStats() *VirtualStatsResource {
	return &ltm.virtualStats
}

func (ltm LTM) Pool() *PoolResource {
	return &ltm.pool
}

func (ltm LTM) SnatPool() *SnatPoolResource {
	return &ltm.snatPool
}

func (ltm LTM) Rule() *RuleResource {
	return &ltm.rule
}

func (ltm LTM) PoolMembers() *PoolMembersResource {
	return &ltm.poolMembers
}

func (ltm LTM) PoolStats() *PoolStatsResource {
	return &ltm.poolStats
}

func (ltm LTM) NodeStats() *NodeStatsResource {
	return &ltm.nodeStats
}
