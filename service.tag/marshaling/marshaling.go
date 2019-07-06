package marshaling

import (
	"github.com/lolibrary/lolibrary/libraries/util"
	"github.com/lolibrary/lolibrary/service.tag/domain"
	tagproto "github.com/lolibrary/lolibrary/service.tag/proto"
)

func TagToProto(tag *domain.Tag) *tagproto.Tag {
	if tag == nil {
		return nil
	}

	return &tagproto.Tag{
		Id:        tag.ID,
		Slug:      tag.Slug,
		Name:      tag.Name,
		CreatedAt: util.TimeToProto(tag.CreatedAt),
		UpdatedAt: util.TimeToProto(tag.UpdatedAt),
	}
}

func TagsToProto(tags []*domain.Tag) []*tagproto.Tag {
	protos := make([]*tagproto.Tag, 0, len(tags))
	for _, tag := range tags {
		protos = append(protos, TagToProto(tag))
	}

	return protos
}
