package model

import (
	"time"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
  Title    string     `json:"photo_title,omitempty" example:"photo title"`
	Caption  string     `json:"photo_caption,omitempty" example:"photo caption"`
	PhotoURL string     `json:"photo_url,omitempty" example:"photo url"`
	UserID   uint       `json:"user_id,omitempty"`
	User     *User      `json:"user,omitempty"`
	Comments []*Comment `json:"comments,omitempty"`
}

func (p Photo) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Title, validation.Required, validation.Length(5, 0)),
		validation.Field(&p.PhotoURL, validation.Required, validation.Length(5, 50), is.URL),
	)
}

func (p *Photo) BeforeCreate(tx *gorm.DB) error {
	err := p.Validate()
	if err != nil {
		return err
	}

	p.CreatedAt = time.Now()
	return nil
}

type SwagPhoto struct {
  Title    string     `json:"photo_title,omitempty" example:"Photo Title"`
	Caption  string     `json:"photo_caption,omitempty" example:"Your caption" `
  PhotoURL string     `json:"photo_url,omitempty" example:"http://img-url.com"`
}
