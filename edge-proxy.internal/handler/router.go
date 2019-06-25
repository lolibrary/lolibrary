package handler

import (
	"strings"

	"github.com/monzo/slog"
	"github.com/monzo/terrors"
	"github.com/monzo/typhon"
)

// Service returns a typhon service.
func Service() typhon.Service {
	return Proxy
}

// Proxy takes an incoming request and proxies it into the cluster.
func Proxy(req typhon.Request) typhon.Response {
	switch {
	case strings.HasPrefix(req.URL.Path, "/service."):
		return handleService(req)
	}

	return typhon.Response{Error: terrors.BadRequest("missing_endpoint", "Don't know how to query that service!", nil)}
}

// handleService takes a request and sends it to an internal service.
func handleService(req typhon.Request) typhon.Response {
	path := "http://s-" + strings.TrimPrefix(req.URL.Path, "/service.")
	slog.Trace(req, "Sending request to: %v", path)

	return typhon.NewRequest(req, req.Method, path, req.Body).Send().Response()
}