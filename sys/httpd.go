package sys

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
)

// HTTPDConfig holds the configuration of a single HTTPD.
type HTTPDConfig struct {
	Kind                     string   `json:"kind"`
	SelfLink                 string   `json:"selfLink"`
	Allow                    []string `json:"allow"`
	AuthName                 string   `json:"authName"`
	AuthPamDashboardTimeout  string   `json:"authPamDashboardTimeout"`
	AuthPamIdleTimeout       int      `json:"authPamIdleTimeout"`
	AuthPamValidateIP        string   `json:"authPamValidateIp"`
	FastcgiTimeout           int      `json:"fastcgiTimeout"`
	FipsCipherVersion        int      `json:"fipsCipherVersion"`
	HostnameLookup           string   `json:"hostnameLookup"`
	LogLevel                 string   `json:"logLevel"`
	MaxClients               int      `json:"maxClients"`
	RedirectHTTPToHTTPS      string   `json:"redirectHttpToHttps"`
	RequestBodyMaxTimeout    int      `json:"requestBodyMaxTimeout"`
	RequestBodyMinRate       int      `json:"requestBodyMinRate"`
	RequestBodyTimeout       int      `json:"requestBodyTimeout"`
	RequestHeaderMaxTimeout  int      `json:"requestHeaderMaxTimeout"`
	RequestHeaderMinRate     int      `json:"requestHeaderMinRate"`
	RequestHeaderTimeout     int      `json:"requestHeaderTimeout"`
	SslCertfile              string   `json:"sslCertfile"`
	SslCertkeyfile           string   `json:"sslCertkeyfile"`
	SslCiphersuite           string   `json:"sslCiphersuite"`
	SslOcspDefaultResponder  string   `json:"sslOcspDefaultResponder"`
	SslOcspEnable            string   `json:"sslOcspEnable"`
	SslOcspOverrideResponder string   `json:"sslOcspOverrideResponder"`
	SslOcspResponderTimeout  int      `json:"sslOcspResponderTimeout"`
	SslOcspResponseMaxAge    int      `json:"sslOcspResponseMaxAge"`
	SslOcspResponseTimeSkew  int      `json:"sslOcspResponseTimeSkew"`
	SslPort                  int      `json:"sslPort"`
	SslProtocol              string   `json:"sslProtocol"`
	SslVerifyClient          string   `json:"sslVerifyClient"`
	SslVerifyDepth           int      `json:"sslVerifyDepth"`
}

// HTTPDEndpoint represents the REST resource for managing HTTPD.
const HTTPDEndpoint = "httpd"

// HTTPDResource provides an API to manage HTTPD configurations.
type HTTPDResource struct {
	b *bigip.BigIP
}

// Get the HTTPD configurations.
func (r *HTTPDResource) Get() (*HTTPDConfig, error) {
	var item HTTPDConfig
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(HTTPDEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}
