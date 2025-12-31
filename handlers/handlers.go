package handlers

import (
	"strconv"
	"net/http"

	"github.com/labstack/echo/v4"
)


func HelloHandler(c echo.Context) error {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}
	return c.String(http.StatusOK, "Hello, ID:" + strconv.Itoa(id) + "!")
}

func GetTodoListHandler(c echo.Context) error {
	return c.String(http.StatusOK,"todo1, todo2, todo3")
}

func GetTodoHandler(c echo.Context) error {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}
	return c.String(http.StatusOK, "todo:" + strconv.Itoa(id))
}

func PostTodoHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Todo created!!")
}

func PutTodoHandler(c echo.Context) error {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}
	return c.String(http.StatusOK, "Todo " + strconv.Itoa(id) + " updated!!")
}

func DeleteTodoHandler(c echo.Context) error {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}
	return c.String(http.StatusOK, "Todo " + strconv.Itoa(id) + " deleted!!")
}
