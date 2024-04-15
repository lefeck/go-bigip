package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

type Request struct {
	C            *RESTClient
	timeout      time.Duration
	maxRetries   int
	verb         string
	PathPrefix   string
	subpath      string
	params       url.Values
	headers      http.Header
	partition    string
	partitionSet bool
	resource     string
	subResource  string
	resourceName string
	body         io.Reader
	bodyBytes    []byte
	err          error
}

func NewRequest(c *RESTClient) *Request {
	var pathPrefix string
	if c.Base != nil {
		pathPrefix = path.Join("/", c.Base.Path, c.baseAPIPath)
	} else {
		pathPrefix = path.Join("/", c.baseAPIPath)
	}
	var timeout time.Duration
	if c.Client != nil {
		timeout = c.Client.Timeout
	}
	r := Request{
		C:          c,
		timeout:    timeout,
		maxRetries: 10,
		PathPrefix: pathPrefix,
	}
	switch {
	case len(c.content.AcceptContentTypes) > 0:
		r.SetHeader("Accept", c.content.AcceptContentTypes)
	case len(c.content.ContentType) > 0:
		r.SetHeader("Accept", c.content.ContentType+", */*")
	}
	return &r
}

func NewRequestWithClient(base *url.URL, baseAPIPath string, content ClientContentConfig, client *http.Client) *Request {
	return NewRequest(&RESTClient{
		Base:        base,
		baseAPIPath: baseAPIPath,
		content:     content,
		Client:      client,
	})
}

func (r *Request) Verb(verb string) *Request {
	r.verb = verb
	return r
}

/*
	{
	   "link": "https://localhost/mgmt/tm/ltm/persistence/sip"
	 },
	 {
	   "link": "https://localhost/mgmt/tm/sys/restricted-module"
	 },

https://IP/mgmt/tm/<module name>/<subresource>
*/
func (r *Request) Prefix(segments ...string) *Request {
	if r.err != nil {
		return r
	}
	r.PathPrefix = path.Join(r.PathPrefix, path.Join(segments...))
	return r
}

func (r *Request) Suffix(segments ...string) *Request {
	if r.err != nil {
		return r
	}
	r.subpath = path.Join(r.subpath, path.Join(segments...))
	return r
}

func (r *Request) Resource(resource string) *Request {
	if r.err != nil {
		return r
	}
	if len(r.resource) != 0 {
		r.err = fmt.Errorf("resource already set to %q, cannot change to %q", r.resource, resource)
		return r
	}
	if msgs := IsValidPathSegmentName(resource); len(msgs) != 0 {
		r.err = fmt.Errorf("invalid resource %q: %v", resource, msgs)
		return r
	}
	r.resource = resource
	return r
}

func (r *Request) SubResource(subResources ...string) *Request {
	if r.err != nil {
		return r
	}
	subResource := path.Join(subResources...)
	if len(r.subResource) != 0 {
		r.err = fmt.Errorf("subresource already set to %q, cannot change to %q", r.subResource, subResource)
		return r
	}
	for _, s := range subResources {
		if msgs := IsValidPathSegmentName(s); len(msgs) != 0 {
			r.err = fmt.Errorf("invalid subresource %q: %v", s, msgs)
			return r
		}
	}
	r.subResource = subResource
	return r
}
func (r *Request) SubResourcePath(parts []string) string {
	var buffer bytes.Buffer
	var lastPath int
	if strings.HasPrefix(parts[len(parts)-1], "?") {
		lastPath = len(parts) - 2
	} else {
		lastPath = len(parts) - 1
	}
	for i, p := range parts {
		buffer.WriteString(strings.Replace(p, "/", "~", -1))
		if i < lastPath {
			buffer.WriteString("/")
		}
	}
	return buffer.String()
}

// Name sets the name of a resource to access (<resource>/[ns/<namespace>/]<name>)
func (r *Request) Name(resourceName string) *Request {
	if r.err != nil {
		return r
	}
	if len(resourceName) == 0 {
		r.err = fmt.Errorf("resource name may not be empty")
		return r
	}
	if len(r.resourceName) != 0 {
		r.err = fmt.Errorf("resource name already set to %q, cannot change to %q", r.resourceName, resourceName)
		return r
	}
	if msgs := IsValidPathSegmentName(resourceName); len(msgs) != 0 {
		r.err = fmt.Errorf("invalid resource name %q: %v", resourceName, msgs)
		return r
	}
	r.resourceName = resourceName
	return r
}

func (r *Request) AbsPath(segments ...string) *Request {
	if r.err != nil {
		return r
	}
	r.PathPrefix = path.Join(r.C.Base.Path, path.Join(segments...))
	if len(segments) == 1 && (len(r.C.Base.Path) > 1 || len(segments) > 1) && strings.HasSuffix(segments[0], "/") {
		r.PathPrefix += "/"
	}
	return r
}

func (r *Request) SetParams(paramName, s string) *Request {
	if r.params != nil {
		return r
	}
	return r.setParams(paramName, s)
}

func (r *Request) setParams(paramName, value string) *Request {
	if r.params == nil {
		r.params = make(url.Values)
	}
	r.params[paramName] = append(r.params[paramName], value)
	return r
}

func (r *Request) SetHeader(key string, values ...string) *Request {
	if r.headers == nil {
		r.headers = http.Header{}
	}
	r.headers.Del(key)
	for _, value := range values {
		r.headers.Add(key, value)
	}
	return r
}

func (r *Request) Timeout(d time.Duration) *Request {
	if r.err != nil {
		return r
	}
	r.timeout = d
	return r
}

func (r *Request) MaxRetries(maxRetries int) *Request {
	if maxRetries < 0 {
		r.maxRetries = 0
	}
	r.maxRetries = maxRetries
	return r
}

func (r *Request) Error() error {
	return r.err
}

// NameMayNotBe specifies strings that cannot be used as names specified as path segments (like the REST API or etcd store)
var NameMayNotBe = []string{".", ".."}

// NameMayNotContain specifies substrings that cannot be used in names specified as path segments (like the REST API or etcd store)
var NameMayNotContain = []string{"/", "%"}

// IsValidPathSegmentName validates the name can be safely encoded as a path segment
func IsValidPathSegmentName(name string) []string {
	for _, illegalName := range NameMayNotBe {
		if name == illegalName {
			return []string{fmt.Sprintf(`may not be '%s'`, illegalName)}
		}
	}

	var errors []string
	for _, illegalContent := range NameMayNotContain {
		if strings.Contains(name, illegalContent) {
			errors = append(errors, fmt.Sprintf(`may not contain '%s'`, illegalContent))
		}
	}

	return errors
}

// IsValidPathSegmentPrefix validates the name can be used as a prefix for a name which will be encoded as a path segment
// It does not check for exact matches with disallowed names, since an arbitrary suffix might make the name valid
func IsValidPathSegmentPrefix(name string) []string {
	var errors []string
	for _, illegalContent := range NameMayNotContain {
		if strings.Contains(name, illegalContent) {
			errors = append(errors, fmt.Sprintf(`may not contain '%s'`, illegalContent))
		}
	}

	return errors
}

// ValidatePathSegmentName validates the name can be safely encoded as a path segment
func ValidatePathSegmentName(name string, prefix bool) []string {
	if prefix {
		return IsValidPathSegmentPrefix(name)
	}
	return IsValidPathSegmentName(name)
}

func (r *Request) Partition(partition string) *Request {
	if r.err != nil {
		return r
	}
	if r.partitionSet {
		r.err = fmt.Errorf("partition already set to %q, cannot change to %q", r.partitionSet, partition)
		return r
	}
	if msgs := IsValidPathSegmentName(partition); len(msgs) != 0 {
		r.err = fmt.Errorf("invalid namespace %q: %v", partition, msgs)
		return r
	}
	r.partitionSet = true
	r.partition = partition
	return r
}

// NamespaceIfScoped is a convenience function to set a namespace if scoped is true
func (r *Request) PartitionIfScoped(partition string, scoped bool) *Request {
	if scoped {
		return r.Partition(partition)
	}
	return r
}

func (r *Request) requestPreflightCheck() error {
	if !r.partitionSet {
		return nil
	}
	if len(r.partition) > 0 {
		return nil
	}

	switch r.verb {
	case "POST":
		return fmt.Errorf("an empty partition may not be set during creation")
	case "GET", "PUT", "DELETE":
		if len(r.resourceName) > 0 {
			return fmt.Errorf("an empty partition may not be set when a resource name is provided")
		}
	}
	return nil
}

func (r *Request) Request(ctx context.Context, data interface{}) (*http.Response, error) {
	client := r.C.Client
	if client == nil {
		client = http.DefaultClient
	}
	if r.timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, r.timeout)
		defer cancel()
	}

	if err := r.requestPreflightCheck(); err != nil {
		return nil, err
	}

	req, err := r.newHTTPRequest(data)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		if strings.HasPrefix(resp.Header.Get("Content-Type"), "application/json") {
			return nil, r.checkError(responseBody)
		}
		return nil, fmt.Errorf("HTTP %d :: %s", resp.StatusCode, string(responseBody))
	}
	return resp, nil
}

func (r *Request) checkError(resp []byte) error {
	if len(resp) == 0 {
		return nil
	}

	var reqError RequestInfo

	err := json.Unmarshal(resp, &reqError)
	if err != nil {
		return fmt.Errorf("%s\n%s", err.Error(), string(resp[:]))
	}
	return nil
}

// URL returns the current working URL. Check the result of Error() to ensure
// that the returned URL is valid.
func (r *Request) URL() *url.URL {
	p := r.PathPrefix
	if r.partitionSet && len(r.partition) > 0 {
		p = path.Join(p, "partition", r.partition)
	}
	if len(r.resource) != 0 {
		p = path.Join(p, strings.ToLower(r.resource))
	}
	// Join trims trailing slashes, so preserve r.pathPrefix's trailing slash for backwards compatibility if nothing was changed
	if len(r.resourceName) != 0 || len(r.subpath) != 0 || len(r.subResource) != 0 {
		p = path.Join(p, r.resourceName, r.subResource, r.subpath)
	}

	finalURL := &url.URL{}
	if r.C.Base != nil {
		*finalURL = *r.C.Base
	}
	finalURL.Path = p

	query := url.Values{}
	for key, values := range r.params {
		for _, value := range values {
			query.Add(key, value)
		}
	}

	// timeout is handled specially here.
	if r.timeout != 0 {
		query.Set("timeout", r.timeout.String())
	}
	finalURL.RawQuery = query.Encode()
	return finalURL
}

func (r *Request) newHTTPRequest(data interface{}) (*http.Request, error) {
	var body io.Reader
	switch {
	case r.body != nil && r.bodyBytes != nil:
		return nil, fmt.Errorf("cannot set both body and bodyBytes")
	case r.body != nil:
		body = r.body
	case r.bodyBytes != nil:
		// Create a new reader specifically for this request.
		// Giving each request a dedicated reader allows retries to avoid races resetting the request body.
		body = bytes.NewReader(r.bodyBytes)
	default:
		if data != nil {
			bodyBytes, err := json.Marshal(data)
			if err != nil {
				return nil, err
			}
			body = bytes.NewReader(bodyBytes)
		}
	}

	url := r.URL().String()
	req, err := http.NewRequest(r.verb, url, body)
	if err != nil {
		return nil, err
	}
	req.Header = r.headers
	return req, nil
}

//func (r *Request) Do(req *http.Request) (*http.Response, error) {
//	resp, err := r.c.Client.Do(req)
//	if err != nil {
//		return nil, err
//	}
//	if resp.StatusCode == 401 {
//		if err := c.authType(req); err != nil {
//			return nil, fmt.Errorf("cannot re-authenticate after 401: %v", err)
//		}
//	}
//	return resp, err
//}

// todo: otherwise remove, waiting for....
func (r *Request) Errors(resp *http.Response) error {
	if resp.StatusCode >= 400 {
		if contentType := resp.Header.Get("Content-Type"); !strings.Contains(contentType, "application/json") {
			return errors.New("http response error: " + resp.Status)
		}
		errResp, err := NewRequestError(resp.Body)
		if err != nil {
			return errors.New("cannot read error message from response body: " + err.Error())
		}
		return errResp
	}
	return nil
}

type RequestInfo struct {
	Code     int      `json:"code,omitempty"`
	Message  string   `json:"message,omitempty"`
	ErrStack []string `json:"errorStack,omitempty"`
}

// NewRequestError unmarshal a RequestError from a HTTP response body.
func NewRequestError(body io.Reader) (*RequestInfo, error) {
	var reqErr RequestInfo
	dec := json.NewDecoder(body)
	if err := dec.Decode(&reqErr); err != nil {
		return nil, fmt.Errorf("failed to decode request error: %v", err)
	}
	return &reqErr, nil
}

// Error implements the errors.Error interface
func (err RequestInfo) Error() string {
	return fmt.Sprintf("%s (code: %d)", err.Message, err.Code)
}

func (err RequestInfo) String() string {
	buf := bytes.NewBufferString(err.Error())
	for _, es := range err.ErrStack {
		buf.WriteString("\n   " + es)
	}
	return buf.String()
}
