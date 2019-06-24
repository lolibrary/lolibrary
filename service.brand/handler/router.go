package handler

import "github.com/monzo/typhon"

var router = typhon.Router{}

func init() {
	router.GET("/read", handleReadBrand)
	router.GET("/list", handleListBrands)
	router.POST("/create", handleCreateBrand)
	router.PUT("/update", handleReadBrand)
	router.DELETE("/delete", handleReadBrand)
}

// Service serves this router as a typhon service.
func Service() typhon.Service {
	return router.Serve()
}