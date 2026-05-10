package storage

import (
	"context"
	"database/sql"
	"sandbox/internal/domain"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) *PostgresRepo {
	return &PostgresRepo{
		db: db,
	}
}

func (r *PostgresRepo) Create(ctx context.Context, todo domain.ToDo) (domain.ToDo, error) {
	query := `INSERT INTO todos (title, completed) VALUES ($1, $2) RETURNING id, title, completed`
	row := r.db.QueryRowContext(ctx, query, todo.Title, todo.Completed)
	if err := row.Scan(&todo.Id, &todo.Title, &todo.Completed); err != nil {
		return domain.ToDo{}, err
	}
	return todo, nil
}

func (r *PostgresRepo) GetById(ctx context.Context, id int) (domain.ToDo, error) {
	query := `SELECT id, title, completed FROM todos WHERE id = $1`

	var todo domain.ToDo
	row := r.db.QueryRowContext(ctx, query, id)
	if err := row.Scan(&todo.Id, &todo.Title, &todo.Completed); err != nil {
		return domain.ToDo{}, err
	}

	return todo, nil
}

func CreateTodoTable(ctx context.Context, db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		completed BOOLEAN DEFAULT FALSE
	);
	`
	_, err := db.ExecContext(ctx, query)

	return err
}
