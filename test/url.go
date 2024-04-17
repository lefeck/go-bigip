package main

import (
	"bytes"
	"fmt"
	"io"
	"k8s.io/klog/v2"
	"net/http"
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

func HTTPWrappersForConfig(config *Config, rt http.RoundTripper) (http.RoundTripper, error) {
	if config.WrapTransport != nil {
		rt = config.WrapTransport(rt)
	}

	// Set authentication wrappers
	switch {
	case config.HasBasicAuth():
		rt = NewBasicAuthRoundTripper(config.Username, config.Password, rt)
	}

	return rt, nil
}

// 我想通过这段代码验证 我写的没问题， 帮我实现？
type basicAuthRoundTripper struct {
	username string
	password string `datapolicy:"password"`
	rt       http.RoundTripper
}

type RoundTripperWrapper interface {
	http.RoundTripper
	WrappedRoundTripper() http.RoundTripper
}

var _ RoundTripperWrapper = &basicAuthRoundTripper{}

// NewBasicAuthRoundTripper will apply a BASIC auth authorization header to a
// request unless it has already been set.
func NewBasicAuthRoundTripper(username, password string, rt http.RoundTripper) http.RoundTripper {
	return &basicAuthRoundTripper{username, password, rt}
}

func (rt *basicAuthRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if len(req.Header.Get("Authorization")) != 0 {
		return rt.rt.RoundTrip(req)
	}
	req = CloneRequest(req)
	req.SetBasicAuth(rt.username, rt.password)
	return rt.rt.RoundTrip(req)
}

func (rt *basicAuthRoundTripper) CancelRequest(req *http.Request) {
	tryCancelRequest(rt.WrappedRoundTripper(), req)
}

func (rt *basicAuthRoundTripper) WrappedRoundTripper() http.RoundTripper { return rt.rt }

func tryCancelRequest(rt http.RoundTripper, req *http.Request) {
	type canceler interface {
		CancelRequest(*http.Request)
	}
	switch rt := rt.(type) {
	case canceler:
		rt.CancelRequest(req)
	case RoundTripperWrapper:
		tryCancelRequest(rt.WrappedRoundTripper(), req)
	default:
		klog.Warningf("Unable to cancel request for %T", rt)
	}
}

// CloneRequest creates a shallow copy of the request along with a deep copy of the Headers.
func CloneRequest(req *http.Request) *http.Request {
	r := new(http.Request)

	// shallow clone
	*r = *req

	// deep copy headers
	r.Header = CloneHeader(req.Header)

	return r
}

// CloneHeader creates a deep copy of an http.Header.
func CloneHeader(in http.Header) http.Header {
	out := make(http.Header, len(in))
	for key, values := range in {
		newValues := make([]string, len(values))
		copy(newValues, values)
		out[key] = newValues
	}
	return out
}

type Config struct {
	// Username and password for basic authentication
	Username string
	Password string

	// Bearer token for authentication
	BearerToken string

	// Path to a file containing a BearerToken.
	// If set, the contents are periodically read.
	// The last successfully read value takes precedence over BearerToken.
	BearerTokenFile string

	// WrapTransport for most client level operations.
	Transport http.RoundTripper

	WrapTransport WrapperFunc
}

type WrapperFunc func(rt http.RoundTripper) http.RoundTripper

func (c *Config) Wrap(fn WrapperFunc) {
	c.WrapTransport = Wrappers(c.WrapTransport, fn)
}

func Wrappers(fns ...WrapperFunc) WrapperFunc {
	if len(fns) == 0 {
		return nil
	}
	// optimize the common case of wrapping a possibly nil transport wrapper
	// with an additional wrapper
	if len(fns) == 2 && fns[0] == nil {
		return fns[1]
	}
	return func(rt http.RoundTripper) http.RoundTripper {
		base := rt
		for _, fn := range fns {
			if fn != nil {
				base = fn(base)
			}
		}
		return base
	}
}

func (c *Config) HasBasicAuth() bool {
	return len(c.Username) != 0
}

// 为啥这种方式不可以呢？
func New(config *Config) (http.RoundTripper, error) {
	if config.Transport != nil {
		return nil, fmt.Errorf("using a custom transport with TLS certificate options or the insecure flag is not allowed")
	}

	//var rt http.RoundTripper
	return HTTPWrappersForConfig(config, http.DefaultTransport)
}

// 这种方式可以呢？
func HTTPWrappersForConfigs(config *Config, rt http.RoundTripper) (http.RoundTripper, error) {
	return HTTPWrappersForConfig(config, rt)
}

func main() {
	// 创建一个简单的HTTP服务器
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, this is a test for basic authentication.")
	})
	go func() {
		http.ListenAndServe(":8089", nil)
	}()

	time.Sleep(time.Second)

	// 创建配置
	config := &Config{
		Username: "admin",
		Password: "admin123",
	}

	// 通过配置创建HTTP 客户端
	transport, err := New(config)
	//transport, err := HTTPWrappersForConfigs(config, http.DefaultTransport)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	client := &http.Client{Transport: transport}

	// 创建HTTP请求并发送
	req, err := http.NewRequest("GET", "http://localhost:8089", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Before sending the request, headers are: ", req.Header)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	// 检查收到的响应的请求头
	fmt.Println("After sending the request, headers of the received response are: ", resp.Request)

	// 我打印发现header是空的，按道理应该有username和password才对呀？
	fmt.Println(req.Header)

	// 读取响应并打印
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response:", string(body))
}
