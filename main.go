package main

import (
	"backend/ent"
	"backend/internal/domain/user"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	client, err := ent.Open("postgres", "postgres://f49e2bf29c4bc8a564588af0b6f7ae1c8e80da2d47a882855d65f55a382f7d8e:sk_t3l6zuh-8gG5tODdYEhEW@db.prisma.io:5432/postgres?sslmode=require")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "", // no password set
	// 	DB:       0,  // use default DB
	// })
	// defer rdb.Close()

	if err != nil {
		panic(e)
	}

	userRepository := user.UserRepository(client)
	userService := user.UserService(userRepository)
	user.NewUserRouter(e, userService)

	e.Logger.Fatal(e.Start(":8080"))
}
