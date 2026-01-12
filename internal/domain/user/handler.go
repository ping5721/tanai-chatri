package user

import  (
	response "backend/internal/infrastructure/http"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct{
	UserService *UserService
}

func (userHandler *UserHandler) getUsers(e echo.Context) error {
		users, err := userHandler.UserService.getUsers(e.Request().Context())
		if err != nil {
			return e.JSON(http.StatusNotFound, nil)
		}
		return response.Success(e, users)
}