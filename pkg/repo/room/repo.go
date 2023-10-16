package room

import (
	cst "Telegram/pkg/constants"
	"Telegram/pkg/storage/room"
)

type Repository interface {
	GetRoomInfo(building, number string) (string, error)
	//GetFreeRooms(building, hasDot, hasProjector, date, interval string) ([]models.Room, error)
}

type repository struct {
	storage room.Storage
}

func NewRepository(storage room.Storage) Repository {
	return &repository{storage: storage}
}

func (r *repository) GetRoomInfo(building, number string) (string, error) {
	name := building + "-" + number

	roomInfo, err := r.storage.GetRoomInfo(name)
	if err != nil {
		return "", err
	}
	res := roomInfo.String()

	if res == "" {
		return cst.RoomDoesntExist, nil
	}

	return res, nil
}

//func (r *repository) GetFreeRooms(building, hasDot, hasProjector string) ([]models.Room, error) {
//	rooms, err := r.storage.GetFreeRooms(building, hasDot, hasProjector)
//	if err != nil {
//		return nil, err
//	}
//
//	return rooms, nil
//}
