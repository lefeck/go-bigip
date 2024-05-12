package profile

type service struct {
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Partition  string `json:"partition"`
	FullPath   string `json:"fullPath"`
	Generation int    `json:"generation"`
	SelfLink   string `json:"selfLink"`
	AppService string `json:"appService"`
	Type       string `json:"type"`
}
