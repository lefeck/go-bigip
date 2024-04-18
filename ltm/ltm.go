// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ltm provides a REST client for the /tm/ltm F5 BigIP API.
package ltm

import (
	"github.com/lefeck/bigip"
)

// BasePath is the base path of the LTM API.

// LTM implements a REST client for the F5 BigIP LTM API.
type LTM struct {
	b *bigip.BigIP

	virtual        VirtualResource
	virtualAddress VirtualAddressResource
}

// New creates a new LTM client.
func New(b *bigip.BigIP) LTM {
	return LTM{
		b:              b,
		virtual:        VirtualResource{b: b},
		virtualAddress: VirtualAddressResource{c: b},
	}
}

// Virtual returns a VirtualResource configured to query tm/ltm/virtual API.
func (ltm LTM) Virtual() *VirtualResource {
	return &ltm.virtual
}

func (ltm LTM) VirtualAddress() *VirtualAddressResource {
	return &ltm.virtualAddress
}
