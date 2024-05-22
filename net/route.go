package net

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// A RouteList holds a list of Route.
type RouteList struct {
	Items    []Route `json:"items,omitempty"`
	Kind     string  `json:"kind,omitempty"`
	SelfLink string  `json:"selfLink,omitempty"`
}

// A Route hold the uration for a route.
type Route struct {
	FullPath   string `json:"fullPath,omitempty"`
	Generation int    `json:"generation,omitempty"`
	Gw         string `json:"gw,omitempty"`
	Kind       string `json:"kind,omitempty"`
	Mtu        int    `json:"mtu,omitempty"`
	Name       string `json:"name,omitempty"`
	Network    string `json:"network,omitempty"`
	Partition  string `json:"partition,omitempty"`
	SelfLink   string `json:"selfLink,omitempty"`
}

// RouteEndpoint represents the REST resource for managing a route.
const RouteEndpoint = "route"

// A RouteResource provides API to manage routes uration.
type RouteResource struct {
	b *bigip.BigIP
}

// List lists all the route urations.
func (rr *RouteResource) List() (*RouteList, error) {
	var rl RouteList
	res, err := rr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).Resource(RouteEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &rl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &rl, nil
}

// Get a single route uration identified by id.
func (rr *RouteResource) Get(fullPathName string) (*Route, error) {
	res, err := rr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).
		Resource(RouteEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	var route Route
	if err := json.Unmarshal(res, &route); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &route, nil
}

// Create a new route uration.
func (rr *RouteResource) Create(item Route) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = rr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).
		Resource(RouteEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Edit a route uration identified by id.
func (rr *RouteResource) Update(name string, item Route) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = rr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).
		Resource(RouteEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single route uration identified by id.
func (rr *RouteResource) Delete(name string) error {
	_, err := rr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(NetManager).
		Resource(RouteEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
