package profile

import "github.com/lefeck/go-bigip"

// MonitorEndpoint is a commonly used basepath, providing a large number of api resource types
const MonitorEndpoint = "monitor"

type ProfileResoucre struct {
	b *bigip.BigIP
}

func NewProfile(b *bigip.BigIP) ProfileResoucre {
	return ProfileResoucre{b: b}
}
