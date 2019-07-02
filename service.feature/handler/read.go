package handler

import (
	"github.com/iancoleman/strcase"
	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.feature/dao"
	"github.com/lolibrary/lolibrary/service.feature/domain"
	"github.com/lolibrary/lolibrary/service.feature/marshaling"
	featureproto "github.com/lolibrary/lolibrary/service.feature/proto"

	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleReadFeature(req typhon.Request) typhon.Response {
	body := &featureproto.GETReadFeatureRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Error decoding body: %v", err)
		return typhon.Response{Error: err}
	}

	switch {
	case !validation.AtLeastOne(body.Id, body.Slug):
		return typhon.Response{Error: validation.ErrMissingOneOf("id", "slug", "short_name")}
	case body.Slug != strcase.ToKebab(body.Slug):
		return typhon.Response{Error: validation.ErrBadParam("slug", "slug should be in kebab-case")}
	}

	var (
		feature *domain.Feature
		err   error
	)

	switch {
	case body.Id != "":
		feature, err = dao.ReadFeature(body.Id)
	case body.Slug != "":
		feature, err = dao.ReadFeatureBySlug(body.Slug)
	}

	if err != nil {
		slog.Error(req, "Error fetching feature: %v", err)
		return typhon.Response{Error: err}
	}

	return req.Response(&featureproto.GETReadFeatureResponse{
		Feature: marshaling.FeatureToProto(feature),
	})
}
