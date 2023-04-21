package service

import (
	"finalproj/model"
)

type PhotoService interface {
	CreatePhoto(model.Photo) (model.Photo, error)
	GetPhoto(uint) (model.Photo, error)
	GetAllPhoto(uint) ([]*model.Photo, error)
	UpdatePhoto(model.Photo, uint, uint) (model.Photo, error)
	DeletePhoto(uint, uint) (string, error)
}

func (s *Service) CreatePhoto(photo model.Photo) (res model.Photo, err error) {

	res, err = s.repo.CreatePhoto(photo)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) GetPhoto(photoId uint) (res model.Photo, err error) {

	res, err = s.repo.GetPhoto(photoId)
	if err != nil {
		return res, err
	}

	res.User.HideSensitive()

	return res, nil
}

func (s *Service) GetAllPhoto(userId uint) (res []*model.Photo, err error) {

	res, err = s.repo.GetAllPhoto(userId)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) UpdatePhoto(photo model.Photo, userId, photoId uint) (res model.Photo, err error) {

	res, err = s.repo.UpdatePhoto(photo, userId, photoId)
	if err != nil {
		return res, err
	}

	res.User.HideSensitive()

	return res, nil
}

func (s *Service) DeletePhoto(userId, photoId uint) (res string, err error) {

	res, err = s.repo.DeletePhoto(userId, photoId)
	if err != nil {
		return res, err
	}

	return res, nil

}
