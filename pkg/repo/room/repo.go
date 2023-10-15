package room

import (
	cst "Telegram/pkg/constants"
	"Telegram/pkg/storage/room"
	"strings"
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

	rooms, err := r.storage.GetRoomInfo(name)
	if err != nil {
		return "", err
	}

	if rooms == nil {
		return cst.RoomDoesntExist, nil
	}

	var res strings.Builder
	res.WriteString(cst.RoomInfo + "\n")
	for _, r := range rooms {
		res.WriteString(r.String() + "\n")
	}

	return res.String(), nil
}

//func (r *repository) GetFreeRooms(building, hasDot, hasProjector string) ([]models.Room, error) {
//	rooms, err := r.storage.GetFreeRooms(building, hasDot, hasProjector)
//	if err != nil {
//		return nil, err
//	}
//
//	return rooms, nil
//}
