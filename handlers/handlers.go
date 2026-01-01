package handlers

import (
	"strconv"
	"net/http"

	"github.com/labstack/echo/v4"
	"time"
)

type GetTodo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateTodo struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

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
	todo := GetTodo{
		ID:        1,
		Title:     "Sample Todo",
		Content:   "This is a sample task",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// JSONレスポンスとして返す
	return c.JSON(http.StatusOK, todo)
}

func PostTodoHandler(c echo.Context) error {
	var todo CreateTodo

	// JSONデータを構造体にバインド
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON format"})
	}

	// デコードされたデータをレスポンスとして返す
	return c.JSON(http.StatusOK, todo)
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
