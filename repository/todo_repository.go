package repository

import(
	"context"
	"database/sql"
	"example.com/GolangAPI2/model"
)

type ToDoRepository interface {
	Save(ctx context.Context, tx *sql.Tx, todo model.ToDo) model.ToDo
	Update(ctx context.Context, tx *sql.Tx, todo model.ToDo) model.ToDo
	Delete(ctx context.Context, tx *sql.Tx, todo model.ToDo)
	FindById(ctx context.Context, tx *sql.Tx, todoId int) (model.ToDo, error)
	GetAll(ctx context.Context, tx *sql.Tx) []model.ToDo
}