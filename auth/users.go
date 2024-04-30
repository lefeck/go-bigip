package auth

import (
	"context"
	"encoding/json"
	"github.com/lefeck/bigip"
	"github.com/lefeck/bigip/ltm"
)

type UsersList struct {
	Item             []User `json:"items,omitempty"`
	Generation       int    `json:"generation"`
	LastUpdateMicros int    `json:"lastUpdateMicros"`
	Kind             string `json:"kind"`
	SelfLink         string `json:"selfLink"`
}

type User struct {
	Name              string `json:"name"`
	DisplayName       string `json:"displayName"`
	EncryptedPassword string `json:"encryptedPassword"`
	Shell             string `json:"shell"`
	Generation        int    `json:"generation"`
	LastUpdateMicros  int    `json:"lastUpdateMicros"`
	Kind              string `json:"kind"`
	SelfLink          string `json:"selfLink"`
}

type UsersResource struct {
	b *bigip.BigIP
}

// UserEndpoint is the base path of the authz API.
const UserEndpoint = "users"

// /mgmt/shared/authz/users
func (ur *UsersResource) List() (*UsersList, error) {
	res, err := ur.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.SHAREResoucre).
		ManagerName(AuthzManager).Resource(UserEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	var ul UsersList
	if err := json.Unmarshal(res, &ul); err != nil {
		panic(err)
	}
	return &ul, nil
}
