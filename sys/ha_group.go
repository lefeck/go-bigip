package sys

import (
	"github.com/lefeck/go-bigip"
)

// HAGroupList holds a list of HAGroup configuration.
type HAGroupList struct {
	Items    []HAGroup `json:"items"`
	Kind     string    `json:"kind"`
	SelfLink string    `json:"selflink"`
}

// HAGroup holds the configuration of a single HAGroup.
type HAGroup struct {
}

// HAGroupEndpoint represents the REST resource for managing HAGroup.
const HAGroupEndpoint = "ha-group"

// HAGroupResource provides an API to manage HAGroup configurations.
type HAGroupResource struct {
	b *bigip.BigIP
}
