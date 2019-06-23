package filters

import (
	"github.com/monzo/typhon"
)

// commonFilters is used in foundation, to serve every service.
var commonFilters = [...]typhon.Filter{
	PingFilter,
	HealthCheckFilter,
	PanicFilter,
	typhon.ErrorFilter,
}

// CommonFilters runs every "common" filter against a service.
// This is used to globally update filters.
func CommonFilters(req typhon.Request, svc typhon.Service) typhon.Response {
	for _, filter := range commonFilters {
		svc = svc.Filter(filter)
	}

	return svc(req)
}