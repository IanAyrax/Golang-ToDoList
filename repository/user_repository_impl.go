package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"example.com/GolangAPI2/model"
	"example.com/GolangAPI2/helper"
	"fmt"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user model.User) model.User {
	SQL := "insert into tb_user(full_name, email, password, id_role, created_at, updated_at) values(?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, user.FullName, user.Password, user.RoleId, time.Now(), time.Now())
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.UserId = int(id)
	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user model.User) model.User {
	SQL := "update tb_user set full_name = ?, email = ?, updated_at = ? where id_user = ?"
	_, err := tx.ExecContext(ctx, SQL, user.FullName, user.Email, time.Now(), user.UserId)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user model.User) {
	SQL := "delete from tb_user where id_user = ?"
	_, err := tx.ExecContext(ctx, SQL, user.UserId)	
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (model.User, error) {
	SQL := "select id_user, full_name, email from tb_user where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()
	
	user := model.User{}
	if rows.Next(){
		err := rows.Scan(&user.UserId, &user.FullName, &user.Email)
		helper.PanicIfError(err)
		return user, nil
	}else{
		return user, errors.New("Not Found")
	}
}

func (repository *UserRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []model.User {
	fmt.Println("GetAll Repository OK")
	SQL := "select id_user, full_name, email from tb_user"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []model.User
	for rows.Next(){
		user := model.User{}
		err := rows.Scan(&user.UserId, &user.FullName, &user.Email)
		helper.PanicIfError(err)
		users = append(users, user)
	}

	return users
}

