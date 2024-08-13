package auth

import (
	"github.com/lefeck/go-bigip"
)

// AuthzManager is a commonly used bigip.GetBaseResource(), providing a large number of api resource types
const AuthManager = "auth"

type Auth struct {
	user      UsersResource
	partition PartitionResource
}

func NewAuth(b *bigip.BigIP) Auth {
	return Auth{
		user:      UsersResource{b: b},
		partition: PartitionResource{b: b},
	}
}

func (auth Auth) User() *UsersResource {
	return &auth.user
}

func (auth Auth) Partition() *PartitionResource {
	return &auth.partition
}
