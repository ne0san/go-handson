package handlers

import (
	"io"
	"net/http"
	"github.com/labstack/echo/v4"
)


func HelloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}

func GetTodoListHandler(c echo.Context) error {
	return c.String(http.StatusOK,"todo1, todo2, todo3")
}

// TODO すべてのハンドラを書き換え

func GetTodoHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "todo1")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


func PostTodoHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Todo created!!")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


func PutTodoHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPut {
		io.WriteString(w, "Todo 2 updated!!")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func DeleteTodoHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodDelete {
		io.WriteString(w, "Todo 3 deleted!!")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
