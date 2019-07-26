package handler

import (
	"github.com/monzo/typhon"
)

var router = typhon.Router{}

func init() {
	router.GET("/", handleListCategories)
	router.GET("/:category", handleReadCategory)
}

func Service() typhon.Service {
	return router.Serve()
}
