package marshaling

import (
	"github.com/lolibrary/lolibrary/libraries/util"
	"github.com/lolibrary/lolibrary/service.item/domain"
	itemproto "github.com/lolibrary/lolibrary/service.item/proto"
)

func ItemToProto(item *domain.Item) *itemproto.Item {
	if item == nil {
		return nil
	}

	return &itemproto.Item{
		Id:            item.ID,
		Slug:          item.Slug,
		BrandId:       item.BrandID,
		CategoryId:    item.CategoryID,
		UserId:        item.UserID,
		ImageId:       item.ImageID,
		PublisherId:   item.PublisherID,
		EnglishName:   item.EnglishName,
		ForeignName:   item.ForeignName,
		ProductNumber: item.ProductNumber,
		Currency:      item.Currency,
		Price:         item.Price,
		Year:          int32(item.Year),
		Notes:         item.Notes,
		Status:        StatusToProto(item.Status),
		Metadata:      item.Metadata,
		Features:      item.Features,
		Colors:        item.Colors,
		Tags:          item.Tags,
		Attributes:    item.Attributes,
		CreatedAt:     util.TimeToProto(item.CreatedAt),
		UpdatedAt:     util.TimeToProto(item.UpdatedAt),
		PublishedAt:   util.TimeToProto(item.PublishedAt),
	}
}

func ItemsToProto(items []*domain.Item) []*itemproto.Item {
	protos := make([]*itemproto.Item, 0, len(items))
	for _, item := range items {
		protos = append(protos, ItemToProto(item))
	}

	return protos
}

func StatusToProto(status int) itemproto.Status {
	switch status {
	case 0:
		return itemproto.Status_DRAFT
	case 1:
		return itemproto.Status_PUBLISHED
	default:
		return itemproto.Status_UNKNOWN
	}
}
