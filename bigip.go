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
	"time"
)

const TimeFormat = "2006-01-02T15:04:05.000-0700"

// DefaultTimeout defines the default timeout for HTTP clients.
var DefaultTimeout time.Duration = 60

// BigIP struct contains a pointer to the RESTClient
type BigIP struct {
	RestClient *rest.RESTClient
}

// NewSession creates a new BigIP structure initialized with a username and password.
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

// NewToken retrieves a login token from a new BigIP structure with token authentication
func NewToken(host, username, password, loginProviderName string, options ...Option) (*BigIP, error) {
	auth := newAuthPayload(host, username, password, loginProviderName, options...)
	token, _, err := auth.generateToken()
	if err != nil {
		log.Fatalf("generation token failed %v\n", err)
	}
	config := &rest.Config{
		Host: host,
		ContentConfig: rest.ContentConfig{
			ContentType: "application/json",
		},
		BearerToken: token,
	}

	restClient, err := restClientFor(config)
	if err != nil {
		return nil, err
	}

	return &BigIP{
		RestClient: restClient,
	}, nil
}

// restClientFor is a helper function that creates a new REST client for the given config.
func restClientFor(config *rest.Config) (*rest.RESTClient, error) {
	httpClient, err := rest.HTTPClientFor(config)
	if err != nil {
		return nil, err
	}
	return rest.RESTClientForConfigAndClient(config, httpClient)
}

// Option is a custom type that handles options
type Option func(payload *authPayload)

// authPayload contains authentication related information such as hostname, username, password, etc.
type authPayload struct {
	Host              string        `json:"host"`
	UserName          string        `json:"username"`
	Password          string        `json:"password"`
	LoginProviderName string        `json:"loginProviderName"`
	Timeout           time.Duration `json:"timeout"`
	token             string
	tokenExpiresAt    time.Time
	Client            *http.Client `json:"client"`
}

// WithTimeout is an Option type function used for setting the timeout
func WithTimeout(timeout time.Duration) Option {
	return func(auth *authPayload) {
		auth.Timeout = timeout
	}
}

// newAuthPayload creates a new authPayload based on the given hostname, username, password, and loginProviderName among other things.
func newAuthPayload(host, username, password, loginProviderName string, options ...Option) *authPayload {
	auth := &authPayload{
		Host:              host,
		UserName:          username,
		Password:          password,
		LoginProviderName: loginProviderName,
		Timeout:           DefaultTimeout,
		Client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
	}

	// Apply any incoming options
	for _, opt := range options {
		opt(auth)
	}

	return auth
}

// newHTTPRequest is a helper function for creating new HTTP requests.
func (auth *authPayload) newHTTPRequest() (*http.Request, error) {
	authz := authPayload{
		Host:              auth.Host,
		UserName:          auth.UserName,
		Password:          auth.Password,
		LoginProviderName: auth.LoginProviderName,
		Timeout:           auth.Timeout,
	}
	data, err := json.Marshal(authz)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON data: %v", err)
	}

	rawURL, basePath, _ := rest.DefaultServerURL(authz.Host, "/mgmt/shared/authn/login")
	fullURL := rawURL.String() + basePath
	parsedURL, err := url.Parse(fullURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, parsedURL.String(), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

// authToken  contains information related to the token
type authToken struct {
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

// generateToken is used to create a new authentication token
func (auth *authPayload) generateToken() (string, time.Time, error) {
	req, err := auth.newHTTPRequest()
	if err != nil {
		return "", time.Time{}, err
	}
	resp, err := auth.Client.Do(req)
	if err != nil {
		return "", time.Time{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", time.Time{}, fmt.Errorf("http response status code error: %s", resp.Status)
	}

	token := authToken{}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&token); err != nil {
		return "", time.Time{}, fmt.Errorf("failed to create token: %v", err)
	}

	startTime, err := time.Parse(TimeFormat, token.Token.StartTime)

	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to parse token start time: %v", err)
	}

	expiresAt := startTime.Add(time.Duration(token.Token.Timeout-token.Token.Timeout+int(auth.Timeout.Seconds())) * time.Second)

	auth.token = token.Token.Token
	auth.tokenExpiresAt = expiresAt

	return token.Token.Token, expiresAt, nil
}
