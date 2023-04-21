package model

import "time"

type GormModel struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoUpdateTime:false"`
	UpdatedAt time.Time `json:"updated_at"`
}
