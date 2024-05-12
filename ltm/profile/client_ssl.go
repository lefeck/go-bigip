package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"github.com/lefeck/go-bigip/ltm"
	"strings"
)

type ClientSSLConfigList struct {
	Items    []ClientSSLConfig `json:"items,omitempty"`
	Kind     string            `json:"kind,omitempty"`
	SelfLink string            `json:"selflink,omitempty"`
}

type ClientSSLConfig struct {
	Kind                     string `json:"kind"`
	Name                     string `json:"name"`
	Partition                string `json:"partition"`
	FullPath                 string `json:"fullPath"`
	Generation               int    `json:"generation"`
	SelfLink                 string `json:"selfLink"`
	AlertTimeout             string `json:"alertTimeout"`
	AllowDynamicRecordSizing string `json:"allowDynamicRecordSizing"`
	AllowExpiredCrl          string `json:"allowExpiredCrl"`
	AllowNonSsl              string `json:"allowNonSsl"`
	AppService               string `json:"appService"`
	Authenticate             string `json:"authenticate"`
	AuthenticateDepth        int    `json:"authenticateDepth"`
	BypassOnClientCertFail   string `json:"bypassOnClientCertFail"`
	BypassOnHandshakeAlert   string `json:"bypassOnHandshakeAlert"`
	C3DClientFallbackCert    string `json:"c3dClientFallbackCert"`
	C3DDropUnknownOcspStatus string `json:"c3dDropUnknownOcspStatus"`
	C3DOcsp                  string `json:"c3dOcsp"`
	CaFile                   string `json:"caFile"`
	CacheSize                int    `json:"cacheSize"`
	CacheTimeout             int    `json:"cacheTimeout"`
	Cert                     string `json:"cert"`
	CertReference            struct {
		Link string `json:"link"`
	} `json:"certReference"`
	CertExtensionIncludes  []string `json:"certExtensionIncludes"`
	CertLifespan           int      `json:"certLifespan"`
	CertLookupByIpaddrPort string   `json:"certLookupByIpaddrPort"`
	Chain                  string   `json:"chain"`
	CipherGroup            string   `json:"cipherGroup"`
	Ciphers                string   `json:"ciphers"`
	ClientCertCa           string   `json:"clientCertCa"`
	Crl                    string   `json:"crl"`
	CrlFile                string   `json:"crlFile"`
	Data0Rtt               string   `json:"data_0rtt"`
	DefaultsFrom           string   `json:"defaultsFrom"`
	DefaultsFromReference  struct {
		Link string `json:"link"`
	} `json:"defaultsFromReference"`
	Description                     string        `json:"description"`
	DestinationIPBlacklist          string        `json:"destinationIpBlacklist"`
	DestinationIPWhitelist          string        `json:"destinationIpWhitelist"`
	ForwardProxyBypassDefaultAction string        `json:"forwardProxyBypassDefaultAction"`
	GenericAlert                    string        `json:"genericAlert"`
	HandshakeTimeout                string        `json:"handshakeTimeout"`
	HelloExtensionIncludes          []interface{} `json:"helloExtensionIncludes"`
	HostnameBlacklist               string        `json:"hostnameBlacklist"`
	HostnameWhitelist               string        `json:"hostnameWhitelist"`
	InheritCaCertkeychain           string        `json:"inheritCaCertkeychain"`
	InheritCertkeychain             string        `json:"inheritCertkeychain"`
	Key                             string        `json:"key"`
	KeyReference                    struct {
		Link string `json:"link"`
	} `json:"keyReference"`
	LogPublisher          string `json:"logPublisher"`
	LogPublisherReference struct {
		Link string `json:"link"`
	} `json:"logPublisherReference"`
	MaxActiveHandshakes                string `json:"maxActiveHandshakes"`
	MaxAggregateRenegotiationPerMinute string `json:"maxAggregateRenegotiationPerMinute"`
	MaxRenegotiationsPerMinute         int    `json:"maxRenegotiationsPerMinute"`
	MaximumRecordSize                  int    `json:"maximumRecordSize"`
	ModSslMethods                      string `json:"modSslMethods"`
	Mode                               string `json:"mode"`
	NotifyCertStatusToVirtualServer    string `json:"notifyCertStatusToVirtualServer"`
	OcspStapling                       string `json:"ocspStapling"`
	TmOptions                          string `json:"tmOptions"`
	PeerCertMode                       string `json:"peerCertMode"`
	PeerNoRenegotiateTimeout           string `json:"peerNoRenegotiateTimeout"`
	ProxyCaCert                        string `json:"proxyCaCert"`
	ProxyCaKey                         string `json:"proxyCaKey"`
	ProxySsl                           string `json:"proxySsl"`
	ProxySslPassthrough                string `json:"proxySslPassthrough"`
	RenegotiateMaxRecordDelay          string `json:"renegotiateMaxRecordDelay"`
	RenegotiatePeriod                  string `json:"renegotiatePeriod"`
	RenegotiateSize                    string `json:"renegotiateSize"`
	Renegotiation                      string `json:"renegotiation"`
	RetainCertificate                  string `json:"retainCertificate"`
	SecureRenegotiation                string `json:"secureRenegotiation"`
	ServerName                         string `json:"serverName"`
	SessionMirroring                   string `json:"sessionMirroring"`
	SessionTicket                      string `json:"sessionTicket"`
	SessionTicketTimeout               int    `json:"sessionTicketTimeout"`
	SniDefault                         string `json:"sniDefault"`
	SniRequire                         string `json:"sniRequire"`
	SourceIPBlacklist                  string `json:"sourceIpBlacklist"`
	SourceIPWhitelist                  string `json:"sourceIpWhitelist"`
	SslC3D                             string `json:"sslC3d"`
	SslForwardProxy                    string `json:"sslForwardProxy"`
	SslForwardProxyBypass              string `json:"sslForwardProxyBypass"`
	SslForwardProxyVerifiedHandshake   string `json:"sslForwardProxyVerifiedHandshake"`
	SslSignHash                        string `json:"sslSignHash"`
	StrictResume                       string `json:"strictResume"`
	UncleanShutdown                    string `json:"uncleanShutdown"`
	CertKeyChain                       []struct {
		Name          string `json:"name"`
		AppService    string `json:"appService"`
		Cert          string `json:"cert"`
		CertReference struct {
			Link string `json:"link"`
		} `json:"certReference"`
		Chain        string `json:"chain"`
		Key          string `json:"key"`
		KeyReference struct {
			Link string `json:"link"`
		} `json:"keyReference"`
		Usage string `json:"usage"`
	} `json:"certKeyChain"`
}

const ClientSSLEndpoint = "client-ssl"

type ProfileClientSSLSoucre struct {
	B *bigip.BigIP
}

func (mir *ProfileClientSSLSoucre) List() (*ClientSSLConfigList, error) {
	var micl ClientSSLConfigList
	res, err := mir.B.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(ClientSSLEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &micl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &micl, nil
}

func (mir *ProfileClientSSLSoucre) Get(fullPathName string) (*ClientSSLConfig, error) {
	var mic ClientSSLConfig
	res, err := mir.B.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(ClientSSLEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mic); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mic, nil
}

func (mir *ProfileClientSSLSoucre) Create(item ClientSSLConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mir.B.RestClient.Post().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(ClientSSLEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mir *ProfileClientSSLSoucre) Update(name string, item ClientSSLConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mir.B.RestClient.Put().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(ClientSSLEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mir *ProfileClientSSLSoucre) Delete(name string) error {
	_, err := mir.B.RestClient.Delete().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(ClientSSLEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
