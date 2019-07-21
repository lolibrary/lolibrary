package handler

import (
	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.tag/dao"
	"github.com/lolibrary/lolibrary/service.tag/domain"
	"github.com/lolibrary/lolibrary/service.tag/marshaling"
	tagproto "github.com/lolibrary/lolibrary/service.tag/proto"

	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleReadTag(req typhon.Request) typhon.Response {
	body := &tagproto.GETReadTagRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Error decoding body: %v", err)
		return typhon.Response{Error: err}
	}

	switch {
	case !validation.AtLeastOne(body.Id, body.Slug):
		return typhon.Response{Error: validation.ErrMissingOneOf("id", "slug", "short_name")}
	case body.Slug != "" && !validation.Slug(body.Slug):
		return typhon.Response{Error: validation.ErrBadSlug("slug", body.Slug)}
	}

	var (
		tag *domain.Tag
		err error
	)

	switch {
	case body.Id != "":
		tag, err = dao.ReadTag(body.Id)
	case body.Slug != "":
		tag, err = dao.ReadTagBySlug(body.Slug)
	}

	if err != nil {
		slog.Error(req, "Error fetching tag: %v", err)
		return typhon.Response{Error: err}
	}

	return req.Response(&tagproto.GETReadTagResponse{
		Tag: marshaling.TagToProto(tag),
	})
}
