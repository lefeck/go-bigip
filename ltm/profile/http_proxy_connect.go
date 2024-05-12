package profile

type HttpProxyConnect struct {
	Kind         string `json:"kind"`
	Name         string `json:"name"`
	Partition    string `json:"partition"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	SelfLink     string `json:"selfLink"`
	AppService   string `json:"appService"`
	DefaultState string `json:"defaultState"`
	DefaultsFrom string `json:"defaultsFrom"`
	Description  string `json:"description"`
}
