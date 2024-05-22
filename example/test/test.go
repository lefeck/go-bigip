package main

import (
	"fmt"
	"net/url"
	"path"
	"strings"
)

func validateCmdForBash(cmd string) string {
	cmd = strings.TrimSpace(cmd)
	if strings.HasPrefix(cmd, "-c ") {
		last := strings.TrimPrefix(cmd, "-c ")
		last = strings.TrimSpace(last)
		if !strings.HasPrefix(last, "'") || !strings.HasSuffix(last, "'") {
			last = "'" + last + "'"
		}
		cmd = "-c " + last
	}
	return cmd
}

func api() {
	apiPath := "/hello/pod"
	fullpath := path.Join("/", apiPath)
	fmt.Println(fullpath)
}

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

func main() {
	rawURL, bigip.GetBaseResource(), _ := DefaultServerURL("192.168.13.91", "/mgmt/shared/authn/login")
	urls := rawURL.String() + bigip.GetBaseResource()
	fmt.Println(rawURL, urls, bigip.GetBaseResource())
}
