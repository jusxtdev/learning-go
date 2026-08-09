package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
)

/*
CLI TO-DO APP

SCHEMA:
	id : int
	title : string
	done : bool

IN MEMORY:
	a single struct

STORAGE:
	json file

FLAGS:
	-add "buy milk"
	-list
	-done <todo_id>
	-delete <todo_id>
	-help
*/

type Todo struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos []Todo

var filename = "data.json"

func main() {

	create_file(filename)

	read_file_in_memory(filename)

	title := flag.String("add", "", "Title of todo")
	list := flag.Bool("list", false, "List all todos")
	mark_done_todoid := flag.Int("done", 0, "Mark todo as done")
	delete_todoid := flag.Int("delete", 0, "Delete todo")

	flag.Parse()

	if *title != "" {
		add_todo(*title)
	}
	if *list {
		list_todos()
	}
	if *mark_done_todoid != 0 {
		mark_as_done(*mark_done_todoid)
	}
	if *delete_todoid != 0 {
		delete(*delete_todoid)
	}
	write_to_file(filename, todos)
}

/* COMMANDS */

func add_todo(title string) {
	id := get_new_id()

	new_todo := Todo{
		Id:    id,
		Title: title,
		Done:  false,
	}

	todos = append(todos, new_todo)
}

func list_todos() {
	for _, todo := range todos {
		status := "[ ]"
		if todo.Done {
			status = "[x]"
		}
		fmt.Printf("%-4d %s %s\n", todo.Id, status, todo.Title)
	}
}

func mark_as_done(id int) {
	for i, todo := range todos {
		if todo.Id == id {
			todos[i].Done = !todo.Done
		}
	}
}

func delete(id int) {
	var delete_index int
	for i, todo := range todos {
		if todo.Id == id {
			delete_index = i
		}
	}

	todos = slices.Delete(todos, delete_index, delete_index+1)
}

/* FILE OPERATIONS */
func read_file_in_memory(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &todos)
	if err != nil {
		log.Fatal(err)
	}
}

func create_file(filename string) {
	// If the file doesn't exist, create it, or append to the file
	var file *os.File
	var err error

	// return if file already exists
	if _, err = os.Stat(filename); !os.IsNotExist(err) {
		return
	}

	// create new file and
	file, err = os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	// initialize it with '[]'
	_, err = file.Write([]byte("[]"))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func write_to_file(filename string, todos []Todo) {
	data, err := json.MarshalIndent(todos, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

/* HELPERS */
func get_new_id() int {
	if len(todos) == 0 {
		return 1
	}
	last_todo := todos[len(todos)-1]
	return last_todo.Id + 1
}
