package sys

// FailoverConfigList holds a list of Failover configuration.
type FailoverConfigList struct {
	Items    []FailoverConfig `json:"items"`
	Kind     string           `json:"kind"`
	SelfLink string           `json:"selflink"`
}

// FailoverConfig holds the configuration of a single Failover.
type FailoverConfig struct {
}

// FailoverEndpoint represents the REST resource for managing Failover.
const FailoverEndpoint = "/failover"
