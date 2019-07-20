package main

import (
	"github.com/lolibrary/lolibrary/foundation"
	"github.com/lolibrary/lolibrary/libraries/filters"
	"github.com/lolibrary/lolibrary/service.sitemap-generator/generator"

	"github.com/lolibrary/lolibrary/service.sitemap-generator/dao"
	"github.com/lolibrary/lolibrary/service.sitemap-generator/handler"
)

func main() {
	svc := handler.Service().Filter(filters.CommonFilters)
	srv := foundation.Init("service.sitemap-generator")

	db := dao.Init()
	defer db.Close()

	generator.Start()
	defer generator.Stop()

	srv.Run(svc)
}
