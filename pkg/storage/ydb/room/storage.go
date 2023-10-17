package room

import (
	cst "Telegram/pkg/constants"
	"Telegram/pkg/models"
	"fmt"
	"gorm.io/gorm"
	"log"
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

	err := s.db.Where("building = ?", building).Where("name = ?", name).First(&room).Error
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

	err := s.db.Table("rooms").Where("Re2::Match(?)(name)", pattern).Scan(&rooms).Error
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
	timeValue, _ := time.Parse(time.RFC3339, defaultDateTime)
	timeValue = timeValue.Add(-3 * time.Hour)

	var roomsInBuilding []models.Room
	err := s.db.Where("building = ?", building).Find(&roomsInBuilding).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	log.Println(len(roomsInBuilding), "rooms in building", building)
	log.Println(day, week, timeValue)

	var freeRooms []models.FreeRoomDto
	for _, room := range roomsInBuilding {
		var lessonsInRoom []models.Lesson
		err = s.db.Where("room_id = ?", room.ID).
			Where("week = 0 OR week = ?", week).
			Where("week_day = ?", day).
			Order("time_from").Find(&lessonsInRoom).Error
		log.Println(len(lessonsInRoom), "lessons in room", room.Name)
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}

		if len(lessonsInRoom) == 0 {
			freeRooms = append(freeRooms, models.FreeRoomDto{
				Building: room.Building,
				RoomName: room.Name,
				Interval: fmt.Sprintf("8:30 - 21:00"),
			})
			continue
		}

		if len(lessonsInRoom) == 1 {
			if timeValue.Before(lessonsInRoom[0].TimeFrom) {
				freeRooms = append(freeRooms, models.FreeRoomDto{
					Building: room.Building,
					RoomName: room.Name,
					Interval: fmt.Sprintf("8:30 - %s", lessonsInRoom[0].TimeFrom.Format("15:04")),
				})
				break
			} else if timeValue.After(lessonsInRoom[0].TimeTo) {
				freeRooms = append(freeRooms, models.FreeRoomDto{
					Building: room.Building,
					RoomName: room.Name,
					Interval: fmt.Sprintf("%s - 21:00", lessonsInRoom[0].TimeTo.Format("15:04")),
				})
				break
			} else {
				continue
			}
		}

		prevTime := ""
		for _, lesson := range lessonsInRoom {
			fmt.Println(lesson.TimeFrom, lesson.TimeTo, timeValue)
			if timeValue.After(lesson.TimeFrom) && timeValue.Before(lesson.TimeTo) {
				fmt.Println("in lesson")
				break
			}
			if timeValue.After(lesson.TimeTo) {
				prevTime = lesson.TimeTo.Format("15:04")
				fmt.Println("prev added")
			}

			if timeValue.Before(lesson.TimeFrom) && prevTime != "" {
				fmt.Println("time added")
				freeRooms = append(freeRooms, models.FreeRoomDto{
					Building: room.Building,
					RoomName: room.Name,
					Interval: fmt.Sprintf("%s - %s", prevTime, lesson.TimeFrom.Format("15:04")),
				})
				break
			}
		}
		if prevTime != "" && timeValue.After(lessonsInRoom[len(lessonsInRoom)-1].TimeTo) {
			fmt.Println("after added")
			freeRooms = append(freeRooms, models.FreeRoomDto{
				Building: room.Building,
				RoomName: room.Name,
				Interval: fmt.Sprintf("%s - 21:00", prevTime),
			})
		}
	}

	return freeRooms, nil
}
