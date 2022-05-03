package sqlstore

import (
	"Diplom/internal/app/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}
	//r.store.db.QueryRow()
	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}
