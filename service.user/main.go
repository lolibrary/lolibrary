package main

import (
	"github.com/lolibrary/lolibrary/foundation"
	"github.com/lolibrary/lolibrary/libraries/filters"

	"github.com/lolibrary/lolibrary/service.user/dao"
	"github.com/lolibrary/lolibrary/service.user/handler"
)

func main() {
	svc := handler.Service().Filter(filters.CommonFilters)
	srv := foundation.Init("service.user")
	db := dao.Init()
	defer db.Close()

	srv.Run(svc)
}
