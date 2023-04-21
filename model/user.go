package model

import (
	"finalproj/helper"
	"time"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string     `json:"username"`
	Email    string     `json:"user_email,omitempty"`
	Password string     `json:"user_password,omitempty"`
	Age      uint8      `json:"user_age,omitempty"`
	Photos   []*Photo   `json:"photos,omitempty" gorm:"OnUpdate:CASCADE,OnDelete:SET NULL"`
	Comments []*Comment `json:"comments,omitempty" gorm:"OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Username, validation.Required, validation.Length(5, 20)),
		validation.Field(&u.Email, validation.Required, validation.Length(5, 30), is.EmailFormat),
		validation.Field(&u.Password, validation.Required, validation.Length(6, 0)),
		validation.Field(&u.Age, validation.Required, validation.Min(uint8(13))),
	)
}

func (u *User) BeforeCreate(tx *gorm.DB) error {

	err := u.Validate()
	if err != nil {
		return err
	}

	u.GormModel.UpdatedAt = time.Now()
	u.Password = helper.HashPassword(u.Password)
	return nil
}

func (u *User) HideSensitive() {
	u.Password = ""
	u.Email = ""
	u.Age = 0
}

type SwagUserRegister struct {
	Username string     `json:"username" example:"my_username"`
  Email    string     `json:"user_email " example:"my_email@address.com"`
	Password string     `json:"user_password" example:"my_password"`
	Age      uint8      `json:"user_age,omitempty" example:"15"`
}

type SwagUserLogin struct {
	Email    string     `json:"user_email" example:"my_email@address.com"`
	Password string     `json:"user_password" example:"my_password"`
}
