package repository

import (
	"twittie"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user twittie.User) (int, error)
	GetUser(username, password string) (twittie.User, error)
}

type Post interface {
	Create(userId int, posts twittie.Post) (int, error)
	GetAll(userId int) ([]twittie.Post, error)
	GetById(userId, listId int) (twittie.Post, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input twittie.UpdatePostInput) error
}

type Repository struct {
	Authorization
	Post
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Post: NewPostPostgres(db),
	}
}