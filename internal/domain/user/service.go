package user

import (
	"context"
	"os/user"
)

type UserServiceInterface interface {
	getUser(c context.Context) ([]user.User, error)
}
type UserService struct {
	userRepository *UserRepository
}

func (u *UserService) getUser(c context.Context) ([]user.User, error) {
	return u.userRepository.getUsers()
}
