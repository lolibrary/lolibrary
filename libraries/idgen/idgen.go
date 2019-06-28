package idgen

import (
	"github.com/monzo/terrors"
	uuid "github.com/nu7hatch/gouuid"
)

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
