package handler

import "github.com/monzo/typhon"

var router = typhon.Router{}

func init() {
	router.GET("/read", handleReadColor)
	router.GET("/list", handleListColors)
	router.POST("/create", handleCreateColor)
	router.PUT("/update", handleUpdateColor)
	router.DELETE("/delete", handleDeleteColor)
}

// Service serves this router as a typhon service.
func Service() typhon.Service {
	return router.Serve()
}
