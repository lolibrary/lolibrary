package handler

import (
	"fmt"

	"github.com/lolibrary/lolibrary/libraries/validation"
	idproto "github.com/lolibrary/lolibrary/service.id/proto"
	"github.com/monzo/slog"
	"github.com/monzo/terrors"
	"github.com/monzo/typhon"
	"github.com/segmentio/ksuid"
)

func handleGenerateFlake(req typhon.Request) typhon.Response {
	body := &idproto.POSTGenerateFlakeRequest{}
	if err := req.Decode(body); err != nil {
		slog.Error(req, "Failed to decode body: %v", err)
		return typhon.Response{Error: err}
	}

	switch {
	case body.Prefix == "":
		return typhon.Response{Error: validation.ErrMissingParam("prefix")}
	}

	// generate KSUID here.
	k, err := ksuid.NewRandom()
	if err != nil {
		return typhon.Response{Error: terrors.Wrap(err, nil)}
	}

	return req.Response(&idproto.POSTGenerateFlakeResponse{
		Id: fmt.Sprintf("%s_%s", body.Prefix, k.String()),
	})
}
