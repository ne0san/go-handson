package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)


func HelloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}

func GetTodoListHandler(c echo.Context) error {
	return c.String(http.StatusOK,"todo1, todo2, todo3")
}

func GetTodoHandler(c echo.Context) error {
	return c.String(http.StatusOK, "todo1")
}

func PostTodoHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Todo created!!")
}

func PutTodoHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Todo 2 updated!!")
}

func DeleteTodoHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Todo 3 deleted!!")
}
