package handler

import (
	"github.com/monzo/typhon"
)

var router = typhon.Router{}

func init() {
	router.GET("/", handleListColors)
	router.GET("/:color", handleReadColor)
}

func Service() typhon.Service {
	return router.Serve()
}
