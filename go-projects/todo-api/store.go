package main

import "sync"

type MemoryStore struct {
	mu    sync.RWMutex
	Todos []Todo
}

type TodoStore interface {
	GetAllTodos(DoneFilter bool, DoneValue bool, title string) []Todo
	GetTodoById(id int) (Todo, error)
	AddTodo(title string) Todo
	UpdateTodo(id int, newTitle string, done bool) (Todo, error)
	DeleteTodo(id int) error
}
