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
		Id:        item.ID,
		Slug:      item.Slug,
		Name:      item.Name,
		CreatedAt: util.TimeToProto(item.CreatedAt),
		UpdatedAt: util.TimeToProto(item.UpdatedAt),
	}
}

func ItemsToProto(items []*domain.Item) []*itemproto.Item {
	protos := make([]*itemproto.Item, 0, len(items))
	for _, item := range items {
		protos = append(protos, ItemToProto(item))
	}

	return protos
}
