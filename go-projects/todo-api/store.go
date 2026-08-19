package main

var Todos []Todo

var TodoStore interface {
	GetAllTodos() []Todo
	GetTodoById(id int) (Todo, error)
	AddTodo(title string) Todo
	UpdateTodo(id int, newTitle string, done bool) (Todo, error)
	DeleteTodo(id int) error
}
