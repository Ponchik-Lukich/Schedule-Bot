package room

import (
	"Telegram/pkg/models"
	"Telegram/pkg/storage/room"
)

type Repository interface {
	GetRoomInfo(number string) ([]models.Room, error)
	//GetFreeRooms(building, hasDot, hasProjector, date, interval string) ([]models.Room, error)
}

type repository struct {
	storage room.Storage
}

func NewRepository(storage room.Storage) Repository {
	return &repository{storage: storage}
}

func (r *repository) GetRoomInfo(number string) ([]models.Room, error) {
	name := ""

	rooms, err := r.storage.GetRoomInfo(name)
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

//func (r *repository) GetFreeRooms(building, hasDot, hasProjector string) ([]models.Room, error) {
//	rooms, err := r.storage.GetFreeRooms(building, hasDot, hasProjector)
//	if err != nil {
//		return nil, err
//	}
//
//	return rooms, nil
//}
