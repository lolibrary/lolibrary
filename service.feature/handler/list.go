package handler

import (
	"github.com/lolibrary/lolibrary/service.feature/dao"
	"github.com/lolibrary/lolibrary/service.feature/marshaling"
	featureproto "github.com/lolibrary/lolibrary/service.feature/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleListFeatures(req typhon.Request) typhon.Response {
	body := &featureproto.GETListFeaturesRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Error decoding request body: %v", err)
		return typhon.Response{Error: err}
	}

	features, err := dao.ListFeatures(req)
	if err != nil {
		slog.Error(req, "Failed to read features: %v", err)
		return typhon.Response{Error: err}
	}

	return req.Response(&featureproto.GETListFeaturesResponse{
		Features: marshaling.FeaturesToProto(features),
	})
}
