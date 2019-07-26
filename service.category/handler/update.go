package handler

import (
	"fmt"
	"time"

	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.category/dao"
	"github.com/lolibrary/lolibrary/service.category/domain"
	"github.com/lolibrary/lolibrary/service.category/marshaling"
	categoryproto "github.com/lolibrary/lolibrary/service.category/proto"
	"github.com/monzo/terrors"

	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleUpdateCategory(req typhon.Request) typhon.Response {
	body := &categoryproto.PUTUpdateCategoryRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Error decoding body: %v", err)
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

	category = handleUserUpdates(category, body)
	category.UpdatedAt = time.Now().UTC()

	if err := dao.UpdateCategory(req, category); err != nil {
		slog.Error(req, "Error updating category: %v", err, slogParams)
		return typhon.Response{Error: err}
	}

	return req.Response(&categoryproto.PUTUpdateCategoryResponse{
		Category: marshaling.CategoryToProto(category),
	})
}

func handleUserUpdates(category *domain.Category, body *categoryproto.PUTUpdateCategoryRequest) *domain.Category {
	if body.Slug != "" {
		category.Slug = body.Slug
	}

	if body.Name != "" {
		category.Name = body.Name
	}

	return category
}
