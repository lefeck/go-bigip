package main

//
//import (
//	"bytes"
//	"crypto/tls"
//	"encoding/json"
//	"errors"
//	"fmt"
//	"github.com/lefeck/bigip"
//	"net/http"
//	"net/url"
//	"os"
//	"strings"
//	"time"
//)
//
//// ErrNoToken is the error returned when the Client does not have a token.
//var ErrNoToken = errors.New("no token")
//
//// DefaultTimeout defines the default timeout for HTTP clients.
//var DefaultTimeout = 5 * time.Second
//
//// F5TimeLayout defines the layout to use for decoding dates returned by the
//// F5 iControl REST API.
//const F5TimeLayout = "2006-01-02T15:04:05.999999999-0700"
//
//// F5Date wraps time.Time in order to override the time layout used during
//// JSON decoding.
//type F5Date struct {
//	time.Time
//}
//
//// UnmarshalJSON overrides time.Time JSON decoding to support F5 time parsing
//// layout.
//func (d *F5Date) UnmarshalJSON(b []byte) error {
//	rawdate := strings.Trim(string(b), `"`)
//	t, err := time.Parse(F5TimeLayout, rawdate)
//	if err != nil {
//		return err
//	}
//	d.Time = t
//	return nil
//}
//
//// An authFunc is function responsible for setting necessary headers to
//// perform authenticated requests.
//type authFunc func(req *http.Request) error
//
//// A Client manages communication with the F5 API.
//type BigIPs struct {
//	c         http.Client
//	baseURL   string
//	authType  authFunc
//	transport *http.Transport
//
//	username, password string
//
//	token          string
//	tokenExpiresAt time.Time
//}
//
//// NewBasicClient creates a new F5 client with HTTP Basic Authentication.
////
//// baseURL is the base URL of the F5 API server.
//func NewCredentials(baseURL, user, password string) (*bigip.BigIP, error) {
//	transport := &http.Transport{}
//	return &bigip.BigIP{
//		c:         http.Client{Transport: transport, Timeout: DefaultTimeout},
//		baseURL:   baseURL,
//		transport: transport,
//		username:  user,
//		password:  password,
//		authType: func(req *http.Request) error {
//			req.SetBasicAuth(user, password)
//			return nil
//		},
//	}, nil
//}
//
//// TokenClientConnection creates a new client with the given token.
//func NewTokenSession(baseURL, token string) (*bigip.BigIP, error) {
//	transport := &http.Transport{}
//	c := &bigip.BigIP{
//		c: http.Client{
//			Transport: transport,
//			Timeout:   DefaultTimeout,
//		}, baseURL: baseURL}
//	c.token = token
//	c.authType = func(req *http.Request) error {
//		req.Header.Add("X-F5-Auth-Token", c.token)
//		return nil
//	}
//
//	return c, nil
//}
//
//// CreateToken creates a new token with the given baseURL, user, password and loginProvName.
//func CreateToken(baseURL, user, password, loginProvName string) (string, time.Time, error) {
//	t := &http.Transport{}
//	c := &bigip.BigIP{c: http.Client{Transport: t, Timeout: DefaultTimeout}, baseURL: baseURL, transport: t}
//	c.transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
//	// Negociate token with a pair of username/password.
//	data, err := json.Marshal(map[string]string{"username": user, "password": password, "loginProviderName": loginProvName})
//	if err != nil {
//		return "", time.Time{}, fmt.Errorf("failed to create token client (cannot marshal user credentials): %v", err)
//	}
//
//	tokReq, err := http.NewRequest("POST", c.makeURL("/mgmt/shared/authn/login"), bytes.NewBuffer(data))
//	if err != nil {
//		return "", time.Time{}, fmt.Errorf("failed to create token client, (cannot create login request): %v", err)
//	}
//
//	tokReq.Header.Add("Content-Type", "application/json")
//
//	resp, err := c.c.Do(tokReq)
//	if err != nil {
//		return "", time.Time{}, fmt.Errorf("failed to create token client (token negociation failed): %v", err)
//	}
//	if resp.StatusCode >= 400 {
//		return "", time.Time{}, fmt.Errorf("failed to create token client (token negociation failed): http status %s", resp.Status)
//	}
//	defer resp.Body.Close()
//
//	tok := struct {
//		Token struct {
//			Token     string `json:"token"`
//			StartTime F5Date `json:"startTime"`
//			Timeout   int    `json:"timeout"`
//		} `json:"token"`
//	}{}
//	dec := json.NewDecoder(resp.Body)
//	if err := dec.Decode(&tok); err != nil {
//		return "", time.Time{}, fmt.Errorf("failed to create token client (cannot decode token): %v", err)
//	}
//
//	// Compate time at which the token will expire (minus a minute).
//	expireAt := tok.Token.StartTime.Add(time.Duration(tok.Token.Timeout-60) * time.Second)
//
//	return tok.Token.Token, expireAt, nil
//}
//
//// NewTokenClient creates a new F5 client with token based authentication.
////
//// baseURL is the base URL of the F5 API server.
//func NewTokenClient(baseURL, user, password, loginProvName string) (*bigip.BigIP, error) {
//	t := &http.Transport{}
//	c := bigip.BigIP{
//		c:         http.Client{Transport: t, Timeout: DefaultTimeout},
//		baseURL:   baseURL,
//		transport: t,
//		username:  user,
//		password:  password,
//	}
//
//	// Create auth function for token based authentication.
//	c.authType = func(req *http.Request) (err error) {
//		if c.token == "" || time.Now().After(c.tokenExpiresAt) {
//			c.token, c.tokenExpiresAt, err = CreateToken(baseURL, user, password, loginProvName)
//			if err != nil {
//				return
//			}
//		}
//		req.Header.Set("X-F5-Auth-Token", c.token)
//		return
//	}
//
//	return &c, nil
//}
//
//// DisableCertCheck disables certificate verification, meaning that insecure
//// certificate will not cause any error.
//func (b *bigip.BigIP) DisableCertCheck() {
//	b.transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
//}
//
//// CheckAuth verifies that the credentials provided at the client initialization
//// are correct.
//func (c *bigip.BigIP) CheckAuth() error {
//	if _, err := c.SendRequest("GET", "/mgmt/tm/ltm/available", nil); err != nil {
//		return fmt.Errorf("authentication verification failed: %v", err)
//	}
//	return nil
//}
//
//// RevokeToken revokes the current token. If the Client has not been initialized
//// with NewTokenClient, ErrNoToken is returned.
//func (c *bigip.BigIP) RevokeToken() error {
//	if c.token == "" {
//		return ErrNoToken
//	}
//
//	resp, err := c.SendRequest("DELETE", "/mgmt/shared/authz/tokens/"+c.token, nil)
//	if err != nil {
//		return errors.New("token revocation request failed: " + err.Error())
//	}
//	defer resp.Body.Close()
//
//	var respData struct {
//		Token string `json:"token"`
//	}
//	dec := json.NewDecoder(resp.Body)
//	if err := dec.Decode(&respData); err != nil {
//		return errors.New("cannot decode token revocation response: " + err.Error())
//	}
//	if respData.Token != c.token {
//		return errors.New("invalid token revocation response")
//	}
//
//	c.token = ""
//
//	return nil
//}
//
//// SetTimeout sets the HTTP timeout for the underlying HTTP client.
//func (c *bigip.BigIP) SetTimeout(timeout time.Duration) {
//	c.c.Timeout = timeout
//}
//
//// SetHTTPClient sets the underlying HTTP used to make requests.
//func (c *bigip.BigIP) SetHTTPClient(client http.Client) {
//	c.c = client
//}
//
//// UseProxy configures a proxy to use for outbound connections
//func (c *bigip.BigIP) UseProxy(proxy string) error {
//	proxyURL, err := url.Parse(proxy)
//	if err != nil {
//		return err
//	}
//	c.transport.Proxy = http.ProxyURL(proxyURL)
//	return nil
//}
//
//// UseSystemProxy configures the client to use the system proxy
//func (c *bigip.BigIP) UseSystemProxy() error {
//	proxy := os.Getenv("HTTP_PROXY")
//	if proxy != "" {
//		proxyURL, err := url.Parse(proxy)
//		if err != nil {
//			return err
//		}
//		c.transport.Proxy = http.ProxyURL(proxyURL)
//	}
//	return nil
//}
//
//// MakeRequest creates a request with headers appropriately set to make
//// authenticated requests. This method must be called for every new request.
//func (c *bigip.BigIP) MakeRequest(method, restPath string, data interface{}) (*http.Request, error) {
//	var (
//		req *http.Request
//		err error
//	)
//	if data != nil {
//		switch v := data.(type) {
//		case string:
//			req, err = http.NewRequest(method, c.makeURL(restPath), strings.NewReader(v))
//		default:
//			bf := bytes.NewBuffer([]byte{})
//			jsonEncoder := json.NewEncoder(bf)
//			jsonEncoder.SetEscapeHTML(false)
//			err := jsonEncoder.Encode(data)
//			if err != nil {
//				return nil, fmt.Errorf("failed to marshal data into json: %v", err)
//			}
//			req, err = http.NewRequest(method, c.makeURL(restPath), bf)
//		}
//	} else {
//		req, err = http.NewRequest(method, c.makeURL(restPath), nil)
//	}
//	if err != nil {
//		return nil, fmt.Errorf("failed to create F5 authenticated request: %v", err)
//	}
//	req.Header.Add("Accept", "application/json")
//	if err := c.authType(req); err != nil {
//		return nil, err
//	}
//	return req, nil
//}
//
//// Do sends an HTTP request and returns an HTTP response. It is just a wrapper
//// arround http.Client Do method.
////
//// Callers should close resp.Body when done reading from it.
////
//// See http package documentation for more information:
////
////	https://golang.org/pkg/net/http/#Client.Do
//func (c *bigip.BigIP) Do(req *http.Request) (*http.Response, error) {
//	resp, err := c.c.Do(req)
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
//
//// SendRequest is a shortcut for MakeRequest() + Do() + ReadError().
//func (c *bigip.BigIP) SendRequest(method, restPath string, data interface{}) (*http.Response, error) {
//	req, err := c.MakeRequest(method, restPath, data)
//	if err != nil {
//		return nil, err
//	}
//	resp, err := c.Do(req)
//	if err != nil {
//		return nil, err
//	}
//	if err := c.ReadError(resp); err != nil {
//		resp.Body.Close()
//		return nil, err
//	}
//	return resp, nil
//}
//
//// ReadQuery performs a GET query and unmarshal the response (from JSON) into outputData.
////
//// outputData must be a pointer.
//func (c *bigip.BigIP) ReadQuery(restPath string, outputData interface{}) error {
//	resp, err := c.SendRequest("GET", restPath, nil)
//	if err != nil {
//		return err
//	}
//	defer resp.Body.Close()
//	dec := json.NewDecoder(resp.Body)
//	if err := dec.Decode(outputData); err != nil {
//		return err
//	}
//	return nil
//}
//
//// ModQuery performs a modification query such as POST, PUT or DELETE.
//func (c *bigip.BigIP) ModQuery(method, restPath string, inputData interface{}) error {
//	if method != "POST" && method != "PUT" && method != "DELETE" && method != "PATCH" {
//		return errors.New("invalid method " + method)
//	}
//	resp, err := c.SendRequest(method, restPath, inputData)
//	if err != nil {
//		return err
//	}
//	resp.Body.Close()
//	return nil
//}
//
//// ReadError checks if a HTTP response contains an error and returns it.
//func (c *bigip.BigIP) ReadError(resp *http.Response) error {
//	if resp.StatusCode >= 400 {
//		if contentType := resp.Header.Get("Content-Type"); !strings.Contains(contentType, "application/json") {
//			return errors.New("http response error: " + resp.Status)
//		}
//		errResp, err := bigip.NewRequestError(resp.Body)
//		if err != nil {
//			return errors.New("cannot read error message from response body: " + err.Error())
//		}
//		return errResp
//	}
//	return nil
//}
//
//// makeURL creates an URL from the client base URL and a given REST path. What
//// this function actually does is to concatenate the base URL and the REST path
//// by handling trailing slashes.
//func (c *bigip.BigIP) makeURL(restPath string) string {
//	return strings.TrimSuffix(c.baseURL, "/") + "/" + strings.TrimPrefix(restPath, "/")
//}
//
//func (c *bigip.BigIP) clone() *bigip.BigIP {
//	return &bigip.BigIP{
//		c:        c.c,
//		baseURL:  c.baseURL,
//		authType: c.authType,
//	}
//}
