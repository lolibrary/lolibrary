package main

import (
	"github.com/lolibrary/lolibrary/foundation"
	"github.com/lolibrary/lolibrary/libraries/filters"
	"github.com/lolibrary/lolibrary/service.api.categories/handler"
)

func main() {
	srv := foundation.Init("service.api.categories")
	svc := handler.Service().Filter(filters.CommonFilters)

	srv.Run(svc)
}
