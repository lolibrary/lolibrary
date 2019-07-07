package handler

import "github.com/monzo/typhon"

var router = typhon.Router{}

func init() {
	router.GET("/read", handleReadAttribute)
	router.GET("/list", handleListAttributes)
	router.POST("/create", handleCreateAttribute)
	router.PUT("/update", handleUpdateAttribute)
	router.DELETE("/delete", handleDeleteAttribute)
}

// Service serves this router as a typhon service.
func Service() typhon.Service {
	return router.Serve()
}
