// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// SNMP holds the configuration of a single SNMP.
type SNMP struct {
	Kind                 string   `json:"kind"`
	SelfLink             string   `json:"selfLink"`
	AgentAddresses       []string `json:"agentAddresses"`
	AgentTrap            string   `json:"agentTrap"`
	AllowedAddresses     []string `json:"allowedAddresses"`
	AuthTrap             string   `json:"authTrap"`
	BigipTraps           string   `json:"bigipTraps"`
	LoadMax1             int      `json:"loadMax1"`
	LoadMax15            int      `json:"loadMax15"`
	LoadMax5             int      `json:"loadMax5"`
	Snmpv1               string   `json:"snmpv1"`
	Snmpv2C              string   `json:"snmpv2c"`
	SysContact           string   `json:"sysContact"`
	SysLocation          string   `json:"sysLocation"`
	SysServices          int      `json:"sysServices"`
	TrapCommunity        string   `json:"trapCommunity"`
	TrapSource           string   `json:"trapSource"`
	CommunitiesReference struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"communitiesReference"`
	DiskMonitors []struct {
		Name         string `json:"name"`
		Minspace     int    `json:"minspace"`
		MinspaceType string `json:"minspaceType"`
		Path         string `json:"path"`
	} `json:"diskMonitors"`
	ProcessMonitors []struct {
		Name         string `json:"name"`
		MaxProcesses string `json:"maxProcesses"`
		MinProcesses int    `json:"minProcesses"`
		Process      string `json:"process"`
	} `json:"processMonitors"`
	TrapsReference struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"trapsReference"`
	UsersReference struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"usersReference"`
}

// SNMPEndpoint represents the REST resource for managing SNMP.
const SNMPEndpoint = "snmp"

// SNMPResource provides an API to manage SNMP configurations.
type SNMPResource struct {
	b *bigip.BigIP
}

// list all the snmp configurations.
func (sr *SNMPResource) Get() (*SNMP, error) {
	var sl SNMP
	res, err := sr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SNMPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &sl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &sl, nil
}

// Create a new snmp configuration.
func (sr *SNMPResource) Create(item SNMP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = sr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SNMPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Edit a snmp configuration identified by name.
func (sr *SNMPResource) Update(name string, item SNMP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = sr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SNMPEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single snmp configuration identified by name.
func (sr *SNMPResource) Delete(name string) error {
	_, err := sr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SNMPEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
