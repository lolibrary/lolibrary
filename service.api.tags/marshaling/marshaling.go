package marshaling

import (
	"github.com/lolibrary/lolibrary/service.api.tags/domain"
	tagproto "github.com/lolibrary/lolibrary/service.tag/proto"
)

func ProtoToTags(protos []*tagproto.Tag) []*domain.Tag {
	tags := make([]*domain.Tag, 0, len(protos))
	for _, tag := range protos {
		tags = append(tags, ProtoToTag(tag))
	}

	return tags
}

func ProtoToTag(tag *tagproto.Tag) *domain.Tag {
	if tag == nil {
		return nil
	}

	return &domain.Tag{
		Slug: tag.Slug,
		Name: tag.Name,
	}
}
