package repository

import(
	"context"
	"database/sql"
	"example.com/GolangAPI2/model"
)

type AuthRepository interface {
	Register(ctx context.Context, tx *sql.Tx, user model.User) error
	Login(ctx context.Context, tx *sql.Tx, user model.User) (model.User, error)
}