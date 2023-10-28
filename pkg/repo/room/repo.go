package room

import (
	cst "Telegram/pkg/constants"
	"Telegram/pkg/storage/room"
	"Telegram/pkg/utils"
	"strings"
)

type Repository interface {
	GetRoomInfo(building, number string) (string, bool, error)
	GetFreeRooms(building, date, time string) (string, error)
}

type repository struct {
	storage room.Storage
}

func NewRepository(storage room.Storage) Repository {
	return &repository{storage: storage}
}

func (r *repository) GetRoomInfo(building, number string) (string, bool, error) {
	name := strings.Split(building, " ")[1] + "-" + number

	roomNames, err := r.storage.GetRoomsByName(name)
	if err != nil {
		return "", false, err
	}
	if roomNames == nil {
		return cst.RoomDoesntExist, false, nil
	}

	if len(roomNames) > 1 {
		var res strings.Builder
		res.WriteString(cst.RoomsFound)
		for _, name := range roomNames {
			res.WriteString(strings.Split(name, "-")[1] + "\n")
		}
		return res.String(), false, nil
	}

	roomInfo, err := r.storage.GetRoomInfo(name, building)
	if err != nil {
		return "", false, err
	}
	res := roomInfo.String(utils.GetCurrentWeek())

	return res, true, nil
}

func (r *repository) GetFreeRooms(building, date, time string) (string, error) {
	rooms, err := r.storage.GetFreeRooms(building, utils.GetWeekDay(date), utils.GetCurrentWeek(), time)
	if err != nil {
		return "", err
	}
	if rooms == nil {
		return cst.NoRoomsFound, nil
	}

	var res strings.Builder
	res.WriteString("<b>" + cst.FreeRooms + "</b>\n\n")
	for _, r := range rooms {
		res.WriteString(r.String() + "\n\n")
	}

	return res.String(), nil
}
