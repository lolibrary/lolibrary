package handler

import (
	"github.com/lolibrary/lolibrary/service.api.brands/domain"
	"github.com/lolibrary/lolibrary/service.api.brands/marshaling"
	brandproto "github.com/lolibrary/lolibrary/service.brand/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleListBrands(req typhon.Request) typhon.Response {
	rsp, err := brandproto.GETListBrandsRequest{}.Send(req).DecodeResponse()
	if err != nil {
		slog.Error(req, "Failed to list brands: %v", err)
		return typhon.Response{Error: err}
	}

	return req.Response(&domain.ListBrandsResponse{
		Brands: marshaling.ProtoToBrands(rsp.Brands),
	})
}
