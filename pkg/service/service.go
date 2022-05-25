package service

import (
	"twittie"
	"twittie/pkg/repository"
)

type Authorization interface {
	CreateUser(user twittie.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Post interface {
	Create(userId int, input twittie.Post) (int, error)
}

type Service struct {
	Authorization
	Post
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Post: NewPostsService(repos.Post),
	}
}