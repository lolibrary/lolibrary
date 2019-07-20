package handler

import (
	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.color/dao"
	"github.com/lolibrary/lolibrary/service.color/domain"
	"github.com/lolibrary/lolibrary/service.color/marshaling"
	colorproto "github.com/lolibrary/lolibrary/service.color/proto"

	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleReadColor(req typhon.Request) typhon.Response {
	body := &colorproto.GETReadColorRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Error decoding body: %v", err)
		return typhon.Response{Error: err}
	}

	switch {
	case !validation.AtLeastOne(body.Id, body.Slug):
		return typhon.Response{Error: validation.ErrMissingOneOf("id", "slug", "short_name")}
	case !validation.Slug(body.Slug):
		return typhon.Response{Error: validation.ErrBadSlug("slug", body.Slug)}
	}

	var (
		color *domain.Color
		err   error
	)

	switch {
	case body.Id != "":
		color, err = dao.ReadColor(body.Id)
	case body.Slug != "":
		color, err = dao.ReadColorBySlug(body.Slug)
	}

	if err != nil {
		slog.Error(req, "Error fetching color: %v", err)
		return typhon.Response{Error: err}
	}

	return req.Response(&colorproto.GETReadColorResponse{
		Color: marshaling.ColorToProto(color),
	})
}
