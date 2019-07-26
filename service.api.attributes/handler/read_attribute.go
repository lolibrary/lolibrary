package handler

import (
	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.api.attributes/marshaling"
	attributeproto "github.com/lolibrary/lolibrary/service.attribute/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleReadAttribute(req typhon.Request) typhon.Response {
	params := router.Params(req)
	attribute, ok := params["attribute"]
	if !ok {
		slog.Error(req, "Missing attribute param somehow?")
		return typhon.Response{Error: validation.ErrBadParam("attribute", "Missing attribute to query")}
	}

	rsp, err := attributeproto.GETReadAttributeRequest{
		Slug: attribute,
	}.Send(req).DecodeResponse()
	if err != nil {
		slog.Error(req, "Error reading attribute: %v", err, params)
		return typhon.Response{Error: err}
	}

	return req.Response(marshaling.ProtoToAttribute(rsp.Attribute))
}
