package handler

import (
	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.attribute/dao"
	"github.com/lolibrary/lolibrary/service.attribute/marshaling"
	attributeproto "github.com/lolibrary/lolibrary/service.attribute/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleReadAttributeValue(req typhon.Request) typhon.Response {
	body := &attributeproto.GETReadAttributeValueRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Error decoding request body: %v", err)
		return typhon.Response{Error: err}
	}

	switch {
	case body.Id == "":
		return typhon.Response{Error: validation.ErrMissingParam("id")}
	}

	slogParams := map[string]string{"id": body.Id}

	value, err := dao.ReadAttributeValue(req, body.Id)
	if err != nil {
		slog.Error(req, "Error reading attribute value: %v", err, slogParams)
		return typhon.Response{Error: err}
	}

	return req.Response(&attributeproto.GETReadAttributeValueResponse{
		Attribute: marshaling.AttributeValueToProto(value),
	})
}
