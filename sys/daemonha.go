package sys

import "github.com/lefeck/go-bigip"

// DaemonHAConfigList holds a list of DaemonHA configuration.
type DaemonHAConfigList struct {
	Items    []DaemonHAConfig `json:"items"`
	Kind     string           `json:"kind"`
	SelfLink string           `json:"selflink"`
}

// DaemonHAConfig holds the configuration of a single DaemonHA.
type DaemonHAConfig struct {
	FullPath         string `json:"fullPath"`
	Generation       int    `json:"generation"`
	Heartbeat        string `json:"heartbeat"`
	HeartbeatAction  string `json:"heartbeatAction"`
	Kind             string `json:"kind"`
	Name             string `json:"name"`
	NotRunningAction string `json:"notRunningAction"`
	Running          string `json:"running"`
	RunningTimeout   int    `json:"runningTimeout"`
	SelfLink         string `json:"selfLink"`
}

// DaemonHAEndpoint represents the REST resource for managing DaemonHA.
const DaemonHAEndpoint = "/daemon-ha"

// DaemonHAResource provides an API to manage DaemonHA configurations.
type DaemonHAResource struct {
	b *bigip.BigIP
}
