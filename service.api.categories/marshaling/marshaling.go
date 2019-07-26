package marshaling

import (
	"github.com/lolibrary/lolibrary/service.api.categories/domain"
	categoryproto "github.com/lolibrary/lolibrary/service.category/proto"
)

func ProtoToCategories(protos []*categoryproto.Category) []*domain.Category {
	categories := make([]*domain.Category, 0, len(protos))
	for _, category := range protos {
		categories = append(categories, ProtoToCategory(category))
	}

	return categories
}

func ProtoToCategory(category *categoryproto.Category) *domain.Category {
	if category == nil {
		return nil
	}

	return &domain.Category{
		Slug: category.Slug,
		Name: category.Name,
	}
}
