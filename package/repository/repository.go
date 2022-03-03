package repository

import (
	"github.com/StepanShevelev/mySite"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user mySite.User) (int, error)
	GetUser(username, password string) (mySite.User, error)
	//ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list mySite.TodoList) (int, error)
	GetAll(userId int) ([]mySite.TodoList, error)
	GetById(userId, listId int) (mySite.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input mySite.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item mySite.TodoItem) (int, error)
	GetAll(userId, listId int) ([]mySite.TodoItem, error)
	GetById(userId, itemId int) (mySite.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input mySite.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
