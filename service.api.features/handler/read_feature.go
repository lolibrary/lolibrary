package handler

import (
	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.api.features/marshaling"
	featureproto "github.com/lolibrary/lolibrary/service.feature/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleReadFeature(req typhon.Request) typhon.Response {
	params := router.Params(req)
	feature, ok := params["feature"]
	if !ok {
		slog.Error(req, "Missing feature param somehow?")
		return typhon.Response{Error: validation.ErrBadParam("feature", "Missing feature to query")}
	}

	rsp, err := featureproto.GETReadFeatureRequest{
		Slug: feature,
	}.Send(req).DecodeResponse()
	if err != nil {
		slog.Error(req, "Error reading feature: %v", err, params)
		return typhon.Response{Error: err}
	}

	return req.Response(marshaling.ProtoToFeature(rsp.Feature))
}
