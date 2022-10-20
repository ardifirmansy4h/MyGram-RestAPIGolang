package cusers

import (
	"final-projek/controllers"
	"final-projek/models/musers"
	"final-projek/service/susers"
	"net/http"

	"github.com/labstack/echo/v4"
)

var authService susers.AuthService = susers.NewUserAuth()

func Register(c echo.Context) error {
	userRegister := new(musers.InputRegister)
	if err := c.Bind(userRegister); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "Bad Request", "")
	}
	err := userRegister.ValidRegister()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "data tidak boleh kosong, age harus lebih dari 8 dan email harus lengkap", "")
	}

	user := authService.Register(*userRegister)

	return controllers.NewResponse(c, http.StatusOK, "success", "succes register", user)

}

func Login(c echo.Context) error {
	userLogin := new(musers.InputLogin)
	if err := c.Bind(userLogin); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "failed login", "")

	}
	err := userLogin.ValidLogin()
	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "data tidak boleh kosong", "")
	}
	token := authService.Login(*userLogin)
	if token == "" {
		return controllers.NewResponse(c, http.StatusUnauthorized, "failed", "failed created token", "")
	}
	return controllers.NewResponse(c, http.StatusAccepted, "success", "token created", token)

}
