package net

import (
	"github.com/lefeck/go-bigip"
)

const NetManager = "net"

type Net struct {
	inet        InetResource
	inetStats   InetStatsResource
	route       RouteResource
	vlan        VlanResource
	self        SelfResource
	routeDomain RouteDomainResource
	trunk       TrunkResource
	address     AddressResource
	port        PortResource
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
		port:        PortResource{b: b},
	}
}

// Inet returns a InetResource used to query tm/net/interface API.
func (net Net) Inet() *InetResource {
	return &net.inet
}

// InetStats returns a InetStatsResource used to query tm/net/interface/stats API.
func (net Net) InetStats() *InetStatsResource {
	return &net.inetStats
}

// Route returns a RouteResource used to query tm/net/route API.
func (net Net) Route() *RouteResource {
	return &net.route
}

// Vlan returns a VlanResource used to query /tm/net/vlan API.
func (net Net) Vlan() *VlanResource {
	return &net.vlan
}

// Self returns a SelfResource used to query /tm/net/self API.
func (net Net) Self() *SelfResource {
	return &net.self
}

// RouteDomain returns a RouteDomainResource used to query /tm/net/route-domain API.
func (net Net) RouteDomain() *RouteDomainResource {
	return &net.routeDomain
}

// Trunk returns a TrunkResource used to query /tm/net/trunk API.
func (net Net) Trunk() *TrunkResource {
	return &net.trunk
}

// AddressList returns a AddressResource used to query /tm/net/address-list API.
func (net Net) AddressList() *AddressResource {
	return &net.address
}

// PortList returns a PortResource used to query /tm/net/port-list API.
func (net Net) PortList() *PortResource {
	return &net.port
}
