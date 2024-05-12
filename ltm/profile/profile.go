package profile

import "github.com/lefeck/go-bigip"

// ProfileEndpoint is a commonly used basepath, providing a large number of api resource types
const ProfileEndpoint = "profile"

type ProfileResoucre struct {
	b *bigip.BigIP
}

func NewProfile(b *bigip.BigIP) ProfileResoucre {
	return ProfileResoucre{b: b}
}
