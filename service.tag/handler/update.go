package handler

import (
	"fmt"
	"time"

	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.tag/dao"
	"github.com/lolibrary/lolibrary/service.tag/domain"
	"github.com/lolibrary/lolibrary/service.tag/marshaling"
	tagproto "github.com/lolibrary/lolibrary/service.tag/proto"
	"github.com/monzo/terrors"

	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleUpdateTag(req typhon.Request) typhon.Response {
	body := &tagproto.PUTUpdateTagRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Error decoding body: %v", err)
		return typhon.Response{Error: err}
	}

	switch {
	case body.Id == "":
		return typhon.Response{Error: validation.ErrMissingParam("id")}
	}

	slogParams := map[string]string{"tag_id": body.Id}

	tag, err := dao.ReadTag(req, body.Id)
	if err != nil {
		slog.Error(req, "Error checking if tag exists: %v", err, slogParams)
		return typhon.Response{Error: err}
	}
	if tag == nil {
		return typhon.Response{Error: terrors.NotFound("tag", fmt.Sprintf("Tag '%s' not found", body.Id), nil)}
	}

	tag = handleUserUpdates(tag, body)
	tag.UpdatedAt = time.Now().UTC()

	if err := dao.UpdateTag(req, tag); err != nil {
		slog.Error(req, "Error updating tag: %v", err, slogParams)
		return typhon.Response{Error: err}
	}

	return req.Response(&tagproto.PUTUpdateTagResponse{
		Tag: marshaling.TagToProto(tag),
	})
}

func handleUserUpdates(tag *domain.Tag, body *tagproto.PUTUpdateTagRequest) *domain.Tag {
	if body.Slug != "" {
		tag.Slug = body.Slug
	}

	if body.Name != "" {
		tag.Name = body.Name
	}

	return tag
}