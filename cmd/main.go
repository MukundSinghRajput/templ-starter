package main

import (
	"fmt"
	"templ-starter/handlers"
	"templ-starter/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.Use(utils.WithUser)
	userHandler := handlers.UserHandler{}
	var context echo.Context
	app.GET("/user", userHandler.HandlerUserShow(context))
	fmt.Println("Started the server")
	app.Start(":8080")
}
