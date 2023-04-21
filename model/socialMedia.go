package model

import (
	"time"

	"github.com/go-ozzo/ozzo-validation/v4"
	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

type SocialMedia struct {
	GormModel
	Name           string `json:"name,omitempty"`
	SocialMediaURL string `json:"social_media_url,omitempty"`
	UserID         uint   `json:"user_id,omitempty"`
	User           *User  `json:"user,omitempty"`
}

func (SocialMedia) TableName() string {
	return "social_medias"
}

func (s SocialMedia) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Name, validation.Required),
		validation.Field(&s.SocialMediaURL, validation.Required),
	)
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) error {
	err := s.Validate()
	if err != nil {
		return err
	}

	s.CreatedAt = time.Now()
	return nil
}

type SwagSocialMedia struct {
	Name           string `json:"name,omitempty" example:"My Name"`
  SocialMediaURL string `json:"social_media_url,omitempty" example:"https://socialurl.com"`
}
