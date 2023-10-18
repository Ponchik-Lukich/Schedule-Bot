package room

import (
	cst "Telegram/pkg/constants"
	"Telegram/pkg/models"
	"Telegram/pkg/storage/room"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) GetRoomInfo(name, building string) (models.RoomInfoDto, error) {
	var room models.Room
	var roomInfo models.RoomInfoDto
	var err error

	if building == cst.Any {
		err = s.db.Where("name = ?", name).First(&room).Error
	} else {
		err = s.db.Where("building = ?", building).Where("name = ?", name).First(&room).Error
	}

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.RoomInfoDto{}, nil
		} else {
			return models.RoomInfoDto{}, err
		}
	}

	var lessons []models.Lesson

	err = s.db.Preload("Tutors").Preload("Groups").Where("room_id = ?", room.ID).
		Order("week_day, time_from").Find(&lessons).Error
	if err != nil {
		return models.RoomInfoDto{}, err
	}

	roomInfo = models.RoomInfoDto{
		RoomName:     room.Name,
		IsAvailable:  room.IsAvailable,
		HasProjector: room.HasProjector,
		Lessons:      lessons,
	}

	return roomInfo, nil
}

func (s *Storage) GetRoomsByName(name string) ([]string, error) {
	var rooms []models.Room
	var roomNames []string
	pattern := strings.Replace(cst.RoomPattern, "number", name, 1)

	err := s.db.Table("rooms").Where("name ~ ?", pattern).Scan(&rooms).Error
	fmt.Println(len(rooms))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, err
		}
	}

	for _, r := range rooms {
		roomNames = append(roomNames, r.Name)
	}

	return roomNames, nil
}

func (s *Storage) GetFreeRooms(building string, day, week int, timeStr string) ([]models.FreeRoomDto, error) {
	defaultDateTime := "2020-01-01T" + timeStr + ":00.000000Z"
	targetTime, _ := time.Parse(time.RFC3339, defaultDateTime)
	var rooms []models.RawRoomDto
	var result []models.FreeRoomDto

	if building == cst.Any {
		if err := s.db.Raw(strings.Replace(room.GetFreeRooms, cst.Building, "", 1), targetTime, week, day).
			Scan(&rooms).Error; err != nil {
			return nil, err
		}
	} else {
		if err := s.db.Raw(room.GetFreeRooms, targetTime, building, week, day).Scan(&rooms).Error; err != nil {
			return nil, err
		}
	}

	for _, r := range rooms {
		result = append(result, models.FreeRoomDto{
			Building: r.Building,
			RoomName: r.RoomName,
			Interval: string(r.TimeFrom.UTC().Add(3*time.Hour).Format("15:04")) + " - " +
				string(r.TimeTo.UTC().Add(3*time.Hour).Format("15:04")),
		})
	}

	return result, nil
}
