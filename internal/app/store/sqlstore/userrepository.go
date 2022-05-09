package sqlstore

import (
	"Diplom/internal/app/model"
	"Diplom/internal/app/store"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	fmt.Println("Create user", *u)
	if err := u.Validate(); err != nil {
		return err
	}
	fmt.Println("Correct validate")
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	fmt.Println("Correct before create")
	result, err := r.store.db.Exec(
		"INSERT INTO user_content_db.users (email, encrypted_password) VALUES (?, ?) ",
		u.Email,
		u.EncryptionPassword,
	)
	fmt.Println("Correct insert", err, len(u.EncryptionPassword))
	id, err := result.LastInsertId()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Correct get new id", id)
	u.ID = int(id)
	fmt.Println("Correct value ")
	return err
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}

	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password "+
			"FROM user_content_db.users "+
			"WHERE email = ?", email).Scan(&u.ID, &u.Email, &u.Password); err != nil {

		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}

func (r *UserRepository) Delete(user *model.User) error {
	if _, err := r.store.db.Exec(
		"DELETE FROM user_content_db.users WHERE email = ? AND encrypted_password = ?;",
		user.Email,
		user.EncryptionPassword); err != nil {
		return err
	}

	return nil
}
