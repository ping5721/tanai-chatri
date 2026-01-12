package user

import (
	"context"
)

type UserServiceInterface interface {
	getUsers(ctx context.Context) ([]User, error)
}
type UserService struct {
	userRepository *UserRepository
}

func (u *UserService) getUsers(ctx context.Context) ([]User, error) {
	return u.userRepository.getUsers(ctx)
}
