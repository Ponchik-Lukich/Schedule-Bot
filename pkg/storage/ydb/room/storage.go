package room

import (
	"Telegram/pkg/models"
	"fmt"
	"gorm.io/gorm"
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

func (s *Storage) GetRoomInfo(name string) (models.RoomInfoDto, error) {
	var room models.Room
	var roomInfo models.RoomInfoDto

	err := s.db.Where("name LIKE ?", name+"%").First(&room).Error
	if err != nil {
		fmt.Println("AAAAAAAAAAAAA")
		return models.RoomInfoDto{}, err
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
		RoomName:       room.Name,
		IsAvailability: room.IsAvailable,
		HasProjector:   room.HasProjector,
		Lessons:        lessonInfos,
	}

	return roomInfo, nil
}
