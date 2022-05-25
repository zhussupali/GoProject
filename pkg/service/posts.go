package service

import (
	"twittie"
	"twittie/pkg/repository"
)

type PostsService struct {
	repo repository.Post
}

func NewPostsService(repo repository.Post) *PostsService {
	return &PostsService{
		repo: repo,
	}
}

func (s *PostsService) Create(userId int, posts twittie.Post) (int, error) {
	return s.repo.Create(userId, posts)
}

func (s *PostsService) GetAll(userId int) ([]twittie.Post, error) {
	return s.repo.GetAll(userId)
}

func (s *PostsService) GetById(userId, listId int) (twittie.Post, error) {
	return s.repo.GetById(userId, listId)
}

func (s *PostsService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}