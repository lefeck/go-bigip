package profile

import "github.com/lefeck/go-bigip"

type StreamList struct {
	Items    []Stream `json:"items,omitempty"`
	Kind     string   `json:"kind,omitempty"`
	SelfLink string   `json:"selflink,omitempty"`
}

type Stream struct {
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

const StreamEndpoint = "stream"

type StreamResource struct {
	b *bigip.BigIP
}
