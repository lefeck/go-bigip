package rest

import (
	"fmt"
	"net/url"
	"path"
)

func DefaultServerURL(host, apiPath string) (*url.URL, string, error) {
	if host == "" {
		return nil, "", fmt.Errorf("host must be a URL or a host:port pair")
	}
	base := host
	hostURL, err := url.Parse(base)
	if err != nil || hostURL.Scheme == "" || hostURL.Host == "" {
		scheme := "https://"
		hostURL, err = url.Parse(scheme + base)
		if err != nil {
			return nil, "", err
		}
		if hostURL.Path != "" && hostURL.Path != "/" {
			return nil, "", fmt.Errorf("host must be a URL or a host:port pair: %q", base)
		}
	}

	baseAPIPath := DefaultAPIPath(apiPath)

	return hostURL, baseAPIPath, nil
}

// DefaultAPIPath constructs the default path, assuming the given
// API path, following the standard conventions of the BigIP API.
func DefaultAPIPath(apiPath string) string {
	baseAPIPath := path.Join("/", apiPath)
	return baseAPIPath
}

// DefaultServerUrlFor is shared RESTClientFor. It requires Host to be set prior to being called.
func DefaultServerUrlFor(config *Config) (*url.URL, string, error) {
	host := config.Host
	if host == "" {
		host = "localhost"
	}
	return DefaultServerURL(host, config.APIPath)
}
