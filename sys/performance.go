package sys

// PerformanceConfigList holds a list of Performance configuration.
type PerformanceConfigList struct {
	Items    []PerformanceConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

// PerformanceConfig holds the configuration of a single Performance.
type PerformanceConfig struct {
}

// PerformanceEndpoint represents the REST resource for managing Performance.
const PerformanceEndpoint = "/performance"
