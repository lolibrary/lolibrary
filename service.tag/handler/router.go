package handler

import "github.com/monzo/typhon"

var router = typhon.Router{}

func init() {
	router.GET("/read", handleReadTag)
	router.GET("/list", handleListTags)
	router.POST("/create", handleCreateTag)
	router.PUT("/update", handleUpdateTag)
	router.DELETE("/delete", handleDeleteTag)
}

// Service serves this router as a typhon service.
func Service() typhon.Service {
	return router.Serve()
}
