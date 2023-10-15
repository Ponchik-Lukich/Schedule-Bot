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

func (s *Storage) GetRoomInfo(name string) ([]models.Room, error) {
	var rooms []models.Room

	err := s.db.Where("name like ?", name+"%").Find(&rooms).Error
	if err != nil {
		return nil, err
	}

	return rooms, nil
}
