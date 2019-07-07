package main

import (
	"github.com/lolibrary/lolibrary/foundation"
	"github.com/lolibrary/lolibrary/libraries/filters"

	"github.com/lolibrary/lolibrary/service.color/dao"
	"github.com/lolibrary/lolibrary/service.color/handler"
)

func main() {
	srv := foundation.Init("service.color")
	db := dao.Init()
	defer db.Close()

	svc := handler.Service().Filter(filters.CommonFilters)
	srv.Run(svc)
}
