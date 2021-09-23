package repository

import(
	"context"
	"database/sql"
	"example.com/GolangAPI2/model"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user model.User) model.User
	Update(ctx context.Context, tx *sql.Tx, user model.User) model.User
	Delete(ctx context.Context, tx *sql.Tx, user model.User)
	FindById(ctx context.Context, tx *sql.Tx, userId int) (model.User, error)
	GetAll(ctx context.Context, tx *sql.Tx) []model.User
}