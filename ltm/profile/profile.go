package profile

import "github.com/lefeck/go-bigip"

// ProfileEndpoint is a commonly used basepath, providing a large number of api resource types
const ProfileEndpoint = "profile"

type ProfileResource struct {
	fasthttp FasthttpResource
}

func NewProfile(b *bigip.BigIP) ProfileResource {
	return ProfileResource{
		fasthttp: FasthttpResource{b: b},
	}
}

func (p ProfileResource) FASTHTTP() *FasthttpResource {
	return &p.fasthttp
}
