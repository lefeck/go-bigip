package transport

import (
	"fmt"
	"golang.org/x/oauth2"
	"net/http"
)

func HTTPWrappersFor(config *Config, rt http.RoundTripper) (http.RoundTripper, error) {
	if config.WrapTransport != nil {
		rt = config.WrapTransport(rt)
	}
	switch {
	case config.HasBasicAuth() && config.HasTokenAuth():
		return nil, fmt.Errorf("username/password or bearer token may be set, but not both")
	case config.HasTokenAuth():
		var err error
		rt = NewTokenAuthRoundTripper(config.BearerToken, rt)
		if err != nil {
			return nil, err
		}
	case config.HasBasicAuth():
		rt = NewBasicAuthRoundTripper(config.Username, config.Password, rt)
	}
	return rt, nil
}

type RoundTripperWrapper interface {
	http.RoundTripper
	WrappedRoundTripper() http.RoundTripper
}

type basicAuthRoundTripper struct {
	username string
	password string
	rt       http.RoundTripper
}

var _ RoundTripperWrapper = &basicAuthRoundTripper{}

// username and password login
func NewBasicAuthRoundTripper(username, password string, rt http.RoundTripper) http.RoundTripper {
	return &basicAuthRoundTripper{username: username, password: password, rt: rt}
}

func (rt *basicAuthRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if len(req.Header.Get("Authorization")) != 0 {
		return rt.rt.RoundTrip(req)
	}
	req = CloneRequest(req)

	req.SetBasicAuth(rt.username, rt.password)
	return rt.rt.RoundTrip(req)
}

func (rt *basicAuthRoundTripper) CancelRequest(req *http.Request) {
	tryCancelRequest(rt.WrappedRoundTripper(), req)
}

func (rt *basicAuthRoundTripper) WrappedRoundTripper() http.RoundTripper {
	return rt.rt
}

// token login
type tokenAuthRoundTripper struct {
	token  string
	source oauth2.TokenSource
	rt     http.RoundTripper
}

var _ RoundTripperWrapper = &tokenAuthRoundTripper{}

// NewTokenAuthRoundTripper adds the provided bearer token to a request
// unless the authorization header has already been set.
func NewTokenAuthRoundTripper(token string, rt http.RoundTripper) http.RoundTripper {
	return &tokenAuthRoundTripper{token, nil, rt}
}

func (rt *tokenAuthRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if len(req.Header.Get("X-F5-Auth-Token")) != 0 {
		return rt.rt.RoundTrip(req)
	}
	req = CloneRequest(req)
	token := rt.token
	if rt.source != nil {
		if refreshedToken, err := rt.source.Token(); err == nil {
			token = refreshedToken.AccessToken
		}
	}

	req.Header.Set("X-F5-Auth-Token", token)
	return rt.rt.RoundTrip(req)
}

func (rt *tokenAuthRoundTripper) WrappedRoundTripper() http.RoundTripper {
	return rt.rt
}

//func NewBearerAuthWithRefreshRoundTripper(bearer string, tokenFile string, rt http.RoundTripper) (http.RoundTripper, error) {
//	if len(tokenFile) == 0 {
//		return &bearerAuthRoundTripper{bearer, nil, rt}, nil
//	}
//	source := NewCachedFileTokenSource(tokenFile)
//	if len(bearer) == 0 {
//		token, err := source.Token()
//		if err != nil {
//			return nil, err
//		}
//		bearer = token.AccessToken
//	}
//	return &bearerAuthRoundTripper{bearer, source, rt}, nil
//}

// CloneRequest creates a shallow copy of the request along with a deep copy of the Headers.
func CloneRequest(req *http.Request) *http.Request {
	r := new(http.Request)

	// shallow clone
	*r = *req

	// deep copy headers
	r.Header = CloneHeader(req.Header)

	return r
}

// CloneHeader creates a deep copy of an http.Header.
func CloneHeader(in http.Header) http.Header {
	out := make(http.Header, len(in))
	for key, values := range in {
		newValues := make([]string, len(values))
		copy(newValues, values)
		out[key] = newValues
	}
	return out
}
