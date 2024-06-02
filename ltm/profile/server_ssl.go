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
	Kind                             string        `json:"kind,omitempty"`
	Name                             string        `json:"name,omitempty"`
	Partition                        string        `json:"partition,omitempty"`
	FullPath                         string        `json:"fullPath,omitempty"`
	Generation                       int           `json:"generation,omitempty"`
	SelfLink                         string        `json:"selfLink,omitempty"`
	AlertTimeout                     string        `json:"alertTimeout,omitempty"`
	AllowExpiredCrl                  string        `json:"allowExpiredCrl,omitempty"`
	AppService                       string        `json:"appService,omitempty"`
	Authenticate                     string        `json:"authenticate,omitempty"`
	AuthenticateDepth                int           `json:"authenticateDepth,omitempty"`
	AuthenticateName                 string        `json:"authenticateName,omitempty"`
	BypassOnClientCertFail           string        `json:"bypassOnClientCertFail,omitempty"`
	BypassOnHandshakeAlert           string        `json:"bypassOnHandshakeAlert,omitempty"`
	C3DCaCert                        string        `json:"c3dCaCert,omitempty"`
	C3DCaKey                         string        `json:"c3dCaKey,omitempty"`
	C3DCertExtensionCustomOids       []interface{} `json:"c3dCertExtensionCustomOids,omitempty"`
	C3DCertExtensionIncludes         []string      `json:"c3dCertExtensionIncludes,omitempty"`
	C3DCertLifespan                  int           `json:"c3dCertLifespan,omitempty"`
	CaFile                           string        `json:"caFile,omitempty"`
	CacheSize                        int           `json:"cacheSize,omitempty"`
	CacheTimeout                     int           `json:"cacheTimeout,omitempty"`
	Cert                             string        `json:"cert,omitempty"`
	Chain                            string        `json:"chain,omitempty"`
	CipherGroup                      string        `json:"cipherGroup,omitempty"`
	Ciphers                          string        `json:"ciphers,omitempty"`
	Crl                              string        `json:"crl,omitempty"`
	CrlFile                          string        `json:"crlFile,omitempty"`
	Data0Rtt                         string        `json:"data_0rtt,omitempty"`
	DefaultsFrom                     string        `json:"defaultsFrom,omitempty"`
	Description                      string        `json:"description,omitempty"`
	ExpireCertResponseControl        string        `json:"expireCertResponseControl,omitempty"`
	GenericAlert                     string        `json:"genericAlert,omitempty"`
	HandshakeTimeout                 string        `json:"handshakeTimeout,omitempty"`
	Key                              string        `json:"key,omitempty"`
	LogPublisher                     string        `json:"logPublisher,omitempty"`
	MaxActiveHandshakes              string        `json:"maxActiveHandshakes,omitempty"`
	ModSslMethods                    string        `json:"modSslMethods,omitempty"`
	Mode                             string        `json:"mode,omitempty"`
	Ocsp                             string        `json:"ocsp,omitempty"`
	TmOptions                        string        `json:"tmOptions,omitempty"`
	PeerCertMode                     string        `json:"peerCertMode,omitempty"`
	ProxySsl                         string        `json:"proxySsl,omitempty"`
	ProxySslPassthrough              string        `json:"proxySslPassthrough,omitempty"`
	RenegotiatePeriod                string        `json:"renegotiatePeriod,omitempty"`
	RenegotiateSize                  string        `json:"renegotiateSize,omitempty"`
	Renegotiation                    string        `json:"renegotiation,omitempty"`
	RetainCertificate                string        `json:"retainCertificate,omitempty"`
	RevokedCertStatusResponseControl string        `json:"revokedCertStatusResponseControl,omitempty"`
	SecureRenegotiation              string        `json:"secureRenegotiation,omitempty"`
	ServerName                       string        `json:"serverName,omitempty"`
	SessionMirroring                 string        `json:"sessionMirroring,omitempty"`
	SessionTicket                    string        `json:"sessionTicket,omitempty"`
	SniDefault                       string        `json:"sniDefault,omitempty"`
	SniRequire                       string        `json:"sniRequire,omitempty"`
	SslC3D                           string        `json:"sslC3d,omitempty"`
	SslForwardProxy                  string        `json:"sslForwardProxy,omitempty"`
	SslForwardProxyBypass            string        `json:"sslForwardProxyBypass,omitempty"`
	SslForwardProxyVerifiedHandshake string        `json:"sslForwardProxyVerifiedHandshake,omitempty"`
	SslSignHash                      string        `json:"sslSignHash,omitempty"`
	StrictResume                     string        `json:"strictResume,omitempty"`
	UncleanShutdown                  string        `json:"uncleanShutdown,omitempty"`
	UnknownCertStatusResponseControl string        `json:"unknownCertStatusResponseControl,omitempty"`
	UntrustedCertResponseControl     string        `json:"untrustedCertResponseControl,omitempty"`
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
