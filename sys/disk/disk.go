package disk

import "github.com/lefeck/go-bigip"

// DiskEndpoint represents the REST resource for managing Disk.
const DiskEndpoint = "disk"
const SysManager = "sys"

// DiskResource provides an API to manage Disk configurations.
type DiskResource struct {
	b *bigip.BigIP

	logicalDisk       LogicalDiskResource
	applicationVolume ApplicationVolumeResource
}

func NewDisk(b *bigip.BigIP) DiskResource {
	return DiskResource{
		logicalDisk:       LogicalDiskResource{b: b},
		applicationVolume: ApplicationVolumeResource{b: b},
	}
}

type Disker interface {
	LogicalDisk() *LogicalDiskResource
	ApplicationVolume() *ApplicationVolumeResource
}

var _ Disker = DiskResource{}

func (m DiskResource) LogicalDisk() *LogicalDiskResource {
	return &m.logicalDisk
}

func (m DiskResource) ApplicationVolume() *ApplicationVolumeResource {
	return &m.applicationVolume
}
