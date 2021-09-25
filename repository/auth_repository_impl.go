package repository

import(
	"context"
	"database/sql"
	"errors"
	//"time"
	"example.com/GolangAPI2/model"
	"example.com/GolangAPI2/helper"
	//"fmt"
)

type AuthRepositoryImpl struct {
}

func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

func (authRepository *AuthRepositoryImpl)Register(ctx context.Context, tx *sql.Tx, user model.User) error {
	// SQL := "insert into tb_user(full_name, email, password, id_role, created_at, updated_at) values(?, ?, ?, ?)"
	// result, err := tx.ExecContext(ctx, SQL, user.FullName, user.Email, user.Password, 2, time.Now(), time.Now())
	// helper.PanicIfError(err)

	// id, err := result.LastInsertId()
	// helper.PanicIfError(err)

	// user.Id = int(id)
	// return todo
	return nil
}

func (authRepository *AuthRepositoryImpl)Login(ctx context.Context, tx *sql.Tx, user model.User) (int, error) {
	SQL := "select id_user, email, password from tb_user where email = ? and password = ?"
	rows, err := tx.QueryContext(ctx, SQL, user.Email, user.Password)
	helper.PanicIfError(err)
	defer rows.Close()
	
	logged_user := model.User{}
	if rows.Next(){
		err := rows.Scan(&logged_user.UserId, &logged_user.Email, &logged_user.Password)
		helper.PanicIfError(err)
		return logged_user.UserId, nil
	}else{
		return 0, errors.New("Email Not Found or the password is wrong !")
	}
}
