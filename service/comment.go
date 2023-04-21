package service

import "finalproj/model"

type CommentService interface {
	CreateComment(model.Comment, uint, uint) (model.Comment, error)
	GetComment(uint) (model.Comment, error)
	GetAllComment(uint) ([]*model.Comment, error)
	UpdateComment(model.Comment, uint, uint) (model.Comment, error)
	DeleteComment(uint, uint) (string, error)
}

func (s *Service) CreateComment(comment model.Comment, userId, photoId uint) (res model.Comment, err error) {

	res, err = s.repo.CreateComment(comment, userId, photoId)
	if err != nil {
		return res, err
	}

	res.User.HideSensitive()

	return res, nil
}

func (s *Service) GetComment(commentId uint) (res model.Comment, err error) {

	res, err = s.repo.GetComment(commentId)
	if err != nil {
		return res, err
	}

	res.User.HideSensitive()
	// res.Photo.User.HideSensitive()

	return res, nil
}

func (s *Service) GetAllComment(photoId uint) (res []*model.Comment, err error) {

	res, err = s.repo.GetAllComment(photoId)
	if err != nil {
		return res, err
	}

	return res, nil

}

func (s *Service) UpdateComment(comment model.Comment, userId, commentId uint) (res model.Comment, err error) {

	res, err = s.repo.UpdateComment(comment, userId, commentId)
	if err != nil {
		return res, err
	}

	res.User.HideSensitive()
	return res, nil

}

func (s *Service) DeleteComment(userId, commentId uint) (res string, err error) {
	res, err = s.repo.DeleteComment(userId, commentId)
	if err != nil {
		return res, err
	}

	return res, nil
}
