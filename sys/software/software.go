package software

import (
	"github.com/lefeck/go-bigip"
)

// SoftwareEndpoint represents the REST resource for managing Software.
const SoftwareEndpoint = "/software"

// SysEndpoint represents the REST resource for managing Disk.
const SysManager = "sys"

// DiskResource provides an API to manage Disk configurations.
type SoftwareResource struct {
	image        ImageResource
	update       UpdateResource
	updateStatus UpdateStatusResource

	volume            VolumeResource
	blockDeviceHotfix BlockDeviceHotfixResource
	blockDeviceImage  BlockDeviceImageResource
	hotfix            HotfixResource
}

func NewSoftware(b *bigip.BigIP) SoftwareResource {
	return SoftwareResource{
		image:             ImageResource{b: b},
		update:            UpdateResource{b: b},
		updateStatus:      UpdateStatusResource{b: b},
		volume:            VolumeResource{b: b},
		blockDeviceHotfix: BlockDeviceHotfixResource{b: b},
		blockDeviceImage:  BlockDeviceImageResource{b: b},
		hotfix:            HotfixResource{b: b},
	}
}

type Software interface {
	Image() *ImageResource
	Update() *UpdateResource
	UpdateStatus() *UpdateStatusResource
	Volume() *VolumeResource
	BlockDeviceHotfix() *BlockDeviceHotfixResource
	BlockDeviceImage() *BlockDeviceImageResource
	Hotfix() *HotfixResource
}

var _ Software = SoftwareResource{}

func (m SoftwareResource) Image() *ImageResource {
	return &m.image
}

func (m SoftwareResource) Update() *UpdateResource {
	return &m.update
}

func (m SoftwareResource) UpdateStatus() *UpdateStatusResource {
	return &m.updateStatus
}

func (m SoftwareResource) Volume() *VolumeResource {
	return &m.volume
}

func (m SoftwareResource) BlockDeviceHotfix() *BlockDeviceHotfixResource {
	return &m.blockDeviceHotfix
}

func (m SoftwareResource) BlockDeviceImage() *BlockDeviceImageResource {
	return &m.blockDeviceImage
}

func (m SoftwareResource) Hotfix() *HotfixResource {
	return &m.hotfix
}
