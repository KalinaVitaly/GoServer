package store

import "Diplom/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	Delete(user *model.User) error
}

type GroupRepository interface {
	Create(u *model.Group) error
	Delete(int, string) error
}

type FileRepository interface {
	Create(*model.File) error
	GetFilePath(string) (string, error)
	Delete(string, int) error
	UpdateAvailableFile(string, bool) error
}
