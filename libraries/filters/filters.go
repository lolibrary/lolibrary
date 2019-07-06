package filters

import (
	"github.com/monzo/typhon"
)

// commonFilters is used in foundation, to serve every service.
var commonFilters = [...]typhon.Filter{
	PingFilter,
	HealthCheckFilter,
	VersionFilter,
	PanicFilter,
	typhon.ErrorFilter,
}

// clientFilters is used for setting up the default typhon.Client
var clientFilters = [...]typhon.Filter{
	// ClientURLFilter,
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

// ClientFilters runs a set of filters to allow sending requests to other services.
func ClientFilters(req typhon.Request, svc typhon.Service) typhon.Response {
	for _, filter := range clientFilters {
		svc = svc.Filter(filter)
	}

	return svc(req)
}
