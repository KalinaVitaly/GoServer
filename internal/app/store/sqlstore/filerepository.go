package sqlstore

import (
	"Diplom/internal/app/model"
	"Diplom/internal/app/store"
	"database/sql"
)

type FileRepository struct {
	store *Store
}

func (r *FileRepository) GetFilePath(fileQuery string) (string, error) {
	var filePath string
	row := r.store.db.QueryRow("SELECT file_path FROM user_content_db.files WHERE file_query = ?)", fileQuery)

	err := row.Scan(&filePath)

	if err != nil {
		return "", err
	}

	return filePath, nil
}

func (r *FileRepository) Create(u *model.File) error {
	result, err := r.store.db.Exec(
		"INSERT INTO user_content_db.files ("+
			"file_owner, "+
			"file_path, "+
			"file_name, "+
			"file_query, "+
			"file_available) "+
			"VALUES (?, ?, ?, ?, ?) ",
		u.FileOwner,
		u.FilePath,
		u.FileName,
		u.FileQuery,
		u.FileAvailable,
	)
	id, _ := result.LastInsertId()
	u.ID = int(id)
	return err
}

func (r *FileRepository) UpdateAvailableFile(filePath string, fileAvailable bool) error {
	_, err := r.store.db.Exec(`UPDATE user_content_db.files
									SET file_available = ?
									WHERE file_path = ?`,
		fileAvailable,
		filePath)

	return err
}

func (r *FileRepository) isOwner(fileQuery string, userID int) (bool, error) {
	var ownerExist bool
	row := r.store.db.QueryRow("SELECT EXISTS (SELECT * FROM user_content_db.files WHERE file_owner = ? AND file_query = ?)", userID, fileQuery)

	err := row.Scan(&ownerExist)

	if err != nil {
		return false, err
	}

	return ownerExist, nil
}

func (r *FileRepository) Delete(fileQuery string, userID int) error {
	is_ok, err := r.isOwner(fileQuery, userID)
	if err != nil {
		return err
	} else if !is_ok {
		return store.ErrAccessRights
	}

	result, err := r.store.db.Exec("DELETE FROM user_content_db.files WHERE file_query = ?", fileQuery)
	if err != nil {
		return err
	}

	if rowsAffected, err := result.RowsAffected(); err != nil || rowsAffected == 0 {
		if err != nil {
			return err
		}

		return store.ErrFoundFile
	}

	return nil
}

func (r *FileRepository) FindByQuery(fileQuery string) (*model.File, error) {
	u := &model.File{}

	if err := r.store.db.QueryRow(
		"SELECT id, file_name, file_path, file_owner, file_available"+
			"FROM user_content_db.files "+
			"WHERE file_query = ?", fileQuery).Scan(&u.ID, &u.FileName, &u.FilePath, &u.FileOwner, &u.FileAvailable); err != nil {

		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}
