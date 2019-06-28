package filters

import (
	"net/http"

	"github.com/monzo/typhon"
)

type pingResponse struct {
	Ping string `json:"ping"`
}

// PingFilter checks if the incoming request is GET /ping and returns a response.
func PingFilter(req typhon.Request, svc typhon.Service) typhon.Response {
	if matchesPingPath(req) {
		return req.Response(&pingResponse{Ping: "pong"})
	}

	return svc(req)
}

func matchesPingPath(req typhon.Request) bool {
	return req.Method == http.MethodGet && req.URL.Path == "/ping"
}
