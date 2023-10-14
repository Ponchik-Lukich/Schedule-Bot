package models

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	gorm.Model
	DeletedAt *time.Time `json:"deleted_at"`
}
