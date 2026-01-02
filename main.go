package main

import (
	handlers "go-handson/handlers"
	"github.com/labstack/echo/v4"
)


func main() {
	
	e := echo.New()

	e.GET("/hello/:id", handlers.HelloHandler)
	e.GET("/todos", handlers.GetTodoListHandler)
	e.GET("/todos/:id", handlers.GetTodoHandler)
	e.POST("/todos", handlers.PostTodoHandler)
	e.PUT("/todos/:id", handlers.PutTodoHandler)
	e.DELETE("/todos/:id", handlers.DeleteTodoHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
