package handler

import (
	"fmt"
	"time"

	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.feature/dao"
	"github.com/lolibrary/lolibrary/service.feature/domain"
	"github.com/lolibrary/lolibrary/service.feature/marshaling"
	featureproto "github.com/lolibrary/lolibrary/service.feature/proto"
	"github.com/monzo/terrors"

	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleUpdateFeature(req typhon.Request) typhon.Response {
	body := &featureproto.PUTUpdateFeatureRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Error decoding body: %v", err)
		return typhon.Response{Error: err}
	}

	switch {
	case body.Id == "":
		return typhon.Response{Error: validation.ErrMissingParam("id")}
	}

	slogParams := map[string]string{"feature_id": body.Id}

	feature, err := dao.ReadFeature(req, body.Id)
	if err != nil {
		slog.Error(req, "Error checking if feature exists: %v", err, slogParams)
		return typhon.Response{Error: err}
	}
	if feature == nil {
		return typhon.Response{Error: terrors.NotFound("feature", fmt.Sprintf("Feature '%s' not found", body.Id), nil)}
	}

	feature = handleUserUpdates(feature, body)
	feature.UpdatedAt = time.Now().UTC()

	if err := dao.UpdateFeature(req, feature); err != nil {
		slog.Error(req, "Error updating feature: %v", err, slogParams)
		return typhon.Response{Error: err}
	}

	return req.Response(&featureproto.PUTUpdateFeatureResponse{
		Feature: marshaling.FeatureToProto(feature),
	})
}

func handleUserUpdates(feature *domain.Feature, body *featureproto.PUTUpdateFeatureRequest) *domain.Feature {
	if body.Slug != "" {
		feature.Slug = body.Slug
	}

	if body.Name != "" {
		feature.Name = body.Name
	}

	return feature
}