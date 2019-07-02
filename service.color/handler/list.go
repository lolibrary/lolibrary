package handler

import (
	"github.com/lolibrary/lolibrary/service.color/dao"
	"github.com/lolibrary/lolibrary/service.color/marshaling"
	colorproto "github.com/lolibrary/lolibrary/service.color/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleListColors(req typhon.Request) typhon.Response {
	body := &colorproto.GETListColorsRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Error decoding request body: %v", err)
		return typhon.Response{Error: err}
	}

	colors, err := dao.ListColors()
	if err != nil {
		slog.Error(req, "Failed to read colors: %v", err)
		return typhon.Response{Error: err}
	}

	return req.Response(&colorproto.GETListColorsResponse{
		Colors: marshaling.ColorsToProto(colors),
	})
}
