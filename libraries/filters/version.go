package filters

import (
	"strings"

	"github.com/monzo/typhon"
)

// Version is also replicated in github.com/lolibrary/lolibrary/foundation and set at compile time.
var Version = "dev"

type versionResponse struct {
	Version string `json:"version"`
}

// VersionFilter returns the git hash of this version.
// It is used to see if this service is up to date.
func VersionFilter(req typhon.Request, svc typhon.Service) typhon.Response {
	if matchesVersionPath(req) {
		return req.Response(&versionResponse{Version: Version})
	}

	return svc(req)
}

func matchesVersionPath(req typhon.Request) bool {
	return strings.ToUpper(req.Method) == "GET" && req.URL.Path == "/foundation/version"
}
