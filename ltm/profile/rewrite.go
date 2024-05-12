package profile

type rewrite struct {
	Kind                  string        `json:"kind"`
	Name                  string        `json:"name"`
	Partition             string        `json:"partition"`
	FullPath              string        `json:"fullPath"`
	Generation            int           `json:"generation"`
	SelfLink              string        `json:"selfLink"`
	AppService            string        `json:"appService"`
	BypassList            []interface{} `json:"bypassList"`
	ClientCachingType     string        `json:"clientCachingType"`
	DefaultsFrom          string        `json:"defaultsFrom"`
	DefaultsFromReference struct {
		Link string `json:"link"`
	} `json:"defaultsFromReference"`
	JavaCaFile          string `json:"javaCaFile"`
	JavaCaFileReference struct {
		Link string `json:"link"`
	} `json:"javaCaFileReference"`
	JavaCrl              string `json:"javaCrl"`
	JavaSignKey          string `json:"javaSignKey"`
	JavaSignKeyReference struct {
		Link string `json:"link"`
	} `json:"javaSignKeyReference"`
	JavaSigner          string `json:"javaSigner"`
	JavaSignerReference struct {
		Link string `json:"link"`
	} `json:"javaSignerReference"`
	LocationSpecific string `json:"locationSpecific"`
	Request          struct {
		InsertXforwardedFor   string `json:"insertXforwardedFor"`
		InsertXforwardedHost  string `json:"insertXforwardedHost"`
		InsertXforwardedProto string `json:"insertXforwardedProto"`
		RewriteHeaders        string `json:"rewriteHeaders"`
	} `json:"request"`
	Response struct {
		RewriteContent string `json:"rewriteContent"`
		RewriteHeaders string `json:"rewriteHeaders"`
	} `json:"response"`
	RewriteList       []interface{} `json:"rewriteList"`
	RewriteMode       string        `json:"rewriteMode"`
	SplitTunneling    string        `json:"splitTunneling"`
	URIRulesReference struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"uriRulesReference"`
}
