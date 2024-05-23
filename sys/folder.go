package sys

// FolderConfigList holds a list of Folder configuration.
type FolderConfigList struct {
	Items    []FolderConfig `json:"items"`
	Kind     string         `json:"kind"`
	SelfLink string         `json:"selflink"`
}

// FolderConfig holds the configuration of a single Folder.
type FolderConfig struct {
	DeviceGroup           string `json:"deviceGroup"`
	FullPath              string `json:"fullPath"`
	Generation            int    `json:"generation"`
	Hidden                string `json:"hidden"`
	InheritedDevicegroup  string `json:"inheritedDevicegroup"`
	InheritedTrafficGroup string `json:"inheritedTrafficGroup"`
	Kind                  string `json:"kind"`
	Name                  string `json:"name"`
	NoRefCheck            string `json:"noRefCheck"`
	SelfLink              string `json:"selfLink"`
	TrafficGroup          string `json:"trafficGroup"`
	TrafficGroupReference struct {
		Link string `json:"link"`
	} `json:"trafficGroupReference"`
}

type NewFolderConfig struct {
	Name      string `json:"name"`
	Partition string `json:"partition"`
}

// FolderEndpoint represents the REST resource for managing Folder.
const FolderEndpoint = "/folder"
