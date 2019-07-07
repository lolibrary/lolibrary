package main

import (
	"github.com/lolibrary/lolibrary/foundation"
	"github.com/lolibrary/lolibrary/libraries/filters"

	"github.com/lolibrary/lolibrary/service.category/dao"
	"github.com/lolibrary/lolibrary/service.category/handler"
)

func main() {
	srv := foundation.Init("service.category")
	db := dao.Init()
	defer db.Close()

	svc := handler.Service().Filter(filters.CommonFilters)
	srv.Run(svc)
}
