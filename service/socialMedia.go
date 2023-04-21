package service

import "finalproj/model"

type SocialMediaService interface {
	CreateSocialMedia(model.SocialMedia, uint) (model.SocialMedia, error)
	GetSocialMedia(uint) (model.SocialMedia, error)
	GetAllSocialMedia() ([]*model.SocialMedia, error)
	UpdateSocialMedia(model.SocialMedia, uint, uint) (model.SocialMedia, error)
	DeleteSocialMedia(uint, uint) (string, error)
}

func (s *Service) CreateSocialMedia(social model.SocialMedia, userId uint) (res model.SocialMedia, err error) {

	res, err = s.repo.CreateSocialMedia(social, userId)
	if err != nil {
		return res, err
	}

	res.User.HideSensitive()
	return res, nil

}

func (s *Service) GetSocialMedia(id uint) (res model.SocialMedia, err error) {

	res, err = s.repo.GetSocialMedia(id)
	if err != nil {
		return res, err
	}

	res.User.HideSensitive()
	return res, nil
}

func (s *Service) GetAllSocialMedia() (res []*model.SocialMedia, err error) {

	res, err = s.repo.GetAllSocialMedia()
	if err != nil {
		return res, err
	}

	return res, nil

}

func (s *Service) UpdateSocialMedia(social model.SocialMedia, userId uint, socialId uint) (res model.SocialMedia, err error) {

	res, err = s.repo.UpdateSocialMedia(social, userId, socialId)
	if err != nil {
		return res, err
	}

	res.User.HideSensitive()
	return res, nil

}

func (s *Service) DeleteSocialMedia(userId uint, socialId uint) (res string, err error) {

	res, err = s.repo.DeleteSocialMedia(userId, socialId)
	if err != nil {
		return res, err
	}

	return res, nil

}
