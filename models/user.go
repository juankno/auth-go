package models

import (
	"time"
)

type User struct {
	ID        uint       `gorm:"column:id;primaryKey"`
	Name      string     `json:"name" validate:"required"`
	Email     string     `json:"email" validate:"required,email" gorm:"unique"`
	Password  string     `json:"password" validate:"required,min=6"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index"`
}
