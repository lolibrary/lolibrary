package handler

import (
	"fmt"
	"regexp"

	"github.com/monzo/slog"
	"github.com/monzo/terrors"
	"github.com/monzo/typhon"
)

var (
	reAPI     = regexp.MustCompile(`^/service.api.([a-z\-]+)/([a-z\-]*)$`)
	reService = regexp.MustCompile(`^/service.([a-z\-]+)/([a-z\-]*)$`)

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
	case reAPI.MatchString(path):
		return handleAPI(req)
	case reService.MatchString(path):
		return handleService(req)
	}

	return typhon.Response{Error: terrors.NotFound("missing_endpoint", "Don't know how to query that service!", nil)}
}

// handle takes a request and sends it to an internal service.
func handle(req typhon.Request, service, path string) typhon.Response {
	url := fmt.Sprintf(requestFormat, service, path)

	// TODO: validate that service exists first via a TCP dial and return a 404, service %s not found.

	slog.Trace(req, "Sending request to: %v", url)

	return typhon.NewRequest(req, req.Method, url, req.Body).Send().Response()
}

// handleAPI takes a request and sends it to an internal API endpoint
func handleAPI(req typhon.Request) typhon.Response {
	parts := reAPI.FindStringSubmatch(req.URL.Path)
	if len(parts) != 3 {
		return typhon.Response{Error: terrors.NotFound("bad_endpoint", "Unable to determine endpoint.", nil)}
	}

	return handle(req, "s-api-"+parts[1], parts[2])
}

// handleService takes a request and sends it to an internal service.
func handleService(req typhon.Request) typhon.Response {
	parts := reService.FindStringSubmatch(req.URL.Path)
	if len(parts) != 3 {
		return typhon.Response{Error: terrors.NotFound("bad_endpoint", "Unable to determine endpoint.", nil)}
	}

	return handle(req, "s-"+parts[1], parts[2])
}
