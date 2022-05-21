package service

import (
	"twittie"
	"twittie/pkg/repository"
)

type Authorization interface {
	CreateUser(user twittie.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type Post interface {

}

type Service struct {
	Authorization
	Post
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}