package handler

import (
	"github.com/monzo/typhon"
)

var router = typhon.Router{}

func init() {
	router.GET("/", handleListFeatures)
	router.GET("/:feature", handleReadFeature)
}

func Service() typhon.Service {
	return router.Serve()
}
