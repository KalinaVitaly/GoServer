package teststore

import (
	"Diplom/internal/app/model"
	"Diplom/internal/app/store"
)

// Store ...
type Store struct {
	userRepository      *UserRepository
	fileRepository      *FileRepository
	groupRepository     *GroupRepository
	userGroupRepository *UserGroupRepository
	fileGroupRepository *FileGroupRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}

func (s *Store) File() store.FileRepository {
	if s.fileRepository != nil {
		return s.fileRepository
	}

	s.fileRepository = &FileRepository{
		store: s,
		files: make(map[int]*model.File),
	}

	return s.fileRepository
}

func (s *Store) Group() store.GroupRepository {
	if s.groupRepository != nil {
		return s.groupRepository
	}

	s.groupRepository = &GroupRepository{
		store: s,
		files: make(map[int]*model.Group),
	}

	return s.groupRepository
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
