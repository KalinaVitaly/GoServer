package sqlstore

import (
	"Diplom/internal/app/store"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	db                  *sql.DB
	userRepository      *UserRepository
	fileRepository      *FileRepository
	groupRepository     *GroupRepository
	userGroupRepository *UserGroupRepository
	fileGroupRepository *FileGroupRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}

func (s *Store) Group() store.GroupRepository {
	if s.userRepository != nil {
		return s.groupRepository
	}

	s.groupRepository = &GroupRepository{
		store: s,
	}
	return s.groupRepository
}

func (s *Store) File() store.FileRepository {
	if s.fileRepository != nil {
		return s.fileRepository
	}

	s.fileRepository = &FileRepository{
		store: s,
	}
	return s.fileRepository
}

func (s *Store) UserInGroup() store.UserGroupRepository {
	if s.userGroupRepository != nil {
		return s.userGroupRepository
	}

	s.userGroupRepository = &UserGroupRepository{
		store: s,
	}
	return s.userGroupRepository
}

func (s *Store) FileInGroup() store.FileGroupRepository {
	if s.fileGroupRepository != nil {
		return s.fileGroupRepository
	}

	s.fileGroupRepository = &FileGroupRepository{
		store: s,
	}
	return s.fileGroupRepository
}
