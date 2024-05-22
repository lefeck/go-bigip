// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"github.com/lefeck/go-bigip"
)

const NetManager = "net"

type Net struct {
	b           *bigip.BigIP
	inet        InetResource
	inetStats   InetStatsResource
	route       RouteResource
	vlan        VlanResource
	self        SelfResource
	routeDomain RouteDomainResource
	trunk       TrunkResource
	address     AddressResource
}

// New creates a new NET client.
func New(b *bigip.BigIP) Net {
	return Net{
		inet:        InetResource{b: b},
		inetStats:   InetStatsResource{b: b},
		route:       RouteResource{b: b},
		vlan:        VlanResource{b: b},
		self:        SelfResource{b: b},
		routeDomain: RouteDomainResource{b: b},
		trunk:       TrunkResource{b: b},
		address:     AddressResource{b: b},
	}
}

// Inet returns a InetResource ured to query tm/net/interface API.
func (net Net) Inet() *InetResource {
	return &net.inet
}

func (net Net) InetStats() *InetStatsResource {
	return &net.inetStats
}

// Route returns a RouteResource ured to query tm/net/route API.
func (net Net) Route() *RouteResource {
	return &net.route
}

// Vlan returns a VlanResource ured to query /tm/net/vlan API.
func (net Net) Vlan() *VlanResource {
	return &net.vlan
}

// Self returns a SelfResource ured to query /tm/net/self API.
func (net Net) Self() *SelfResource {
	return &net.self
}

// RouteDomain returns a RouteDomainResource ured to query /tm/net/route-domain API.
func (net Net) RouteDomain() *RouteDomainResource {
	return &net.routeDomain
}

// RouteDomain returns a RouteDomainResource ured to query /tm/net/route-domain API.
func (net Net) Trunk() *TrunkResource {
	return &net.trunk
}

// RouteDomain returns a RouteDomainResource ured to query /tm/net/route-domain API.
func (net Net) AddressList() *AddressResource {
	return &net.address
}
