package handler

import (
	"fmt"
	"time"

	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.color/dao"
	"github.com/lolibrary/lolibrary/service.color/domain"
	"github.com/lolibrary/lolibrary/service.color/marshaling"
	colorproto "github.com/lolibrary/lolibrary/service.color/proto"
	"github.com/monzo/terrors"

	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleUpdateColor(req typhon.Request) typhon.Response {
	body := &colorproto.PUTUpdateColorRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Error decoding body: %v", err)
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

	color = handleUserUpdates(color, body)
	color.UpdatedAt = time.Now().UTC()

	if err := dao.UpdateColor(req, color); err != nil {
		slog.Error(req, "Error updating color: %v", err, slogParams)
		return typhon.Response{Error: err}
	}

	return req.Response(&colorproto.PUTUpdateColorResponse{
		Color: marshaling.ColorToProto(color),
	})
}

func handleUserUpdates(color *domain.Color, body *colorproto.PUTUpdateColorRequest) *domain.Color {
	if body.Slug != "" {
		color.Slug = body.Slug
	}

	if body.Name != "" {
		color.Name = body.Name
	}

	return color
}