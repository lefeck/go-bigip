package sys

import "github.com/lefeck/go-bigip"

// FixConnectionConfigList holds a list of FixConnection configuration.
type FixConnectionConfigList struct {
	Items    []FixConnectionConfig `json:"items"`
	Kind     string                `json:"kind"`
	SelfLink string                `json:"selflink"`
}

// FixConnectionConfig holds the configuration of a single FixConnection.
type FixConnectionConfig struct {
}

// FixConnectionEndpoint represents the REST resource for managing FixConnection.
const FixConnectionEndpoint = "/fix-connection"

// FixConnectionResource provides an API to manage FixConnection configurations.
type FixConnectionResource struct {
	b *bigip.BigIP
}
