package handler

import (
	"github.com/monzo/typhon"
)

var router = typhon.Router{}

func init() {
	router.GET("/", handleListTags)
	router.GET("/:tag", handleReadTag)
}

func Service() typhon.Service {
	return router.Serve()
}
