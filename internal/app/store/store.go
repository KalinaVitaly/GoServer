package store

type Store interface {
	User() UserRepository
	Group() GroupRepository
	File() FileRepository
}
