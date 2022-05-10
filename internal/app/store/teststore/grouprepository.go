package teststore

import "Diplom/internal/app/model"

type GroupRepository struct {
	store *Store
	files map[int]*model.Group
}

//type GroupRepository interface {
//	Create(u *model.Group) error
//	Delete(int, string) error
//	FindGroupByName(string) (*model.Group, error)
//}
func (r *GroupRepository) Create(u *model.Group) error {

	return nil
}

func (r *GroupRepository) Delete(int, string) error {

	return nil
}

func (r *GroupRepository) FindGroupByName(string) (*model.Group, error) {

	return nil, nil
}
