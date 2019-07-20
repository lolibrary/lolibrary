package handler

import "github.com/monzo/typhon"

var router = typhon.Router{}

func init() {
	router.GET("/read", handleReadItem)
	// router.PUT("/update", handleUpdateItem)
	// router.DELETE("/delete", handleDeleteItem)
}

// Service serves this router as a typhon service.
func Service() typhon.Service {
	return router.Serve()
}
