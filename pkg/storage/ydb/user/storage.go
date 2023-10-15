package user

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

func (s *Storage) CreateUser(id int64) error {
	if err := s.db.Table("users").Create(&models.User{ChatID: id, State: "wait"}).Error; err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetUserState(id int64) (string, error) {
	var state string

	if err := s.db.Table("users").Where("chat_id = ?", id).Select("state").Scan(&state).Error; err != nil {
		return "", err
	}

	return state, nil
}

func (s *Storage) SetUserState(id int64, state string) error {
	if err := s.db.Table("users").Where("chat_id = ?", id).Update("state", state).Error; err != nil {
		return err
	}

	return nil
}
