package main

import (
	"github.com/lolibrary/lolibrary/image"
	"github.com/lolibrary/lolibrary/pkg/service"
)

func main() {
	service.Serve(image.NewService())
}
