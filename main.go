package main

import (
	"final-projek/app/config"
	route "final-projek/app/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	config.ConnectDB()
	e := echo.New()
	route.SetRoute(e)
	e.Logger.Fatal(e.Start(":1323"))
}
