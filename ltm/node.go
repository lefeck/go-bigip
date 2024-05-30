package ltm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// NodeList is a list containing multiple Node objects.
type NodeList struct {
	Kind     string `json:"kind"`
	SelfLink string `json:"selfLink"`
	Items    []Node `json:"items"`
}

// Node represents an F5 BIG-IP LTM Node configuration.
type Node struct {
	Kind            string `json:"kind"`
	Name            string `json:"name"`
	Partition       string `json:"partition"`
	FullPath        string `json:"fullPath"`
	Generation      int    `json:"generation"`
	SelfLink        string `json:"selfLink"`
	Address         string `json:"address"`
	ConnectionLimit int    `json:"connectionLimit"`
	DynamicRatio    int    `json:"dynamicRatio"`
	Ephemeral       string `json:"ephemeral"`
	Fqdn            struct {
		AddressFamily string `json:"addressFamily"`
		Autopopulate  string `json:"autopopulate"`
		DownInterval  int    `json:"downInterval"`
		Interval      string `json:"interval"`
	} `json:"fqdn"`
	Logging   string `json:"logging"`
	string    `json:""`
	RateLimit string `json:"rateLimit"`
	Ratio     int    `json:"ratio"`
	Session   string `json:"session"`
	State     string `json:"state"`
}

// NodeEndpoint represents the REST resource for managing Node.
const NodeEndpoint = "node"

// NodeResource provides an API to manage node object.
type NodeResource struct {
	b *bigip.BigIP
}

// List all node details
func (nr *NodeResource) List() (*NodeList, error) {
	var nl NodeList
	res, err := nr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(NodeEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &nl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &nl, nil
}

// Get a single node details by the node name
func (nr *NodeResource) Get(name string) (*Node, error) {
	var node Node
	res, err := nr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(NodeEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &node); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &node, nil
}

// Create a new node item
func (nr *NodeResource) Create(item Node) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = nr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(NodeEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update the node item identified by the node name, otherwise an error will be reported.
func (nr *NodeResource) Update(name string, item Node) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = nr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(NodeEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Enable a node identified by the node name.
func (nr *NodeResource) Enable(name string) error {
	item := Node{Session: "user-enabled", State: "user-up"}
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = nr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(NodeEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Disable a node identified by the node name.
func (nr *NodeResource) Disable(name string) error {
	item := Node{Session: "user-disabled", State: "user-up"}
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = nr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(NodeEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// ForceOffline a node identified by the node name.
func (nr *NodeResource) ForceOffline(name string) error {
	item := Node{Session: "user-disabled", State: "user-down"}
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = nr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(NodeEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single node identified by the node name. if it is not exist return error
func (nr *NodeResource) Delete(name string) error {
	_, err := nr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(NodeEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}

//func (nr *NodeResource) ShowStats(id string) (*NodeStatsList, error) {
//	var item NodeStatsList
//	if err := nr.c.ReadQuery(bigip.GetBaseResource()+NodeEndpoint+"/"+id+"/stats", &item); err != nil {
//		return nil, err
//	}
//	return &item, nil
//}
