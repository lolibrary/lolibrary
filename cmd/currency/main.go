package main

import (
	"github.com/lolibrary/lolibrary/currency"
	"github.com/lolibrary/lolibrary/pkg/service"
)

func main() {
	service.Serve(currency.Service)
}