// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ecm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// Config holds the configuration of a single Config.
type Config struct {
	Kind                           string `json:"kind"`
	SelfLink                       string `json:"selfLink"`
	MaxNumNodes                    int    `json:"maxNumNodes"`
	NodesDynDisc                   string `json:"nodesDynDisc"`
	NodesDynDiscUpd                string `json:"nodesDynDiscUpd"`
	SeedIP                         string `json:"seedIp"`
	SslKeyPairName                 string `json:"sslKeyPairName"`
	TenantID                       string `json:"tenantId"`
	TransportManagerMaxConnections int    `json:"transportManagerMaxConnections"`
}

// ConfigEndpoint represents the REST resource for managing Config.
const ConfigEndpoint = "config"

// ConfigResource provides an API to manage Config configurations.
type ConfigResource struct {
	b *bigip.BigIP
}

// Get retrieves the details of a single Config by node name.
func (r *ConfigResource) Get(name string) (*Config, error) {
	var item Config
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ECMEndpoint).SubResource(ConfigEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new Config item.
func (r *ConfigResource) Create(item Config) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ECMEndpoint).SubResource(ConfigEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the Config item identified by the Config name.
func (r *ConfigResource) Update(name string, item Config) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ECMEndpoint).SubResource(ConfigEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single Config identified by the Config name. If it does not exist, return an error.
func (r *ConfigResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ECMEndpoint).SubResource(ConfigEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
