package store

import "Diplom/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	//Delete(user *model.User) error
}
