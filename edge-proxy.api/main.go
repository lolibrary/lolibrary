// Package main (edge-proxy.api) is an external API proxy used to route to services starting with "service.api.".
// It is meant to be accessible at https://api.lolibrary.org/<service>
package main

import (
	"github.com/lolibrary/lolibrary/edge-proxy.api/handler"
	"github.com/lolibrary/lolibrary/foundation"
	"github.com/lolibrary/lolibrary/libraries/filters"
)

func main() {
	service := handler.Service().Filter(filters.CommonFilters)
	srv := foundation.Init("edge-proxy.api")
	srv.Run(service)
}
