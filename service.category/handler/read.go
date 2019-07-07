package handler

import (
	"github.com/iancoleman/strcase"
	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.category/dao"
	"github.com/lolibrary/lolibrary/service.category/domain"
	"github.com/lolibrary/lolibrary/service.category/marshaling"
	categoryproto "github.com/lolibrary/lolibrary/service.category/proto"

	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleReadCategory(req typhon.Request) typhon.Response {
	body := &categoryproto.GETReadCategoryRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Error decoding body: %v", err)
		return typhon.Response{Error: err}
	}

	switch {
	case !validation.AtLeastOne(body.Id, body.Slug):
		return typhon.Response{Error: validation.ErrMissingOneOf("id", "slug", "short_name")}
	case body.Slug != strcase.ToKebab(body.Slug):
		return typhon.Response{Error: validation.ErrBadParam("slug", "slug should be in kebab-case")}
	}

	var (
		category *domain.Category
		err   error
	)

	switch {
	case body.Id != "":
		category, err = dao.ReadCategory(body.Id)
	case body.Slug != "":
		category, err = dao.ReadCategoryBySlug(body.Slug)
	}

	if err != nil {
		slog.Error(req, "Error fetching category: %v", err)
		return typhon.Response{Error: err}
	}

	return req.Response(&categoryproto.GETReadCategoryResponse{
		Category: marshaling.CategoryToProto(category),
	})
}
