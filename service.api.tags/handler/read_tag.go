package handler

import (
	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.api.tags/marshaling"
	tagproto "github.com/lolibrary/lolibrary/service.tag/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleReadTag(req typhon.Request) typhon.Response {
	params := router.Params(req)
	tag, ok := params["tag"]
	if !ok {
		slog.Error(req, "Missing tag param somehow?")
		return typhon.Response{Error: validation.ErrBadParam("tag", "Missing tag to query")}
	}

	rsp, err := tagproto.GETReadTagRequest{
		Slug: tag,
	}.Send(req).DecodeResponse()
	if err != nil {
		slog.Error(req, "Error reading tag: %v", err, params)
		return typhon.Response{Error: err}
	}

	return req.Response(marshaling.ProtoToTag(rsp.Tag))
}
