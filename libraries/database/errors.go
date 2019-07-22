package database

import (
	"fmt"

	"github.com/monzo/terrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NotFound(err error) bool {
	return status.Code(err) == codes.NotFound
}

func ErrNotFound(name, key, id string) *terrors.Error {
	return terrors.NotFound(key, fmt.Sprintf("%s with %s '%s' not found", name, key, id), map[string]string{
		key: id,
	})
}

func AlreadyExists(err error) bool {
	return status.Code(err) == codes.AlreadyExists
}

func ErrAlreadyExists(name, key, id string) *terrors.Error {
	return terrors.BadRequest("unique", fmt.Sprintf("%s with %s '%s' already exists", name, key, id), map[string]string{
		key: id,
	})
}
