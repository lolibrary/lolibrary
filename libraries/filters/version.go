package filters

import (
	"github.com/lolibrary/lolibrary/foundation"
	"github.com/monzo/typhon"
)

type versionResponse struct {
	Version string `json:"version"`
}

// VersionFilter returns the git hash of this version.
// It is used to see if this service is up to date.
func VersionFilter(req typhon.Request, svc typhon.Service) typhon.Response {
	if matchesVersionPath(req) {
		return req.Response(&versionResponse{Version: foundation.Version})
	}

	return svc(req)
}

func matchesVersionPath(req typhon.Request) bool {
	return req.Method == "GET" && req.URL.Path == "/foundation/version"
}
