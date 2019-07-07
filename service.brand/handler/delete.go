package handler

import (
	"fmt"

	"github.com/lolibrary/lolibrary/libraries/validation"
	"github.com/lolibrary/lolibrary/service.brand/dao"
	brandproto "github.com/lolibrary/lolibrary/service.brand/proto"
	"github.com/monzo/slog"
	"github.com/monzo/terrors"
	"github.com/monzo/typhon"
)

func handleDeleteBrand(req typhon.Request) typhon.Response {
	body := &brandproto.DELETERemoveBrandRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Failed to decode body: %v", err)
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

	if err := dao.DeleteBrand(brand.ID); err != nil {
		slog.Error(req, "Error deleting brand: %v", err, slogParams)
		return typhon.Response{Error: err}
	}

	return req.Response(&brandproto.DELETERemoveBrandResponse{})
}
