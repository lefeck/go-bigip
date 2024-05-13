package transport

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
)

func New(config *Config) (http.RoundTripper, error) {
	if config.Transport != nil {
		return nil, fmt.Errorf("using a custom transport with TLS certificate options or the insecure flag is not allowed")
	}

	// clone a new http.Transport connect, and settting InsecureSkipVerify is true.

	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// Use customTransport instead of http.DefaultTransport
	return HTTPWrappersFor(config, customTransport)
}

// WrapperFunc wraps an http.RoundTripper when a new transport
// is created for a client, allowing per connection behavior
// to be injected.
type WrapperFunc func(rt http.RoundTripper) http.RoundTripper

// Wrappers accepts any number of wrappers and returns a wrapper
// function that is the equivalent of calling each of them in order. Nil
// values are ignored, which makes this function convenient for incrementally
// wrapping a function.
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

// ContextCanceller prevents new requests after the provided context is finished.
// err is returned when the context is closed, allowing the caller to provide a context
// appropriate error.
func ContextCanceller(ctx context.Context, err error) WrapperFunc {
	return func(rt http.RoundTripper) http.RoundTripper {
		return &contextCanceller{
			ctx: ctx,
			rt:  rt,
			err: err,
		}
	}
}

type contextCanceller struct {
	ctx context.Context
	rt  http.RoundTripper
	err error
}

func (b *contextCanceller) RoundTrip(req *http.Request) (*http.Response, error) {
	select {
	case <-b.ctx.Done():
		return nil, b.err
	default:
		return b.rt.RoundTrip(req)
	}
}

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
		fmt.Printf("Unable to cancel request for %T", rt)
	}
}
