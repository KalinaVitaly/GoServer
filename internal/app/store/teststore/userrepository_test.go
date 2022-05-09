package teststore

import (
	"Diplom/internal/app/model"
	"Diplom/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := TestDB(t, databaseURL)
	defer teardown("user_content_db.users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u.ID)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := TestDB(t, databaseURL)
	defer teardown("user_content_db.users")

	s := sqlstore.New(db)
	u1 := model.TestUser(t)
	user, _ := s.User().FindByEmail(u1.Email)
	assert.NotNil(t, user)
}
