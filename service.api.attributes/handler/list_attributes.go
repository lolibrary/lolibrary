package handler

import (
	"github.com/lolibrary/lolibrary/service.api.attributes/domain"
	"github.com/lolibrary/lolibrary/service.api.attributes/marshaling"
	attributeproto "github.com/lolibrary/lolibrary/service.attribute/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleListAttributes(req typhon.Request) typhon.Response {
	rsp, err := attributeproto.GETListAttributesRequest{}.Send(req).DecodeResponse()
	if err != nil {
		slog.Error(req, "Failed to list attributes: %v", err)
		return typhon.Response{Error: err}
	}

	return req.Response(&domain.ListAttributesResponse{
		Attributes: marshaling.ProtoToAttributes(rsp.Attributes),
	})
}
