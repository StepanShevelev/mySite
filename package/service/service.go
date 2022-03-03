package service

import (
	"github.com/StepanShevelev/mySite"
	"github.com/StepanShevelev/mySite/package/repository"
	_ "github.com/StepanShevelev/mySite/package/repository"
)

type Authorization interface {
	CreateUser(user mySite.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list mySite.TodoList) (int, error)
	GetAll(userId int) ([]mySite.TodoList, error)
	GetById(userId, listId int) (mySite.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input mySite.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item mySite.TodoItem) (int, error)
	GetAll(userId, listId int) ([]mySite.TodoItem, error)
	GetById(userId, itemId int) (mySite.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input mySite.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		//	TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
