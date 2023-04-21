package repository

import (
	"finalproj/helper"
	"finalproj/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type CommentRepo interface {
	CreateComment(model.Comment, uint, uint) (model.Comment, error)
	GetComment(uint) (model.Comment, error)
	GetAllComment(uint) ([]*model.Comment, error)
	UpdateComment(model.Comment, uint, uint) (model.Comment, error)
	DeleteComment(uint, uint) (string, error)
}

func (r Repo) CreateComment(comment model.Comment, userId uint, photoId uint) (res model.Comment, err error) {

	comment.UserID = userId
	comment.PhotoID = photoId

	check := r.gorm.Debug().Find(&model.Photo{}, "id", photoId)
	if check.RowsAffected < 1 {
		return res, gorm.ErrRecordNotFound
	}

	err = r.gorm.Create(&comment).Scan(&res).Error
	if err != nil {
		return res, err
	}

	err = r.gorm.Preload("User").First(&res).Error
	if err != nil {
		return res, err
	}

	err = r.gorm.Preload("Photo").First(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetComment(commentId uint) (res model.Comment, err error) {

	err = r.gorm.First(&res, "id = ?", commentId).Error
	if err != nil {
		return res, err
	}

	err = r.gorm.Preload("User").Preload("Photo").First(&res).Error
	if err != nil {
		return res, err
	}

	/* err = r.gorm.Preload("User").First(&res.Photo).Error
	if err != nil {
		return res, err
	} */

	return res, nil
}

func (r Repo) GetAllComment(photoId uint) (res []*model.Comment, err error) {
	qry := r.gorm.Debug().Find(&res, "photo_id = ?", photoId)
	err = qry.Error
	if err != nil {
		return res, err
	}

	if qry.RowsAffected < 1 {
		return res, gorm.ErrRecordNotFound
	}

	return res, nil
}

func (r Repo) UpdateComment(comment model.Comment, userId, commentId uint) (res model.Comment, err error) {
	qry := r.gorm.Debug().Find(&res, "id = ?", commentId)
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

	err = comment.Validate()
	if err != nil {
		return res, err
	}
	comment.UpdatedAt = time.Now()
	err = qry.Debug().Updates(comment).Error
	if err != nil {
		return res, err
	}

	err = r.gorm.Preload("User").Preload("Photo").First(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) DeleteComment(userId, commentId uint) (res string, err error) {

	comment := model.Comment{}
	qry := r.gorm.Debug().Find(&comment, "id = ?", commentId)
	err = qry.Error

	if qry.RowsAffected < 1 {
		return res, gorm.ErrRecordNotFound
	}

	if comment.UserID != userId {
		return res, helper.ErrForbiddenAccess
	}

	err = r.gorm.Debug().Delete(model.Comment{}, "id = ?", commentId).Error
	if err != nil {
		return res, err
	}

	return fmt.Sprintf("Comment with id: %v deleted successfully", commentId), nil
}
