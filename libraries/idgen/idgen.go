package idgen

import (
	"context"
	"fmt"

	idproto "github.com/lolibrary/lolibrary/service.id/proto"
	"github.com/monzo/terrors"
	uuid "github.com/nu7hatch/gouuid"
)

var mocked = false

// New creates a new UUID.
func New() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", terrors.Wrap(err, nil)
	}

	return id.String(), nil
}

// RandomTestID returns a random ID, panicking if it can't generate one.
func RandomTestID() string {
	id, err := New()
	if err != nil {
		panic(err)
	}

	return id
}

// Flake grabs a generated flake ID from service.id.
func Flake(ctx context.Context, prefix string) (string, error) {
	if mocked {
		return fmt.Sprintf("%s_%s", prefix, RandomTestID()), nil
	}

	rsp, err := idproto.POSTGenerateFlakeRequest{
		Prefix: prefix,
	}.Send(ctx).DecodeResponse()
	if err != nil {
		return "", err
	}

	return rsp.Id, nil
}

func Mock() {
	mocked = true
}