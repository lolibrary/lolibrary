package handler

import (
	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.api.colors/marshaling"
	colorproto "github.com/lolibrary/lolibrary/service.color/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleReadColor(req typhon.Request) typhon.Response {
	params := router.Params(req)
	color, ok := params["color"]
	if !ok {
		slog.Error(req, "Missing color param somehow?")
		return typhon.Response{Error: validation.ErrBadParam("color", "Missing color to query")}
	}

	rsp, err := colorproto.GETReadColorRequest{
		Slug: color,
	}.Send(req).DecodeResponse()
	if err != nil {
		slog.Error(req, "Error reading color: %v", err, params)
		return typhon.Response{Error: err}
	}

	return req.Response(marshaling.ProtoToColor(rsp.Color))
}
