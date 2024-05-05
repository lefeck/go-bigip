package rest

import (
	"github.com/lefeck/go-bigip/transport"
	"net/http"
	"time"
)

type Config struct {
	Host    string
	APIPath string
	ContentConfig
	Username        string
	Password        string
	BearerToken     string
	BearerTokenFile string
	// The maximum length of time to wait before giving up on a server request. A value of zero means no timeout.
	Timeout       time.Duration
	Transport     http.RoundTripper
	WrapTransport transport.WrapperFunc
}

type ContentConfig struct {
	// ContentType specifies the wire format used to communicate with the server.
	// This value will be set as the Accept header on requests made to the server, and
	// as the default content type on any object sent to the server. If not set,
	// "application/json" is used.
	ContentType string
}

func RESTClientForConfigAndClient(config *Config, httpClient *http.Client) (*RESTClient, error) {
	// 对url做检验和处理
	baseURL, baseAPIPath, err := DefaultServerUrlFor(config)
	if err != nil {
		return nil, err
	}
	//添加content
	clientContent := ClientContentConfig{
		ContentType: config.ContentType,
	}
	// 初始化http工作,为下一步做处理
	restClient, err := NewRESTClient(baseURL, baseAPIPath, clientContent, httpClient)

	return restClient, err
}
