package handler

import (
	"github.com/lolibrary/lolibrary/service.api.features/domain"
	"github.com/lolibrary/lolibrary/service.api.features/marshaling"
	featureproto "github.com/lolibrary/lolibrary/service.feature/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleListFeatures(req typhon.Request) typhon.Response {
	rsp, err := featureproto.GETListFeaturesRequest{}.Send(req).DecodeResponse()
	if err != nil {
		slog.Error(req, "Failed to list features: %v", err)
		return typhon.Response{Error: err}
	}

	return req.Response(&domain.ListFeaturesResponse{
		Features: marshaling.ProtoToFeatures(rsp.Features),
	})
}
