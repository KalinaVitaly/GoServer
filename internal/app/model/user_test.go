package model_test

import (
	"Diplom/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptionPassword)
}

func TestUser_Validate(t *testing.T) {
	u := model.TestUser(t)

	assert.NoError(t, u.Validate())

	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "vlados",
			u: func() *model.User {
				return model.TestUser(t)
			},

			isValid: true,
		},
		{
			name: "egor",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""

				return u
			},

			isValid: false,
		},
		{
			name: "egor krasavchik",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "not real email format	"

				return u
			},

			isValid: false,
		},
		{
			name: "egor krasavchik without password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""

				return u
			},

			isValid: false,
		},
		{
			name: "egor krasavchik with short password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "1"

				return u
			},

			isValid: false,
		},
		{
			name: "egor krasavchik with encrypt password and without password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				u.EncryptionPassword = "encryptionpassword"
				return u
			},

			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}
