package marshaling

import (
	"github.com/lolibrary/lolibrary/service.api.brands/domain"
	brandproto "github.com/lolibrary/lolibrary/service.brand/proto"
)

func ProtoToBrands(protos []*brandproto.Brand) []*domain.Brand {
	brands := make([]*domain.Brand, 0, len(protos))
	for _, brand := range protos {
		brands = append(brands, ProtoToBrand(brand))
	}

	return brands
}

func ProtoToBrand(brand *brandproto.Brand) *domain.Brand {
	if brand == nil {
		return nil
	}

	return &domain.Brand{
		Slug:      brand.Slug,
		ShortName: brand.ShortName,
		Name:      brand.Name,
	}
}
