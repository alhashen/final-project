package model

import (
	"time"

	"github.com/go-ozzo/ozzo-validation/v4"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	Message string `json:"comment_message"`
	UserID  uint   `json:"user_id,omitempty"`
	User    *User  `json:"user,omitempty"`
	PhotoID uint   `json:"photo_id,omitempty"`
	Photo   *Photo `json:"photo,omitempty"`
}

func (c Comment) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Message, validation.Required),
	)
}

func (c *Comment) BeforeCreate(tx *gorm.DB) error {
	err := c.Validate()
	if err != nil {
		return err
	}

	c.CreatedAt = time.Now()
	return nil
}

type SwagComment struct{
	Message string `json:"comment_message" example:"Insert you message here"`
}
