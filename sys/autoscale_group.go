package sys

import "github.com/lefeck/go-bigip"

// AutoscaleGroupConfigList holds a list of AutoscaleGroup configuration.
type AutoscaleGroupConfigList struct {
	Items    []AutoscaleGroupConfig `json:"items"`
	Kind     string                 `json:"kind"`
	SelfLink string                 `json:"selflink"`
}

// AutoscaleGroupConfig holds the configuration of a single AutoscaleGroup.
type AutoscaleGroupConfig struct {
}

// AutoscaleGroupEndpoint represents the REST resource for managing AutoscaleGroup.
const AutoscaleGroupEndpoint = "/autoscale-group"

// AutoscaleGroupResource provides an API to manage AutoscaleGroup configurations.
type AutoscaleGroupResource struct {
	b *bigip.BigIP
}
