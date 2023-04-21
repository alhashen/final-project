package repository

import (
	"finalproj/helper"
	"finalproj/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type PhotoRepo interface {
	CreatePhoto(model.Photo) (model.Photo, error)
	GetPhoto(uint) (model.Photo, error)
	GetAllPhoto(uint) ([]*model.Photo, error)
	UpdatePhoto(model.Photo, uint, uint) (model.Photo, error)
	DeletePhoto(uint, uint) (string, error)
}

func (r Repo) CreatePhoto(photo model.Photo) (res model.Photo, err error) {

	err = r.gorm.Create(&photo).Scan(&res).Error
	if err != nil {
		return res, err
	}

	err = r.gorm.Preload("User").First(&res).Error
	if err != nil {
		return res, err
	}

	res.User.HideSensitive()

	return res, nil

}

func (r Repo) GetPhoto(photoId uint) (res model.Photo, err error) {

	err = r.gorm.Preload("Comments").First(&res, "id = ?", photoId).Error
	if err != nil {
		return res, err
	}

	err = r.gorm.Preload("User").First(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetAllPhoto(userId uint) (res []*model.Photo, err error) {

	err = r.gorm.Find(&res, "user_id = ?", userId).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) UpdatePhoto(photo model.Photo, userId, photoId uint) (res model.Photo, err error) {

	qry := r.gorm.Debug().Find(&res, "id = ?", photoId)
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

	//Validates then updates
	err = photo.Validate()
	if err != nil {
		return res, err
	}
	photo.UpdatedAt = time.Now()
	err = qry.Debug().Updates(photo).Error
	if err != nil {
		return res, err
	}

	err = r.gorm.Preload("User").Preload("Comments").First(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) DeletePhoto(userId, photoId uint) (res string, err error) {

	photo := model.Photo{}
	qry := r.gorm.Debug().Find(&photo, "id = ?", photoId)
	err = qry.Error

	if qry.RowsAffected < 1 {
		return res, gorm.ErrRecordNotFound
	}

	if photo.UserID != userId {
		return res, helper.ErrForbiddenAccess
	}

	err = qry.Debug().Delete(model.Photo{}).Error
	if err != nil {
		return res, err
	}

	return fmt.Sprintf("Photo with id: %v deleted successfully", photoId), nil

}
