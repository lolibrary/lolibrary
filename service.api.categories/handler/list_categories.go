package handler

import (
	"github.com/lolibrary/lolibrary/service.api.categories/domain"
	"github.com/lolibrary/lolibrary/service.api.categories/marshaling"
	categoryproto "github.com/lolibrary/lolibrary/service.category/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleListCategories(req typhon.Request) typhon.Response {
	rsp, err := categoryproto.GETListCategoriesRequest{}.Send(req).DecodeResponse()
	if err != nil {
		slog.Error(req, "Failed to list categories: %v", err)
		return typhon.Response{Error: err}
	}

	return req.Response(&domain.ListCategoriesResponse{
		Categories: marshaling.ProtoToCategories(rsp.Categories),
	})
}
