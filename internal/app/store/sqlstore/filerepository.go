package sqlstore

import (
	"Diplom/internal/app/model"
)

type FileRepository struct {
	store *Store
}

func (r *FileRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}
	result, err := r.store.db.Exec(
		"INSERT INTO user_content_db.files (email, encrypted_password) VALUES (?, ?) ",
		u.Email,
		u.EncryptionPassword,
	)
	id, _ := result.LastInsertId()
	u.ID = int(id)
	return err
}

//func (r *FileRepository) FindByEmail(email string) (*model.User, error) {
//	u := &model.User{}
//
//	if err := r.store.db.QueryRow(
//		"SELECT id, email, encrypted_password "+
//			"FROM user_content_db.users "+
//			"WHERE email = ?", email).Scan(&u.ID, &u.Email, &u.Password); err != nil {
//
//		if err == sql.ErrNoRows {
//			return nil, store.ErrRecordNotFound
//		}
//
//		return nil, err
//	}
//
//	return u, nil
//}
