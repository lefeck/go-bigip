// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// A RouteDomainList includes a list of RouteDomain.
type RouteDomainList struct {
	Items    []RouteDomain `json:"items,omitempty"`
	Kind     string        `json:"kind,omitempty"`
	SelfLink string        `json:"selfLink,omitempty"`
}

// A RouteDomain includes the configuration for a route domain.
type RouteDomain struct {
	ConnectionLimit int      `json:"connectionLimit,omitempty"`
	FullPath        string   `json:"fullPath,omitempty"`
	Generation      int      `json:"generation,omitempty"`
	ID              int      `json:"id,omitempty"`
	Kind            string   `json:"kind,omitempty"`
	Name            string   `json:"name,omitempty"`
	SelfLink        string   `json:"selfLink,omitempty"`
	Strict          string   `json:"strict,omitempty"`
	Vlans           []string `json:"vlans,omitempty"`
	VlansReference  []struct {
		Link string `json:"link,omitempty"`
	} `json:"vlansReference,omitempty"`
}

// RouteDomainEndpoint represents the REST resource for managing a route domain.
const RouteDomainEndpoint = "route-domain"

// A RouteDomainResource provides API to manage route domain configuration.
type RouteDomainResource struct {
	b *bigip.BigIP
}

// List lists all the route domain configurations.
func (rdr *RouteDomainResource) List() (*RouteDomainList, error) {
	var items RouteDomainList
	res, err := rdr.b.RestClient.Get().Prefix(bigip.BasePath).ResourceCategory(bigip.TMResource).ManagerName(NetManager).Resource(RouteDomainEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get a single route domain configuration identified by name.
func (rdr *RouteDomainResource) Get(fullPathName string) (*RouteDomain, error) {
	var item RouteDomain
	res, err := rdr.b.RestClient.Get().Prefix(bigip.BasePath).ResourceCategory(bigip.TMResource).ManagerName(NetManager).
		Resource(RouteDomainEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create a new route domain configuration.
func (rdr *RouteDomainResource) Create(item RouteDomain) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = rdr.b.RestClient.Post().Prefix(bigip.BasePath).ResourceCategory(bigip.TMResource).ManagerName(NetManager).
		Resource(RouteDomainEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update a route domain configuration identified by name.
func (rdr *RouteDomainResource) Edit(fullPathName string, item RouteDomain) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = rdr.b.RestClient.Put().Prefix(bigip.BasePath).ResourceCategory(bigip.TMResource).ManagerName(NetManager).
		Resource(RouteDomainEndpoint).ResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single route domain configuration identified by name.
func (rdr *RouteDomainResource) Delete(fullPathName string) error {
	_, err := rdr.b.RestClient.Delete().Prefix(bigip.BasePath).ResourceCategory(bigip.TMResource).ManagerName(NetManager).
		Resource(RouteDomainEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
