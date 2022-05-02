package store

import "Diplom/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}
	//r.store.db.QueryRow()
	return nil, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}
