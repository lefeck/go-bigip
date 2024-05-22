package auth

import (
	"github.com/lefeck/go-bigip"
)

type Authz struct {
	b     *bigip.BigIP
	users UsersResource
}

func NewAuth(b *bigip.BigIP) Authz {
	return Authz{
		users: UsersResource{b: b},
	}
}

func (auth Authz) Users() *UsersResource {
	return &auth.users
}

// AuthzManager is a commonly used bigip.GetBaseResource(), providing a large number of api resource types
const AuthzManager = "authz"
