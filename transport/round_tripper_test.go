package transport

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBasicAuthRoundTripper(t *testing.T) {
	username := "admin"
	password := "admin23423"
	expectedAuthHeader := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(username+":"+password)))

	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader != expectedAuthHeader {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	client := &http.Client{
		Transport: NewBasicAuthRoundTripper(username, password, ts.Client().Transport),
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error performing request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestBearerAuthRoundTripper(t *testing.T) {
	bearerToken := "123josd235l0o2lf;235rj"
	expectedAuthHeader := fmt.Sprintf("Bearer %s", bearerToken)

	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader != expectedAuthHeader {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	client := &http.Client{
		Transport: NewTokenAuthRoundTripper(bearerToken, ts.Client().Transport),
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error performing request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
