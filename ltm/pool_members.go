package ltm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// A PoolMembersList represents a list of pool members.
type PoolMembersList struct {
	Kind     string        `json:"kind,omitempty"`
	SelfLink string        `json:"selfLink,omitempty"`
	Items    []PoolMembers `json:"items,omitempty"`
}

// A PoolMembers represents the members of a pool.
type PoolMembers struct {
	Kind            string `json:"kind,omitempty"`
	Name            string `json:"name,omitempty"`
	Partition       string `json:"partition,omitempty"`
	FullPath        string `json:"fullPath,omitempty"`
	Generation      int64  `json:"generation,omitempty"`
	SelfLink        string `json:"selfLink,omitempty"`
	Address         string `json:"address,omitempty"`
	ConnectionLimit int64  `json:"connectionLimit,omitempty"`
	DynamicRatio    int64  `json:"dynamicRatio,omitempty"`
	Ephemeral       string `json:"ephemeral,omitempty"`
	Fqdn            struct {
		Autopopulate string `json:"autopopulate,omitempty"`
	} `json:"fqdn,omitempty"`
	InheritProfile string `json:"inheritProfile,omitempty"`
	Logging        string `json:"logging,omitempty"`
	Monitor        string `json:"monitor,omitempty"`
	PriorityGroup  int64  `json:"priorityGroup,omitempty"`
	RateLimit      string `json:"rateLimit,omitempty"`
	Ratio          int64  `json:"ratio,omitempty"`
	Session        string `json:"session,omitempty"`
	State          string `json:"state,omitempty"`
}

// PoolMembersEndpoint represents the REST resource for managing pool members.
const poolMembersEndpoint = "members"

// PoolMembersResource provides an API to manage pool members object.
type PoolMembersResource struct {
	b *bigip.BigIP
}

// lists all the pool members.
func (pmr *PoolMembersResource) List(pool string) (*PoolMembersList, error) {
	var pml PoolMembersList
	res, err := pmr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(PoolEndpoint).ResourceInstance(pool).SubResource(poolMembersEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &pml); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &pml, nil
}

// Get a single pool members identified by pool name and member name.
func (pmr *PoolMembersResource) Get(poolName string, memberName string) (*PoolMembers, error) {
	var pm PoolMembers
	res, err := pmr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(PoolEndpoint).ResourceInstance(poolName).SubResource(poolMembersEndpoint).SubResourceInstance(memberName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &pm); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &pm, nil
}

// Create a new pool members.
func (pmr *PoolMembersResource) Create(pool string, item PoolMembers) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = pmr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(PoolEndpoint).ResourceInstance(pool).SubResource(poolMembersEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update a pool members indentified by pool name and member name.
func (pmr *PoolMembersResource) Update(poolName string, memberName string, item PoolMembers) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = pmr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(PoolEndpoint).ResourceInstance(poolName).SubResource(poolMembersEndpoint).SubResourceInstance(memberName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single pool members identified by pool name and member name.
func (pmr *PoolMembersResource) Delete(poolName string, memberName string) error {
	_, err := pmr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(PoolEndpoint).ResourceInstance(poolName).SubResource(poolMembersEndpoint).SubResourceInstance(memberName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
