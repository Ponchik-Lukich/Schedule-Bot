package room

import (
	cst "Telegram/pkg/constants"
	"Telegram/pkg/models"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) *Storage {
	return &Storage{db: db}
}

//func (s *Storage) GetRoomInfo(name string) ([]models.Room, error) {
//	var rooms []models.Room
//
//	err := s.db.Where("name like ?", name+"%").Find(&rooms).Error
//	if err != nil {
//		return nil, err
//	}
//
//	return rooms, nil
//}

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
	var lessonInfos []models.LessonInfoDto

	err = s.db.Preload("Tutors").Preload("Groups").Where("room_id = ?", room.ID).Order("week_day, time_from").Find(&lessons).Error
	if err != nil {
		return models.RoomInfoDto{}, err
	}

	for _, lesson := range lessons {
		groupNames := make([]string, len(lesson.Groups))
		for i, group := range lesson.Groups {
			groupNames[i] = group.Name
		}

		tutorNames := make([]string, len(lesson.Tutors))
		for i, tutor := range lesson.Tutors {
			tutorNames[i] = tutor.ShortName
		}

		lessonInfo := models.LessonInfoDto{
			TimeFrom:   lesson.TimeFrom,
			TimeTo:     lesson.TimeTo,
			TutorNames: tutorNames,
			GroupNames: groupNames,
		}

		lessonInfos = append(lessonInfos, lessonInfo)
	}

	roomInfo = models.RoomInfoDto{
		RoomName:     room.Name,
		IsAvailable:  room.IsAvailable,
		HasProjector: room.HasProjector,
		Lessons:      lessonInfos,
	}

	return roomInfo, nil
}

func (s *Storage) GetRoomsByName(building, name string) ([]string, error) {
	var rooms []models.Room
	var roomNames []string
	pattern := strings.Replace(cst.RoomPattern, "number", name, 1)

	fmt.Println(name, building, pattern)
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
