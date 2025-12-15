package main

import (
	"github.com/labstack/echo/v4"
	handlers "go-handson/handlers"
)

func main() {
	e := echo.New()

	e.GET("/hello", handlers.HelloHandler)
	e.GET("/todos/list", handlers.GetTodoListHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

