package registry

import "jinshiho/domain/repo"

type Registry interface {
	GetUserRepository() repo.UserRepository
}

type registry struct{}

func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) GetUserRepository() repo.UserRepository {
	return nil
}
