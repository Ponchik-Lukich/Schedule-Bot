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

//func (s *Storage) GetFreeRooms(building, date, timeStr string) ([]models.FreeRoomDto, error) {
//	var conflictingLessons []models.Lesson
//	defaultDateTime := "2020-01-01T" + timeStr + ":00.000000Z"
//	timeValue, _ := time.Parse(time.RFC3339, defaultDateTime)
//
//	err := s.db.Joins("JOIN rooms ON lessons.room_id = rooms.id").
//		Where("time_from <= ? AND time_to >= ?", timeValue, timeValue).
//		Where("rooms.building = ?", building).
//		Find(&conflictingLessons).Error
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return nil, err
//	}
//
//	var conflictingRoomIDs []string
//	for _, lesson := range conflictingLessons {
//		conflictingRoomIDs = append(conflictingRoomIDs, lesson.RoomID)
//	}
//	fmt.Println(len(conflictingLessons))
//
//	var availableRooms []models.Room
//	err = s.db.Table("rooms").Where("building = ?", building).Not("id IN ?", conflictingRoomIDs).Find(&availableRooms).Error
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return nil, err
//	}
//
//	var freeRooms []models.FreeRoomDto
//	for _, room := range availableRooms {
//		freeRooms = append(freeRooms, models.FreeRoomDto{
//			Building: room.Building,
//			RoomName: room.Name,
//			Interval: fmt.Sprintf("%s - next lesson time", timeStr),
//		})
//	}
//
//	return freeRooms, nil
//}

func (s *Storage) GetFreeRooms(building string, day, week int, timeStr string) ([]models.FreeRoomDto, error) {
	defaultDateTime := "2020-01-01T" + timeStr + ":00.000000Z"
	timeValue, _ := time.Parse(time.RFC3339, defaultDateTime)

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

		var prevEndTime time.Time
		for _, lesson := range lessonsInRoom {
			if !prevEndTime.IsZero() && prevEndTime.Before(lesson.TimeFrom) && prevEndTime.Before(timeValue) && lesson.TimeFrom.After(timeValue) {
				freeRooms = append(freeRooms, models.FreeRoomDto{
					Building: room.Building,
					RoomName: room.Name,
					Interval: fmt.Sprintf("%s-%s", prevEndTime.Format("15:04"), lesson.TimeFrom.Format("15:04")),
				})
			}
			prevEndTime = lesson.TimeTo
		}

		if (len(lessonsInRoom) == 0) || (prevEndTime.Before(timeValue)) {
			freeRooms = append(freeRooms, models.FreeRoomDto{
				Building: room.Building,
				RoomName: room.Name,
				Interval: fmt.Sprintf("%s-∞", prevEndTime.Format("15:04")),
			})
		}
	}

	return freeRooms, nil
}
