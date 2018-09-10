package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserAnonymizing(t *testing.T) {
	user := &User{
		ID:       "testing value",
		Name:     "Jane Doe",
		Email:    "jane@email.example",
		Username: "jane-doe",
	}

	anon := user.Anonymize()

	assert.Equal(t, user.Name, anon.Name)
	assert.Equal(t, user.Username, anon.Username)
}
