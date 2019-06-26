package handler

import (
	"fmt"
	"time"

	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.brand/dao"
	"github.com/lolibrary/lolibrary/service.brand/marshaling"
	brandproto "github.com/lolibrary/lolibrary/service.brand/proto"
	"github.com/monzo/terrors"

	"github.com/monzo/slog"
	"github.com/monzo/typhon"
)

func handleUpdateBrand(req typhon.Request) typhon.Response {
	body := &brandproto.PUTUpdateBrandRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Error decoding body: %v", err)
		return typhon.Response{Error: err}
	}

	switch {
	case body.Id == "":
		return typhon.Response{Error: validation.ErrMissingParam("id")}
	}

	slogParams := map[string]string{"brand_id": body.Id}

	brand, err := dao.ReadBrand(body.Id)
	if err != nil {
		slog.Error(req, "Error checking if brand exists: %v", err, slogParams)
		return typhon.Response{Error: err}
	}
	if brand == nil {
		return typhon.Response{Error: terrors.NotFound("brand", fmt.Sprintf("Brand '%s' not found", body.Id), nil)}
	}

	switch {
	case body.ImageId != "":
		brand.ImageID = body.ImageId
		fallthrough
	case body.Name != "":
		brand.Name = body.Name
		fallthrough
	case body.Description != "":
		brand.Description = body.Description
		fallthrough
	default:
		brand.UpdatedAt = time.Now().UTC()
	}

	if err := dao.UpdateBrand(brand); err != nil {
		slog.Error(req, "Error updating brand: %v", err, slogParams)
		return typhon.Response{Error: err}
	}

	return req.Response(&brandproto.PUTUpdateBrandResponse{
		Brand: marshaling.BrandToProto(brand),
	})
}
