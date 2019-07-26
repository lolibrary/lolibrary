package handler

import (
	"github.com/lolibrary/lolibrary/service.api.tags/domain"
	"github.com/lolibrary/lolibrary/service.api.tags/marshaling"
	tagproto "github.com/lolibrary/lolibrary/service.tag/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleListTags(req typhon.Request) typhon.Response {
	rsp, err := tagproto.GETListTagsRequest{}.Send(req).DecodeResponse()
	if err != nil {
		slog.Error(req, "Failed to list tags: %v", err)
		return typhon.Response{Error: err}
	}

	return req.Response(&domain.ListTagsResponse{
		Tags: marshaling.ProtoToTags(rsp.Tags),
	})
}
