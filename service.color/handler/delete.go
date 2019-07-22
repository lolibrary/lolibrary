package handler

import (
	"fmt"

	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.color/dao"
	colorproto "github.com/lolibrary/lolibrary/service.color/proto"
	"github.com/monzo/slog"
	"github.com/monzo/terrors"
	"github.com/monzo/typhon"
)

func handleDeleteColor(req typhon.Request) typhon.Response {
	body := &colorproto.DELETERemoveColorRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Failed to decode body: %v", err)
		return typhon.Response{Error: err}
	}

	switch {
	case body.Id == "":
		return typhon.Response{Error: validation.ErrMissingParam("id")}
	}

	slogParams := map[string]string{"color_id": body.Id}

	color, err := dao.ReadColor(req, body.Id)
	if err != nil {
		slog.Error(req, "Error checking if color exists: %v", err, slogParams)
		return typhon.Response{Error: err}
	}
	if color == nil {
		return typhon.Response{Error: terrors.NotFound("color", fmt.Sprintf("Color '%s' not found", body.Id), nil)}
	}

	if err := dao.DeleteColor(req, color); err != nil {
		slog.Error(req, "Error deleting color: %v", err, slogParams)
		return typhon.Response{Error: err}
	}

	return req.Response(&colorproto.DELETERemoveColorResponse{})
}
