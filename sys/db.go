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

// A DBList holds a list of DB.
type DBList struct {
	Items    []DB   `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selfLink,omitempty"`
}

// A DB holds the configuration for a db.
type DB struct {
	DefaultValue string `json:"defaultValue"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	Kind         string `json:"kind"`
	Name         string `json:"name"`
	ScfConfig    string `json:"scfConfig"`
	SelfLink     string `json:"selfLink"`
	Value        string `json:"value"`
	ValueRange   string `json:"valueRange"`
}

const DBEndpoint = "db"

type DBResource struct {
	b *bigip.BigIP
}

// ListAll lists all the db configurations.
func (dr *DBResource) List() (*DBList, error) {
	var dl DBList
	res, err := dr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(DBEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &dl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &dl, nil
}

// Get a single db configuration identified by fullPathName.
func (dr *DBResource) Get(fullPathName string) (*DB, error) {
	var db DB
	res, err := dr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(DBEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &db); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &db, nil
}

// Create a new db configuration.
func (dr *DBResource) Create(item DB) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = dr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(DBEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Edit a db configuration identified by name.
func (dr *DBResource) Update(name string, item DB) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = dr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(DBEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single db configuration identified by name.
func (dr *DBResource) Delete(name string) error {
	_, err := dr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(DBEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
