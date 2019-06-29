package marshaling

import (
	"github.com/lolibrary/lolibrary/libraries/util"
	"github.com/lolibrary/lolibrary/service.category/domain"
	categoryproto "github.com/lolibrary/lolibrary/service.category/proto"
)

func CategoryToProto(category *domain.Category) *categoryproto.Category {
	if category == nil {
		return nil
	}

	return &categoryproto.Category{
		Id:        category.ID,
		Slug:      category.Slug,
		Name:      category.Name,
		CreatedAt: util.TimeToProto(category.CreatedAt),
		UpdatedAt: util.TimeToProto(category.UpdatedAt),
	}
}

func CategoriesToProto(categories []*domain.Category) []*categoryproto.Category {
	protos := make([]*categoryproto.Category, 0, len(categories))
	for _, category := range categories {
		protos = append(protos, CategoryToProto(category))
	}

	return protos
}
