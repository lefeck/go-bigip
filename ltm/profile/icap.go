package profile

type icap struct {
	Kind          string `json:"kind"`
	Name          string `json:"name"`
	Partition     string `json:"partition"`
	FullPath      string `json:"fullPath"`
	Generation    int    `json:"generation"`
	SelfLink      string `json:"selfLink"`
	AppService    string `json:"appService"`
	DefaultsFrom  string `json:"defaultsFrom"`
	HeaderFrom    string `json:"headerFrom"`
	Host          string `json:"host"`
	PreviewLength int    `json:"previewLength"`
	Referer       string `json:"referer"`
	URI           string `json:"uri"`
	UserAgent     string `json:"userAgent"`
}
