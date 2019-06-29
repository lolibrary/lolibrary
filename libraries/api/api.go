package api

import (
	"fmt"

	"github.com/lolibrary/lolibrary/libraries/portforward"
	"github.com/monzo/typhon"
)

var endpoint = defaultEndpoint

const (
	defaultEndpoint    = "localhost:8080"
	defaultEndpointFmt = "localhost:%d"
)

// SetEndpoint sets the current endpoint for the edge proxy.
func SetEndpoint(url string) {
	endpoint = url
}

// Endpoint fetches the current edge proxy endpoint.
func Endpoint() string {
	return endpoint
}

// Run opens a port-forward to the edge proxy and runs f with this open.
func Run(f func() error) error {
	port, c := portforward.Enable()
	defer func() {
		c.Close()
		SetEndpoint(defaultEndpoint)
		typhon.Client = typhon.BareClient
	}()

	SetEndpoint(fmt.Sprintf(defaultEndpointFmt, port))
	typhon.Client = Client

	return f()
}

// Client is a wrapped typhon client
func Client(req typhon.Request) typhon.Response {
	host := req.URL.Host

	req.URL.Host = Endpoint()
	req.URL.Path = fmt.Sprintf("/%s%s", host, req.URL.Path)

	return typhon.HttpService(typhon.RoundTripper)(req)
}
