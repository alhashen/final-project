package service

import (
	"finalproj/helper"
	"finalproj/model"
)

type UserService interface {
	RegisterUser(model.User) (model.User, error)
	LogInUser(map[string]any) error
}

func (s *Service) RegisterUser(user model.User) (model.User, error) {

	res, err := s.repo.RegisterUser(user)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) LogInUser(user map[string]any) error {
	pass := user["user_password"]
	err := s.repo.LogInUser(user)
	if err != nil {
		return err
	}

	valid := helper.ComparePassword(user["user_password"].(string), pass.(string))
	if !valid {
		return helper.ErrInvalidAuth
	}

	return nil
}
