package handler

import (
	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.brand/dao"
	"github.com/lolibrary/lolibrary/service.brand/domain"
	"github.com/lolibrary/lolibrary/service.brand/marshaling"
	brandproto "github.com/lolibrary/lolibrary/service.brand/proto"

	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleReadBrand(req typhon.Request) typhon.Response {
	body := &brandproto.GETReadBrandRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Error decoding body: %v", err)
		return typhon.Response{Error: err}
	}

	switch {
	case !validation.AtLeastOne(body.Id, body.Slug, body.ShortName):
		return typhon.Response{Error: validation.ErrMissingOneOf("id", "slug", "short_name")}
	case body.Slug != "" && !validation.Slug(body.Slug):
		return typhon.Response{Error: validation.ErrBadSlug("slug", body.Slug)}
	case body.ShortName != "" && !validation.Slug(body.ShortName):
		return typhon.Response{Error: validation.ErrBadSlug("short_name", body.ShortName)}
	}

	var (
		brand *domain.Brand
		err   error
	)

	switch {
	case body.Id != "":
		brand, err = dao.ReadBrand(req, body.Id)
	case body.Slug != "":
		brand, err = dao.ReadBrandBySlug(req, body.Slug)
	case body.ShortName != "":
		brand, err = dao.ReadBrandByShortName(req, body.ShortName)
	}

	if err != nil {
		slog.Error(req, "Error fetching brand: %v", err)
		return typhon.Response{Error: err}
	}

	return req.Response(&brandproto.GETReadBrandResponse{
		Brand: marshaling.BrandToProto(brand),
	})
}
