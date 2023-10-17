package room

import (
	cst "Telegram/pkg/constants"
	"Telegram/pkg/models"
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

	err := s.db.Table("rooms").Where("name REGEXP ?", pattern).Scan(&rooms).Error
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
	targetTime = targetTime.Add(-3 * time.Hour)
	var rooms []models.RawRoomDto
	var result []models.FreeRoomDto

	s.db.Raw(`
$r1 = (SELECT *
FROM rooms
WHERE rooms.building = ?
  AND rooms.is_available = true
  AND rooms.is_laboratory = false);

$r2 = (SELECT r1.*
FROM
    $r1 AS r1
    LEFT JOIN
    lessons
ON r1.id = lessons.room_id
WHERE
    ((lessons.week = ?
   OR lessons.week = 0)
  AND lessons.week_day = ?)
  AND
    (
    lessons.time_from <= ?
  AND lessons.time_to >= ?
    ));

$r3 = (SELECT r1.*
FROM $r1 as r1 LEFT ONLY JOIN $r2 as r2
ON r1.id = r2.id);

$r4 = (SELECT r3.id as id, r3.name as name, r3.building as building, MAX(lessons.time_to) as time_from
FROM $r3 AS r3 LEFT JOIN lessons
ON r3.id = lessons.room_id
WHERE 
((lessons.week = ?
   OR lessons.week = 0)
  AND lessons.week_day = ?)
AND lessons.time_to < ? OR lessons.id IS NULL
GROUP BY r3.id, r3.name, r3.building);

SELECT r4.id                  as id,
       r4.name                as name,
       r4.building            as building,
       r4.time_from           as time_from,
       MIN(lessons.time_from) as time_to
FROM $r4 AS r4 LEFT JOIN lessons
ON r4.id = lessons.room_id
WHERE
((lessons.week = ?
 OR lessons.week = 0)
AND lessons.week_day = ?)
AND lessons.time_from
> ?
   OR lessons.id IS NULL
GROUP BY r4.id, r4.name, r4.building, r4.time_from;
`, building, week, day, targetTime, targetTime, week, day, targetTime, week, day, targetTime).Scan(&rooms)

	for _, room := range rooms {
		result = append(result, models.FreeRoomDto{
			Building: room.Building,
			RoomName: room.RoomName,
			Interval: string(room.TimeFrom.Format("15:04")) + " - " + string(room.TimeTo.Format("15:04")),
		})
	}

	return result, nil
}
