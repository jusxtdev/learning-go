package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	Health(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 ok, got %d", w.Code)
	}
}

func TestHealthWrongMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	w := httptest.NewRecorder()

	Health(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405 Method not allowed, got %d", w.Code)
	}
}

func getMockHandler(t *testing.T) TodoHandler {
	t.Helper()
	store := MemoryStore{
		Todos: []Todo{
			{Id: 1, Title: "test 1", Done: false},
			{Id: 2, Title: "test 2", Done: true},
			{Id: 3, Title: "test 3", Done: true},
			{Id: 4, Title: "test 4", Done: false},
		},
	}
	handler := TodoHandler{store: &store}
	return handler
}

func TestGETTodo(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	w := httptest.NewRecorder()
	h := getMockHandler(t)

	h.GETTodos(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 ok, got %d", w.Code)
	}
}

func TestGETTodoWrongMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/todos", nil)
	w := httptest.NewRecorder()
	h := getMockHandler(t)

	h.GETTodos(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405 Method not allowed, got %d", w.Code)
	}
}

func TestGETTodoById(t *testing.T) {
	h := getMockHandler(t)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /todos/{id}", h.GETTodoById)

	req := httptest.NewRequest(http.MethodGet, "/todos/1", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	type GetByIdResponse struct {
		Msg  string `json:"msg"`
		Data Todo   `json:"data"`
	}

	var got GetByIdResponse
	err := json.NewDecoder(w.Body).Decode(&got)
	if err != nil {
		t.Fatal(err)
	}

	if got.Data.Id != 1 {
		t.Fatalf("expected todo id 1, got %d", got.Data.Id)
	}
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 ok, got %d", w.Code)
	}
}

func TestGETTodoByIdWrongMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/todos/1", nil)
	w := httptest.NewRecorder()
	h := getMockHandler(t)

	h.GETTodoById(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405 Method not allowed, got %d", w.Code)
	}
}

func TestGETTodoByIdInvalidId(t *testing.T) {
	h := getMockHandler(t)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /todos/{id}", h.GETTodoById)

	req := httptest.NewRequest(http.MethodGet, "/todos/999", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404 not found, got %d", w.Code)
	}
}

func TestPOSTTodo(t *testing.T) {
	h := getMockHandler(t)

	reqBody := `{"title" : "test title"}`

	req := httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(reqBody))
	w := httptest.NewRecorder()

	req.Header.Set("Content-Type", "application/json")

	h.POSTTodo(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201 created, got %d", w.Code)
	}
}

func TestPOSTTodoInvalidReqBody(t *testing.T) {
	h := getMockHandler(t)

	reqBody := `{"titleWrondg" : "test title"}`

	req := httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(reqBody))
	w := httptest.NewRecorder()

	req.Header.Set("Content-Type", "application/json")

	h.POSTTodo(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 422 created, got %d", w.Code)
	}
}

func TestPOSTTodoByIdWrongMethod(t *testing.T) {
	h := getMockHandler(t)
	reqBody := `{"title" : "test title"}`

	req := httptest.NewRequest(http.MethodGet, "/todos/1", strings.NewReader(reqBody))
	w := httptest.NewRecorder()

	h.POSTTodo(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405 Method not allowed, got %d", w.Code)
	}
}

func TestPUTTodo(t *testing.T) {
	h := getMockHandler(t)
	reqBody := `{"done" : true}`

	mux := http.NewServeMux()
	mux.HandleFunc("PUT /todos/{id}", h.PUTTodo)

	req := httptest.NewRequest(http.MethodPut, "/todos/1", strings.NewReader(reqBody))
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 ok, got %d", w.Code)
	}
}

func TestPUTTodoIdNotFound(t *testing.T) {
	h := getMockHandler(t)
	reqBody := `{"done" : true}`

	mux := http.NewServeMux()
	mux.HandleFunc("PUT /todos/{id}", h.PUTTodo)

	req := httptest.NewRequest(http.MethodPut, "/todos/999", strings.NewReader(reqBody))
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404 not found, got %d", w.Code)
	}
}

func TestPUTTodoIdInvalid(t *testing.T) {
	h := getMockHandler(t)
	reqBody := `{"done" : true}`

	mux := http.NewServeMux()
	mux.HandleFunc("PUT /todos/{id}", h.PUTTodo)

	req := httptest.NewRequest(http.MethodPut, "/todos/abd", strings.NewReader(reqBody))
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusConflict {
		t.Fatalf("expected 409 conflict, got %d", w.Code)
	}
}

func TestPUTTodoByIdWrongMethod(t *testing.T) {
	h := getMockHandler(t)
	reqBody := `{"done" : true}`

	mux := http.NewServeMux()
	mux.HandleFunc("PUT /todos/{id}", h.PUTTodo)

	req := httptest.NewRequest(http.MethodGet, "/todos/1", strings.NewReader(reqBody))
	w := httptest.NewRecorder()

	h.POSTTodo(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405 Method not allowed, got %d", w.Code)
	}
}

func TestDELTodo(t *testing.T) {
	h := getMockHandler(t)

	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /todos/{id}", h.DELTodo)

	req := httptest.NewRequest(http.MethodDelete, "/todos/1", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatalf("expected 204 no content, got %d", w.Code)
	}
}

func TestDELTodoNotFound(t *testing.T) {
	h := getMockHandler(t)

	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /todos/{id}", h.DELTodo)

	req := httptest.NewRequest(http.MethodDelete, "/todos/999", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404 not found, got %d", w.Code)
	}
}

func TestDELTodoIdInvalid(t *testing.T) {
	h := getMockHandler(t)

	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /todos/{id}", h.DELTodo)

	req := httptest.NewRequest(http.MethodDelete, "/todos/ads", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusConflict {
		t.Fatalf("expected 409 confilct, got %d", w.Code)
	}
}

func TestDELTodoWrongMethod(t *testing.T) {
	h := getMockHandler(t)

	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /todos/{id}", h.DELTodo)

	req := httptest.NewRequest(http.MethodGet, "/todos/1", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405 method not allowed, got %d", w.Code)
	}
}
