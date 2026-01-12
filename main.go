package main

import (
	"backend/internal/domain/user"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer rdb.Close()

	if err != nil {
		panic(e)
	}

	userRepository := user.UserRepository(*db, rdb)
	userService := user.UserService(*userRepository)
	user.NewUserRouter(e, userService)

	e.Logger.Fatal(e.Start(":8080"))
}
