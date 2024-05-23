package raid

import "github.com/lefeck/go-bigip"

const SysManager = "sys"

// RAIDEndpoint represents the REST resource for managing RAID.
const RAIDEndpoint = "raid"

type RAIDResource struct {
	b    *bigip.BigIP
	disk DiskResource
}

func NewRAID(b *bigip.BigIP) RAIDResource {
	return RAIDResource{
		disk: DiskResource{b: b},
	}
}

func (raid *RAIDResource) Disk() *DiskResource {
	return &raid.disk
}
