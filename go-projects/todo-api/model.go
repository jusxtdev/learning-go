package main

import (
	"errors"
	"slices"
	"strings"
)

type Todo struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func (s *MemoryStore) Seed() {
	s.AddTodo("buy milk")
	s.AddTodo("Test api")
	s.AddTodo("Make report")
	s.AddTodo("complete assignments")
}

const ()

func (s *MemoryStore) GetAllTodos(DoneFilter bool, DoneValue bool, title string, page int, limit int) ([]Todo, PaginationData) {
	/*
		Return a copy of the Todos
		cuz if you return the reference to the original slice and function ends (hence slice is locked)
		the caller may use that reference to read the data / json serialize it
		AND if some other routine in modifying the data, you might get DATA RACE

		Hence, return a copy of Todos

		IF DoneFilter is given, then only do filtering of "done" == DoneSearch, otherwise no filtering on that
		IF title != "" then search for the given title
	*/
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []Todo
	for _, todo := range s.Todos {
		if title != "" && !strings.Contains(todo.Title, title) {
			continue
		}

		if DoneFilter && todo.Done != DoneValue {
			continue
		}

		result = append(result, todo)
	}

	var pd PaginationData

	if page == 0 || limit == 0 {
		pd.Page = 1
		pd.Limit = len(s.Todos)
		return result, pd
	}

	start := (page - 1) * limit
	end := start + limit

	if start >= len(s.Todos) {
		pd.Page = 1
		pd.Limit = len(s.Todos)
		return result, pd
	} else {
		if end > len(s.Todos) {
			end = len(s.Todos)
		}
		pd.Page = page
		pd.Limit = limit
		result = result[start:end]
		return result, pd
	}
}

func (s *MemoryStore) GetTodoById(id int) (Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, v := range s.Todos {
		if v.Id == id {
			return v, nil
		}
	}
	return Todo{}, errors.New("todo not found")
}

func (s *MemoryStore) AddTodo(title string) Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	// create new todo
	t := Todo{
		Id:    s.getNewTodoId(),
		Title: title,
		Done:  false,
	}
	s.Todos = append(s.Todos, t)

	return t
}

func (s *MemoryStore) UpdateTodo(id int, newTitle string, done bool) (Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// get index of todo to update
	index, err := s.getIndexById(id)
	if err != nil {
		return Todo{}, err
	}

	// update todo based on fields given
	if newTitle != "" {
		s.Todos[index].Title = newTitle
	}
	s.Todos[index].Done = done

	// respond
	return s.Todos[index], nil
}

func (s *MemoryStore) DeleteTodo(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	index, err := s.getIndexById(id)
	if err != nil {
		return err
	}

	// delete
	s.Todos = slices.Delete(s.Todos, index, index+1)
	return nil
}

/* HELPER */
// they don't need the locks because their Callers (methods above) have the locks
// hence no lock in helpers to prevent deadlock

func (s *MemoryStore) getNewTodoId() int {
	if len(s.Todos) == 0 {
		return 1
	}
	last_todo := s.Todos[len(s.Todos)-1]
	return last_todo.Id + 1
}

func (s *MemoryStore) getIndexById(id int) (int, error) {
	// using index based for because in `for i, v := range Todos`,
	// v is a copy of todo hence changing v doesn't change original todo
	for index := range s.Todos {
		if s.Todos[index].Id == id {
			return index, nil
		}
	}
	return 0, errors.New("not found")
}
