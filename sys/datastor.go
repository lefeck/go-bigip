package sys

import "github.com/lefeck/go-bigip"

// DataStorConfigList holds a list of DataStor configuration.
type DataStorConfigList struct {
	Items    []DataStorConfig `json:"items"`
	Kind     string           `json:"kind"`
	SelfLink string           `json:"selflink"`
}

// DataStorConfig holds the configuration of a single DataStor.
type DataStorConfig struct {
}

// DataStorEndpoint represents the REST resource for managing DataStor.
const DataStorEndpoint = "/datastor"

// DataStorResource provides an API to manage DataStor configurations.
type DataStorResource struct {
	b *bigip.BigIP
}
