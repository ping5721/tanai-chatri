package user

import (
	response "backend/internal/infrastructure/http"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewUserRouter(r *echo.Echo, userService *UserService) {
	userGroup := r.Group("/users")
	{
		userGroup.GET("/", func(c echo.Context) error {
			users, err := userService.getUsers(c.Request().Context())
			if err != nil {
				c.JSON(http.StatusNotFound, nil)
			}
			return response.Success(c, users)
		})
	}
}
