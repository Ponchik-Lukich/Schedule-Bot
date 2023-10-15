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

func (s *Storage) GetUser(id int64) (models.User, error) {
	var user models.User

	if err := s.db.Table("users").Where("chat_id = ?", id).Select("state").Scan(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s *Storage) UpdateUser(id int64, updates map[string]any) error {
	if err := s.db.Table("users").Where("chat_id = ?", id).Updates(updates).Error; err != nil {
		return err
	}

	return nil
}
