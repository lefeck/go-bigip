package transport

import (
	"net/http"
)

// Config holds various options for establishing a transport.
type Config struct {

	// Username and password for basic authentication
	Username string
	Password string `datapolicy:"password"`

	// Bearer token for authentication
	BearerToken string `datapolicy:"token"`

	// Path to a file containing a BearerToken.
	// If set, the contents are periodically read.
	// The last successfully read value takes precedence over BearerToken.
	BearerTokenFile string

	// WrapTransport for most client level operations.
	Transport http.RoundTripper

	WrapTransport WrapperFunc
}

func (c *Config) Wrap(fn WrapperFunc) {
	c.WrapTransport = Wrappers(c.WrapTransport, fn)
}

func (c *Config) HasBasicAuth() bool {
	return len(c.Username) != 0
}

func (c *Config) HasTokenAuth() bool {
	return len(c.BearerToken) != 0 || len(c.BearerTokenFile) != 0
}
