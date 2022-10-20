package routes

import (
	"final-projek/app/middlewares"
	"final-projek/controllers/ccomments"
	"final-projek/controllers/cphotos"
	"final-projek/controllers/csocialmedias"
	cUser "final-projek/controllers/cusers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetRoute(e *echo.Echo) {
	e.GET("/users", cUser.GetAllUser)
	e.POST("/users", cUser.CreateUser)
	e.POST("/users/register", cUser.Register)
	e.POST("/users/login", cUser.Login)
	e.GET("/users/:id", cUser.GetByIDUser)

	e.GET("/socialmedia/:id", csocialmedias.GetByIDSocialMedia)

	e.GET("/photo/:id", cphotos.GetByIDPhoto)

	e.GET("/comment/:id", ccomments.GetByIDComment)

	e.DELETE("/comment/:id", ccomments.DeleteComment)
	privateRoutes := e.Group("")

	privateRoutes.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secretkey"),
	}))

	privateRoutes.Use(middlewares.CheckTokenMiddleware)
	privateRoutes.PUT("/users/:id", cUser.UpdateUser)
	privateRoutes.DELETE("/users/:id", cUser.DeleteUser)
	privateRoutes.POST("/photo", cphotos.CreatePhoto)
	privateRoutes.GET("/photo", cphotos.GetAllPhoto)
	privateRoutes.PUT("/photo/:id", cphotos.UpdatePhoto)
	privateRoutes.DELETE("/photo/:id", cphotos.DeletePhoto)
	privateRoutes.GET("/comment", ccomments.GetAllComment)
	privateRoutes.PUT("/comment/:id", ccomments.UpdateComment)
	privateRoutes.POST("/comment", ccomments.CreateCommment)
	privateRoutes.POST("/socialmedia", csocialmedias.CreateSocialMedia)
	privateRoutes.PUT("/socialmedia/:id", csocialmedias.UpdateSocialMedia)
	privateRoutes.DELETE("/socialmedia/:id", csocialmedias.DeleteSocialMedia)
	privateRoutes.GET("/socialmedia", csocialmedias.GetAllSocialMedia)
}
