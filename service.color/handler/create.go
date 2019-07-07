package handler

import (
	"time"

	"github.com/iancoleman/strcase"
	"github.com/lolibrary/lolibrary/libraries/idgen"
	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.color/dao"
	"github.com/lolibrary/lolibrary/service.color/domain"
	"github.com/lolibrary/lolibrary/service.color/marshaling"
	colorproto "github.com/lolibrary/lolibrary/service.color/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleCreateColor(req typhon.Request) typhon.Response {
	body := &colorproto.POSTCreateColorRequest{}
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
	case body.Slug != strcase.ToKebab(body.Slug):
		return typhon.Response{Error: validation.ErrBadParam("slug", "slug should be in kebab-case")}
	}

	if body.Id == "" {
		id, err := idgen.New()
		if err != nil {
			slog.Error(req, "Error generating ID: %v", err)
			return typhon.Response{Error: err}
		}

		body.Id = id
	}

	color := &domain.Color{
		ID:        body.Id,
		Slug:      body.Slug,
		Name:      body.Name,
		CreatedAt: time.Now().UTC(),
	}

	if err := dao.CreateColor(color); err != nil {
		slog.Error(req, "Failed to create color entry: %v", err)
		return typhon.Response{Error: err}
	}

	return req.Response(&colorproto.POSTCreateColorResponse{
		Color: marshaling.ColorToProto(color),
	})
}
