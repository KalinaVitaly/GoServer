package sqlstore

type FileGroupRepository struct {
	store *Store
}

func (r *FileGroupRepository) AddFileInGroup(fileId, groupId int) error {

	_, err := r.store.db.Exec(
		"INSERT INTO user_content_db.files_in_group (file_id, file_group_id) VALUES (?, ?) ",
		fileId,
		groupId,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *FileGroupRepository) DeleteFileFromGroup(fileId, groupId int) error {
	if _, err := r.store.db.Exec(
		"DELETE FROM user_content_db.files_in_groups WHERE file_id = ? AND file_group_id = ?;",
		fileId,
		groupId); err != nil {
		return err
	}

	return nil
}
