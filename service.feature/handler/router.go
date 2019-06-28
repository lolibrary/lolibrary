package handler

import "github.com/monzo/typhon"

var router = typhon.Router{}

func init() {
	// add routes
}

// Service serves this router as a typhon service.
func Service() typhon.Service {
	return router.Serve()
}
