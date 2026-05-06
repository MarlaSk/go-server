package storage

import (
	"errors"
	"sandbox/internal/domain"
	"sync"
)

type TodoStorage struct {
	todos  map[int]domain.ToDo
	nextID int
	mu     sync.Mutex
}

func NewTodoStorage() *TodoStorage {
  m := make(map[int]domain.ToDo)
	return &TodoStorage{
		todos: m,
		nextID: 1,
	}
}

func (t *TodoStorage) Create(todo domain.ToDo) domain.ToDo {
	t.mu.Lock()
	defer t.mu.Unlock()

	todo.Id = t.nextID
	t.nextID++

	t.todos[todo.Id] = todo

	return todo
}

func (t *TodoStorage) GetById(id int) (domain.ToDo, error) {
	todo, ok := t.todos[id]
	if !ok {
		return domain.ToDo{}, errors.New("Такого id не существует")
	}

	return todo, nil
}
