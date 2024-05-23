// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ecm

import "github.com/lefeck/go-bigip"

const SysManager = "sys"

// ECMEndpoint represents the REST resource for managing ECM.
const ECMEndpoint = "ecm"

type ECMResource struct {
	config ConfigResource
}

func NewECM(b *bigip.BigIP) ECMResource {
	return ECMResource{
		config: ConfigResource{b: b},
	}
}

func (ecm ECMResource) Element() *ConfigResource {
	return &ecm.config
}
