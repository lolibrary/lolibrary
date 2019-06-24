package main

import (
	"github.com/lolibrary/lolibrary/foundation"
	"github.com/lolibrary/lolibrary/libraries/filters"

	"github.com/lolibrary/lolibrary/service.brand/dao"
	"github.com/lolibrary/lolibrary/service.brand/handler"
)

func main() {
	svc := handler.Service().Filter(filters.CommonFilters)
	db := dao.Init()
	defer db.Close()

	srv := foundation.Init("service.brand")
	srv.Run(svc)
}
