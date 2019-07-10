package main

import (
	"github.com/lolibrary/lolibrary/foundation"
	"github.com/lolibrary/lolibrary/libraries/filters"

	"github.com/lolibrary/lolibrary/service.item/dao"
	"github.com/lolibrary/lolibrary/service.item/handler"
)

func main() {
	svc := handler.Service().Filter(filters.CommonFilters)
	srv := foundation.Init("service.item")
	db := dao.Init()
	defer db.Close()

	srv.Run(svc)
}