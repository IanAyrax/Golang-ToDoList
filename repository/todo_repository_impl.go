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

type ToDoRepositoryImpl struct {
}

func NewToDoRepository() ToDoRepository {
	return &ToDoRepositoryImpl{}
}

func (repository *ToDoRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, todo model.ToDo) model.ToDo {
	fmt.Println("ToDoRepository Ok!")
	SQL := "insert into tb_todo(user_id, title, created_at, updated_at) values(?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, todo.UserId, todo.Title, time.Now(), time.Now())
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	todo.Id = int(id)
	return todo
}

func (repository *ToDoRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, todo model.ToDo) model.ToDo {
	SQL := "update tb_todo set user_id = ?, title = ?, updated_at = ? where id_todo = ?"
	_, err := tx.ExecContext(ctx, SQL, todo.UserId, todo.Title, time.Now(), todo.Id)
	helper.PanicIfError(err)

	return todo
}

func (repository *ToDoRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, todo model.ToDo) {
	SQL := "delete from tb_todo where id_todo = ?"
	_, err := tx.ExecContext(ctx, SQL, todo.Id)	
	helper.PanicIfError(err)
}

func (repository *ToDoRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, todoId int) (model.ToDo, error) {
	SQL := "select id_todo, user_id, title from tb_todo where id_todo = ?"
	rows, err := tx.QueryContext(ctx, SQL, todoId)
	helper.PanicIfError(err)
	defer rows.Close()
	
	todo := model.ToDo{}
	if rows.Next(){
		err := rows.Scan(&todo.Id, &todo.UserId, &todo.Title)
		helper.PanicIfError(err)
		return todo, nil
	}else{
		return todo, errors.New("Not Found")
	}
}

func (repository *ToDoRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []model.ToDo {
	fmt.Println("GetAll Repository OK")
	SQL := "select id_todo, user_id, title from tb_todo"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var todos []model.ToDo
	for rows.Next(){
		todo := model.ToDo{}
		err := rows.Scan(&todo.Id, &todo.UserId, &todo.Title)
		helper.PanicIfError(err)
		todos = append(todos, todo)
	}

	return todos
}

