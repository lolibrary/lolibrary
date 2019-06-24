// Package main (edge-proxy.internal) is an internal proxy used to route to services for dev and debug.
// It is not actually used by services internally (which rely on cluster DNS and services),
// and just provides a way to access internal services to use RPC (e.g. with the flower CLI)
package main

import (
	"github.com/lolibrary/lolibrary/edge-proxy.internal/handler"
	"github.com/lolibrary/lolibrary/foundation"
	"github.com/lolibrary/lolibrary/libraries/filters"
)

func main() {
	service := handler.Service().Filter(filters.CommonFilters)
	srv := foundation.Init("edge-proxy.internal")
	srv.Run(service)
}
