package cusers

import (
	"final-projek/controllers"
	"final-projek/models/musers"
	sUser "final-projek/service/susers"
	"net/http"

	"github.com/labstack/echo/v4"
)

var userService sUser.UserService = sUser.NewUserService()

func GetAllUser(c echo.Context) error {
	users := userService.GetAllUser()
	res := controllers.NewResponse(c, http.StatusOK, "success", "Get All Users", users)
	return res

}

func GetByIDUser(c echo.Context) error {
	userID := c.Param("id")

	user := userService.GetByIDUser(userID)
	if user.ID == 0 {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "GetByID", "")
	} else {
		return controllers.NewResponse(c, http.StatusOK, "Success", "GetByID", user)
	}

}

func CreateUser(c echo.Context) error {
	input := new(musers.InputUser)
	if err := c.Bind(input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "CreatedUser", "")
	}

	user := userService.CreateUser(*input)

	return controllers.NewResponse(c, http.StatusCreated, "success", "CreateUser", user)
}


func UpdateUser(c echo.Context) error {
	input := new(musers.InputUser)

	if err := c.Bind(input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "UpdateUser", "")
	}

	userID := c.Param("id")

	user := userService.UpdateUser(userID, *input)

	return controllers.NewResponse(c, http.StatusCreated, "success", "your account has been succesfully updated", user)
}

func DeleteUser(c echo.Context) error {
	userID := c.Param("id")

	dUser := userService.DeleteUser(userID)

	if !dUser {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "your account has been failed deleted", "")
	} else {
		return controllers.NewResponse(c, http.StatusOK, "success", "your account has been succesfully deleted", "")
	}
}
