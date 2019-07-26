package handler

import (
	"time"

	"github.com/lolibrary/lolibrary/libraries/idgen"
	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.item/dao"
	"github.com/lolibrary/lolibrary/service.item/domain"
	"github.com/lolibrary/lolibrary/service.item/marshaling"
	itemproto "github.com/lolibrary/lolibrary/service.item/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleCreateItem(req typhon.Request) typhon.Response {
	body := &itemproto.POSTCreateItemRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Error decoding request body: %v", err)
		return typhon.Response{Error: err}
	}

	// validate body.
	if err := validateCreateRequest(body); err != nil {
		return typhon.Response{Error: err}
	}

	if body.Id == "" {
		id, err := idgen.New()
		if err != nil {
			slog.Error(req, "Error generating ID: %v", err)
			return typhon.Response{Error: err}
		}

		body.Id = id
	}

	item := &domain.Item{
		ID:            body.Id,
		Slug:          body.Slug,
		BrandID:       body.BrandId,
		CategoryID:    body.CategoryId,
		UserID:        body.UserId,
		ImageID:       body.ImageId,
		EnglishName:   body.EnglishName,
		ForeignName:   body.ForeignName,
		ProductNumber: body.ProductNumber,
		Currency:      body.Currency,
		Price:         body.Price,
		Year:          int(body.Year),
		Notes:         body.Notes,
		Status:        int(itemproto.Status_DRAFT),
		Metadata:      body.Metadata,
		Features:      dedupe(body.Features),
		Colors:        dedupe(body.Colors),
		Tags:          dedupe(body.Tags),
		Attributes:    body.Attributes,
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
	}

	if err := dao.CreateItem(req, item); err != nil {
		slog.Error(req, "Error creating item: %v", err)
		return typhon.Response{Error: err}
	}

	// todo: emit an event

	return req.Response(&itemproto.POSTCreateItemResponse{
		Item: marshaling.ItemToProto(item),
	})
}

func validateCreateRequest(body *itemproto.POSTCreateItemRequest) error {
	switch {
	case body.Slug == "":
		return validation.ErrMissingParam("slug")
	case body.BrandId == "":
		return validation.ErrMissingParam("brand_id")
	case body.CategoryId == "":
		return validation.ErrMissingParam("category_id")
	case body.UserId == "":
		return validation.ErrMissingParam("user_id")
	case body.EnglishName == "":
		return validation.ErrMissingParam("english_name")
	case body.ForeignName == "":
		return validation.ErrMissingParam("foreign_name")
	case !validation.Slug(body.Slug):
		return validation.ErrBadSlug("slug", body.Slug)
	}

	return nil
}

func dedupe(args []string) []string {
	m := make(map[string]bool)
	for _, arg := range args {
		m[arg] = true
	}

	deduped := make([]string, 0, len(m))
	for k, _ := range m {
		deduped = append(deduped, k)
	}

	return deduped
}
