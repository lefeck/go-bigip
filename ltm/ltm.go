// Package ltm provides a REST client for the /tm/ltm F5 BigIP API.
package ltm

import (
	"github.com/lefeck/go-bigip"
	"github.com/lefeck/go-bigip/ltm/monitor"
	"github.com/lefeck/go-bigip/ltm/profile"
)

// LtmManager is a commonly used bigip.GetBaseResource(), providing a large number of api resource types
const LtmManager = "ltm"

// LTM implements a REST client for the F5 BigIP LTM API.
type LTM struct {
	virtual             VirtualResource
	virtualStats        VirtualStatsResource
	virtualAddress      VirtualAddressResource
	virtualAddressStats VirtualAddressStatsResource
	pool                PoolResource
	rule                RuleResource
	poolMembers         PoolMembersResource
	poolStats           PoolStatsResource
	snatPool            SnatPoolResource
	node                NodeResource
	nodeStats           NodeStatsResource

	iFile             IFileResource
	dataGroupInternal DataGroupInternalResource

	// Provide a public entry point for  monitor resources
	monitor monitor.MonitorResource

	// Provide a public entry point for profile resources
	profile profile.ProfileResource
}

// New creates a new LTM client.
func New(b *bigip.BigIP) LTM {
	return LTM{
		virtual:             VirtualResource{b: b},
		virtualAddress:      VirtualAddressResource{b: b},
		virtualAddressStats: VirtualAddressStatsResource{b: b},
		virtualStats:        VirtualStatsResource{b: b},
		pool:                PoolResource{b: b},
		snatPool:            SnatPoolResource{b: b},
		poolStats:           PoolStatsResource{b: b},
		poolMembers:         PoolMembersResource{b: b},
		rule:                RuleResource{b: b},
		node:                NodeResource{b: b},
		nodeStats:           NodeStatsResource{b: b},

		iFile:             IFileResource{b: b},
		dataGroupInternal: DataGroupInternalResource{b: b},

		// monitor
		monitor: monitor.NewMonitor(b),
		// profile
		profile: profile.NewProfile(b),
	}
}

// Virtual returns a VirtualResource ured to query tm/ltm/virtual API.
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

// IFile returns an IFileResource ured to query /tm/ltm/ifile API.
func (ltm LTM) IFile() *IFileResource {
	return &ltm.iFile
}

// DataGroupInternal returns a DataGroupInternalResource ured to query /tm/ltm/data-group/internal API.
func (ltm LTM) DataGroupInternal() *DataGroupInternalResource {
	return &ltm.dataGroupInternal
}

func (ltm LTM) Monitor() *monitor.MonitorResource {
	return &ltm.monitor
}

func (ltm LTM) Profile() *profile.ProfileResource {
	return &ltm.profile
}
