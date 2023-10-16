package room

import "Telegram/pkg/models"

type Storage interface {
	GetRoomInfo(name, building string) (models.RoomInfoDto, error)
	GetRoomsByName(name string) ([]string, error)
}
