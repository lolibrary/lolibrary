package handler

import (
	"time"

	"github.com/iancoleman/strcase"
	"github.com/lolibrary/lolibrary/libraries/idgen"
	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.brand/dao"
	"github.com/lolibrary/lolibrary/service.brand/domain"
	"github.com/lolibrary/lolibrary/service.brand/marshaling"
	brandproto "github.com/lolibrary/lolibrary/service.brand/proto"
	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleCreateBrand(req typhon.Request) typhon.Response {
	body := &brandproto.POSTCreateBrandRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Error decoding request body: %v", err)
		return typhon.Response{Error: err}
	}

	switch {
	case body.Slug == "":
		return typhon.Response{Error: validation.ErrMissingParam("slug")}
	case body.ShortName == "":
		return typhon.Response{Error: validation.ErrMissingParam("short_name")}
	case body.Name == "":
		return typhon.Response{Error: validation.ErrMissingParam("name")}
	case body.Description == "":
		return typhon.Response{Error: validation.ErrMissingParam("description")}
	case body.ImageId == "":
		return typhon.Response{Error: validation.ErrMissingParam("image_id")}
	case body.Slug != strcase.ToKebab(body.Slug):
		return typhon.Response{Error: validation.ErrBadParam("slug", "slug should be in kebab-case")}
	case body.ShortName != strcase.ToKebab(body.ShortName):
		return typhon.Response{Error: validation.ErrBadParam("short_name", "short_name should be in kebab-case")}
	}

	id, err := idgen.New()
	if err != nil {
		slog.Error(req, "Error generating ID: %v", err)
		return typhon.Response{Error: err}
	}

	brand := &domain.Brand{
		ID:          id,
		ImageID:     body.ImageId,
		Slug:        body.Slug,
		ShortName:   body.ShortName,
		Name:        body.Name,
		Description: body.Description,
		CreatedAt:   time.Now().UTC(),
	}

	if err := dao.CreateBrand(brand); err != nil {
		slog.Error(req, "Failed to create brand entry: %v", err)
		return typhon.Response{Error: err}
	}

	return req.Response(&brandproto.POSTCreateBrandResponse{
		Brand: marshaling.BrandToProto(brand),
	})
}
