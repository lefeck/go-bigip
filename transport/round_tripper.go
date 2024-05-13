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
		rt, err = NewBearerAuthWithRefreshRoundTripper(config.BearerToken, config.BearerTokenFile, rt)
		if err != nil {
			return nil, err
		}
	case config.HasBasicAuth():
		rt = NewBasicAuthRoundTripper(config.Username, config.Password, rt)
	}
	return rt, nil
}

type basicAuthRoundTripper struct {
	username string
	password string
	rt       http.RoundTripper
}

var _ RoundTripperWrapper = &basicAuthRoundTripper{}

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

type bearerAuthRoundTripper struct {
	bearer string
	source oauth2.TokenSource
	rt     http.RoundTripper
}

func (rt *bearerAuthRoundTripper) WrappedRoundTripper() http.RoundTripper {
	return rt.rt
}

var _ RoundTripperWrapper = &bearerAuthRoundTripper{}

// NewBearerAuthRoundTripper adds the provided bearer token to a request
// unless the authorization header has already been set.
func NewBearerAuthRoundTripper(bearer string, rt http.RoundTripper) http.RoundTripper {
	return &bearerAuthRoundTripper{bearer, nil, rt}
}

func (rt *bearerAuthRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if len(req.Header.Get("X-F5-Auth-Token")) != 0 {
		return rt.rt.RoundTrip(req)
	}
	req = CloneRequest(req)
	token := rt.bearer
	if rt.source != nil {
		if refreshedToken, err := rt.source.Token(); err == nil {
			token = refreshedToken.AccessToken
		}
	}
	req.Header.Set("X-F5-Auth-Token", token)
	return rt.rt.RoundTrip(req)
}

//func (rt *bearerAuthRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
//	if len(req.Header.Get("Authorization")) != 0 {
//		return rt.rt.RoundTrip(req)
//	}
//
//	// 创建一个新的请求
//	newReq, err := http.NewRequest(req.Method, req.URL.String(), req.Body)
//	if err != nil {
//		return nil, err
//	}
//	newReq.Header = req.Header.Clone()
//	token := rt.bearer
//
//	if rt.source != nil {
//		if refreshedToken, err := rt.source.Token(); err == nil {
//			token = refreshedToken.AccessToken
//		}
//	}
//	newReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
//	fmt.Println(token)
//	return rt.rt.RoundTrip(newReq)
//}

func NewBearerAuthWithRefreshRoundTripper(bearer string, tokenFile string, rt http.RoundTripper) (http.RoundTripper, error) {
	if len(tokenFile) == 0 {
		return &bearerAuthRoundTripper{bearer, nil, rt}, nil
	}
	source := NewCachedFileTokenSource(tokenFile)
	if len(bearer) == 0 {
		token, err := source.Token()
		if err != nil {
			return nil, err
		}
		bearer = token.AccessToken
	}
	return &bearerAuthRoundTripper{bearer, source, rt}, nil
}

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
