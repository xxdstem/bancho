package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

// NewProxy takes target host and creates a reverse proxy
func NewProxy(targetHost string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	return httputil.NewSingleHostReverseProxy(url), nil
}

// ProxyRequestHandler handles the http request using proxy
func ProxyRequestHandler(proxy *httputil.ReverseProxy, handler *http.Handler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}

func startProxy(handler http.Handler) error {
	// initialize a reverse proxy and pass the actual backend server url here

	// proxy, err := NewProxy("https://ce.ppy.sh")
	// if err != nil {
	// 	panic(err)
	// }

	return http.ListenAndServeTLS(":443", "cuttingedge.crt", "cuttingedge.key", handler)
}
