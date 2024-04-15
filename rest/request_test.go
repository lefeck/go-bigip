package rest

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"

	"testing"
)

func TestNewRESTClient(t *testing.T) {
	baseUrl := "http://localhost:8080"
	baseAPIPath := "/api/v1"
	baseURL, _ := url.Parse(baseUrl)
	contentConfig := ClientContentConfig{}
	client := http.DefaultClient
	restClientInstance, _ := NewRESTClient(baseURL, baseAPIPath, contentConfig, client)
	if restClientInstance == nil {
		t.Errorf("Expected non-nil RESTClient instance, but got nil")
	}
}

func TestRequest(t *testing.T) {
	// 创建一个测试服务器
	testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/api/v1/resource/subresource" {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		writer.WriteHeader(http.StatusOK)

	}))
	defer testServer.Close()

	baseUrl := testServer.URL
	baseAPIPath := "/api/v1"
	baseURL, _ := url.Parse(baseUrl)
	contentConfig := ClientContentConfig{}
	client := http.DefaultClient

	restClient, _ := NewRESTClient(baseURL, baseAPIPath, contentConfig, client)

	response, err := restClient.Get().Resource("resource").SubResource("subresource").Request(context.Background(), nil)
	if err != nil {
		t.Errorf("Unexpected error in Request: %v", err)
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	expectedResponseBody := `{"message":"success","code":200}`
	if string(responseBody) != expectedResponseBody {
		t.Errorf("Expected response body: %s, but got: %s", expectedResponseBody, string(responseBody))
	}
}