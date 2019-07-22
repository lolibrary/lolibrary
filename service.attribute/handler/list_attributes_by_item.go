package handler

import (
	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.attribute/dao"
	"github.com/lolibrary/lolibrary/service.attribute/marshaling"
	attributeproto "github.com/lolibrary/lolibrary/service.attribute/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleListAttributesByItem(req typhon.Request) typhon.Response {
	body := &attributeproto.GETListAttributesByItemRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Error decoding request body: %v", err)
		return typhon.Response{Error: err}
	}

	switch {
	case body.ItemId == "":
		return typhon.Response{Error: validation.ErrMissingParam("item_id")}
	}

	values, err := dao.ListAttributesByItem(req, body.ItemId)
	if err != nil {
		slog.Error(req, "Failed to list attributes by item ID: %v", err, map[string]string{
			"item_id": body.ItemId,
		})
		return typhon.Response{Error: err}
	}

	return req.Response(&attributeproto.GETListAttributesByItemResponse{
		Attributes: marshaling.AttributeValuesToProto(values),
	})
}