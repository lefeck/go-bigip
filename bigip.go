package bigip

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip/rest"
	"log"
	"net/http"
	"net/url"
)

type BigIP struct {
	RestClient *rest.RESTClient
}

func NewSession(host, username, password string) (*BigIP, error) {
	config := &rest.Config{
		Host:     host,
		Username: username,
		Password: password,
		ContentConfig: rest.ContentConfig{
			ContentType: "application/json",
		},
	}

	restClient, err := restClientFor(config)
	if err != nil {
		return nil, err
	}

	return &BigIP{
		RestClient: restClient,
	}, nil
}

func NewToken(host, username, password, loginProviderName string) (*BigIP, error) {
	auth := authProvider{}

	token, err := auth.generateToken(host, username, password, loginProviderName)
	if err != nil {
		log.Fatalf("generation token failed %v\n", err)
	}
	config := &rest.Config{
		Host: host,
		ContentConfig: rest.ContentConfig{
			ContentType: "application/json",
		},
		BearerToken: token,
		//Timeout:     10 * time.Second,
	}

	restClient, err := restClientFor(config)
	if err != nil {
		return nil, err
	}

	return &BigIP{
		RestClient: restClient,
	}, nil
}

func restClientFor(config *rest.Config) (*rest.RESTClient, error) {
	// Setup HTTP client with authentication and custom transport
	httpClient, err := rest.HTTPClientFor(config)
	if err != nil {
		return nil, err
	}
	return rest.RESTClientForConfigAndClient(config, httpClient)
}

type authProvider struct {
	UserName      string `json:"username"`
	Password      string `json:"password"`
	LoginProvider string `json:"loginProviderName"`
	client        *http.Client
	transport     *http.Transport
}

func (a *authProvider) newHTTPRequest(host, username, password, loginProviderName string) (*http.Request, error) {
	t := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	a.transport = t
	a.client = &http.Client{
		Transport: a.transport,
	}

	auth := authProvider{
		UserName:      username,
		Password:      password,
		LoginProvider: loginProviderName,
	}
	data, err := json.Marshal(auth)
	if err != nil {
		return nil, fmt.Errorf("failed to create token client (cannot marshal user credentials): %v", err)
	}

	rawURL, basePath, _ := rest.DefaultServerURL(host, "/mgmt/shared/authn/login")
	fullURL := rawURL.String() + basePath
	parsedURL, err := url.Parse(fullURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %v", err)
	}
	fmt.Println(parsedURL)
	req, err := http.NewRequest(http.MethodPost, parsedURL.String(), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	return req, nil
}

type Login struct {
	Username          string `json:"username"`
	LoginProviderName string `json:"loginProviderName"`
	Token             struct {
		Token            string        `json:"token"`
		Name             string        `json:"name"`
		UserName         string        `json:"userName"`
		AuthProviderName string        `json:"authProviderName"`
		GroupReferences  []interface{} `json:"groupReferences"`
		Timeout          int           `json:"timeout"`
		StartTime        string        `json:"startTime"`
		Address          string        `json:"address"`
		Partition        string        `json:"partition"`
		Generation       int           `json:"generation"`
		LastUpdateMicros int64         `json:"lastUpdateMicros"`
		ExpirationMicros int64         `json:"expirationMicros"`
		Kind             string        `json:"kind"`
		SelfLink         string        `json:"selfLink"`
	} `json:"token"`
	Generation       int `json:"generation"`
	LastUpdateMicros int `json:"lastUpdateMicros"`
}

func (a *authProvider) generateToken(host, username, password, loginProviderName string) (string, error) {
	req, err := a.newHTTPRequest(host, username, password, loginProviderName)
	if err != nil {
		return "", err
	}
	resp, err := a.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	login := Login{}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&login); err != nil {
		return "", fmt.Errorf("failed to create token client (cannot decode token): %v", err)
	}

	// Compate time at which the token will expire (minus a minute).
	//expireAt := token.StartTime.Add(time.Duration(token.Timeout-60) * time.Second)
	return login.Token.Name, nil
}
