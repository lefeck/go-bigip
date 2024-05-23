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

// NTPConfig holds the configuration of a single NTP.
type NTP struct {
	Kind              string `json:"kind,omitempty"`
	RestrictReference struct {
		IsSubcollection bool   `json:"isSubcollection,omitempty"`
		Link            string `json:"link,omitempty"`
	} `json:"restrictReference,omitempty"`
	SelfLink string   `json:"selfLink,omitempty"`
	Servers  []string `json:"servers,omitempty"`
	Timezone string   `json:"timezone,omitempty"`
}

// NTPEndpoint represents the REST resource for managing NTP.
const NTPEndpoint = "ntp"

// NTPResource provides an API to manage NTP configurations.
type NTPResource struct {
	b *bigip.BigIP
}

// Get a single ntp configuration identified by fullPathName.
func (nr *NTPResource) Get() (*NTP, error) {
	var ntp NTP
	res, err := nr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(NTPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &ntp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &ntp, nil
}

// Create a new ntp configuration.
func (nr *NTPResource) Create(item NTP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = nr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(NTPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Edit a ntp configuration identified by name.
func (nr *NTPResource) Update(name string, item NTP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = nr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(NTPEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (nr *NTPResource) AddServersForNTP(rs ...string) error {
	if len(rs) == 0 {
		return nil
	}
	item, err := nr.Get()
	if err != nil {
		return err
	}
	item.Servers = append(item.Servers, rs...)

	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = nr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(NTPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}

// Delete a single ntp configuration identified by name.
func (nr *NTPResource) Delete(name string) error {
	_, err := nr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(NTPEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
