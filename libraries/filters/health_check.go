package filters

import (
	"net/http"

	"github.com/monzo/typhon"
)

type healthCheckResponse struct {
	OK bool `json:"ok"`
}

// HealthCheckFilter checks if the incoming request is GET /healthz and returns a response.
// TODO: implement proper health checking by polling the service for a non-200.
func HealthCheckFilter(req typhon.Request, svc typhon.Service) typhon.Response {
	if matchesHealthCheckPath(req) {
		return req.Response(&healthCheckResponse{OK: true})
	}

	return svc(req)
}

func matchesHealthCheckPath(req typhon.Request) bool {
	return req.Method == http.MethodGet && req.URL.Path == "/healthz"
}
