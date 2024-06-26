package rest

import (
	"net/http"
	"net/url"
	"strings"
)

// Interface captures the set of operations for generically interacting with BigIP REST apis.
type Interface interface {
	Verb(verb string) *Request
	Post() *Request
	Put() *Request
	Patch() *Request
	Get() *Request
	Delete() *Request
}

// ClientContent controls how RESTClient communicates with the server.
type ClientContentConfig struct {
	AcceptContentTypes string
	ContentType        string
}
type RESTClient struct {
	Base        *url.URL
	baseAPIPath string
	// content describes how a RESTClient encodes and decodes responses.
	content ClientContentConfig
	// Set specific behavior of the client.  If not set http.DefaultClient will be used.
	Client *http.Client
}

var _ Interface = &RESTClient{}

func NewRESTClient(baseURL *url.URL, baseAPIPath string, config ClientContentConfig, client *http.Client) (*RESTClient, error) {
	if len(config.ContentType) == 0 {
		config.ContentType = "application/json"
	}
	base := *baseURL
	if !strings.HasSuffix(base.Path, "/") {
		base.Path += "/"
	}

	base.RawQuery = ""
	base.Fragment = ""

	return &RESTClient{
		Base:        &base,
		baseAPIPath: baseAPIPath,
		content:     config,
		Client:      client,
	}, nil
}

// Verb begins a request with a verb (GET, POST, PUT, DELETE).
//
// Example usage of RESTClient's request building interface:
// c, err := NewRESTClient(...)
// if err != nil { ... }
// resp, err := c.Verb("GET").
//
//	Path("virtual").
//	Timeout(10*time.Second).
//	Do()
//
// if err != nil { ... }
// list, ok := resp.(*api.List)
func (c *RESTClient) Verb(verb string) *Request {
	return NewRequest(c).Verb(verb)
}

// Post begins a POST request. Short for c.Verb("POST").
func (c *RESTClient) Post() *Request {
	return c.Verb(http.MethodPost)
}

// Put begins a PUT request. Short for c.Verb("PUT").
func (c *RESTClient) Put() *Request {
	return c.Verb(http.MethodPut)
}

// Patch begins a PATCH request. Short for c.Verb("Patch").
func (c *RESTClient) Patch() *Request {
	return c.Verb(http.MethodPatch)
}

// Get begins a GET request. Short for c.Verb("GET").
func (c *RESTClient) Get() *Request {
	return c.Verb(http.MethodGet)
}

// Delete begins a DELETE request. Short for c.Verb("DELETE").
func (c *RESTClient) Delete() *Request {
	return c.Verb(http.MethodDelete)
}

//func (c *RESTClient) Patch(pt PatchType) *Request {
//	return c.Verb(http.MethodPatch).SetHeader("Content-Type", string(pt))
//}

// TODO: specify header type for url
type PatchType string

const (
	JSONPatchType  PatchType = "application/json-patch+json"
	MergePatchType PatchType = "application/merge-patch+json"
)
