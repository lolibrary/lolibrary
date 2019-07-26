package handler

import (
	"fmt"

	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.category/dao"
	categoryproto "github.com/lolibrary/lolibrary/service.category/proto"
	"github.com/monzo/slog"
	"github.com/monzo/terrors"
	"github.com/monzo/typhon"
)

func handleDeleteCategory(req typhon.Request) typhon.Response {
	body := &categoryproto.DELETERemoveCategoryRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Failed to decode body: %v", err)
		return typhon.Response{Error: err}
	}

	switch {
	case body.Id == "":
		return typhon.Response{Error: validation.ErrMissingParam("id")}
	}

	slogParams := map[string]string{"category_id": body.Id}

	category, err := dao.ReadCategory(req, body.Id)
	if err != nil {
		slog.Error(req, "Error checking if category exists: %v", err, slogParams)
		return typhon.Response{Error: err}
	}
	if category == nil {
		return typhon.Response{Error: terrors.NotFound("category", fmt.Sprintf("Category '%s' not found", body.Id), nil)}
	}

	if err := dao.DeleteCategory(req, category); err != nil {
		slog.Error(req, "Error deleting category: %v", err, slogParams)
		return typhon.Response{Error: err}
	}

	return req.Response(&categoryproto.DELETERemoveCategoryResponse{})
}
