package sqlstore

type UserGroupRepository struct {
	store *Store
}

func (r *UserGroupRepository) AddUserInGroup(userId, groupId int) error {

	_, err := r.store.db.Exec(
		"INSERT INTO user_content_db.users_in_groups (user_id, group_id) VALUES (?, ?) ",
		userId,
		groupId,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserGroupRepository) DeleteUserFromGroup(userId, groupId int) error {
	if _, err := r.store.db.Exec(
		"DELETE FROM user_content_db.users_in_groups WHERE user_id = ? AND group_id = ?;",
		userId,
		groupId); err != nil {
		return err
	}

	return nil
}
