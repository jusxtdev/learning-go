package main

import "testing"

func newTodoMemoryStore(t *testing.T) *MemoryStore {
	t.Helper()
	s := MemoryStore{}
	s.Seed()
	return &s
}

func TestGetAllTodos(t *testing.T) {
	s := newTodoMemoryStore(t)

	if got, _ := s.GetAllTodos(false, false, "", 0, 0); len(got) != 4 {
		t.Fatalf("expected 4 todos got %d", len(got))
	}
}

func TestGetTodoById(t *testing.T) {
	s := newTodoMemoryStore(t)

	todo, err := s.GetTodoById(1)
	if err != nil {
		t.Fatalf("unexpected error : %s", err)
	}
	if todo.Id != 1 {
		t.Fatalf("expected todo with id 1, got %d", todo.Id)
	}
}

func TestGetTodoByIdNotFound(t *testing.T) {
	s := newTodoMemoryStore(t)

	_, err := s.GetTodoById(999)
	if err == nil {
		t.Fatalf("expected error for missing todo")
	}
}

func TestAddTodo(t *testing.T) {
	s := newTodoMemoryStore(t)

	testTitle := "test_title"
	s.AddTodo(testTitle)

	// as 4 todos are already seeded, the new todo id should be 5
	todo, err := s.GetTodoById(5)
	if err != nil {
		t.Fatalf("unexpected error : %s", err)
	}
	if todo.Title != testTitle {
		t.Fatalf("expected todo with title %s, got %s", testTitle, todo.Title)
	}
}

func TestUpdateTodo(t *testing.T) {
	s := newTodoMemoryStore(t)

	newTitle := "abc"
	newDone := true
	updatedTodo, err := s.UpdateTodo(1, newTitle, newDone)
	if err != nil {
		t.Fatalf("unexpected error : %s", err)
	}
	if updatedTodo.Title != newTitle || updatedTodo.Done != newDone {
		t.Fatalf("expected : {%s, %t}, got : {%s, %t}", newTitle, newDone, updatedTodo.Title, updatedTodo.Done)
	}
}

func TestUpdateTodoPartial(t *testing.T) {
	s := newTodoMemoryStore(t)

	// old todo
	oldTodo, err := s.GetTodoById(1)
	if err != nil {
		t.Fatalf("unexpected error : %s", err)
	}

	newTitle := "abc"
	oldDone := oldTodo.Done

	updatedTodo, err := s.UpdateTodo(1, newTitle, oldDone)
	if err != nil {
		t.Fatalf("unexpected error : %s", err)
	}
	if updatedTodo.Title != newTitle || updatedTodo.Done != oldDone {
		t.Fatalf("expected : {%s, %t}, got : {%s, %t}", newTitle, oldDone, updatedTodo.Title, updatedTodo.Done)
	}
}

func TestUpdateTodoNotFound(t *testing.T) {
	s := newTodoMemoryStore(t)

	_, err := s.UpdateTodo(999, "title", false)
	if err == nil {
		t.Fatalf("expected error for missing todo")
	}
}

func TestDeleteTodo(t *testing.T) {
	s := newTodoMemoryStore(t)

	err := s.DeleteTodo(1)
	if err != nil {
		t.Fatalf("unexpected error : %s", err)
	}
	if got, _ := s.GetAllTodos(false, false, "", 0, 0); len(got) != (3) {
		t.Fatalf("expected 3 total todos, got %d", len(got))
	}
}

func TestDeleteTodoNotFound(t *testing.T) {
	s := newTodoMemoryStore(t)

	err := s.DeleteTodo(999)
	if err == nil {
		t.Fatalf("expected error for missing todo")
	}
}
