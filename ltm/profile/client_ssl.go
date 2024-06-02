package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type ClientSSLList struct {
	Items    []ClientSSL `json:"items,omitempty"`
	Kind     string      `json:"kind,omitempty"`
	SelfLink string      `json:"selflink,omitempty"`
}

type ClientSSL struct {
	Kind                     string `json:"kind,omitempty"`
	Name                     string `json:"name,omitempty"`
	Partition                string `json:"partition,omitempty"`
	FullPath                 string `json:"fullPath,omitempty"`
	Generation               int    `json:"generation,omitempty"`
	SelfLink                 string `json:"selfLink,omitempty"`
	AlertTimeout             string `json:"alertTimeout,omitempty"`
	AllowDynamicRecordSizing string `json:"allowDynamicRecordSizing,omitempty"`
	AllowExpiredCrl          string `json:"allowExpiredCrl,omitempty"`
	AllowNonSsl              string `json:"allowNonSsl,omitempty"`
	AppService               string `json:"appService,omitempty"`
	Authenticate             string `json:"authenticate,omitempty"`
	AuthenticateDepth        int    `json:"authenticateDepth,omitempty"`
	BypassOnClientCertFail   string `json:"bypassOnClientCertFail,omitempty"`
	BypassOnHandshakeAlert   string `json:"bypassOnHandshakeAlert,omitempty"`
	C3DClientFallbackCert    string `json:"c3dClientFallbackCert,omitempty"`
	C3DDropUnknownOcspStatus string `json:"c3dDropUnknownOcspStatus,omitempty"`
	C3DOcsp                  string `json:"c3dOcsp,omitempty"`
	CaFile                   string `json:"caFile,omitempty"`
	CacheSize                int    `json:"cacheSize,omitempty"`
	CacheTimeout             int    `json:"cacheTimeout,omitempty"`
	Cert                     string `json:"cert,omitempty"`
	CertReference            struct {
		Link string `json:"link,omitempty"`
	} `json:"certReference,omitempty"`
	CertExtensionIncludes  []string `json:"certExtensionIncludes,omitempty"`
	CertLifespan           int      `json:"certLifespan,omitempty"`
	CertLookupByIpaddrPort string   `json:"certLookupByIpaddrPort,omitempty"`
	Chain                  string   `json:"chain,omitempty"`
	CipherGroup            string   `json:"cipherGroup,omitempty"`
	Ciphers                string   `json:"ciphers,omitempty"`
	ClientCertCa           string   `json:"clientCertCa,omitempty"`
	Crl                    string   `json:"crl,omitempty"`
	CrlFile                string   `json:"crlFile,omitempty"`
	Data0Rtt               string   `json:"data_0rtt,omitempty"`
	DefaultsFrom           string   `json:"defaultsFrom,omitempty"`
	DefaultsFromReference  struct {
		Link string `json:"link,omitempty"`
	} `json:"defaultsFromReference,omitempty"`
	Description                     string        `json:"description,omitempty"`
	DestinationIPBlacklist          string        `json:"destinationIpBlacklist,omitempty"`
	DestinationIPWhitelist          string        `json:"destinationIpWhitelist,omitempty"`
	ForwardProxyBypassDefaultAction string        `json:"forwardProxyBypassDefaultAction,omitempty"`
	GenericAlert                    string        `json:"genericAlert,omitempty"`
	HandshakeTimeout                string        `json:"handshakeTimeout,omitempty"`
	HelloExtensionIncludes          []interface{} `json:"helloExtensionIncludes,omitempty"`
	HostnameBlacklist               string        `json:"hostnameBlacklist,omitempty"`
	HostnameWhitelist               string        `json:"hostnameWhitelist,omitempty"`
	InheritCaCertkeychain           string        `json:"inheritCaCertkeychain,omitempty"`
	InheritCertkeychain             string        `json:"inheritCertkeychain,omitempty"`
	Key                             string        `json:"key,omitempty"`
	KeyReference                    struct {
		Link string `json:"link,omitempty"`
	} `json:"keyReference,omitempty"`
	LogPublisher          string `json:"logPublisher,omitempty"`
	LogPublisherReference struct {
		Link string `json:"link,omitempty"`
	} `json:"logPublisherReference,omitempty"`
	MaxActiveHandshakes                string `json:"maxActiveHandshakes,omitempty"`
	MaxAggregateRenegotiationPerMinute string `json:"maxAggregateRenegotiationPerMinute,omitempty"`
	MaxRenegotiationsPerMinute         int    `json:"maxRenegotiationsPerMinute,omitempty"`
	MaximumRecordSize                  int    `json:"maximumRecordSize,omitempty"`
	ModSslMethods                      string `json:"modSslMethods,omitempty"`
	Mode                               string `json:"mode,omitempty"`
	NotifyCertStatusToVirtualServer    string `json:"notifyCertStatusToVirtualServer,omitempty"`
	OcspStapling                       string `json:"ocspStapling,omitempty"`
	TmOptions                          string `json:"tmOptions,omitempty"`
	PeerCertMode                       string `json:"peerCertMode,omitempty"`
	PeerNoRenegotiateTimeout           string `json:"peerNoRenegotiateTimeout,omitempty"`
	ProxyCaCert                        string `json:"proxyCaCert,omitempty"`
	ProxyCaKey                         string `json:"proxyCaKey,omitempty"`
	ProxySsl                           string `json:"proxySsl,omitempty"`
	ProxySslPassthrough                string `json:"proxySslPassthrough,omitempty"`
	RenegotiateMaxRecordDelay          string `json:"renegotiateMaxRecordDelay,omitempty"`
	RenegotiatePeriod                  string `json:"renegotiatePeriod,omitempty"`
	RenegotiateSize                    string `json:"renegotiateSize,omitempty"`
	Renegotiation                      string `json:"renegotiation,omitempty"`
	RetainCertificate                  string `json:"retainCertificate,omitempty"`
	SecureRenegotiation                string `json:"secureRenegotiation,omitempty"`
	ServerName                         string `json:"serverName,omitempty"`
	SessionMirroring                   string `json:"sessionMirroring,omitempty"`
	SessionTicket                      string `json:"sessionTicket,omitempty"`
	SessionTicketTimeout               int    `json:"sessionTicketTimeout,omitempty"`
	SniDefault                         string `json:"sniDefault,omitempty"`
	SniRequire                         string `json:"sniRequire,omitempty"`
	SourceIPBlacklist                  string `json:"sourceIpBlacklist,omitempty"`
	SourceIPWhitelist                  string `json:"sourceIpWhitelist,omitempty"`
	SslC3D                             string `json:"sslC3d,omitempty"`
	SslForwardProxy                    string `json:"sslForwardProxy,omitempty"`
	SslForwardProxyBypass              string `json:"sslForwardProxyBypass,omitempty"`
	SslForwardProxyVerifiedHandshake   string `json:"sslForwardProxyVerifiedHandshake,omitempty"`
	SslSignHash                        string `json:"sslSignHash,omitempty"`
	StrictResume                       string `json:"strictResume,omitempty"`
	UncleanShutdown                    string `json:"uncleanShutdown,omitempty"`
	CertKeyChain                       []struct {
		Name          string `json:"name,omitempty"`
		AppService    string `json:"appService,omitempty"`
		Cert          string `json:"cert,omitempty"`
		CertReference struct {
			Link string `json:"link,omitempty"`
		} `json:"certReference,omitempty"`
		Chain        string `json:"chain,omitempty"`
		Key          string `json:"key,omitempty"`
		KeyReference struct {
			Link string `json:"link,omitempty"`
		} `json:"keyReference,omitempty"`
		Usage string `json:"usage,omitempty"`
	} `json:"certKeyChain,omitempty"`
}

const ClientSSLEndpoint = "client-ssl"

type ClientSSLResource struct {
	b *bigip.BigIP
}

func (mir *ClientSSLResource) List() (*ClientSSLList, error) {
	var items ClientSSLList
	res, err := mir.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ClientSSLEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

func (mir *ClientSSLResource) Get(fullPathName string) (*ClientSSL, error) {
	var item ClientSSL
	res, err := mir.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ClientSSLEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

func (mir *ClientSSLResource) Create(item ClientSSL) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mir.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ClientSSLEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mir *ClientSSLResource) Update(fullPathName string, item ClientSSL) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mir.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ClientSSLEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mir *ClientSSLResource) Delete(fullPathName string) error {
	_, err := mir.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ClientSSLEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
