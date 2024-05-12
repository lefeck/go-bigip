package profile

type stream struct {
	Kind         string `json:"kind"`
	Name         string `json:"name"`
	Partition    string `json:"partition"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	SelfLink     string `json:"selfLink"`
	AppService   string `json:"appService"`
	ChunkSize    int    `json:"chunkSize"`
	Chunking     string `json:"chunking"`
	DefaultsFrom string `json:"defaultsFrom"`
	Description  string `json:"description"`
	Source       string `json:"source"`
	TmTarget     string `json:"tmTarget"`
}
