package handler

import (
	"github.com/lolibrary/lolibrary/service.tag/dao"
	"github.com/lolibrary/lolibrary/service.tag/marshaling"
	tagproto "github.com/lolibrary/lolibrary/service.tag/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleListTags(req typhon.Request) typhon.Response {
	body := &tagproto.GETListTagsRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Error decoding request body: %v", err)
		return typhon.Response{Error: err}
	}

	tags, err := dao.ListTags()
	if err != nil {
		slog.Error(req, "Failed to read tags: %v", err)
		return typhon.Response{Error: err}
	}

	return req.Response(&tagproto.GETListTagsResponse{
		Tags: marshaling.TagsToProto(tags),
	})
}
