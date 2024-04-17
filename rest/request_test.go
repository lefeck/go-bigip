package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// TestRequest verifies that http requests are properly made to specified path and method.
func TestRequest(t *testing.T) {
	testRequest(t, "GET", "/test", "name: jok", "/api/v1/test")
	testRequest(t, "PUT", "/test/put", "name: jok", "/api/v1/test/put")
	testRequest(t, "POST", "/test/post", "name: jok", "/api/v1/test/post")
	testRequest(t, "DELETE", "/test/delete", "name: jok", "/api/v1/test/delete")
}

func testRequest(t *testing.T, method, path, body string, expectedURL string) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			t.Errorf("Expected method %s, got %s", method, r.Method)
		}

		if r.URL.String() != expectedURL {
			t.Errorf("Expected URL %s, got %s", expectedURL, r.URL.String())
		}

		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"fake_key": "fake_value"}`)
	}

	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	baseURL, _ := url.Parse(ts.URL)
	content := ClientContentConfig{}
	client := ts.Client()
	req := NewRequestWithClient(baseURL, "/api/v1", content, client)

	req.Verb(method).Prefix(path)

	resp, err := req.Request(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("Error closing response body: %v", err)
		}
	}()

	var responseBody map[string]string
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&responseBody)
	if err != nil {
		t.Fatal(err)
	}

	if responseBody["fake_key"] != "fake_value" {
		t.Errorf("Expected fake_key value to be fake_value, got: %v", responseBody["fake_key"])
	}
}

func TestRequestWithBody(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"fake_key": "fake_value"}`)
	}

	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	baseURL, _ := url.Parse(ts.URL)
	content := ClientContentConfig{}
	client := ts.Client()
	req := NewRequestWithClient(baseURL, "/api/v1", content, client)

	req.Verb("POST").Prefix("/test")

	data := map[string]string{
		"key": "value",
	}

	req.body = bytes.NewReader([]byte(`{"key": "value"}`))
	req.bodyBytes, _ = json.Marshal(data)

	resp, err := req.Request(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}

	var responseBody map[string]string
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&responseBody)
	if err != nil {
		t.Fatal(err)
	}

	if responseBody["fake_key"] != "fake_value" {
		t.Errorf("Expected fake_key value to be fake_value, got: %v", responseBody["fake_key"])
	}
}
