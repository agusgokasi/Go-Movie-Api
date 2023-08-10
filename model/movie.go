package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Movie struct {
	ID          uint64    `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
	Title       string    `gorm:"not null" json:"title" form:"title" valid:"required~Title is required"`
	Description string    `json:"description" form:"description"`
	Rating      float64   `json:"rating" form:"rating"`
	Image       string    `json:"image" form:"image"`
}

func (m *Movie) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(m)
	if err != nil {
		return err
	}

	return
}
