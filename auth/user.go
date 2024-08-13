package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type UsersList struct {
	Item     []User `json:"items,omitempty"`
	Kind     string `json:"kind"`
	SelfLink string `json:"selfLink"`
}

// NameReference represents the reference link in partitionAccess
type NameReference struct {
	Link string `json:"link,omitempty"`
}

// PartitionAccess represents the access details for a partition
type PartitionAccess struct {
	Name          string        `json:"name,omitempty"`
	Role          string        `json:"role,omitempty"`
	NameReference NameReference `json:"nameReference,omitempty"`
}

// User represents the user information
type User struct {
	Kind            string            `json:"kind,omitempty"`
	Name            string            `json:"name,omitempty"`
	FullPath        string            `json:"fullPath,omitempty"`
	Generation      int               `json:"generation,omitempty"`
	SelfLink        string            `json:"selfLink,omitempty"`
	Description     string            `json:"description,omitempty"`
	Password        string            `json:"password,omitempty"`
	SessionLimit    int               `json:"sessionLimit,omitempty"`
	Shell           string            `json:"shell,omitempty"`
	PartitionAccess []PartitionAccess `json:"partitionAccess,omitempty"`
}

type UsersResource struct {
	b *bigip.BigIP
}

// UserEndpoint is the base path of the authz API.
const UserEndpoint = "user"

func (ur *UsersResource) List() (*UsersList, error) {
	res, err := ur.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).
		ManagerName(AuthManager).Resource(UserEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	var ul UsersList
	if err := json.Unmarshal(res, &ul); err != nil {
		panic(err)
	}
	return &ul, nil
}

// Get a single User configuration identified by name.
func (r *UsersResource) Get(name string) (*User, error) {
	var item User
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(AuthManager).
		Resource(UserEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create a new User configuration.
func (r *UsersResource) Create(item User) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(AuthManager).
		Resource(UserEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Edit a User configuration identified by name.
func (r *UsersResource) Update(name string, item User) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(AuthManager).
		Resource(UserEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single User configuration identified by name.
func (r *UsersResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(AuthManager).
		Resource(UserEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
