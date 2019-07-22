package handler

import "github.com/monzo/typhon"

var router = typhon.Router{}

func init() {
	router.GET("/read", handleReadAttribute)
	router.GET("/list", handleListAttributes)
	router.POST("/create", handleCreateAttribute)
	router.PUT("/update", handleUpdateAttribute)
	router.DELETE("/delete", handleDeleteAttribute)

	router.GET("/values-by-item/list", handleListAttributesByItem)
	router.GET("/values/read", handleReadAttributeValue)
	router.POST("/values/create", handleCreateAttributeValue)
	router.PUT("/values/update", handleUpdateAttributeValue)
	router.DELETE("/values/delete", handleDeleteAttributeValue)
}

// Service serves this router as a typhon service.
func Service() typhon.Service {
	return router.Serve()
}
