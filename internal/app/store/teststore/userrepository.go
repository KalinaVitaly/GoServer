package teststore

import (
	"Diplom/internal/app/model"
	"Diplom/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[int]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	u.ID = len(r.users) + 1
	r.users[u.ID] = u
	return nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

func (r *UserRepository) Delete(user *model.User) error {
	userModel, err := r.FindByEmail(user.Email)

	if err != nil {
		return store.ErrRecordNotFound
	}

	if err := user.BeforeCreate(); err != nil {
		return err
	}

	if _, isOk := r.users[userModel.ID]; isOk {
		delete(r.users, userModel.ID)
		return nil
	}

	return store.ErrRecordNotFound
}
