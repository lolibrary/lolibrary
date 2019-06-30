package filters

import (
	"regexp"
	"strings"

	"github.com/monzo/typhon"
)

var reService = regexp.MustCompile(`^/([a-z.\-]+)/(.*)?$`)

func ClientURLFilter(req typhon.Request, svc typhon.Service) typhon.Response {
	if req.URL == nil || req.URL.Host != "" {
		return svc(req)
	}

	matches := reService.FindStringSubmatch(req.URL.Path)
	if matches != nil {
		req.URL.Scheme = "http"
		req.URL.Host = serviceToHost(matches[1])
		req.URL.Path = "/" + matches[2]

		req.Host = req.URL.Host
	}

	return svc(req)
}

// serviceToHost turns a service.foo string into a routable hostname.
func serviceToHost(service string) string {
	host := strings.ReplaceAll(service, ".", "-")
	host = strings.Replace(host, "service", "s", 1)

	return host
}
