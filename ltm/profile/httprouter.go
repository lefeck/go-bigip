package profile

type httprouter struct {
	Kind     string `json:"kind"`
	SelfLink string `json:"selfLink"`
	Items    []struct {
		Kind         string `json:"kind"`
		Name         string `json:"name"`
		Partition    string `json:"partition"`
		FullPath     string `json:"fullPath"`
		Generation   int    `json:"generation"`
		SelfLink     string `json:"selfLink"`
		AppService   string `json:"appService"`
		DefaultsFrom string `json:"defaultsFrom"`
		Description  string `json:"description"`
	} `json:"items"`
}
