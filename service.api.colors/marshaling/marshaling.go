package marshaling

import (
	"github.com/lolibrary/lolibrary/service.api.colors/domain"
	colorproto "github.com/lolibrary/lolibrary/service.color/proto"
)

func ProtoToColors(protos []*colorproto.Color) []*domain.Color {
	colors := make([]*domain.Color, 0, len(protos))
	for _, color := range protos {
		colors = append(colors, ProtoToColor(color))
	}

	return colors
}

func ProtoToColor(color *colorproto.Color) *domain.Color {
	if color == nil {
		return nil
	}

	return &domain.Color{
		Slug: color.Slug,
		Name: color.Name,
	}
}
