package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:              "kalina.vitaly2016@yandex.com",
		Password:           "password",
		EncryptionPassword: "encrypted_password",
	}
}
