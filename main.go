package main

import (
	"github.com/labstack/echo/v4"
	handlers "go-handson/handlers"
)

func main() {
	e := echo.New()

	e.GET("/hello", handlers.HelloHandler)
	e.GET("/todos/list", handlers.GetTodoListHandler)
	e.GET("/todos/1", handlers.GetTodoHandler)
	e.POST("/todo/", handlers.PostTodoHandler)
	e.PUT("/todo/2", handlers.PutTodoHandler)
	e.DELETE("/todo/3", handlers.DeleteTodoHandler)	

	e.Logger.Fatal(e.Start(":8080"))
}

