package sys

import (
	"github.com/lefeck/go-bigip"
)

// RestrictedModuleConfigList holds a list of RestrictedModule configuration.
type RestrictedModuleConfigList struct {
	Items    []RestrictedModuleConfig `json:"items"`
	Kind     string                   `json:"kind"`
	SelfLink string                   `json:"selflink"`
}

// RestrictedModuleConfig holds the configuration of a single RestrictedModule.
type RestrictedModuleConfig struct {
}

// RestrictedModuleEndpoint represents the REST resource for managing RestrictedModule.
const RestrictedModuleEndpoint = "restricted-module"

// RestrictedModuleResource provides an API to manage RestrictedModule configurations.
type RestrictedModuleResource struct {
	b *bigip.BigIP
}
