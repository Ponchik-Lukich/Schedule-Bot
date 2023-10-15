package room

import "Telegram/pkg/models"

type Storage interface {
	GetRoomInfo(name string) ([]models.Room, error)
}
