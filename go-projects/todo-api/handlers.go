package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Response struct {
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

/* GET / */
func Health(w http.ResponseWriter, r *http.Request) {
	// check method
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

	response := Response{
		Msg:  "Alive",
		Data: nil,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		fmt.Println(err)
	}
}

/* GET /todos  */
func GETTodos(w http.ResponseWriter, r *http.Request) {
	// check method
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

	// get all todos
	all_todo := GetAllTodos()
	data := map[string]any{
		"count": len(all_todo),
		"todos": all_todo,
	}

	// respond with todo, 200
	response := Response{
		Msg:  "All todos",
		Data: data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		fmt.Println(err)
	}
}

/* GET /todos/{id} */
func GETTodoById(w http.ResponseWriter, r *http.Request) {
	// check method
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

	// extract id path parameter, 409 if invalid
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusConflict)
		return
	}

	// get todo by id, 404 if not found
	todo, err := GetTodoById(id)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	// respond with todo, 200
	response := Response{
		Msg:  "Found todo",
		Data: todo,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		fmt.Println(err)
	}
}

/* POST /todos */
type POSTRequestBody struct {
	Title string `json:"title"`
}

func POSTTodo(w http.ResponseWriter, r *http.Request) {
	// check method
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

	// decode request body
	var body POSTRequestBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Failed to decode response", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// add todo
	newTodo := AddTodo(body.Title)

	// resond, 201
	response := Response{
		Msg:  "Todo added successfully",
		Data: newTodo,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		fmt.Println(err)
	}
}

/* PUT /todos/{id} */
type PUTRequestBody struct {
	Title *string `json:"title"`
	Done  *bool   `json:"done"`
}

func PUTTodo(w http.ResponseWriter, r *http.Request) {
	// check method
	if r.Method != http.MethodPut {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

	// extract todo id from path, 409 on Invalid id
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusConflict)
		return
	}

	// decode body
	var body PUTRequestBody
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Failed to decode response", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// update todo, 404 if not found
	var title string
	var done bool

	if body.Title != nil {
		title = *body.Title
	}
	if body.Done != nil {
		done = true
	}

	updated, err := UpdateTodo(id, title, done)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	// respond, 200
	response := Response{
		Msg:  "Updated successfully",
		Data: updated,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		fmt.Println(err)
	}
}

/* DELETE /todos/{id} */
func DELTodo(w http.ResponseWriter, r *http.Request) {
	// check method
	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

	// extract id
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusConflict)
		return
	}

	// delete todo
	err = DeleteTodo(id)
	if err != nil {
		http.Error(w, "todo not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
