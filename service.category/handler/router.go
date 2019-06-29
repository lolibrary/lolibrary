package handler

import "github.com/monzo/typhon"

var router = typhon.Router{}

func init() {
	router.GET("/read", handleReadCategory)
	router.GET("/list", handleListCategories)
	router.POST("/create", handleCreateCategory)
	router.PUT("/update", handleUpdateCategory)
	router.DELETE("/delete", handleDeleteCategory)
}

// Service serves this router as a typhon service.
func Service() typhon.Service {
	return router.Serve()
}
