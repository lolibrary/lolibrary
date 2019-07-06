package handler

import (
	"fmt"

	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.tag/dao"
	tagproto "github.com/lolibrary/lolibrary/service.tag/proto"
	"github.com/monzo/slog"
	"github.com/monzo/terrors"
	"github.com/monzo/typhon"
)

func handleDeleteTag(req typhon.Request) typhon.Response {
	body := &tagproto.DELETERemoveTagRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Failed to decode body: %v", err)
		return typhon.Response{Error: err}
	}

	switch {
	case body.Id == "":
		return typhon.Response{Error: validation.ErrMissingParam("id")}
	}

	slogParams := map[string]string{"tag_id": body.Id}

	tag, err := dao.ReadTag(body.Id)
	if err != nil {
		slog.Error(req, "Error checking if tag exists: %v", err, slogParams)
		return typhon.Response{Error: err}
	}
	if tag == nil {
		return typhon.Response{Error: terrors.NotFound("tag", fmt.Sprintf("Tag '%s' not found", body.Id), nil)}
	}

	if err := dao.DeleteTag(tag.ID); err != nil {
		slog.Error(req, "Error deleting tag: %v", err, slogParams)
		return typhon.Response{Error: err}
	}

	return req.Response(&tagproto.DELETERemoveTagResponse{})
}
