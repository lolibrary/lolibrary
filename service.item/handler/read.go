package handler

import (
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
		return typhon.Response{Error: validation.ErrMissingOneOf("id", "slug")}
	case body.Slug != "" && !validation.Slug(body.Slug):
		return typhon.Response{Error: validation.ErrBadSlug("slug", body.Slug)}
	}

	var (
		item *domain.Item
		err  error
	)

	switch {
	case body.Id != "":
		item, err = dao.ReadItem(req, body.Id)
	case body.Slug != "":
		item, err = dao.ReadItemBySlug(req, body.Slug)
	}

	if err != nil {
		slog.Error(req, "Error fetching item: %v", err)
		return typhon.Response{Error: err}
	}

	return req.Response(&itemproto.GETReadItemResponse{
		Item: marshaling.ItemToProto(item),
	})
}
