package filters

import (
	"fmt"

	"github.com/monzo/typhon"
)

// EdgeProxyFilter proxies a request via the edge proxy by inserting the given hostname if missing.
func EdgeProxyFilter(endpoint string) typhon.Filter {
	return func(req typhon.Request, svc typhon.Service) typhon.Response {
		// request comes in as http://s-foo/path
		// split it up to http://endpoint/s-foo/path
		service := req.URL.Host

		req.Host = endpoint
		req.URL.Host = endpoint
		req.URL.Path = fmt.Sprintf("/%s%s", service, req.URL.Path)

		// put a default in here
		if req.URL.Scheme == "" {
			req.URL.Scheme = "http"
		}

		return svc(req)
	}
}
