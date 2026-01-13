package main

import (
	"backend/ent"
	"backend/internal/domain/user"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/

//	@securityDefinitions.basic	NoAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	client, err := ent.Open("postgres", "postgres://f49e2bf29c4bc8a564588af0b6f7ae1c8e80da2d47a882855d65f55a382f7d8e:sk_t3l6zuh-8gG5tODdYEhEW@db.prisma.io:5432/postgres?sslmode=require")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer client.Close()

	// MongoDb Connection
	uri := os.Getenv("MONGODB_URI")
	mongoDb, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "", // no password set
	// 	DB:       0,  // use default DB
	// })
	// defer rdb.Close()

	e.GET("healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hihi")
	})

	if err != nil {
		panic(e)
	}

	userRepository := user.UserRepository{MongoDb: mongoDb}
	userService := user.UserService{UserRepository: &userRepository}
	userHandler := user.UserHandler{UserService: &userService}
	user.NewUserRouter(e, &userHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
