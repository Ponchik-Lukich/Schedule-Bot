package room

import (
	"Telegram/pkg/models"
	"Telegram/pkg/storage/room"
)

type Repository interface {
	GetRoomInfo(building, number string) ([]models.Room, error)
}

type repository struct {
	storage room.Storage
}

func NewRepository(storage room.Storage) Repository {
	return &repository{storage: storage}
}

func (r *repository) GetRoomInfo(building, number string) ([]models.Room, error) {
	name := building + "-" + number

	rooms, err := r.storage.GetRoomInfo(name)
	if err != nil {
		return nil, err
	}

	return rooms, nil
}
