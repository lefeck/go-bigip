package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func iControlPath(parts []string) string {
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
func main() {
	parts := []string{"/Core/1.1.1.1?ver=17.0.0.1", "/Common/1.2.3.4"}
	p := iControlPath(parts)
	fmt.Println(p)
}


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
 // 我想把data传到body，如何操作？
func (r *Request) newHTTPRequest(ctx context.Context, data interface{}) (*http.Request, error) {
	var body io.Reader
	switch data.(type){
	case r.body != nil && r.bodyBytes != nil:
		return nil, fmt.Errorf("cannot set both body and bodyBytes")
	case r.body != nil:
		body = r.body
	case r.bodyBytes != nil:
		// Create a new reader specifically for this request.
		// Giving each request a dedicated reader allows retries to avoid races resetting the request body.
		body = bytes.NewReader(r.bodyBytes)

	}
	url := r.URL().String()
	req, err := http.NewRequest(r.verb, url, body)
	if err != nil {
		return nil, err
	}
	req.Header = r.headers
	return req, nil
}



func (r *Request) Request(ctx context.Context,data interface{}) (*http.Response, error) {
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

	req, err := r.newHTTPRequest(ctx,data)
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
