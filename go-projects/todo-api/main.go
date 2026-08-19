package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	// in memory store
	store := MemoryStore{}
	store.Seed()
	h := TodoHandler{store: &store}

	mux := http.NewServeMux()

	// root handler
	mux.HandleFunc("GET /", Health)

	// todos handler
	mux.HandleFunc("GET /todos", h.GETTodos)
	mux.HandleFunc("GET /todos/", h.GETTodos)

	mux.HandleFunc("GET /todos/{id}", h.GETTodoById)
	mux.HandleFunc("GET /todos/{id}/", h.GETTodoById)

	mux.HandleFunc("POST /todos", h.POSTTodo)
	mux.HandleFunc("POST /todos/", h.POSTTodo)

	mux.HandleFunc("PUT /todos/{id}", h.PUTTodo)
	mux.HandleFunc("PUT /todos/{id}/", h.PUTTodo)

	mux.HandleFunc("DELETE /todos/{id}", h.DELTodo)
	mux.HandleFunc("DELETE /todos/{id}/", h.DELTodo)

	err := http.ListenAndServe(":8080", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server Closed")
	} else if err != nil {
		fmt.Printf("Server Crashed!\nError : %s\n", err)
	}
}
