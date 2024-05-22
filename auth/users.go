package auth

import (
	"context"
	"encoding/json"
	"github.com/lefeck/go-bigip"
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

func (ur *UsersResource) List() (*UsersList, error) {
	res, err := ur.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetShareResource()).
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
