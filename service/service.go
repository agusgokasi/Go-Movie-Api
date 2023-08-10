package service

import "MovieApi/repository"

type Service struct {
	repo repository.RepoInterface
}

type ServiceInterface interface {
	MovieService
}

func NewService(repo repository.RepoInterface) ServiceInterface {
	return &Service{repo: repo}
}
