package handler

import (
	"github.com/monzo/typhon"
)

// Service returns a typhon service.
func Service() typhon.Service {
	return handleAPI
}