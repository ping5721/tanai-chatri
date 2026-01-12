package user

import (
	"context"
	"encoding/json"
	"os/user"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func (userRepository *UserRepository) getUsers() ([]user.User, error) {
	users := []user.User{}
	result := userRepository.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (userRepository *UserRepository) getUser(id int, ctx context.Context) (user.User, error) {
	userCache := userRepository.redis.Get(ctx, "user:1")
	user := user.User{}
	if userCache.Err() == nil {
		err := json.Unmarshal([]byte(userCache.Val()), &user)
		if err != nil {
			return user, err
		}
		return user, nil
	}
	result := userRepository.db.First(&user, id)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
