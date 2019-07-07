package handler

import (
	"github.com/lolibrary/lolibrary/service.attribute/dao"
	"github.com/lolibrary/lolibrary/service.attribute/marshaling"
	attributeproto "github.com/lolibrary/lolibrary/service.attribute/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleListAttributes(req typhon.Request) typhon.Response {
	body := &attributeproto.GETListAttributesRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Error decoding request body: %v", err)
		return typhon.Response{Error: err}
	}

	attributes, err := dao.ListAttributes()
	if err != nil {
		slog.Error(req, "Failed to read attributes: %v", err)
		return typhon.Response{Error: err}
	}

	return req.Response(&attributeproto.GETListAttributesResponse{
		Attributes: marshaling.AttributesToProto(attributes),
	})
}
