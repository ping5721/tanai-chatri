package user

import (
	"backend/ent"
	"context"
)

type UserRepository struct {
	Db *ent.Client
	// redis *redis.Client
}

func (userRepository *UserRepository) getUsers(ctx context.Context) ([]User, error) {
	users, err := userRepository.Db.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]User, len(users))
	for i, u := range users {
		result[i] = User{
			ID:        u.ID,
			Age:       u.Age,
			Name:      u.Name,
			Username:  u.Username,
			CreatedAt: u.CreatedAt,
			Premium:   u.Premium,
		}
	}
	return result, nil
}

// func (userRepository *UserRepository) getUser(id int, ctx context.Context) (user.User, error) {
// 	// userCache := userRepository.redis.Get(ctx, cache.RedisKey.UserKey(id))
// 	// user := user.User{}
// 	// if userCache.Err() == nil {
// 	// 	err := json.Unmarshal([]byte(userCache.Val()), &user)
// 	// 	if err != nil {
// 	// 		return user, err
// 	// 	}
// 	// 	return user, nil
// 	// }
// 	result := userRepository.Db.First(&user, id)
// 	if result.Error != nil {
// 		return user, result.Error
// 	}
// 	return user, nil
// }
