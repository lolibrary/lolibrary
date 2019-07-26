package main

import (
	"github.com/lolibrary/lolibrary/foundation"
	"github.com/lolibrary/lolibrary/service.api.tags/handler"
)

func main() {
	srv := foundation.Init("service.api.tags")
	svc := handler.Service()

	srv.Run(svc)
}
