package rest

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"

	"testing"
)

func TestVerb(t *testing.T) {
	req := Request{}
	req.Verb("GET")
	if req.verb != "GET" {
		t.Errorf("Expected verb GET, got %s", req.verb)
	}
}

func TestPrefix(t *testing.T) {
	req := &Request{}
	req.pathPrefix = "/api"
	req.Prefix("v1")
	expectedPath := "/api/v1"
	if req.pathPrefix != expectedPath {
		t.Errorf("Expected path prefix %s, got %s", expectedPath, req.pathPrefix)
	}
}

func TestSuffix(t *testing.T) {
	req := &Request{}
	req.subpath = "/api"
	req.Suffix("test")
	expectedSubpath := "/api/test"
	if req.subpath != expectedSubpath {
		t.Errorf("Expected subpath prefix %s, got %s", expectedSubpath, req.subpath)
	}
}

func TestSubResource(t *testing.T) {
	resource := "testresource"
	req := &Request{}
	req.SubResource(resource)

	if req.subResource != resource {
		t.Errorf("Expected subResource %s, got %s", resource, req.subResource)
	}
}

func TestName(t *testing.T) {
	resourceName := "test"
	req := &Request{}
	req.ManagerName(resourceName)

	if req.managerName != resourceName {
		t.Errorf("Expected resourceName %s, got %s", resourceName, req.managerName)
	}
}

func TestAbsPath(t *testing.T) {
	baseURL, _ := url.Parse("http://localhost:8080/")
	restClient := &RESTClient{
		Base:        baseURL,
		Client:      http.DefaultClient,
		content:     ClientContent{},
		baseAPIPath: "/",
	}

	tests := []struct {
		segments []string
		expected string
	}{
		{[]string{"test"}, "/test"},
		{[]string{"test", "sub"}, "/test/sub"},
		{[]string{"test/"}, "/test/"},
		{[]string{"test", "sub/"}, "/test/sub/"},
	}

	for _, tc := range tests {
		req := NewRequest(restClient)
		req.AbsPath(tc.segments...)
		if req.pathPrefix != tc.expected {
			t.Errorf("Expected path prefix %q, got %q for segments %v", tc.expected, req.pathPrefix, tc.segments)
		}
	}
}

func TestSetParams(t *testing.T) {
	req := &Request{}
	req.params = nil
	req.setParams("testKey", "testValue")

	if req.params["testKey"][0] != "testValue" {
		t.Errorf("Expected param value is testValue, got %s", req.params["testKey"][0])
	}
}
func TestBody(t *testing.T) {
	testByte := []byte("test byte")
	testFile := "test.txt"

	// create a test file
	err := os.WriteFile(testFile, []byte("test"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(testFile)

	tests := []struct {
		input       interface{}
		expectedErr error
	}{
		{"test.txt", nil},
		{testByte, nil},
		{bytes.NewReader(testByte), nil},
		{123, fmt.Errorf("unknown type used for body: %+v", 123)},
	}

	for _, tc := range tests {
		req := &Request{}
		req.Body(tc.input)

		if req.err != nil && tc.expectedErr != nil {
			if req.err.Error() != tc.expectedErr.Error() {
				t.Errorf("Unexpected error, expected %v, got %v", tc.expectedErr, req.err)
			}
		} else if req.err != tc.expectedErr {
			t.Errorf("Unexpected error, expected %v, got %v", tc.expectedErr, req.err)
		}
	}
}

func TestDoRaw(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}))
	defer server.Close()

	baseURL, _ := url.Parse(server.URL)
	req := NewRequestWithClient(baseURL, "/test", ClientContent{}, http.DefaultClient).Verb("GET")
	ctx := context.Background()
	resp, err := req.DoRaw(ctx)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedResponse := "Hello, World!"
	if string(resp) != expectedResponse {
		t.Fatalf("Expected response %q, got %q", expectedResponse, resp)
	}
}
