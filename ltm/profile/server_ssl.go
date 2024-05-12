package profile

type ServerSsl struct {
	Kind                       string        `json:"kind"`
	Name                       string        `json:"name"`
	Partition                  string        `json:"partition"`
	FullPath                   string        `json:"fullPath"`
	Generation                 int           `json:"generation"`
	SelfLink                   string        `json:"selfLink"`
	AlertTimeout               string        `json:"alertTimeout"`
	AllowExpiredCrl            string        `json:"allowExpiredCrl"`
	AppService                 string        `json:"appService"`
	Authenticate               string        `json:"authenticate"`
	AuthenticateDepth          int           `json:"authenticateDepth"`
	AuthenticateName           string        `json:"authenticateName"`
	BypassOnClientCertFail     string        `json:"bypassOnClientCertFail"`
	BypassOnHandshakeAlert     string        `json:"bypassOnHandshakeAlert"`
	C3DCaCert                  string        `json:"c3dCaCert"`
	C3DCaKey                   string        `json:"c3dCaKey"`
	C3DCertExtensionCustomOids []interface{} `json:"c3dCertExtensionCustomOids"`
	C3DCertExtensionIncludes   []string      `json:"c3dCertExtensionIncludes"`
	C3DCertLifespan            int           `json:"c3dCertLifespan"`
	CaFile                     string        `json:"caFile"`
	CacheSize                  int           `json:"cacheSize"`
	CacheTimeout               int           `json:"cacheTimeout"`
	Cert                       string        `json:"cert"`
	CertReference              struct {
		Link string `json:"link"`
	} `json:"certReference"`
	Chain                 string `json:"chain"`
	CipherGroup           string `json:"cipherGroup"`
	Ciphers               string `json:"ciphers"`
	Crl                   string `json:"crl"`
	CrlFile               string `json:"crlFile"`
	Data0Rtt              string `json:"data_0rtt"`
	DefaultsFrom          string `json:"defaultsFrom"`
	DefaultsFromReference struct {
		Link string `json:"link"`
	} `json:"defaultsFromReference"`
	Description               string `json:"description"`
	ExpireCertResponseControl string `json:"expireCertResponseControl"`
	GenericAlert              string `json:"genericAlert"`
	HandshakeTimeout          string `json:"handshakeTimeout"`
	Key                       string `json:"key"`
	KeyReference              struct {
		Link string `json:"link"`
	} `json:"keyReference"`
	LogPublisher          string `json:"logPublisher"`
	LogPublisherReference struct {
		Link string `json:"link"`
	} `json:"logPublisherReference"`
	MaxActiveHandshakes              string `json:"maxActiveHandshakes"`
	ModSslMethods                    string `json:"modSslMethods"`
	Mode                             string `json:"mode"`
	Ocsp                             string `json:"ocsp"`
	TmOptions                        string `json:"tmOptions"`
	PeerCertMode                     string `json:"peerCertMode"`
	ProxySsl                         string `json:"proxySsl"`
	ProxySslPassthrough              string `json:"proxySslPassthrough"`
	RenegotiatePeriod                string `json:"renegotiatePeriod"`
	RenegotiateSize                  string `json:"renegotiateSize"`
	Renegotiation                    string `json:"renegotiation"`
	RetainCertificate                string `json:"retainCertificate"`
	RevokedCertStatusResponseControl string `json:"revokedCertStatusResponseControl"`
	SecureRenegotiation              string `json:"secureRenegotiation"`
	ServerName                       string `json:"serverName"`
	SessionMirroring                 string `json:"sessionMirroring"`
	SessionTicket                    string `json:"sessionTicket"`
	SniDefault                       string `json:"sniDefault"`
	SniRequire                       string `json:"sniRequire"`
	SslC3D                           string `json:"sslC3d"`
	SslForwardProxy                  string `json:"sslForwardProxy"`
	SslForwardProxyBypass            string `json:"sslForwardProxyBypass"`
	SslForwardProxyVerifiedHandshake string `json:"sslForwardProxyVerifiedHandshake"`
	SslSignHash                      string `json:"sslSignHash"`
	StrictResume                     string `json:"strictResume"`
	UncleanShutdown                  string `json:"uncleanShutdown"`
	UnknownCertStatusResponseControl string `json:"unknownCertStatusResponseControl"`
	UntrustedCertResponseControl     string `json:"untrustedCertResponseControl"`
}
