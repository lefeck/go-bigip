package rest

import (
	"testing"
)

func TestDefaultServerURL(t *testing.T) {
	tests := []struct {
		name           string
		host           string
		apiPath        string
		defaultTLS     bool
		expectedURL    string
		expectedAPI    string
		expectingError bool
	}{
		{
			name:        "Valid non-TLS server",
			host:        "localhost:8080",
			apiPath:     "api/v1",
			defaultTLS:  false,
			expectedURL: "http://localhost:8080",
			expectedAPI: "/api/v1",
		},
		{
			name:        "Valid TLS server",
			host:        "https://localhost:8443",
			apiPath:     "api/v1",
			defaultTLS:  true,
			expectedURL: "https://localhost:8443",
			expectedAPI: "/api/v1",
		},
		{
			name:        "Missing host",
			host:        "",
			apiPath:     "api/v1",
			defaultTLS:  false,
			expectedURL: "",
			expectedAPI: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			hostURL, apiPath, err := DefaultServerURL(test.host, test.apiPath)

			if test.expectingError && err == nil {
				t.Error("Expected error, but got none")
			} else if !test.expectingError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if test.expectingError {
				return
			}

			if hostURL.String() != test.expectedURL {
				t.Errorf("Expected URL %s, got %s", test.expectedURL, hostURL.String())
			}

			if apiPath != test.expectedAPI {
				t.Errorf("Expected API path %s, got %s", test.expectedAPI, apiPath)
			}
		})
	}
}
