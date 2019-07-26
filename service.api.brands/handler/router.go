package handler

import (
	"github.com/monzo/typhon"
)

var router = typhon.Router{}

func init() {
	router.GET("/", handleListBrands)
	router.GET("/:brand", handleReadBrand)
}

func Service() typhon.Service {
	return router.Serve()
}
