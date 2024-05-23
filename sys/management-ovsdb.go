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
	"time"
)

// ManagementOVSDB holds the configuration of a single ManagementOVSDB.
type ManagementOVSDB struct {
	Kind         string `json:"kind"`
	Generation   int    `json:"generation"`
	APIRawValues struct {
		BaseBuild       string    `json:"base_build"`
		Build           string    `json:"build"`
		Built           string    `json:"built"`
		Changelist      string    `json:"changelist"`
		Edition         string    `json:"edition"`
		Encrypted       string    `json:"encrypted"`
		FileCreatedDate time.Time `json:"file_created_date"`
		FileSize        string    `json:"file_size"`
		Filename        string    `json:"filename"`
		InstallDate     string    `json:"install_date"`
		JobID           string    `json:"job_id"`
		Product         string    `json:"product"`
		Sequence        string    `json:"sequence"`
		Version         string    `json:"version"`
	} `json:"apiRawValues"`
}

// ManagementOVSDBEndpoint represents the REST resource for managing ManagementOVSDB.
const ManagementOVSDBEndpoint = "management-ovsdb"

// ManagementOVSDBResource provides an API to manage ManagementOVSDB configurations.
type ManagementOVSDBResource struct {
	b *bigip.BigIP
}

// List all management OVSDB details
func (r *ManagementOVSDBResource) List() (*ManagementOVSDB, error) {
	var items ManagementOVSDB
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ManagementOVSDBEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Update the management OVSDB item identified by the management OVSDB name, otherwise an error will be reported.
func (r *ManagementOVSDBResource) Update(name string, item ManagementOVSDB) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ManagementOVSDBEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single management OVSDB identified by the management OVSDB name. if it is not exist return error
func (r *ManagementOVSDBResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ManagementOVSDBEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
