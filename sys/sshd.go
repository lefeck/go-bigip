// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
)

// SSHD holds the configuration of a single SSHD.
type SSHD struct {
	Kind              string   `json:"kind"`
	SelfLink          string   `json:"selfLink"`
	Allow             []string `json:"allow"`
	Banner            string   `json:"banner"`
	FipsCipherVersion int      `json:"fipsCipherVersion"`
	InactivityTimeout int      `json:"inactivityTimeout"`
	LogLevel          string   `json:"logLevel"`
	Login             string   `json:"login"`
	Port              int      `json:"port"`
}

// SSHDEndpoint represents the REST resource for managing SSHD.
const SSHDEndpoint = "sshd"

// SSHDResource provides an API to manage SSHD configurations.
type SSHDResource struct {
	b *bigip.BigIP
}

// ListAll  lists all the SSHD configurations.
func (r *SSHDResource) List() (*SSHD, error) {
	var item SSHD
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SSHDEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}
