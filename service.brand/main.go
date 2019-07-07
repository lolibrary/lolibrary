package main

import (
	"github.com/lolibrary/lolibrary/foundation"
	"github.com/lolibrary/lolibrary/libraries/filters"

	"github.com/lolibrary/lolibrary/service.brand/dao"
	"github.com/lolibrary/lolibrary/service.brand/handler"
)

func main() {
	srv := foundation.Init("service.brand")
	db := dao.Init()
	defer db.Close()

	svc := handler.Service().Filter(filters.CommonFilters)
	srv.Run(svc)
}
