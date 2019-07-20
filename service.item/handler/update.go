package handler

import (
	"fmt"
	"time"

	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.item/dao"
	"github.com/lolibrary/lolibrary/service.item/domain"
	"github.com/lolibrary/lolibrary/service.item/marshaling"
	itemproto "github.com/lolibrary/lolibrary/service.item/proto"
	"github.com/monzo/slog"
	"github.com/monzo/terrors"
	"github.com/monzo/typhon"
)

func handleUpdateItem(req typhon.Request) typhon.Response {
	body := &itemproto.PUTUpdateItemRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Error decoding body: %v", err)
		return typhon.Response{Error: err}
	}

	switch {
	case body.Id == "":
		return typhon.Response{Error: validation.ErrMissingParam("id")}
	}

	slogParams := map[string]string{"item_id": body.Id}

	item, err := dao.ReadItem(body.Id)
	if err != nil {
		slog.Error(req, "Error checking if item exists: %v", err, slogParams)
		return typhon.Response{Error: err}
	}
	if item == nil {
		return typhon.Response{Error: terrors.NotFound("item", fmt.Sprintf("Item '%s' not found", body.Id), nil)}
	}

	item = handleUserUpdates(body, item)
	item.UpdatedAt = time.Now().UTC()

	if err := dao.UpdateItem(item); err != nil {
		slog.Error(req, "Error updating feature: %v", err, slogParams)
		return typhon.Response{Error: err}
	}

	return req.Response(&itemproto.PUTUpdateItemResponse{
		Item: marshaling.ItemToProto(item),
	})
}

func handleUserUpdates(body *itemproto.PUTUpdateItemRequest, item *domain.Item) *domain.Item {
	if body.Slug != "" {
		item.Slug = body.Slug
	}

	if body.ForeignName != "" {
		item.ForeignName = body.ForeignName
	}

	if body.EnglishName != "" {
		item.EnglishName = body.EnglishName
	}

	if body.ProductNumber != "" {
		item.ProductNumber = body.ProductNumber
	}

	if body.BrandId != "" {
		item.BrandID = body.BrandId
	}

	if body.CategoryId != "" {
		item.CategoryID = body.CategoryId
	}

	if body.Currency != "" {
		item.Currency = body.Currency
	}

	if body.Price != "" {
		item.Price = body.Price
	}

	if body.Year > 0 {
		item.Year = int(body.Year)
	}

	if body.Notes != "" {
		item.Notes = body.Notes
	}

	if body.Metadata != nil {
		item.Metadata = body.Metadata
	}

	return item
}
