package teststore

import (
	"Diplom/internal/app/model"
	"Diplom/internal/app/store"
	"fmt"
)

type GroupRepository struct {
	store  *Store
	groups map[int]*model.Group
}

func (r *GroupRepository) Create(u *model.Group) error {

	for _, value := range r.groups {
		if value.GroupName == u.GroupName {
			return store.ErrInputValues
		}
	}

	u.ID = len(r.groups) + 1
	r.groups[u.ID] = u

	return nil
}

func (r *GroupRepository) Delete(userID int, groupName string) error {
	groupModel, err := r.FindGroupByName(groupName)

	if err != nil {
		return store.ErrRecordNotFound
	}

	fmt.Println(userID, groupModel.GroupOwner)
	if userID != groupModel.GroupOwner {
		fmt.Println("Invalid user id")
		return store.ErrAccessRights
	}

	if _, isOk := r.groups[groupModel.ID]; isOk {
		delete(r.groups, groupModel.ID)
		return nil
	}

	return store.ErrRecordNotFound
}

func (r *GroupRepository) FindGroupByName(groupName string) (*model.Group, error) {
	for _, value := range r.groups {
		if value.GroupName == groupName {
			return value, nil
		}
	}
	return nil, store.ErrRecordNotFound
}
