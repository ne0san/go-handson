package main

import (
	"log"
	"net/http"
	handlers "go-handson/handlers"
)

func main() {
	http.HandleFunc("/hello", handlers.HelloHandler) // hello/をhelloHandlerにハンドリング
	http.HandleFunc("/todos/list", handlers.GetTodoListHandler)
	http.HandleFunc("/todos/1", handlers.GetTodoHandler)
	http.HandleFunc("/todos", handlers.PostTodoHandler)
	http.HandleFunc("/todos/2", handlers.PutTodoHandler)
	http.HandleFunc("/todos/3", handlers.DeleteTodoHandler)
	log.Println("server start at port 8080")
	err := http.ListenAndServe(":8080", nil) // 8080で待機
	if err != nil {
		log.Fatal(err)
	}
}

