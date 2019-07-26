package handler

import (
	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.api.brands/marshaling"
	brandproto "github.com/lolibrary/lolibrary/service.brand/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleReadBrand(req typhon.Request) typhon.Response {
	params := router.Params(req)
	brand, ok := params["brand"]
	if !ok {
		slog.Error(req, "Missing brand param somehow?")
		return typhon.Response{Error: validation.ErrBadParam("brand", "Missing brand to query")}
	}

	rsp, err := brandproto.GETReadBrandRequest{
		Slug: brand,
	}.Send(req).DecodeResponse()
	if err != nil {
		slog.Error(req, "Error reading brand: %v", err, params)
		return typhon.Response{Error: err}
	}

	return req.Response(marshaling.ProtoToBrand(rsp.Brand))
}
