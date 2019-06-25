package main

import (
	"github.com/lolibrary/lolibrary/foundation"
	"github.com/lolibrary/lolibrary/libraries/filters"

	"github.com/lolibrary/lolibrary/service.image/dao"
	"github.com/lolibrary/lolibrary/service.image/handler"
)

func main() {
	svc := handler.Service().Filter(filters.CommonFilters)
	srv := foundation.Init("service.image")
	db := dao.Init()
	defer db.Close()

	srv.Run(svc)
}
