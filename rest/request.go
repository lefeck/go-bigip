package rest

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"k8s.io/klog/v2"

	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"
)

type Request struct {
	c                *RESTClient
	timeout          time.Duration
	maxRetries       int
	verb             string
	pathPrefix       string
	subpath          string
	params           url.Values
	headers          http.Header
	fullPath         string
	resource         string
	resourceCategory string
	subResource      string
	managerName      string
	body             io.Reader
	bodyBytes        []byte
	err              error
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
		c:          c,
		timeout:    timeout,
		maxRetries: 10,
		pathPrefix: pathPrefix,
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
	r.pathPrefix = path.Join(r.pathPrefix, path.Join(segments...))
	return r
}

func (r *Request) Suffix(segments ...string) *Request {
	if r.err != nil {
		return r
	}
	r.subpath = path.Join(r.subpath, path.Join(segments...))
	return r
}

func (r *Request) ResourceCategory(resourceCategory string) *Request {
	if r.err != nil {
		return r
	}
	if len(r.resourceCategory) != 0 {
		r.err = fmt.Errorf("resourceCategory already set to %q, cannot change to %q", r.resourceCategory, resourceCategory)
		return r
	}
	if msgs := IsValidPathSegmentName(resourceCategory); len(msgs) != 0 {
		r.err = fmt.Errorf("invalid resourceCategory %q: %v", resourceCategory, msgs)
		return r
	}
	r.resourceCategory = resourceCategory
	return r
}

/*
/<api-prefix>/<resource-category>/<manager>/<resource-type>/<resource-instance>
*/
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

func (r *Request) ResourceInstance(fullPaths ...string) *Request {
	if r.err != nil {
		return r
	}
	fullPath := path.Join(fullPaths...)
	if len(r.fullPath) != 0 {
		fmt.Errorf("fullPath already set to %q, cannot change to %q", r.fullPath, fullPath)
		return r
	}
	newfullPath := replaceSlashInSubPath(fullPath)
	r.fullPath = newfullPath

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

func replaceSlashInSubPath(path string) string {
	parts := strings.Split(path, "/")
	for i := 1; i < len(parts); i++ {
		parts[i] = strings.ReplaceAll(parts[i], "/", "~")
	}
	return strings.Join(parts, "~")
}

// Name sets the name of a resource to access (<resource>/[ns/<namespace>/]<name>)
func (r *Request) ManagerName(managerName string) *Request {
	if r.err != nil {
		return r
	}
	if len(managerName) == 0 {
		r.err = fmt.Errorf("resource name may not be empty")
		return r
	}
	if len(r.managerName) != 0 {
		r.err = fmt.Errorf("resource name already set to %q, cannot change to %q", r.managerName, managerName)
		return r
	}
	if msgs := IsValidPathSegmentName(managerName); len(msgs) != 0 {
		r.err = fmt.Errorf("invalid resource name %q: %v", managerName, msgs)
		return r
	}
	r.managerName = managerName
	return r
}

func (r *Request) AbsPath(segments ...string) *Request {
	if r.err != nil {
		return r
	}
	r.pathPrefix = path.Join(r.c.Base.Path, path.Join(segments...))
	if len(segments) == 1 && (len(r.c.Base.Path) > 1 || len(segments) > 1) && strings.HasSuffix(segments[0], "/") {
		r.pathPrefix += "/"
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

// ReadError checks if a HTTP response contains an error and returns it.
func (r *Request) ReadError(resp *http.Response) error {
	if resp.StatusCode < http.StatusOK || resp.StatusCode > http.StatusPartialContent {
		if contentType := resp.Header.Get("Content-Type"); !strings.Contains(contentType, "application/json") {
			return fmt.Errorf("The http response error status code: %s \n", resp.Status)
		}
		errResp, err := NewRequestError(resp.Body)
		if err != nil {
			return errors.New("cannot read error message from response body: " + err.Error())
		}
		return errResp
	}
	return nil
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

func (r *Request) request(ctx context.Context, fn func(req *http.Request, resp *http.Response)) error {
	client := r.c.Client
	if client == nil {
		client = http.DefaultClient
	}
	if r.timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, r.timeout)
		defer cancel()
	}

	req, err := r.newHTTPRequest(ctx)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := r.ReadError(resp); err != nil {
		return err
	}

	fn(req, resp)

	return nil
}

// Body makes the request use obj as the body. Optional.
// If obj is a string, try to read a file of that name.
// If obj is a []byte, send it directly.
// If obj is an io.Reader, use it directly.
// Otherwise, set an error.
func (r *Request) Body(obj interface{}) *Request {
	if r.err != nil {
		return r
	}
	switch t := obj.(type) {
	case string:
		data, err := os.ReadFile(t)
		if err != nil {
			r.err = err
			return r
		}
		r.body = nil
		r.bodyBytes = data
	case []byte:
		r.body = nil
		r.bodyBytes = t
	case io.Reader:
		r.body = t
		r.bodyBytes = nil
	default:
		r.err = fmt.Errorf("unknown type used for body: %+v", obj)
	}
	return r
}

func (r *Request) Do(ctx context.Context) Result {
	var result Result
	err := r.request(ctx, func(req *http.Request, resp *http.Response) {
		result = r.transformResponse(resp, req)
	})
	if err != nil {
		return Result{Err: err}
	}
	return result
}

// DoRaw executes the request but does not process the response body.
func (r *Request) DoRaw(ctx context.Context) ([]byte, error) {
	var result Result
	err := r.request(ctx, func(req *http.Request, resp *http.Response) {
		result.Body, result.Err = io.ReadAll(resp.Body)
	})
	if err != nil {
		return nil, err
	}
	return result.Body, result.Err
}

// transformResponse converts an API response into a structured API object
func (r *Request) transformResponse(resp *http.Response, req *http.Request) Result {
	var body []byte
	if resp.Body != nil {
		data, err := io.ReadAll(resp.Body)
		switch err.(type) {
		case nil:
			body = data
		default:
			klog.Errorf("Unexpected error when reading response body: %v", err)
			unexpectedErr := fmt.Errorf("unexpected error when reading response body. Please retry. Original error: %w", err)
			return Result{
				Err: unexpectedErr,
			}
		}
	}

	contentType := resp.Header.Get("Content-Type")
	if len(contentType) == 0 {
		contentType = r.c.content.ContentType
	}
	if len(contentType) > 0 {
		var err error
		if err != nil {
			// if we fail to negotiate a decoder, treat this as an unstructured error
			switch {
			case resp.StatusCode == http.StatusSwitchingProtocols:
				// no-op, we've been upgraded
			case resp.StatusCode < http.StatusOK || resp.StatusCode > http.StatusPartialContent:
				return Result{}
			}
			return Result{
				Body:        body,
				ContentType: contentType,
				Code:        resp.StatusCode,
			}
		}
	}

	switch {
	case resp.StatusCode == http.StatusSwitchingProtocols:
		// no-op, we've been upgraded
	case resp.StatusCode < http.StatusOK || resp.StatusCode > http.StatusPartialContent:

		return Result{
			Body:        body,
			ContentType: contentType,
			Code:        resp.StatusCode,
		}
	}

	return Result{
		Body:        body,
		ContentType: contentType,
		Code:        resp.StatusCode,
	}
}

// Result contains the result of calling Request.Do().
type Result struct {
	Body        []byte
	ContentType string
	Err         error
	Code        int
}

// URL returns the current working URL. Check the result of Error() to ensure
// that the returned URL is valid.
func (r *Request) URL() *url.URL {
	p := r.pathPrefix

	// Join trims trailing slashes, so preserve r.pathPrefix's trailing slash for backwards compatibility if nothing was changed
	// TODO: anything else
	if len(r.resourceCategory) != 0 || len(r.managerName) != 0 || len(r.resource) != 0 || len(r.subpath) != 0 || len(r.subResource) != 0 {
		p = path.Join(p, r.resourceCategory, r.managerName, r.resource, r.fullPath)
	}

	finalURL := &url.URL{}
	if r.c.Base != nil {
		*finalURL = *r.c.Base
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

func (r *Request) newHTTPRequest(ctx context.Context) (*http.Request, error) {
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
		body = nil
	}
	url := r.URL().String()
	req, err := http.NewRequestWithContext(ctx, r.verb, url, body)
	if err != nil {
		return nil, err
	}
	req.Header = r.headers
	return req, nil
}
