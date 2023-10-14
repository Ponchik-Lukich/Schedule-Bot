package room

import (
	"Telegram/pkg/models"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) Init() *gorm.DB {
	return s.db
}

func (s *Storage) GetRoomInfo(name string) []models.Room {
	var rooms []models.Room

	db := s.Init()
	db.Where("name like ?", name+"%").Find(&rooms)

	return rooms
}
