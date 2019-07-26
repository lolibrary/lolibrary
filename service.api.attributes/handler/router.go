package handler

import (
	"github.com/monzo/typhon"
)

var router = typhon.Router{}

func init() {
	router.GET("/", handleListAttributes)
	router.GET("/:attribute", handleReadAttribute)
}

func Service() typhon.Service {
	return router.Serve()
}
