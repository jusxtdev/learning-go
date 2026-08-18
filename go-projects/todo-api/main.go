package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// root handler
	mux.HandleFunc("GET /", Health)

	// todos handler
	mux.HandleFunc("GET /todos", GETTodos)
	mux.HandleFunc("GET /todos/", GETTodos)

	mux.HandleFunc("GET /todos/{id}", GETTodoById)
	mux.HandleFunc("GET /todos/{id}/", GETTodoById)

	mux.HandleFunc("POST /todos", POSTTodo)
	mux.HandleFunc("POST /todos/", POSTTodo)

	mux.HandleFunc("PUT /todos/{id}", PUTTodo)
	mux.HandleFunc("PUT /todos/{id}/", PUTTodo)

	mux.HandleFunc("DELETE /todos/{id}", DELTodo)
	mux.HandleFunc("DELETE /todos/{id}/", DELTodo)

	err := http.ListenAndServe(":8080", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server Closed")
	} else if err != nil {
		fmt.Printf("Server Crashed!\nError : %s\n", err)
	}
}
