package handler

import (
	"fmt"

	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.attribute/dao"
	attributeproto "github.com/lolibrary/lolibrary/service.attribute/proto"
	"github.com/monzo/slog"
	"github.com/monzo/terrors"
	"github.com/monzo/typhon"
)

func handleDeleteAttribute(req typhon.Request) typhon.Response {
	body := &attributeproto.DELETERemoveAttributeRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Failed to decode body: %v", err)
		return typhon.Response{Error: err}
	}

	switch {
	case body.Id == "":
		return typhon.Response{Error: validation.ErrMissingParam("id")}
	}

	slogParams := map[string]string{"attribute_id": body.Id}

	attribute, err := dao.ReadAttribute(body.Id)
	if err != nil {
		slog.Error(req, "Error checking if attribute exists: %v", err, slogParams)
		return typhon.Response{Error: err}
	}
	if attribute == nil {
		return typhon.Response{Error: terrors.NotFound("attribute", fmt.Sprintf("Attribute '%s' not found", body.Id), nil)}
	}

	if err := dao.DeleteAttribute(attribute.ID); err != nil {
		slog.Error(req, "Error deleting attribute: %v", err, slogParams)
		return typhon.Response{Error: err}
	}

	return req.Response(&attributeproto.DELETERemoveAttributeResponse{})
}
