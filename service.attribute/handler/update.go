package handler

import (
	"fmt"
	"time"

	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.attribute/dao"
	"github.com/lolibrary/lolibrary/service.attribute/domain"
	"github.com/lolibrary/lolibrary/service.attribute/marshaling"
	attributeproto "github.com/lolibrary/lolibrary/service.attribute/proto"
	"github.com/monzo/terrors"

	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleUpdateAttribute(req typhon.Request) typhon.Response {
	body := &attributeproto.PUTUpdateAttributeRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Error decoding body: %v", err)
		return typhon.Response{Error: err}
	}

	switch {
	case body.Id == "":
		return typhon.Response{Error: validation.ErrMissingParam("id")}
	}

	slogParams := map[string]string{"attribute_id": body.Id}

	attribute, err := dao.ReadAttribute(req, body.Id)
	if err != nil {
		slog.Error(req, "Error checking if attribute exists: %v", err, slogParams)
		return typhon.Response{Error: err}
	}
	if attribute == nil {
		return typhon.Response{Error: terrors.NotFound("attribute", fmt.Sprintf("Attribute '%s' not found", body.Id), nil)}
	}

	attribute = handleUserUpdates(attribute, body)
	attribute.UpdatedAt = time.Now().UTC()

	if err := dao.UpdateAttribute(req, attribute); err != nil {
		slog.Error(req, "Error updating attribute: %v", err, slogParams)
		return typhon.Response{Error: err}
	}

	return req.Response(&attributeproto.PUTUpdateAttributeResponse{
		Attribute: marshaling.AttributeToProto(attribute),
	})
}

func handleUserUpdates(attribute *domain.Attribute, body *attributeproto.PUTUpdateAttributeRequest) *domain.Attribute {
	if body.Slug != "" {
		attribute.Slug = body.Slug
	}

	if body.Name != "" {
		attribute.Name = body.Name
	}

	return attribute
}
