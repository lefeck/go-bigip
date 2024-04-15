package transport

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type testRoundTripper struct {
	Request  *http.Request
	Response *http.Response
	Err      error
}

func (rt *testRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	rt.Request = req
	return rt.Response, rt.Err
}

func TestBasicAuthRoundTripper(t *testing.T) {
	for n, tc := range map[string]struct {
		user string
		pass string
	}{
		"basic":   {user: "user", pass: "pass"},
		"no pass": {user: "user"},
	} {
		rt := &testRoundTripper{}
		req := &http.Request{}
		NewBasicAuthRoundTripper(tc.user, tc.pass, rt).RoundTrip(req)
		if rt.Request == nil {
			t.Fatalf("%s: unexpected nil request: %v", n, rt)
		}
		if rt.Request == req {
			t.Fatalf("%s: round tripper should have copied request object: %#v", n, rt.Request)
		}
		if user, pass, found := rt.Request.BasicAuth(); !found || user != tc.user || pass != tc.pass {
			t.Errorf("%s: unexpected authorization header: %#v", n, rt.Request)
		}
	}
}

func TestBearerAuthRoundTripper(t *testing.T) {
	rt := &testRoundTripper{}
	req := &http.Request{}
	NewBearerAuthRoundTripper("test", rt).RoundTrip(req)
	if rt.Request == nil {
		t.Fatalf("unexpected nil request: %v", rt)
	}
	if rt.Request == req {
		t.Fatalf("round tripper should have copied request object: %#v", rt.Request)
	}
	if rt.Request.Header.Get("Authorization") != "Bearer test" {
		t.Errorf("unexpected authorization header: %#v", rt.Request)
	}
}

func TestHTTPWrappersForConfig(t *testing.T) {
	tests := []struct {
		name       string
		config     Config
		expectedRT string
	}{
		{
			name: "Bearer Auth",
			config: Config{
				BearerToken: "testtoken",
			},
			expectedRT: "bearerAuthRoundTripper",
		},
		{
			name: "Basic Auth",
			config: Config{
				Username: "testuser",
				Password: "testpassword",
			},
			expectedRT: "basicAuthRoundTripper",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rt, err := HTTPWrappersForConfig(&test.config, http.DefaultTransport)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			assert.NotNil(t, rt)

			switch test.expectedRT {
			case "bearerAuthRoundTripper":
				assert.IsType(t, &bearerAuthRoundTripper{}, rt)
			case "basicAuthRoundTripper":
				assert.IsType(t, &basicAuthRoundTripper{}, rt)
			default:
				t.Errorf("unknown RoundTripper type")
			}
		})
	}
}
