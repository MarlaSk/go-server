package service

import (
	"context"
	"errors"
	"sandbox/internal/domain"
	"strings"
)

type TodoRepo interface {
	Create(ctx context.Context, todo domain.ToDo)(domain.ToDo, error)
	GetById(ctx context.Context, id int)(domain.ToDo, error)
}

type TodoService struct {
	repo  TodoRepo
}

func NewTodoService(r TodoRepo) *TodoService {
	return &TodoService{
		repo: r,
	}
}

func (s *TodoService) Create(ctx context.Context, title string) (domain.ToDo, error) {
	title = strings.TrimSpace(title)
	if len(title) < 3 {
		return domain.ToDo{}, errors.New("Слишком мало символов")
	}
	todo := domain.ToDo {
		Title: title,
	}
	return s.repo.Create(ctx, todo)
}

func (s *TodoService) GetById(ctx context.Context, id int) (domain.ToDo, error) {
	todo, err := s.repo.GetById(ctx, id)
	if err != nil {
		return domain.ToDo{}, err
	}
	return todo, nil
}