package handler

import (
	"github.com/lolibrary/lolibrary/service.brand/dao"
	"github.com/lolibrary/lolibrary/service.brand/marshaling"
	brandproto "github.com/lolibrary/lolibrary/service.brand/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleListBrands(req typhon.Request) typhon.Response {
	body := &brandproto.GETListBrandsRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Error decoding request body: %v", err)
		return typhon.Response{Error: err}
	}

	brands, err := dao.ListBrands()
	if err != nil {
		slog.Error(req, "Failed to read brands: %v", err)
		return typhon.Response{Error: err}
	}

	return req.Response(&brandproto.GETListBrandsResponse{
		Brands: marshaling.BrandsToProto(brands),
	})
}
