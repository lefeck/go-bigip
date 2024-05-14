package ltm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// A PoolList holds a list of Pool.
type PoolList struct {
	Items    []Pool `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty" pretty:",expanded"`
	SelfLink string `json:"selfLink,omitempty" pretty:",expanded"`
}

// A Pool hold the uration for a pool.
type Pool struct {
	AllowNat              string   `json:"allowNat,omitempty" pretty:",expanded"`
	AllowSnat             string   `json:"allowSnat,omitempty" pretty:",expanded"`
	Description           string   `json:"description,omitempty"`
	FullPath              string   `json:"fullPath,omitempty" pretty:",expanded"`
	Generation            int64    `json:"generation,omitempty" pretty:",expanded"`
	IgnorePersistedWeight string   `json:"ignorePersistedWeight,omitempty" pretty:",expanded"`
	IPTosToClient         string   `json:"ipTosToClient,omitempty" pretty:",expanded"`
	IPTosToServer         string   `json:"ipTosToServer,omitempty" pretty:",expanded"`
	Kind                  string   `json:"kind,omitempty" pretty:",expanded"`
	LinkQosToClient       string   `json:"linkQosToClient,omitempty" pretty:",expanded"`
	LinkQosToServer       string   `json:"linkQosToServer,omitempty" pretty:",expanded"`
	LoadBalancingMode     string   `json:"loadBalancingMode,omitempty"`
	Members               []string `json:"members"`
	MembersReference      struct {
		IsSubcollection bool          `json:"isSubcollection,omitempty"`
		Link            string        `json:"link,omitempty"`
		Members         []PoolMembers `json:"items,omitempty"`
	} `json:"membersReference,omitempty"`
	MinActiveMembers       int64  `json:"minActiveMembers,omitempty"`
	MinUpMembers           int64  `json:"minUpMembers,omitempty"`
	MinUpMembersAction     string `json:"minUpMembersAction,omitempty"`
	MinUpMembersChecking   string `json:"minUpMembersChecking,omitempty"`
	Monitor                string `json:"monitor,omitempty"`
	Name                   string `json:"name,omitempty"`
	QueueDepthLimit        int64  `json:"queueDepthLimit,omitempty" pretty:",expanded"`
	QueueOnConnectionLimit string `json:"queueOnConnectionLimit,omitempty" pretty:",expanded"`
	QueueTimeLimit         int64  `json:"queueTimeLimit,omitempty" pretty:",expanded"`
	ReselectTries          int64  `json:"reselectTries,omitempty"`
	SelfLink               string `json:"selfLink,omitempty" pretty:",expanded"`
	ServiceDownAction      string `json:"serviceDownAction,omitempty"`
	SlowRampTime           int64  `json:"slowRampTime,omitempty" pretty:",expanded"`
	Partition              string `json:"partition,omitempty"`
}

// PoolEndpoint represents the REST resource for managing a pool.
const PoolEndpoint = "pool"

// A PoolResource provides API to manage pool uration.
type PoolResource struct {
	b *bigip.BigIP
}

// lists all the pool instances.
func (pr *PoolResource) List() (*PoolList, error) {
	var pl PoolList
	res, err := pr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(PoolEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &pl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &pl, nil
}

// List all the details of the pool, including: profile, policy, etc.
func (vr *PoolResource) ListDetail() (*PoolList, error) {
	res, err := vr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(PoolEndpoint).SetParams("expandSubcollections", "true").DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	var pl PoolList
	if err := json.Unmarshal(res, &pl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &pl, nil
}

// ListVirtualServerName get all virtual server names
func (vr *PoolResource) ListPoolName() ([]string, error) {
	pl, err := vr.List()
	if err != nil {
		return nil, err
	}
	var items []string
	for _, pool := range pl.Items {
		fullPathName := pool.FullPath
		items = append(items, fullPathName)
	}
	return items, nil
}

// Get a single pool identified by name.
func (pr *PoolResource) Get(fullPathName string) (*Pool, error) {
	var pool Pool
	res, err := pr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(PoolEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &pool); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &pool, nil
}

// Create a new pool instance. If you just create a new pool, for example:
/*
	item := Pool{
		Name:              "pool-demo",
		LoadBalancingMode: "round-robin",
		:           "http",
        .....
	}
*/
// If you create a new pool and also add member， for example:
/*
	item := Pool{
		Name:              "pool-demo",
		LoadBalancingMode: "round-robin",
		Members:           []string{"192.13.23.1:90", "128.3.2.53:90"},
		:           "http",
        ......
	}
*/
func (pr *PoolResource) Create(item Pool) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = pr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(PoolEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// update a pool instance identified by name.
func (pr *PoolResource) Update(fullPathName string, item Pool) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = pr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(PoolEndpoint).ResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single pool instance identified by name.
func (pr *PoolResource) Delete(name string) error {
	_, err := pr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(PoolEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// // GetMembers gets all the members associated to the pool identified by id.
//
//	func (pr *PoolResource) GetMembers(id string) (*PoolMembersList, error) {
//		var poolMembers PoolMembersList
//		if err := pr.c.ReadQuery(BasePath+PoolEndpoint+"/"+id+"/members", &poolMembers); err != nil {
//			return nil, err
//		}
//		return &poolMembers, nil
//	}
//
//	func (pr *PoolResource) AddMember(id string, poolMember PoolMembers) error {
//		if err := pr.c.ModQuery("POST", BasePath+PoolEndpoint+"/"+id+"/members", poolMember); err != nil {
//			return err
//		}
//		return nil
//	}

//func (pr *PoolResource) ShowStats(id string) (*PoolStatsList, error) {
//	var item PoolStatsList
//	if err := pr.c.ReadQuery(BasePath+PoolEndpoint+"/"+id+"/stats", &item); err != nil {
//		return nil, err
//	}
//	return &item, nil
//}
