package main

import (
	"github.com/lolibrary/lolibrary/foundation"
	"github.com/lolibrary/lolibrary/libraries/filters"

	"github.com/lolibrary/lolibrary/service.color/dao"
	"github.com/lolibrary/lolibrary/service.color/handler"
)

func main() {
	srv := foundation.Init("service.color")
	svc := handler.Service().Filter(filters.CommonFilters)
	db := dao.Init()
	defer db.Close()

	srv.Run(svc)
}
