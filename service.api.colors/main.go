package main

import (
	"github.com/lolibrary/lolibrary/foundation"
	"github.com/lolibrary/lolibrary/libraries/filters"
	"github.com/lolibrary/lolibrary/service.api.colors/handler"
)

func main() {
	srv := foundation.Init("service.api.colors")
	svc := handler.Service().Filter(filters.CommonFilters)

	srv.Run(svc)
}
