package api

import (
	"fmt"
	"strings"

	"github.com/lolibrary/lolibrary/libraries/portforward"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

var endpoint = defaultEndpoint

const (
	defaultEndpoint    = "127.0.0.1:8080"
	defaultEndpointFmt = "127.0.0.1:%d"
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
func Run(f func()) {
	port, c := portforward.Enable()
	defer func() {
		c.Close()
		SetEndpoint(defaultEndpoint)
		typhon.Client = typhon.BareClient
	}()

	SetEndpoint(fmt.Sprintf(defaultEndpointFmt, port))
	typhon.Client = Client

	f()
}

// Client is a wrapped typhon client that will send all requests via the edge proxy instead.
func Client(req typhon.Request) typhon.Response {
	url := fmt.Sprintf("http://%s/%s/%s",
		Endpoint(),
		req.URL.Host,
		strings.TrimPrefix(req.URL.Path, "/"))

	req = typhon.NewRequest(req, req.Method, url, req.Body)
	slog.Trace(req, "Sending request via edge proxy: %v", req)

	return typhon.HttpService(typhon.RoundTripper)(req)
}
