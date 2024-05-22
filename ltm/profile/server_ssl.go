package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type ServerSSLList struct {
	Items    []ServerSSL `json:"items,omitempty"`
	Kind     string      `json:"kind,omitempty"`
	SelfLink string      `json:"selflink,omitempty"`
}

type ServerSSL struct {
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

const ServerSSLEndpoint = "server-ssl"

type ServerSSLResource struct {
	b *bigip.BigIP
}

// List retrieves a list of ServerSSL resources.
func (cr *ServerSSLResource) List() (*ServerSSLList, error) {
	var items ServerSSLList
	// Perform a GET request to retrieve a list of ServerSSL resource objects
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ServerSSLEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into ServerSSLList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves a ServerSSL resource by its full path name.
func (cr *ServerSSLResource) Get(fullPathName string) (*ServerSSL, error) {
	var item ServerSSL
	// Perform a GET request to retrieve a specific ServerSSL resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ServerSSLEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into ServerSSL struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new ServerSSL resource using the provided ServerSSL item.
func (cr *ServerSSLResource) Create(item ServerSSL) error {
	// Marshal the ServerSSL struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new ServerSSL resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ServerSSLEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies a ServerSSL resource identified by its full path name using the provided ServerSSL item.
func (cr *ServerSSLResource) Update(fullPathName string, item ServerSSL) error {
	// Marshal the ServerSSL struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified ServerSSL resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ServerSSLEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a ServerSSL resource by its full path name.
func (cr *ServerSSLResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified ServerSSL resource
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ServerSSLEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
