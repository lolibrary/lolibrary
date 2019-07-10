package handler

import (
	"time"

	"github.com/lolibrary/lolibrary/libraries/idgen"
	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.attribute/dao"
	"github.com/lolibrary/lolibrary/service.attribute/domain"
	"github.com/lolibrary/lolibrary/service.attribute/marshaling"
	attributeproto "github.com/lolibrary/lolibrary/service.attribute/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleCreateAttribute(req typhon.Request) typhon.Response {
	body := &attributeproto.POSTCreateAttributeRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Error decoding request body: %v", err)
		return typhon.Response{Error: err}
	}

	switch {
	case body.Id != "" && !validation.UUID(body.Id):
		return typhon.Response{Error: validation.ErrBadParam("id", "id should be a valid uuid")}
	case body.Slug == "":
		return typhon.Response{Error: validation.ErrMissingParam("slug")}
	case body.Name == "":
		return typhon.Response{Error: validation.ErrMissingParam("name")}
	case !validation.Slug(body.Slug):
		return typhon.Response{Error: validation.ErrBadSlug("slug", body.Slug)}
	}

	if body.Id == "" {
		id, err := idgen.New()
		if err != nil {
			slog.Error(req, "Error generating ID: %v", err)
			return typhon.Response{Error: err}
		}

		body.Id = id
	}

	attribute := &domain.Attribute{
		ID:        body.Id,
		Slug:      body.Slug,
		Name:      body.Name,
		CreatedAt: time.Now().UTC(),
	}

	if err := dao.CreateAttribute(attribute); err != nil {
		slog.Error(req, "Failed to create attribute entry: %v", err)
		return typhon.Response{Error: err}
	}

	return req.Response(&attributeproto.POSTCreateAttributeResponse{
		Attribute: marshaling.AttributeToProto(attribute),
	})
}
