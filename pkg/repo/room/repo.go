package room

import (
	cst "Telegram/pkg/constants"
	"Telegram/pkg/storage/room"
	"regexp"
	"strings"
)

type Repository interface {
	GetRoomInfo(building, number string) (string, bool, error)
	//GetFreeRooms(building, hasDot, hasProjector, date, interval string) ([]models.Room, error)
}

type repository struct {
	storage room.Storage
}

func NewRepository(storage room.Storage) Repository {
	return &repository{storage: storage}
}

func (r *repository) GetRoomInfo(building, number string) (string, bool, error) {
	name := strings.Split(building, " ")[1] + "-" + number
	newNames := make([]string, 0)

	roomNames, err := r.storage.GetRoomsByName(building, name)
	if err != nil {
		return "", false, err
	}
	if roomNames == nil {
		return cst.RoomDoesntExist, false, nil
	}

	if len(roomNames) > 1 {
		for _, roomName := range roomNames {
			r, _ := regexp.Compile(cst.RoomPattern)
			if r.MatchString(roomName) {
				newNames = append(newNames, roomName)
			}
		}
	}

	if len(newNames) > 1 {
		var res strings.Builder
		res.WriteString(cst.RoomsFound)
		for _, name := range newNames {
			res.WriteString(name + "\n")
		}
		return res.String(), false, nil
	}

	roomInfo, err := r.storage.GetRoomInfo(name, building)
	if err != nil {
		return "", false, err
	}
	res := roomInfo.String()

	return res, true, nil
}

//func (r *repository) GetFreeRooms(building, hasDot, hasProjector string) ([]models.Room, error) {
//	rooms, err := r.storage.GetFreeRooms(building, hasDot, hasProjector)
//	if err != nil {
//		return nil, err
//	}
//
//	return rooms, nil
//}
