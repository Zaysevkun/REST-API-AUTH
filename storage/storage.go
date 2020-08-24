package storage

// using this interface we can alter witch storage implementation of storage we want to use
type Storage interface {
	User() UserRepository
}
