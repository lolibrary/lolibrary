package handler

import (
	"time"

	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.attribute/dao"
	"github.com/lolibrary/lolibrary/service.attribute/domain"
	"github.com/lolibrary/lolibrary/service.attribute/marshaling"
	attributeproto "github.com/lolibrary/lolibrary/service.attribute/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleCreateAttributeValue(req typhon.Request) typhon.Response {
	body := &attributeproto.POSTCreateAttributeValueRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Error decoding request body: %v", err)
		return typhon.Response{Error: err}
	}

	switch {
	case body.AttributeId == "":
		return typhon.Response{Error: validation.ErrMissingParam("attribute_id")}
	case body.ItemId == "":
		return typhon.Response{Error: validation.ErrMissingParam("item_id")}
	case body.Value == "":
		return typhon.Response{Error: validation.ErrMissingParam("value")}
	case !validation.UUID(body.AttributeId):
		return typhon.Response{Error: validation.ErrBadParam("attribute_id", "attribute_id should be a valid uuid")}
	case !validation.UUID(body.ItemId):
		return typhon.Response{Error: validation.ErrBadParam("item_id", "item_id should be a valid uuid")}
	}

	value := &domain.AttributeValue{
		AttributeID: body.AttributeId,
		ItemID:      body.ItemId,
		Value:       body.Value,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	if err := dao.CreateAttributeValue(req, value); err != nil {
		slog.Error(req, "Failed to create attribute entry: %v", err)
		return typhon.Response{Error: err}
	}

	return req.Response(&attributeproto.POSTCreateAttributeValueResponse{
		Attribute: marshaling.AttributeValueToProto(value),
	})
}
