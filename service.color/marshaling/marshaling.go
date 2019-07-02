package marshaling

import (
	"github.com/lolibrary/lolibrary/libraries/util"
	"github.com/lolibrary/lolibrary/service.color/domain"
	colorproto "github.com/lolibrary/lolibrary/service.color/proto"
)

func ColorToProto(color *domain.Color) *colorproto.Color {
	if color == nil {
		return nil
	}

	return &colorproto.Color{
		Id:        color.ID,
		Slug:      color.Slug,
		Name:      color.Name,
		CreatedAt: util.TimeToProto(color.CreatedAt),
		UpdatedAt: util.TimeToProto(color.UpdatedAt),
	}
}

func ColorsToProto(colors []*domain.Color) []*colorproto.Color {
	protos := make([]*colorproto.Color, 0, len(colors))
	for _, color := range colors {
		protos = append(protos, ColorToProto(color))
	}

	return protos
}
