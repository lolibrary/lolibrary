package main

import (
	"github.com/lolibrary/lolibrary/foundation"
	"github.com/lolibrary/lolibrary/libraries/filters"

	"github.com/lolibrary/lolibrary/service.feature/dao"
	"github.com/lolibrary/lolibrary/service.feature/handler"
)

func main() {
	srv := foundation.Init("service.feature")
	svc := handler.Service().Filter(filters.CommonFilters)
	db := dao.Init()
	defer db.Close()

	srv.Run(svc)
}
