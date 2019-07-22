package handler

import (
	"time"

	"github.com/lolibrary/lolibrary/libraries/idgen"
	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.feature/dao"
	"github.com/lolibrary/lolibrary/service.feature/domain"
	"github.com/lolibrary/lolibrary/service.feature/marshaling"
	featureproto "github.com/lolibrary/lolibrary/service.feature/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleCreateFeature(req typhon.Request) typhon.Response {
	body := &featureproto.POSTCreateFeatureRequest{}
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

	feature := &domain.Feature{
		ID:        body.Id,
		Slug:      body.Slug,
		Name:      body.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	if err := dao.CreateFeature(req, feature); err != nil {
		slog.Error(req, "Failed to create feature entry: %v", err)
		return typhon.Response{Error: err}
	}

	return req.Response(&featureproto.POSTCreateFeatureResponse{
		Feature: marshaling.FeatureToProto(feature),
	})
}
