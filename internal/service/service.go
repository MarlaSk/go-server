package service

import (
	"errors"
	"sandbox/internal/domain"
	"sandbox/internal/storage"
	"strings"
)

type TodoService struct {
	repo *storage.TodoStorage
}

func NewTodoService(r *storage.TodoStorage) *TodoService {
	return &TodoService{
		repo: r,
	}
}

func (s *TodoService) Create(title string) (domain.ToDo, error) {
	title = strings.TrimSpace(title)
	if len(title) < 3 {
		return domain.ToDo{}, errors.New("Слишком мало символов")
	}
	todo := domain.ToDo {
		Title: title,
	}
	return s.repo.Create(todo), nil
}

func (s *TodoService) GetById(id int) (domain.ToDo, error) {
	todo, err := s.repo.GetById(id)
	if err != nil {
		return domain.ToDo{}, err
	}
	return todo, nil
}