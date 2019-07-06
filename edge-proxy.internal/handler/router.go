package handler

import (
	"fmt"
	"net"
	"regexp"
	"strings"

	"github.com/monzo/slog"
	"github.com/monzo/terrors"
	"github.com/monzo/typhon"
)

var (
	reAPI          = regexp.MustCompile(`^/service\.api\.([a-z\-]+)/([a-z\-]*)$`)
	reService      = regexp.MustCompile(`^/service\.([a-z\-]+)/([a-z\-]*)$`)
	reShortService = regexp.MustCompile(`^/s-([a-z\-]+)/([a-z\-]*)$`)

	// requestFormat is the format we'll send to upstream services.
	requestFormat = "http://%s/%s"
)

// Service returns a typhon service.
func Service() typhon.Service {
	return Proxy
}

// Proxy takes an incoming request and proxies it into the cluster.
func Proxy(req typhon.Request) typhon.Response {
	path := req.URL.Path

	slog.Trace(req, "Incoming request to: %v", path)

	switch {
	case strings.HasPrefix(path, "/s-"):
		return handleShortService(req)
	case strings.HasPrefix(path, "/service.api."):
		return handleAPI(req)
	case strings.HasPrefix(path, "/service."):
		return handleService(req)
	}

	return typhon.Response{Error: terrors.NotFound("missing_endpoint", "Don't know how to query that service!", nil)}
}

// handle takes a request and sends it to an internal service.
func handle(req typhon.Request, service, path string) typhon.Response {
	url := fmt.Sprintf(requestFormat, service, path)

	slog.Trace(req, "Handling parsed URL: %v", url)

	if _, err := net.Dial("tcp", fmt.Sprintf("%s:80", service)); err != nil {
		slog.Error(req, "Unable to connect to %s: %v", service, err)
		return typhon.Response{Error: terrors.NotFound("service", fmt.Sprintf("Unable to connect to %v", service), nil)}
	}

	req.Host = service
	req.URL.Scheme = "http"
	req.URL.Path = "/" + strings.TrimPrefix(path, "/")
	req.URL.Host = service

	return req.Send().Response()
}

// handleAPI takes a request and sends it to an internal API endpoint
func handleAPI(req typhon.Request) typhon.Response {
	parts := reAPI.FindStringSubmatch(req.URL.Path)
	if len(parts) != 3 {
		return typhon.Response{Error: terrors.NotFound("bad_endpoint", "Unable to determine API endpoint.", nil)}
	}

	return handle(req, "s-api-"+parts[1], parts[2])
}

// handleService takes a request and sends it to an internal service.
func handleService(req typhon.Request) typhon.Response {
	parts := reService.FindStringSubmatch(req.URL.Path)
	if len(parts) != 3 {
		return typhon.Response{Error: terrors.NotFound("bad_endpoint", "Unable to determine service endpoint.", nil)}
	}

	return handle(req, "s-"+parts[1], parts[2])
}

// handleShortService lets you send requests directly to the service name of a service.
func handleShortService(req typhon.Request) typhon.Response {
	parts := reShortService.FindStringSubmatch(req.URL.Path)
	if len(parts) != 3 {
		return typhon.Response{Error: terrors.NotFound("bad_endpoint", "Unable to determine short service endpoint.", nil)}
	}

	return handle(req, "s-"+parts[1], parts[2])
}
