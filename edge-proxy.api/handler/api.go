package handler

import (
	"fmt"
	"net"

	"github.com/monzo/slog"
	"github.com/monzo/terrors"
	"github.com/monzo/typhon"
)

func handleAPI(req typhon.Request) typhon.Response {
	service, path, err := parsePath(req.URL.Path)
	if err != nil {
		slog.Debug(req, "Error parsing path: %v", err)
		return typhon.Response{Error: err}
	}

	if err := validateService(service); err != nil {
		slog.Error(req, "Unable to connect to %s: %v", service, err)
		return typhon.Response{Error: err}
	}

	req.Host = service
	req.URL.Host = service
	req.URL.Path = path

	return typhon.Send(req).Response()
}

func validateService(service string) error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:80", service))
	if err != nil {
		return terrors.NotFound("service", "Service not found", nil)
	}
	defer conn.Close()

	return nil
}
