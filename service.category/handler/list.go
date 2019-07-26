package handler

import (
	"github.com/lolibrary/lolibrary/service.category/dao"
	"github.com/lolibrary/lolibrary/service.category/marshaling"
	categoryproto "github.com/lolibrary/lolibrary/service.category/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleListCategories(req typhon.Request) typhon.Response {
	body := &categoryproto.GETListCategoriesRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Error decoding request body: %v", err)
		return typhon.Response{Error: err}
	}

	categories, err := dao.ListCategories(req)
	if err != nil {
		slog.Error(req, "Failed to read categories: %v", err)
		return typhon.Response{Error: err}
	}

	return req.Response(&categoryproto.GETListCategoriesResponse{
		Categories: marshaling.CategoriesToProto(categories),
	})
}
