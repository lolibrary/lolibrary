package main

import (
	"github.com/lolibrary/lolibrary/foundation"
	"github.com/lolibrary/lolibrary/libraries/filters"
	"github.com/lolibrary/lolibrary/service.item/dao"
	"github.com/lolibrary/lolibrary/service.item/handler"
)

func main() {
	srv := foundation.Init("service.item")
	dao.Init()

	svc := handler.Service().Filter(filters.CommonFilters)
	srv.Run(svc)
}
