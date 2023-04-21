package repository

import (
	"finalproj/helper"
	"finalproj/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type SocialMediaRepo interface {
	CreateSocialMedia(model.SocialMedia, uint) (model.SocialMedia, error)
	GetSocialMedia(uint) (model.SocialMedia, error)
	GetAllSocialMedia() ([]*model.SocialMedia, error)
	UpdateSocialMedia(model.SocialMedia, uint, uint) (model.SocialMedia, error)
	DeleteSocialMedia(uint, uint) (string, error)
}

func (r Repo) CreateSocialMedia(social model.SocialMedia, userId uint) (res model.SocialMedia, err error) {

	social.UserID = userId

	err = r.gorm.Create(&social).Scan(&res).Error
	if err != nil {
		return res, err
	}

	err = r.gorm.Preload("User").First(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetSocialMedia(id uint) (res model.SocialMedia, err error) {
	err = r.gorm.Preload("User").First(&res, "id = ?", id).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetAllSocialMedia() (res []*model.SocialMedia, err error) {
	err = r.gorm.Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) UpdateSocialMedia(social model.SocialMedia, userId uint, socialId uint) (res model.SocialMedia, err error) {

	qry := r.gorm.Debug().Find(&res, "id = ?", socialId)
	err = qry.Error
	if err != nil {
		return res, err
	}

	if qry.RowsAffected < 1 {
		return res, gorm.ErrRecordNotFound
	}

	if res.UserID != userId {
		return res, helper.ErrForbiddenAccess
	}

	err = social.Validate()
	if err != nil {
		return res, err
	}
	social.UpdatedAt = time.Now()
	err = qry.Debug().Updates(social).Error
	if err != nil {
		return res, err
	}

	err = r.gorm.Preload("User").First(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) DeleteSocialMedia(userId, socialId uint) (res string, err error) {

	social := model.SocialMedia{}
	qry := r.gorm.Debug().Find(&social, "id = ?", socialId)
	err = qry.Error

	if qry.RowsAffected < 1 {
		return res, gorm.ErrRecordNotFound
	}

	if social.UserID != userId {
		return res, helper.ErrForbiddenAccess
	}

	err = r.gorm.Debug().Delete(&model.SocialMedia{}, "id = ?", socialId).Error
	if err != nil {
		return res, err
	}

	return fmt.Sprintf("Social media with id: %v deleted successfully", socialId), nil
}
