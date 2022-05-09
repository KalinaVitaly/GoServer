package store

type Store interface {
	User() UserRepository
	Group() GroupRepository
}
