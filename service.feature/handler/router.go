package handler

import "github.com/monzo/typhon"

var router = typhon.Router{}

func init() {
	router.GET("/read", handleReadFeature)
	router.GET("/list", handleListFeatures)
	router.POST("/create", handleCreateFeature)
	router.PUT("/update", handleUpdateFeature)
	router.DELETE("/delete", handleDeleteFeature)
}

// Service serves this router as a typhon service.
func Service() typhon.Service {
	return router.Serve()
}
