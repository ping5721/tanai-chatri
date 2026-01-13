package user

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	// Db      *ent.Client
	MongoDb *mongo.Client
	// redis *redis.Client
}

func (userRepository *UserRepository) getUsers(ctx context.Context) ([]User, error) {
	collection := userRepository.MongoDb.Database("tanai-chatri").Collection("users")
	cursor, err := collection.Find(ctx, map[string]interface{}{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []User
	for cursor.Next(ctx) {
		var u struct {
			ID        int       `bson:"id"`
			Age       int       `bson:"age"`
			Name      string    `bson:"name"`
			Username  string    `bson:"username"`
			CreatedAt time.Time `bson:"created_at"`
			Premium   bool      `bson:"premium"`
		}
		if err := cursor.Decode(&u); err != nil {
			fmt.Println(err)
			continue
		}
		result = append(result, User{
			ID:        u.ID,
			Age:       u.Age,
			Name:      u.Name,
			Username:  u.Username,
			CreatedAt: u.CreatedAt,
			Premium:   u.Premium,
		})
	}
	if err := cursor.Err(); err != nil {
		return nil, err
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
