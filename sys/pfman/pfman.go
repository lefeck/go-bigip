package pfman

import "github.com/lefeck/go-bigip"

// PFManEndpoint represents the REST resource for managing PFMan.
const PFManEndpoint = "pfman"
const SysManager = "sys"

// PFManResource provides an API to manage PFMan configurations.
type PFManResource struct {
	device   DeviceResource
	consumer ConsumerResource
}

func NewPFMan(b *bigip.BigIP) PFManResource {
	return PFManResource{
		device:   DeviceResource{b: b},
		consumer: ConsumerResource{b: b},
	}
}

func (pfm *PFManResource) Device() *DeviceResource {
	return &pfm.device
}

func (pfm *PFManResource) Consumer() *ConsumerResource {
	return &pfm.consumer
}
