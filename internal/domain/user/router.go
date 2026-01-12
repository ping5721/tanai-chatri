package user

import (
	response "backend/internal/infrastructure/http"
	"github.com/labstack/echo/v4"
	"net/http"
)

func NewUserRouter(r *echo.Echo, userService *UserService) {

	userGroup := r.Group("/users")
	
	// GetUser godoc
	// @Summary Get user
	// @Tags user
	// @Produce json
	// @Param id path string true "User ID"
	// @Success 200 {object} UserResponse
	// @Failure 404 {object} map[string]string
	// @Router /user/{id} [get]
	userGroup.GET("/", func(c echo.Context) error {
		users, err := userService.getUsers(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusNotFound, nil)
		}
		return response.Success(c, users)
	})
}
