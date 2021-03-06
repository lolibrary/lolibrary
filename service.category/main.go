package main

import (
	"github.com/lolibrary/lolibrary/foundation"
	"github.com/lolibrary/lolibrary/libraries/filters"

	"github.com/lolibrary/lolibrary/service.category/dao"
	"github.com/lolibrary/lolibrary/service.category/handler"
)

func main() {
	srv := foundation.Init("service.category")
	svc := handler.Service().Filter(filters.CommonFilters)
	db := dao.Init()
	defer db.Close()

	srv.Run(svc)
}
