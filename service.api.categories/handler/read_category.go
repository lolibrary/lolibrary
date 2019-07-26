package handler

import (
	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.api.categories/marshaling"
	categoryproto "github.com/lolibrary/lolibrary/service.category/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleReadCategory(req typhon.Request) typhon.Response {
	params := router.Params(req)
	category, ok := params["category"]
	if !ok {
		slog.Error(req, "Missing category param somehow?")
		return typhon.Response{Error: validation.ErrBadParam("category", "Missing category to query")}
	}

	rsp, err := categoryproto.GETReadCategoryRequest{
		Slug: category,
	}.Send(req).DecodeResponse()
	if err != nil {
		slog.Error(req, "Error reading category: %v", err, params)
		return typhon.Response{Error: err}
	}

	return req.Response(marshaling.ProtoToCategory(rsp.Category))
}
