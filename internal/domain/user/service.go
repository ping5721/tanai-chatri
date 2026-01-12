package user

import (
	"context"
)

type UserServiceInterface interface {
	getUsers(ctx context.Context) ([]User, error)
}
type UserService struct {
	UserRepository *UserRepository
}

func (u *UserService) getUsers(ctx context.Context) ([]User, error) {
	return u.UserRepository.getUsers(ctx)
}
