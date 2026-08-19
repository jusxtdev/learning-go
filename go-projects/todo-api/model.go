package main

import (
	"errors"
	"slices"
)

type Todo struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func GetAllTodos() []Todo {
	return Todos
}

func GetTodoById(id int) (Todo, error) {
	for _, v := range Todos {
		if v.Id == id {
			return v, nil
		}
	}
	return Todo{}, errors.New("todo not found")
}

func AddTodo(title string) Todo {
	// create new todo
	t := Todo{
		Id:    getNewTodoId(),
		Title: title,
		Done:  false,
	}
	Todos = append(Todos, t)

	return t
}

func UpdateTodo(id int, newTitle string, done bool) (Todo, error) {
	// get index of todo to update
	index, err := getIndexById(id)
	if err != nil {
		return Todo{}, err
	}

	// update todo based on fields given
	if newTitle != "" {
		Todos[index].Title = newTitle
	}
	Todos[index].Done = done

	// respond
	return Todos[index], nil
}

func DeleteTodo(id int) error {
	index, err := getIndexById(id)
	if err != nil {
		return err
	}

	// delete
	Todos = slices.Delete(Todos, index, index+1)
	return nil
}

/* HELPER */
func getNewTodoId() int {
	if len(Todos) == 0 {
		return 1
	}
	last_todo := Todos[len(Todos)-1]
	return last_todo.Id + 1
}

func getIndexById(id int) (int, error) {
	// using index based for because in `for i, v := range Todos`,
	// v is a copy of todo hence changing v doesn't change original todo
	for index := range Todos {
		if Todos[index].Id == id {
			return index, nil
		}
	}
	return 0, errors.New("not found")
}
