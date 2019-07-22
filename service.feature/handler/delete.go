package handler

import (
	"fmt"

	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.feature/dao"
	featureproto "github.com/lolibrary/lolibrary/service.feature/proto"
	"github.com/monzo/slog"
	"github.com/monzo/terrors"
	"github.com/monzo/typhon"
)

func handleDeleteFeature(req typhon.Request) typhon.Response {
	body := &featureproto.DELETERemoveFeatureRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Failed to decode body: %v", err)
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

	if err := dao.DeleteFeature(req, feature); err != nil {
		slog.Error(req, "Error deleting feature: %v", err, slogParams)
		return typhon.Response{Error: err}
	}

	return req.Response(&featureproto.DELETERemoveFeatureResponse{})
}
