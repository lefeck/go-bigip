package bigip

import (
	"github.com/lefeck/bigip/rest"
	"time"
)

type BigIP struct {
	RestClient *rest.RESTClient
}

func NewSession(host, username, password string) (*BigIP, error) {
	config := &rest.Config{
		Host:     host,
		Username: username,
		Password: password,
		ContentConfig: rest.ContentConfig{
			AcceptContentTypes: "application/json",
			ContentType:        "application/json",
		},
	}

	restClient, err := restClientFor(config)
	if err != nil {
		return nil, err
	}

	return &BigIP{
		RestClient: restClient,
	}, nil
}

func NewToken(host, token string) (*BigIP, error) {
	config := &rest.Config{
		Host: host,
		ContentConfig: rest.ContentConfig{
			AcceptContentTypes: "application/json",
			ContentType:        "application/json",
		},
		BearerToken: token,
		Timeout:     10 * time.Second,
	}

	restClient, err := restClientFor(config)
	if err != nil {
		return nil, err
	}

	return &BigIP{
		RestClient: restClient,
	}, nil
}

func restClientFor(config *rest.Config) (*rest.RESTClient, error) {
	// Setup HTTP client with authentication and custom transport
	httpClient, err := rest.HTTPClientFor(config)
	if err != nil {
		return nil, err
	}

	return rest.RESTClientForConfigAndClient(config, httpClient)
}

//type TLSTransport struct {
//	transport *http.Transport
//}
//
//func (tlst *TLSTransport) DisableCertCheck() {
//	// 我想添加https登陆不检查证书， 在之前的http.RoundTripper实现这个功能，给出代码
//	tlst.transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
//}
