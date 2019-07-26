package handler

import (
	"github.com/lolibrary/lolibrary/service.api.colors/domain"
	"github.com/lolibrary/lolibrary/service.api.colors/marshaling"
	colorproto "github.com/lolibrary/lolibrary/service.color/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleListColors(req typhon.Request) typhon.Response {
	rsp, err := colorproto.GETListColorsRequest{}.Send(req).DecodeResponse()
	if err != nil {
		slog.Error(req, "Failed to list colors: %v", err)
		return typhon.Response{Error: err}
	}

	return req.Response(&domain.ListColorsResponse{
		Colors: marshaling.ProtoToColors(rsp.Colors),
	})
}
