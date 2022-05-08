package sqlstore

import (
	"Diplom/internal/app/model"
)

type GroupRepository struct {
	store *Store
}

func (r *GroupRepository) Create(u *model.Group) error {
	result, err := r.store.db.Exec(
		"INSERT INTO user_content_db.group (group_name, group_owner) VALUES (?, ?) ",
		u.GroupName,
		u.GroupOwner,
	)
	id, _ := result.LastInsertId()
	u.ID = int(id)
	return err
}

func (r *GroupRepository) isOwnerGroup(userID int, groupName string) (bool, error) {
	var ownerExist bool
	row := r.store.db.QueryRow(
		"SELECT EXISTS (SELECT * FROM user_content_db.group WHERE group_owner = ? AND group_name = ?)",
		userID,
		groupName)

	err := row.Scan(&ownerExist)

	if err != nil {
		return false, err
	}

	return ownerExist, nil
}

//func (r *GroupRepository) Delete()
