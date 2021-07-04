package store_test

import (
	"github.com/Sereggan/simpleWebApp/internal/app/model"
	"github.com/Sereggan/simpleWebApp/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "user@exampe.org"
	_, err := s.User().FindByEmail(email)

	assert.Error(t, err)
}
