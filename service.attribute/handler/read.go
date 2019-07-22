package handler

import (
	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.attribute/dao"
	"github.com/lolibrary/lolibrary/service.attribute/domain"
	"github.com/lolibrary/lolibrary/service.attribute/marshaling"
	attributeproto "github.com/lolibrary/lolibrary/service.attribute/proto"

	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleReadAttribute(req typhon.Request) typhon.Response {
	body := &attributeproto.GETReadAttributeRequest{}
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
		attribute *domain.Attribute
		err       error
	)

	switch {
	case body.Id != "":
		attribute, err = dao.ReadAttribute(req, body.Id)
	case body.Slug != "":
		attribute, err = dao.ReadAttributeBySlug(req, body.Slug)
	}

	if err != nil {
		slog.Error(req, "Error fetching attribute: %v", err)
		return typhon.Response{Error: err}
	}

	return req.Response(&attributeproto.GETReadAttributeResponse{
		Attribute: marshaling.AttributeToProto(attribute),
	})
}
