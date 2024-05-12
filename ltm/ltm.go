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

	//monitor   monitor.MonitorResoucre
	//clientSSL profile.ProfileClientSSLSoucre
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
		////monitor: monitor.MonitorResoucre{B: b},
		//monitor: monitor.NewMonitor(b ),
		//
		//// profile
		//clientSSL: profile.ProfileClientSSLSoucre{B: b},
	}
}

type LTMBuilder interface {
	Virtual() *VirtualResource
	VirtualAddress() *VirtualAddressResource
	VirtualAddressStats() *VirtualAddressStatsResource
	VirtualStats() *VirtualStatsResource
	Pool() *PoolResource
	SnatPool() *SnatPoolResource
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

//func (ltm LTM) ProfileCLientSSL() *profile.ProfileClientSSLSoucre {
//	return &ltm.clientSSL
//}

// IFile returns an IFileResource configured to query /tm/ltm/ifile API.
func (ltm LTM) IFile() *IFileResource {
	return &ltm.iFile
}

// DataGroupInternal returns a DataGroupInternalResource configured to query /tm/ltm/data-group/internal API.
func (ltm LTM) DataGroupInternal() *DataGroupInternalResource {
	return &ltm.dataGroupInternal
}

//func (ltm LTM) Monitor() monitor.MonitorResoucre {
//	return ltm.monitor
//}
