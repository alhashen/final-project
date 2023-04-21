package service

import "finalproj/repository"

type Service struct {
	repo repository.RepoInterface
}

type ServiceInterface interface {
	UserService
	PhotoService
	CommentService
	SocialMediaService
}

func NewService(repo repository.RepoInterface) *Service {
	return &Service{repo: repo}
}
