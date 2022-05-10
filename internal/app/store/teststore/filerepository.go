package teststore

import (
	"Diplom/internal/app/model"
)

type FileRepository struct {
	store *Store
	files map[int]*model.File
}

//type FileRepository interface {
//	Create(*model.File) error
//	GetFilePath(string) (string, error)
//	Delete(string, int) error
//	UpdateAvailableFile(string, bool) error
//	FindByQuery(fileQuery string) (*model.File, error)
//}

// Create ...
func (r *FileRepository) Create(u *model.File) error {
	u.ID = len(r.files) + 1
	r.files[u.ID] = u

	return nil
}

func (r *FileRepository) GetFilePath(string) (string, error) {
	//u, ok := r.files[id]
	//if !ok {
	//	return nil, store.ErrRecordNotFound
	//}

	return "u", nil
}

func (r *FileRepository) Delete(string, int) error {

	return nil
}

func (r *FileRepository) UpdateAvailableFile(string, bool) error {
	return nil
}

func (r *FileRepository) FindByQuery(fileQuery string) (*model.File, error) {
	return nil, nil
}
