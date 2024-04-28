package auth

import (
	"github.com/lefeck/bigip"
)

type Auth struct {
	b     *bigip.BigIP
	users UsersResource
}

func NewAuth(b *bigip.BigIP) Auth {
	return Auth{
		b:     b,
		users: UsersResource{b: b},
	}
}

// /mgmt/shared/authz/users
func (auth Auth) Users() *UsersResource {
	return &auth.users
}
