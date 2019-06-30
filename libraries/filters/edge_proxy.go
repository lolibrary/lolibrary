package filters

import (
	"github.com/monzo/typhon"
)

// EdgeProxyFilter proxies a request via the edge proxy by inserting the given hostname if missing.
func EdgeProxyFilter(endpoint string) typhon.Filter {
	return func(req typhon.Request, svc typhon.Service) typhon.Response {
		// request coming into /service.bar/path -> endpoint + path
		if req.Host == "" {
			req.Host = endpoint
			req.URL.Host = endpoint
			req.URL.Scheme = "http"
		}

		return svc(req)
	}
}