package repository

import (
	"twittie"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user twittie.User) (int, error)
}

type Post interface {

}

type Repository struct {
	Authorization
	Post
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}