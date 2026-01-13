package user

import (
	"github.com/labstack/echo/v4"
)

func NewUserRouter(r *echo.Echo, userHandler *UserHandler) {

	userGroup := r.Group("/users")

	// GetUser godoc
	// @Summary Get user
	// @Tags user
	// @Produce json
	// @Param id path string true "User ID"
	// @Success 200 {object} UserResponse
	// @Failure 404 {object} map[string]string
	// @Router /user/{id} [get]
	userGroup.GET("/", userHandler.getUsers)
}
