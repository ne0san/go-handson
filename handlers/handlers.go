package handlers

import (
	"io"
	"fmt"
	"net/http"
)


func HelloHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {

		io.WriteString(w, "Hello, world!")
		fmt.Println("request: ", req.Method)

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


func GetTodoListHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "todo1, todo2, todo3")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


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
