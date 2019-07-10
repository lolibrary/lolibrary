package handler

import (
	"github.com/iancoleman/strcase"
	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.item/dao"
	"github.com/lolibrary/lolibrary/service.item/domain"
	"github.com/lolibrary/lolibrary/service.item/marshaling"
	itemproto "github.com/lolibrary/lolibrary/service.item/proto"

	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleReadItem(req typhon.Request) typhon.Response {
	body := &itemproto.GETReadItemRequest{}
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
		item *domain.Item
		err   error
	)

	switch {
	case body.Id != "":
		item, err = dao.ReadItem(body.Id)
	case body.Slug != "":
		item, err = dao.ReadItemBySlug(body.Slug)
	}

	if err != nil {
		slog.Error(req, "Error fetching item: %v", err)
		return typhon.Response{Error: err}
	}

	return req.Response(&itemproto.GETReadItemResponse{
		Item: marshaling.ItemToProto(item),
	})
}
