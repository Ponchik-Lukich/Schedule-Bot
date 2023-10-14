package room

import (
	"Telegram/pkg/models"
	"Telegram/pkg/storage/room"
)

type Repository interface {
	GetRoomInfo(building, number string) []models.Room
}

type repository struct {
	storage room.Storage
}

func NewRepository(storage room.Storage) Repository {
	return &repository{storage: storage}
}

func (r *repository) GetRoomInfo(building, number string) []models.Room {
	name := building + "-" + number
	return r.storage.GetRoomInfo(name)
}
