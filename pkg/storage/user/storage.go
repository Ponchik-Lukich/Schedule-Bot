package user

import "Telegram/pkg/models"

type Storage interface {
	GetUser(id int64) (models.User, error)
	CreateUser(id int64) error
	UpdateUser(id int64, updates map[string]any) error
}
